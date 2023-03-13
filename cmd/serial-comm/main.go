package main

import (
	"fmt"

	"github.com/tarm/serial" // 串口库
)

func main() {
	// 配置串口参数
	config := &serial.Config{
		Name:     "/dev/ttyS0",      // 串口设备名
		Baud:     9600,              // 波特率
		Parity:   serial.ParityNone, // 校验位
		StopBits: serial.Stop1,      // 停止位
		Size:     8,                 // 数据位
	}

	// 打开串口
	port, err := serial.OpenPort(config)
	if err != nil {
		fmt.Println("无法打开串口：", err)
		return
	}

	// 循环读取串口数据
	for {
		// 读取数据
		buf := make([]byte, 128)
		n, err := port.Read(buf)
		if err != nil {
			fmt.Println("读取数据出错：", err)
			continue
		}

		// 解析数据
		data := buf[:n]
		fmt.Printf("收到数据：% X\n", data)
	}

	// 关闭串口
	// port.Close()
}
