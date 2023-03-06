package main

import (
	"fmt"
	"github.com/timjiaming/jiazhen-api/okx"
	"github.com/timjiaming/jiazhen-api/okx/models"
)

func main() {
	channels := []map[string]string{
		//okx.NewBooksChannel("BTC-USD-SWAP"),
		okx.NewTickersChannel("BTC-USD-SWAP"),
	}
	pss := okx.NewOkxPubStreamServe("")
	pss.Channels = channels
	pss.OnDepthUpdate = func(book string) {
		fmt.Println(book)
	}
	pss.OnTickerUpdate = func(ticker models.Ticker) {
		fmt.Println(ticker)
	}
	pss.OnClose = func(serve *okx.OkxPubStreamServe) {
		fmt.Println("serve closed")
	}
	pss.OnReadErr = func(err error) {
		fmt.Println(err)
	}
	err := pss.Serv()
	if err != nil {
		fmt.Println(err)
		return
	}
	pss.Waite()
}
