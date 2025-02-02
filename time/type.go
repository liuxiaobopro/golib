package time

import (
	"fmt"
	"time"
)

var (
	TimeZero = Time(time.Time{})
	TimeNow  = Time(time.Now())
)

type JsonTime time.Time

func (t JsonTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(t).Format(FormatDateTime))
	return []byte(stamp), nil
}

type Time time.Time // 和JsonTime一样，只是为了方便使用

func (t Time) MarshalJSON() ([]byte, error) {
	var stamp = ""
	if t.IsZero() {
		stamp = "\"\""
	} else {
		stamp = fmt.Sprintf("\"%s\"", time.Time(t).Format(FormatDateTime))
	}

	return []byte(stamp), nil
}

func (t Time) IsZero() bool {
	return time.Time(t).IsZero()
}

func (t Time) Format(layout string) string {
	if t.IsZero() {
		return ""
	}

	return time.Time(t).Format(layout)
}

func (t Time) Unix() int64 {
	if t.IsZero() {
		return 0
	}

	return t.Raw().Unix()
}

// Raw 返回原始time.Time
func (t Time) Raw() time.Time {
	return time.Time(t)
}

// IsToday 判断是否今天
func (t Time) IsToday() bool {
	return t.Format(FormatDate) == time.Now().Format(FormatDate)
}

// IsYesterday 判断是否昨天
func (t Time) IsYesterday() bool {
	return t.Format(FormatDate) == time.Now().AddDate(0, 0, -1).Format(FormatDate)
}

func Now() Time {
	return Time(time.Now())
}
