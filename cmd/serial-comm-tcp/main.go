package main

import (
	"fmt"

	"github.com/asialeaf/serial-comm-test/internal/parse"

	"github.com/asialeaf/serial-comm-test/pkg/clients/tcp"
	"github.com/asialeaf/serial-comm-test/pkg/utils"
)

func main() {
	// Create Tcp Client
	tcpClient := tcp.NewTCPClient("192.168.20.187:5022", 5, 5)

	// Request ANALOG DATA(PV, SV, TIME, ETC), SIGNAL SYMBOL '01'
	msg := "@010140*\r"
	fmt.Printf("Send msg: %s\n", msg)

	b, err := tcpClient.Send([]byte(msg))
	if err == nil {
		fmt.Printf("Recive msg: %s\n", b)
	} else {
		fmt.Printf("Failed to get msg: %v\n", err)
	}

	// Parse Data（Temperature, Humidity）
	temp, humi := parse.ParseTempHumi(b)
	fmt.Printf("设备当前温度：%.2f 摄氏度, 当前湿度：%.2f %%\n", utils.ShiftDecimal(int(temp), 2), utils.ShiftDecimal(int(humi), 2))

}
