package parse

import "github.com/asialeaf/serial-comm-test/pkg/utils"

func parseData(data []byte, start, end int) []byte {
	// TODO: 实现数据解析逻辑，例如根据协议解析数据包内容
	// fmt.Printf("收到数据：% X\n", data)
	return data[start:end]

}

func ParseTempHumi(data []byte) (temp, humi uint16) {
	temp = utils.HexToInt(string(parseData(data, 5, 9)))
	humi = utils.HexToInt(string(parseData(data, 9, 13)))
	return temp, humi
}
