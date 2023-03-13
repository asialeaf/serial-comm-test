package utils

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math"
	"strconv"
)

func HexToInt(hexStr string) uint16 {
	bs, _ := hex.DecodeString(hexStr)
	num := binary.BigEndian.Uint16(bs[:2])
	return num
}

func DivideWithPrecision(dividend, divisor float64, precision int) float64 {
	quotient := dividend / divisor
	format := fmt.Sprintf("%%.%df", precision)
	quotientFormatted, _ := strconv.ParseFloat(fmt.Sprintf(format, quotient), 64)
	return quotientFormatted
}

func ShiftDecimal(num int, places int) float64 {
	return float64(num) / math.Pow10(places)
}
