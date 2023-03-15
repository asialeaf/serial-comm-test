package main

import (
	"fmt"

	"github.com/asialeaf/serial-comm-test/internal/parse"
	"github.com/asialeaf/serial-comm-test/pkg/utils"
	"github.com/tarm/serial" // 串口库
)

func main() {
	// 配置串口参数
	config := &serial.Config{
		Name:     "/dev/ttyUSB0",    // 串口设备名
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

	// 发送串口数据
	// Request ANALOG DATA(PV, SV, TIME, ETC), SIGNAL SYMBOL '01'
	msg := "@010140*\r"
	if _, err := port.Write([]byte(msg)); err != nil {
		fmt.Println("发送数据失败：", err)
	}

	fmt.Printf("Send msg: %s\n", msg)

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
		fmt.Printf("Recive msg: %s\n", data)
		temp, humi := parse.ParseTempHumi(data)
		fmt.Printf("设备当前温度：%.2f 摄氏度, 当前湿度：%.2f %%\n", utils.ShiftDecimal(int(temp), 2), utils.ShiftDecimal(int(humi), 2))
	}

	// 关闭串口
	// port.Close()
}
