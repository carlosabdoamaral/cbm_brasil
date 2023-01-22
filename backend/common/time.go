package common

import (
	"time"

	"github.com/golang/protobuf/ptypes"
	googleProtobuf "github.com/golang/protobuf/ptypes/timestamp"
)

func TimestampToTime(ts *googleProtobuf.Timestamp) time.Time {
	if ts == nil {
		return time.Time{}
	}

	a, err := ptypes.Timestamp(ts)

	if err != nil {
		return time.Time{}
	}

	return a
}

func TimeToTimestamp(t time.Time) *googleProtobuf.Timestamp {
	ts, err := ptypes.TimestampProto(t)

	if err != nil {
		return ptypes.TimestampNow()
	}
	return ts
}
