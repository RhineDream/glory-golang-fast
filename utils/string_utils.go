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
