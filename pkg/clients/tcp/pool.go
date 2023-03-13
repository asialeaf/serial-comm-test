package tcp

import (
	"net"
	"sync"
)

// ConnPool 连接池结构体
type ConnPool struct {
	addr     string
	maxConns int
	conns    chan net.Conn
	mu       sync.Mutex
}

// NewConnPool 创建新的连接池
func NewConnPool(addr string, maxConns int) *ConnPool {
	return &ConnPool{
		addr:     addr,
		maxConns: maxConns,
		conns:    make(chan net.Conn, maxConns),
	}
}

// Get 从连接池中获取连接
func (p *ConnPool) Get() (net.Conn, error) {
	select {
	case conn := <-p.conns:
		return conn, nil
	default:
		// 连接池中没有可用连接，创建新连接
		conn, err := net.Dial("tcp", p.addr)
		if err != nil {
			return nil, err
		}
		return conn, nil
	}
}

// Put 将连接放回连接池中
func (p *ConnPool) Put(conn net.Conn) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.conns == nil {
		conn.Close()
		return
	}

	select {
	case p.conns <- conn:
	default:
		conn.Close()
	}
}
