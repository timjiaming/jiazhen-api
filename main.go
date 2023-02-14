package main

import (
	"fmt"
	"jiazhen-api/okx"
	"jiazhen-api/okx/models"
)

func main() {
	//GetTickers()
	//GetTicker()
	//GetIndexTickers()
	//GetBooks()
	//GetBookLite()
	//GetCandles()
	//GetHistoryCandles()
	//GetMarkPriceCandles()
	//GetHistoryMarketPriceCandles()
	//GetTrades()
	//GetHistoryTrades()
	//GetPlatform24Volume()

	//GetInstruments()
	//GetDeliveryExerciseHistory()
	//GetFundingrate()
	//GetFundingrateHistory()
	//GetPriceLimit()
	GetEstimatedPrice()
}

// 行情数据
func GetTickers() {
	acc := okx.OkxAccount{
		Testnet: true,
	}
	api := okx.NewOkxRestApi(&acc, "")
	req := models.GetTickersReq{
		InstType: "SWAP",
	}
	message, err := api.GetTickers(req)
	if err != nil {
		return
	}
	fmt.Println(message)
}
func GetTicker() {
	acc := okx.OkxAccount{
		Testnet: true,
	}
	api := okx.NewOkxRestApi(&acc, "")
	req := models.GetTickerReq{
		InstId: "BTC-USD-SWAP",
	}
	message, err := api.GetTicker(req)
	if err != nil {
		return
	}
	fmt.Println(message)
}
func GetIndexTickers() {
	acc := okx.OkxAccount{
		Testnet: true,
	}
	api := okx.NewOkxRestApi(&acc, "")
	req := models.GetIndexTickersReq{
		QuoteCcy: "USD",
	}
	message, err := api.GetIndexTickers(req)
	if err != nil {
		return
	}
	fmt.Println(message)
}
func GetBooks() {
	acc := okx.OkxAccount{
		Testnet: true,
	}
	api := okx.NewOkxRestApi(&acc, "")
	req := models.GetBooksReq{
		InstId: "BTC-USDT",
		SZ:     "400",
	}
	message, err := api.GetBooks(req)
	if err != nil {
		return
	}
	fmt.Println(message)
}
func GetBookLite() {
	acc := okx.OkxAccount{
		Testnet: true,
	}
	api := okx.NewOkxRestApi(&acc, "")
	req := models.GetBookLiteReq{InstId: "BTC-USDT"}
	message, err := api.GetBooksLite(req)
	if err != nil {
		return
	}
	fmt.Println(message)
}
func GetCandles() {
	acc := okx.OkxAccount{
		Testnet: true,
	}
	api := okx.NewOkxRestApi(&acc, "")
	req := models.GetCandlesReq{
		InstId: "BTC-USDT-SWAP",
		Bar:    "",
	}
	message, err := api.GetCandles(req)
	if err != nil {
		return
	}
	fmt.Println(message)
}
func GetHistoryCandles() {
	acc := okx.OkxAccount{
		Testnet: true,
	}
	api := okx.NewOkxRestApi(&acc, "")
	req := models.GetHistoryCandlesReq{
		InstId: "BTC-USDT-SWAP",
		After:  "",
		Before: "",
		Bar:    "",
		Limit:  "",
	}
	message, err := api.GetHistoryCandles(req)
	if err != nil {
		return
	}
	fmt.Println(message)
}
func GetMarkPriceCandles() {
	acc := okx.OkxAccount{
		Testnet: true,
	}
	api := okx.NewOkxRestApi(&acc, "")
	req := models.GetMarketPriceCandlesReq{
		InstId: "BTC-USDT-SWAP",
		After:  "",
		Before: "",
		Bar:    "",
		Limit:  "",
	}
	message, err := api.GetMarketPriceCandles(req)
	if err != nil {
		return
	}
	fmt.Println(message)
}
func GetHistoryMarketPriceCandles() {
	acc := okx.OkxAccount{
		Testnet: true,
	}
	api := okx.NewOkxRestApi(&acc, "")
	req := models.GetHistoryMarketPriceReq{
		InstId: "BTC-USDT-SWAP",
		After:  "",
		Before: "",
		Bar:    "",
		Limit:  "",
	}
	message, err := api.GetHistoryMarketPriceCandles(req)
	if err != nil {
		return
	}
	fmt.Println(message)
}
func GetTrades() {
	acc := okx.OkxAccount{
		Testnet: true,
	}
	api := okx.NewOkxRestApi(&acc, "")
	req := models.GetTradesReq{
		InstId: "BTC-USDT-SWAP",
		Limit:  "",
	}
	message, err := api.GetTrades(req)
	if err != nil {
		return
	}
	fmt.Println(message)
}
func GetHistoryTrades() {
	acc := okx.OkxAccount{
		Testnet: true,
	}
	api := okx.NewOkxRestApi(&acc, "")
	req := models.GetHistoryTradesReq{
		InstId: "BTC-USDT-SWAP",
		After:  "",
		Before: "",
		Bar:    "",
		Limit:  "",
	}
	message, err := api.GetHistoryTrades(req)
	if err != nil {
		return
	}
	fmt.Println(message)
}
func GetPlatform24Volume() {
	acc := okx.OkxAccount{
		Testnet: true,
	}
	api := okx.NewOkxRestApi(&acc, "")
	message, err := api.GetPlatform24Volume()
	if err != nil {
		return
	}
	fmt.Println(message)
}

// 公共数据
func GetInstruments() {
	acc := okx.OkxAccount{
		Testnet: true,
	}
	api := okx.NewOkxRestApi(&acc, "")
	req := models.GetInstrumentsReq{
		InstType:   "SWAP",
		Uly:        "",
		InstFamily: "",
		InstId:     "",
	}
	message, err := api.GetInstruments(req)
	if err != nil {
		return
	}
	fmt.Println(message)
}

func GetDeliveryExerciseHistory() {
	acc := okx.OkxAccount{
		Testnet: true,
	}
	api := okx.NewOkxRestApi(&acc, "")
	req := models.GetDeliveryExerciseHistoryReq{
		InstType:   "FUTURES",
		Uly:        "",
		InstFamily: "BTC-USD",
		After:      "",
		Before:     "",
		Limit:      "",
	}
	message, err := api.GetDeliveryExerciseHistory(req)
	if err != nil {
		return
	}
	fmt.Println(message)
}
func GetOpenInterest() {
	acc := okx.OkxAccount{
		Testnet: true,
	}
	api := okx.NewOkxRestApi(&acc, "")
	req := models.GetOpenInterestReq{
		InstType:   "SWAP",
		Uly:        "",
		InstFamily: "BTC-USD",
		Limit:      "",
	}
	message, err := api.GetOpenInterest(req)
	if err != nil {
		return
	}
	fmt.Println(message)
}
func GetFundingrate() {
	acc := okx.OkxAccount{
		Testnet: true,
	}
	api := okx.NewOkxRestApi(&acc, "")
	req := models.GetFundingrateReq{InstId: "BTC-USD-SWAP"}
	message, err := api.GetFundingrate(req)
	if err != nil {
		return
	}
	fmt.Println(message)
}
func GetFundingrateHistory() {
	acc := okx.OkxAccount{
		Testnet: true,
	}
	api := okx.NewOkxRestApi(&acc, "")
	req := models.GetFundingrateHistoryReq{
		InstId: "BTC-USD-SWAP",
		Before: "",
		After:  "",
		Limit:  "",
	}
	message, err := api.GetFundingrateHistory(req)
	if err != nil {
		return
	}
	fmt.Println(message)
}
func GetPriceLimit() {
	acc := okx.OkxAccount{
		Testnet: true,
	}
	api := okx.NewOkxRestApi(&acc, "")
	req := models.GetPriceLimitReq{InstId: "BTC-USDT-SWAP"}
	message, err := api.GetPriceLimit(req)
	if err != nil {
		return
	}
	fmt.Println(message)
}
func GetEstimatedPrice() {
	acc := okx.OkxAccount{
		Testnet: true,
	}
	api := okx.NewOkxRestApi(&acc, "")
	req := models.GetEstimatedPriceReq{InstId: "BTC-USDT-SWAP"}
	message, err := api.GetEstimatedPrice(req)
	if err != nil {
		return
	}
	fmt.Println(message)
}
