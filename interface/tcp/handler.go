package tcp

import (
	"context"
	"net"
)

// Handler 业务逻辑的处理接口
type Handler interface {
	// Handler 处理连接
	Handler(ctx context.Context, conn net.Conn)
	Close() error
}
