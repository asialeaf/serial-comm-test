package main

import (
	"fmt"

	"github.com/asialeaf/serial-comm-test/pkg/clients/tcp"
)

const (
	addr     = "192.168.20.187:5022" // 服务器地址
	maxConns = 10                    // 连接池中的最大连接数
)

func main() {
	pool := tcp.NewConnPool(addr, maxConns) // 创建连接池

	// 从连接池中获取连接
	conn, err := pool.Get()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer pool.Put(conn) // 使用完毕后将连接放回池中

	// 向服务器发送数据
	_, err = conn.Write([]byte("@010140*\r"))
	if err != nil {
		fmt.Println(err)
		return
	}

	// 从服务器读取数据
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 打印服务器返回的数据
	fmt.Println(string(buffer[:n]))
}
