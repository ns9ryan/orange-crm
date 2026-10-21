package utils

import (
	"github.com/gogf/gf/v2/errors/gerror"
	"time"
)

const (
	TIME_DEFAULT       = "2006-01-02 15:04:05"
	TIME_DEFAULT_YMD   = "2006-01-02"
	TIME_DEFAULT_YM    = "2006-01"
	TIME_DEFAULT_YMDHI = "2006-01-02 15:04"
	TIME_RFC3339Milli  = "2006-01-02T15:04:05.000Z07:00"
)

// TimestampToDate 时间戳转字符串
func TimestampToDate(timestamp int64, layout string) string {
	if layout == "" {
		return time.Unix(timestamp, 0).Format(TIME_DEFAULT)
	}
	return time.Unix(timestamp, 0).Format(layout)
}

// DateToTime 解析字符串格式的时间
// 返回时间格式是值(time.Time) 而不是指针(*time.Time)
func DateToTime(date, layout string) (*time.Time, error) {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return nil, err
	}
	obj, err := time.ParseInLocation(layout, date, loc)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

// TimeRangeType 查询时间范围类型枚举
type TimeRangeType int

const (
	TimeRangeToday   TimeRangeType = iota // 仅包含今天的数据
	TimeRangeHistory                      // 仅包含历史数据（不含今天）
	TimeRangeMixed                        // 同时包含历史和今天的数据
)

// DetectTimeRangeType 判断查询时间属于哪种类型（今天、历史、混合）
//
// 输入：
//   - startStr：开始时间字符串，格式为 "2006-01-02 15:04:05"
//   - endStr：结束时间字符串，格式为 "2006-01-02 15:04:05"
//
// 返回：
//   - TimeRangeType：时间范围类型（枚举）
//   - error：如时间格式错误，返回错误
func DetectTimeRangeType(startStr, endStr string) (rangeType TimeRangeType, startTime, endTime time.Time, err error) {
	// 解析字符串为 time.Time 类型
	startTime, err = time.ParseInLocation(time.DateTime, startStr, time.Local)
	if err != nil {
		return
	}
	endTime, err = time.ParseInLocation(time.DateTime, endStr, time.Local)
	if err != nil {
		return
	}

	// 获取今天的开始时间和结束时间
	// Time: 当前时间
	now := time.Now()
	// Time: 今天的零点 2025-04-09 00:00:00.000
	todayStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	// 只查询今天
	if !startTime.Before(todayStart) && endTime.After(todayStart) {
		rangeType = TimeRangeToday
		return
	}

	// 只查询之前
	if endTime.Before(todayStart) {
		rangeType = TimeRangeHistory
		return
	}

	// 查询今天和之前
	if startTime.Before(todayStart) && endTime.After(todayStart) {
		rangeType = TimeRangeMixed
		return
	}

	// 理论不会走到这里
	err = gerror.Newf(`{#query.type.exception}`) // 查询类型异常
	return
}
