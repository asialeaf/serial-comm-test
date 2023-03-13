package tcp

import (
	"fmt"
	"time"
)

type TCPClient struct {
	addr       string
	connPool   *ConnPool
	maxConns   int
	timeoutSec int
}

// NewTCPClient 创建TCPClient
func NewTCPClient(addr string, maxConns int, timeoutSec int) *TCPClient {
	return &TCPClient{
		addr:       addr,
		maxConns:   maxConns,
		timeoutSec: timeoutSec,
	}
}

func (c *TCPClient) Send(data []byte) ([]byte, error) {
	c.connPool = NewConnPool(c.addr, c.maxConns)
	conn, err := c.connPool.Get()
	if err != nil {
		return nil, fmt.Errorf("failed to get connection from pool")
	}
	defer c.connPool.Put(conn)

	conn.SetDeadline(time.Now().Add(time.Duration(c.timeoutSec) * time.Second))

	// 向服务器发送数据
	_, err = conn.Write(data)
	if err != nil {
		return nil, err
	}

	// 从服务器读取数据
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		return nil, err
	}

	return buf[:n], nil
}
