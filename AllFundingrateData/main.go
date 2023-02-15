package main

import "jiazhen-api/AllFundingrateData/Exchange"

func main() {
	instId := "BTC-USDT-SWAP"
	before := "1672502400000"
	after := ""
	limit := "100"
	Exchange.GetFundingrateHistory(instId, before, after, limit)
}
