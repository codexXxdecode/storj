// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package segments

import (
	"context"
	"io"
	"math/rand"
	"sync"
	"time"

	"github.com/vivint/infectious"
	"gopkg.in/spacemonkeygo/monkit.v2"

	"storj.io/common/pb"
	"storj.io/common/ranger"
	"storj.io/common/storj"
	"storj.io/storj/uplink/ecclient"
	"storj.io/storj/uplink/eestream"
	"storj.io/storj/uplink/metainfo"
)

var (
	mon = monkit.Package()
)

// Meta info about a segment
type Meta struct {
	Modified   time.Time
	Expiration time.Time
	Size       int64
	Data       []byte
}

// ListItem is a single item in a listing
type ListItem struct {
	Path     storj.Path
	Meta     Meta
	IsPrefix bool
}

// Store for segments
type Store interface {
	// Ranger creates a ranger for downloading erasure codes from piece store nodes.
	Ranger(ctx context.Context, info storj.SegmentDownloadInfo, limits []*pb.AddressedOrderLimit, objectRS storj.RedundancyScheme) (ranger.Ranger, error)
	Put(ctx context.Context, data io.Reader, expiration time.Time, limits []*pb.AddressedOrderLimit, piecePrivateKey storj.PiecePrivateKey) (_ []*pb.SegmentPieceUploadResult, size int64, err error)
	Delete(ctx context.Context, streamID storj.StreamID, segmentIndex int32) (err error)
}

type segmentStore struct {
	metainfo *metainfo.Client
	ec       ecclient.Client
	rs       eestream.RedundancyStrategy
	rngMu    sync.Mutex
	rng      *rand.Rand
}

// NewSegmentStore creates a new instance of segmentStore
func NewSegmentStore(metainfo *metainfo.Client, ec ecclient.Client, rs eestream.RedundancyStrategy) Store {
	return &segmentStore{
		metainfo: metainfo,
		ec:       ec,
		rs:       rs,
		rng:      rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// Put uploads a segment to an erasure code client
func (s *segmentStore) Put(ctx context.Context, data io.Reader, expiration time.Time, limits []*pb.AddressedOrderLimit, piecePrivateKey storj.PiecePrivateKey) (_ []*pb.SegmentPieceUploadResult, size int64, err error) {
	defer mon.Task()(&ctx)(&err)

	sizedReader := SizeReader(NewPeekThresholdReader(data))
	successfulNodes, successfulHashes, err := s.ec.Put(ctx, limits, piecePrivateKey, s.rs, sizedReader, expiration)
	if err != nil {
		return nil, size, Error.Wrap(err)
	}

	uploadResults := make([]*pb.SegmentPieceUploadResult, 0, len(successfulNodes))
	for i := range successfulNodes {
		if successfulNodes[i] == nil {
			continue
		}
		uploadResults = append(uploadResults, &pb.SegmentPieceUploadResult{
			PieceNum: int32(i),
			NodeId:   successfulNodes[i].Id,
			Hash:     successfulHashes[i],
		})
	}

	if l := len(uploadResults); l < s.rs.OptimalThreshold() {
		return nil, size, Error.New("uploaded results (%d) are below the optimal threshold (%d)", l, s.rs.OptimalThreshold())
	}

	return uploadResults, sizedReader.Size(), nil
}

// Ranger creates a ranger for downloading erasure codes from piece store nodes.
func (s *segmentStore) Ranger(
	ctx context.Context, info storj.SegmentDownloadInfo, limits []*pb.AddressedOrderLimit, objectRS storj.RedundancyScheme,
) (rr ranger.Ranger, err error) {
	defer mon.Task()(&ctx, info, limits, objectRS)(&err)

	// no order limits also means its inline segment
	if len(info.EncryptedInlineData) != 0 || len(limits) == 0 {
		return ranger.ByteRanger(info.EncryptedInlineData), nil
	}

	needed := CalcNeededNodes(objectRS)
	selected := make([]*pb.AddressedOrderLimit, len(limits))
	s.rngMu.Lock()
	perm := s.rng.Perm(len(limits))
	s.rngMu.Unlock()

	for _, i := range perm {
		limit := limits[i]
		if limit == nil {
			continue
		}

		selected[i] = limit

		needed--
		if needed <= 0 {
			break
		}
	}

	fc, err := infectious.NewFEC(int(objectRS.RequiredShares), int(objectRS.TotalShares))
	if err != nil {
		return nil, err
	}
	es := eestream.NewRSScheme(fc, int(objectRS.ShareSize))
	redundancy, err := eestream.NewRedundancyStrategy(es, int(objectRS.RepairShares), int(objectRS.OptimalShares))
	if err != nil {
		return nil, err
	}

	rr, err = s.ec.Get(ctx, selected, info.PiecePrivateKey, redundancy, info.Size)
	return rr, Error.Wrap(err)
}

// Delete requests the satellite to delete a segment and tells storage nodes
// to delete the segment's pieces.
func (s *segmentStore) Delete(ctx context.Context, streamID storj.StreamID, segmentIndex int32) (err error) {
	defer mon.Task()(&ctx)(&err)

	_, limits, privateKey, err := s.metainfo.BeginDeleteSegment(ctx, metainfo.BeginDeleteSegmentParams{
		StreamID: streamID,
		Position: storj.SegmentPosition{
			Index: segmentIndex,
		},
	})
	if err != nil {
		return Error.Wrap(err)
	}

	if len(limits) != 0 {
		// remote segment - delete the pieces from storage nodes
		err = s.ec.Delete(ctx, limits, privateKey)
		if err != nil {
			return Error.Wrap(err)
		}
	}

	// don't do FinishDeleteSegment at the moment to avoid satellite round trip
	// FinishDeleteSegment doesn't implement any specific logic at the moment

	return nil
}

// CalcNeededNodes calculate how many minimum nodes are needed for download,
// based on t = k + (n-o)k/o
func CalcNeededNodes(rs storj.RedundancyScheme) int32 {
	extra := int32(1)

	if rs.OptimalShares > 0 {
		extra = int32(((rs.TotalShares - rs.OptimalShares) * rs.RequiredShares) / rs.OptimalShares)
		if extra == 0 {
			// ensure there is at least one extra node, so we can have error detection/correction
			extra = 1
		}
	}

	needed := int32(rs.RequiredShares) + extra

	if needed > int32(rs.TotalShares) {
		needed = int32(rs.TotalShares)
	}
	return needed
}
