package utils

import "time"

/**
 * 获取当前时间
 */
func GetCurrentTime() time.Time {
	currentTime := time.Now() //获取当前时间，类型是Go的时间类型Time
	return currentTime
}

/**
 * 获取当前时间字符串
 */
func GetCurrentTimeString() string {
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	return timeStr
}
