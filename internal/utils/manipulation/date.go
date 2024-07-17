package manipulation

import (
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func StartOfToday() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
}

func StartOfYesterday() time.Time {
	return StartOfToday().AddDate(0, 0, -1)
}

func IsToday(t time.Time) bool {
	today := StartOfToday()
	return t.Year() == today.Year() && t.Month() == today.Month() && t.Day() == today.Day()
}

func IsYesterday(t time.Time) bool {
	yesterday := StartOfYesterday()
	return t.Year() == yesterday.Year() && t.Month() == yesterday.Month() && t.Day() == yesterday.Day()
}

func StartOfDate(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func EndOfDate(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
}

func ConvertToProtoTimestamp(t time.Time) *timestamp.Timestamp {
	return timestamppb.New(t)
}

func FormatDDMM(t time.Time) string {
	return t.Format("02/01")
}
