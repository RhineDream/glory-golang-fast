package utils

import (
	"github.com/satori/go.uuid"
	"strings"
)

/**
获取uuid
*/
func GetUUID() string {
	// 创建UUID
	u1 := uuid.Must(uuid.NewV4(), nil).String()
	uuidStr := strings.Replace(u1, "-", "", -1)
	return uuidStr
}

/**
 * 切割字符串
 */
func SubString(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}
	return string(rs[start:end])
}
