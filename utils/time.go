package utils

import (
	"fmt"
	"strconv"
	"time"
)

// GetDate 返回 20060102
func GetDate(tm time.Time) (r int) {
	r, _ = strconv.Atoi(tm.Format("20060102"))
	return
}

// GetDateHour 获取日期到小时 return 2022122609
func GetDateHour(tm time.Time) (r int) {
	r, _ = strconv.Atoi(tm.Format("2006010215"))
	return
}

// GetWeek 获取周次本年度的第几周 返回2023.6
func GetWeek(tm time.Time) string {
	year, week := tm.ISOWeek()
	return fmt.Sprintf("%d.%d", year, week)
}
