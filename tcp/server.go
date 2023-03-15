package tcp

import (
	"context"
	"go-redis/interface/tcp"
	"go-redis/lib/logger"
	"net"
)

// Config 启动tcp服务器的配置
type Config struct {
	// Address 监听地址
	Address string
}

func ListenAndServeWithSignal(cfg *Config,
	handler tcp.Handler) error {
	closeChan := make(chan struct{})
	listen, err := net.Listen("tcp", cfg.Address)
	if err != nil {
		return err
	}
	logger.Info("start listen")
	ListenAndServe(listen, handler, closeChan)
	return nil
}

func ListenAndServe(listener net.Listener,
	handler tcp.Handler,
	closeChan <-chan struct{}) {
	ctx := context.Background()
	for true {
		conn, err := listener.Accept()
		if err != nil {
			break
		}
		// 接收到新连接
		logger.Info("accept link")
		// 一个协程一个连接
		go func() {
			handler.Handler(ctx, conn)
		}()
	}
}
