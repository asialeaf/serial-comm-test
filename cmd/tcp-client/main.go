package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

type TCPClient struct {
	address    string
	connPool   chan net.Conn
	poolSize   int
	mutex      sync.Mutex
	timeoutSec int
}

func NewTCPClient(address string, poolSize int, timeoutSec int) *TCPClient {
	return &TCPClient{
		address:    address,
		connPool:   make(chan net.Conn, poolSize),
		poolSize:   poolSize,
		timeoutSec: timeoutSec,
	}
}

func (c *TCPClient) Dial() (net.Conn, error) {
	select {
	case conn := <-c.connPool:
		if c.isClosed(conn) {
			return c.Dial()
		}
		return conn, nil
	default:
		return c.createNewConn()
	}
}

func (c *TCPClient) Release(conn net.Conn) error {
	if c.isClosed(conn) {
		return nil
	}

	c.mutex.Lock()
	defer c.mutex.Unlock()

	select {
	case c.connPool <- conn:
		return nil
	default:
		return conn.Close()
	}
}

func (c *TCPClient) createNewConn() (net.Conn, error) {
	conn, err := net.DialTimeout("tcp", c.address, time.Duration(c.timeoutSec)*time.Second)
	if err != nil {
		return nil, err
	}

	c.mutex.Lock()
	defer c.mutex.Unlock()

	if len(c.connPool) >= c.poolSize {
		conn.Close()
		return <-c.connPool, nil
	}

	return conn, nil
}

func (c *TCPClient) isClosed(conn net.Conn) bool {
	tcpConn, ok := conn.(*net.TCPConn)
	if !ok {
		return true
	}

	_, err := tcpConn.File()
	return err == nil
}

func main() {
	client := NewTCPClient("localhost:8080", 5, 10)

	conn, err := client.Dial()
	if err != nil {
		fmt.Println("Dial error:", err)
		return
	}

	defer client.Release(conn)

	// 在连接上执行 TCP 操作...
}
