package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func ClientHandleError(err error, when string) {
	if err != nil {
		fmt.Println(err, when)
		os.Exit(1)
	}
}

func main() {

	conn, err := net.Dial("tcp", "192.168.20.187:5022")
	ClientHandleError(err, "client conn error")

	//预先准备消息缓冲区
	buffer := make([]byte, 1024)

	//准备命令行标准输入
	reader := bufio.NewReader(os.Stdin)

	for {
		lineBytes, _, _ := reader.ReadLine()
		conn.Write(lineBytes)
		n, err := conn.Read(buffer)
		ClientHandleError(err, "client read error")

		serverMsg := string(buffer[0:n])
		fmt.Printf("服务端msg: %s", serverMsg)
		if serverMsg == "bye" {
			break
		}
	}

}
