package service

import (
	"time"
)

type ServerTool struct{}

// 获取系统版本号
func (ServerTool) GetVersion() string {
	return "1.0.0"
}

// 获取服务器时间
func (ServerTool) GetServerDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
