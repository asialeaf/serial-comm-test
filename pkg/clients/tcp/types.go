package tcp

type Message struct {
	UnitNum      string
	SignalSymbol string
	Val          string
}

type Resp struct {
	UnitNum      string
	SignalSymbol string
	Val          string
	Ts           string
}
