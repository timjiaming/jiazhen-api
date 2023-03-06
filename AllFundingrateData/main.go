package main

import (
	"fmt"
	"github.com/timjiaming/jiazhen-api/AllFundingrateData/Exchange"
	"github.com/timjiaming/jiazhen-api/AllFundingrateData/ORM"
	"strings"
	"time"
)

// 批量存储
func fundingrates(symbols []string, dborm ORM.DbOrm) {
	instId := ""
	before := ""
	after := ""
	limit := "100"
	for _, v := range symbols {
		instId = v
		dataList, err := Exchange.GetFundingrateHistory(instId, before, after, limit)
		if err != nil {
			return
		}
		go func() {
			fmt.Println(dataList)
			dborm.WriteList(dataList)
		}()
	}
}

// 单个存储
func fundingrate(symbols []string, dborm ORM.DbOrm) {
	instId := ""
	before := ""
	after := ""
	limit := "10"
	for _, v := range symbols {
		instId = v
		dataList, err := Exchange.GetFundingrateHistory(instId, before, after, limit)
		if err != nil {
			return
		}
		for _, v1 := range dataList {
			fun := v1
			go func() {
				fmt.Println(fun)
				dborm.Write(fun)
			}()
		}
		time.Sleep(100 * time.Millisecond)
	}
}
func CoinSymgol(symbols []string) []string {
	coinStr := "-USD-SWAP"
	var coinSlice []string
	for _, v := range symbols {
		if strings.Index(v, coinStr) != -1 {
			coinSlice = append(coinSlice, v)
		}
	}
	return coinSlice
}
func UsdtSymgol(symbols []string) []string {
	coinStr := "-USDT-SWAP"
	var usdtSlice []string
	for _, v := range symbols {
		if strings.Index(v, coinStr) != -1 {
			usdtSlice = append(usdtSlice, v)
		}
	}
	return usdtSlice
}
func UsdcSymgol(symbols []string) []string {
	coinStr := "-USDC-SWAP"
	var usdcSlice []string
	for _, v := range symbols {
		if strings.Index(v, coinStr) != -1 {
			usdcSlice = append(usdcSlice, v)
		}
	}
	return usdcSlice
}
func main() {
	symbols, _ := Exchange.GetInstruments()
	//coinSlice := CoinSymgol(symbols)
	//usdtSlice := UsdtSymgol(symbols)
	//usdcSlice := UsdcSymgol(symbols)
	//for _, v := range coinSlice {
	//	println(v)
	//}
	//fmt.Println(usdtSlice)
	//fmt.Println(usdcSlice)
	dborm := ORM.DbOrm{}
	dborm.InitDb()
	//fundingrates(symbols, dborm)
	fundingrate(symbols, dborm)
}
