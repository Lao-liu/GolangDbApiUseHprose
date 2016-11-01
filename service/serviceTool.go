package service

import (
	"time"
)

type ServiceTool struct{}

// 获取系统版本号
func (ServiceTool) GetVersion() string {
	return "1.0.0"
}

// 获取服务器时间
func (ServiceTool) GetServerDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
