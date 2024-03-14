package jisuapi

import (
	"go.dtapp.net/golog"
)

// ClientConfig 实例配置
type ClientConfig struct {
}

// Client 实例
type Client struct {
	config struct {
	}
	gormLog struct {
		status bool           // 状态
		client *golog.ApiGorm // 日志服务
	}
}

// NewClient 创建实例化
func NewClient() (*Client, error) {
	c := &Client{}
	return c, nil
}
