package manipulation

import (
	"fmt"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func getLocation(tz string) *time.Location {
	if tz == "" {
		return time.UTC
	}

	loc, err := time.LoadLocation(tz)
	if err != nil {
		fmt.Printf("invalid timezone '%s', defaulting to UTC: %v\n", tz, err)
		return time.UTC
	}
	return loc
}

func Now(tz string) time.Time {
	loc := getLocation(tz)
	return time.Now().In(loc)
}

func NowUTC() time.Time {
	return time.Now().UTC()
}

func StartOfToday(tz string) time.Time {
	loc := getLocation(tz)
	now := time.Now().In(loc)
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)
}

func StartOfYesterday(tz string) time.Time {
	today := StartOfToday(tz)
	return today.AddDate(0, 0, -1)
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

func FormatDDMM(t time.Time, tz string) string {
	return t.In(getLocation(tz)).Format("02/01")
}

func IsToday(t time.Time, tz string) bool {
	today := StartOfToday(tz)
	t = t.In(today.Location())
	return t.Year() == today.Year() && t.Month() == today.Month() && t.Day() == today.Day()
}

func IsYesterday(t time.Time, tz string) bool {
	yesterday := StartOfYesterday(tz)
	t = t.In(yesterday.Location())
	return t.Year() == yesterday.Year() && t.Month() == yesterday.Month() && t.Day() == yesterday.Day()
}

func ToSQLTimestamp(t time.Time, tz string) string {
	loc := getLocation(tz)
	t = t.In(loc)
	return t.Format("2006-01-02 15:04:05.999999-07:00")
}

func ToSQLDate(t time.Time, tz string) string {
	loc := getLocation(tz)
	t = t.In(loc)
	if t.Hour() > 0 || t.Minute() > 0 || t.Second() > 0 {
		t = t.Add(24 * time.Hour).Truncate(24 * time.Hour)
	}
	return t.Format("2006-01-02")
}

func ToSQLDateFrom(t time.Time, tz string) string {
	loc := getLocation(tz)
	t = t.In(loc)
	if t.Hour() > 0 || t.Minute() > 0 || t.Second() > 0 || t.Nanosecond() > 0 {
		t = t.Add(24 * time.Hour).Truncate(24 * time.Hour)
	} else {
		t = t.Truncate(24 * time.Hour)
	}
	return t.Format("2006-01-02")
}

func ToSQLDateTo(t time.Time, tz string) string {
	loc := getLocation(tz)
	t = t.In(loc)
	t = t.Truncate(24 * time.Hour)
	return t.Format("2006-01-02")
}
