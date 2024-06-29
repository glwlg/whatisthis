package utils

import (
	"fmt"
	"log"
	"strings"
	"time"
)

// LogError 记录错误日志
func LogError(message string, err error) {
	log.Printf("ERROR: %s - %v", message, err)
}

// LogInfo 记录信息日志
func LogInfo(message string) {
	log.Println("INFO: ", message)
	fmt.Println("fmt INFO: ", message)
}

func Timestamp() int64 {
	now := time.Now()
	ts := now.UnixMilli()
	return ts
}

func IsBlankOrSpaces(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}