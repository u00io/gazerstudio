package utils

import "time"

func FormatDateTime(dt int64) string {
	return time.UnixMilli(dt).Format("2006-01-02 15:04:05")
}
