package main

import (
	"fmt"
	"github.com/timjiaming/jiazhen-api/okx"
	"github.com/timjiaming/jiazhen-api/okx/models"
)

const (
	apikey     = "025a736a-6c95-40ee-b1a0-a455c5feede8"
	sercet     = "57F2EDB74E21A7E8AD5788894AE50000"
	passphrase = "Jiazhen9527!@#"
)

func main() {
	//GetBookLite()
	//GetTickers()
	//GetTicker()
	//GetIndexTickers()
	//GetBooks()
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
	//GetEstimatedPrice()

	PlaceOrder()
	//PlaceBatchOrders()
	//CancelOrder()
	//CancelBatchOrders()
	//AmendOrder()
	//ClosePosition()
	//GetOrder()
	//GerOrdersHistory()
	//GetOrdersHistoryArchive()
	//GetFills()
	//GetFillsHistory()
	//Currencies()
	//Balances()
	//NonTradableAssets()
	//AssetValuation()
	//Transfer()
	//TransferState()
	//AccountBalance()
	//AccountPosition()
	//AccountConfig()
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
		Testnet: false,
	}
	api := okx.NewOkxRestApi(&acc, "")
	req := models.GetBookLiteReq{InstId: "BTC-USD-SWAP"}
	message, err := api.GetBooksLite(req)
	if err != nil {
		fmt.Println(err)
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

// 交易
func PlaceOrder() {
	acc := okx.OkxAccount{
		ApiKey:     apikey,
		ApiSecret:  sercet,
		Passphrase: passphrase,
		BaseURL:    "",
		Testnet:    true,
	}
	api := okx.NewOkxRestApi(&acc, "")
	//req := models.PlaceOrderReq{
	//	InstId:  "ZEC-USDT",
	//	TdMode:  okx.Cross,
	//	Side:    okx.Buy,
	//	OrdType: okx.Market,
	//	Sz:      "100",
	//}
	req := models.PlaceOrderReq{
		InstId:  "ZEC-USD-SWAP",
		TdMode:  okx.Cross,
		Side:    okx.Sell,
		OrdType: okx.Market,
		Sz:      "10",
	}
	message, err := api.PlaceOrder(req)
	if err != nil {
		return
	}
	fmt.Println(message)
}
func PlaceBatchOrders() {
	acc := okx.OkxAccount{
		ApiKey:     apikey,
		ApiSecret:  sercet,
		Passphrase: passphrase,
		BaseURL:    "",
		Testnet:    true,
	}
	api := okx.NewOkxRestApi(&acc, "")
	req := models.PlaceBatchOrdersReq{
		models.PlaceOrderReq{
			InstId:  "ETH-USDT-SWAP",
			TdMode:  okx.Cross,
			Side:    okx.Buy,
			OrdType: okx.Market,
			Sz:      "1",
		}, models.PlaceOrderReq{
			InstId:  "ETH-USDT-SWAP",
			TdMode:  okx.Cross,
			Side:    okx.Buy,
			OrdType: okx.Market,
			Sz:      "1",
		},
	}
	message, err := api.PlaceBatchOrder(req)
	if err != nil {
		return
	}
	fmt.Println(message)
}
func CancelOrder() {
	acc := okx.OkxAccount{
		ApiKey:     apikey,
		ApiSecret:  sercet,
		Passphrase: passphrase,
		BaseURL:    "",
		Testnet:    true,
	}
	api := okx.NewOkxRestApi(&acc, "")
	req := models.CancelOrderReq{
		InstId:  "BTC-USDT-SWAP",
		OrdId:   "548217645183639552",
		ClOrdId: "",
	}
	message, err := api.CancelOrder(req)
	if err != nil {
		return
	}
	fmt.Println(message)
}
func CancelBatchOrders() {
	acc := okx.OkxAccount{
		ApiKey:     apikey,
		ApiSecret:  sercet,
		Passphrase: passphrase,
		BaseURL:    "",
		Testnet:    true,
	}
	api := okx.NewOkxRestApi(&acc, "")
	req := models.CancelBatchOrdersReq{
		models.CancelOrderReq{
			InstId:  "BTC-USDT-SWAP",
			OrdId:   "548308930246582272",
			ClOrdId: "",
		}, models.CancelOrderReq{
			InstId:  "BTC-USDT-SWAP",
			OrdId:   "548308962785992708",
			ClOrdId: "",
		},
	}
	message, err := api.CancelBatchOrder(req)
	if err != nil {
		return
	}
	fmt.Println(message)
}
func AmendOrder() {
	acc := okx.OkxAccount{
		ApiKey:     apikey,
		ApiSecret:  sercet,
		Passphrase: passphrase,
		BaseURL:    "",
		Testnet:    true,
	}
	api := okx.NewOkxRestApi(&acc, "")
	req := models.AmendOrderReq{
		InstId: "BTC-USDT-SWAP",
		OrdId:  "548310785836032003",
		NewPx:  "22999",
	}
	message, err := api.AmendOrder(req)
	if err != nil {
		return
	}
	fmt.Println(message)
}
func AmendBathOrders() {
	acc := okx.OkxAccount{
		ApiKey:     apikey,
		ApiSecret:  sercet,
		Passphrase: passphrase,
		BaseURL:    "",
		Testnet:    true,
	}
	api := okx.NewOkxRestApi(&acc, "")
	req := models.AmendBatchOrdersReq{}
	message, err := api.AmendBatchOrder(req)
	if err != nil {
		return
	}
	fmt.Println(message)
}
func ClosePosition() {
	acc := okx.OkxAccount{
		ApiKey:     apikey,
		ApiSecret:  sercet,
		Passphrase: passphrase,
		BaseURL:    "",
		Testnet:    true,
	}
	api := okx.NewOkxRestApi(&acc, "")
	req := models.ClosePositionReq{
		InstId:  "BTC-USDT-SWAP",
		PosSide: okx.Cross,
		MgnMode: "",
		Ccy:     "",
		AutoCxl: "",
		ClOrdId: "",
		Tag:     "",
	}
	message, err := api.ClosePosition(req)
	if err != nil {
		return
	}
	fmt.Println(message)
}
func GetOrder() {
	acc := okx.OkxAccount{
		ApiKey:     apikey,
		ApiSecret:  sercet,
		Passphrase: passphrase,
		BaseURL:    "",
		Testnet:    true,
	}
	api := okx.NewOkxRestApi(&acc, "")
	req := models.OrderReq{
		InstId:  "BTC-USDT-SWAP",
		OrdId:   "548512580147458048",
		ClOrdId: "",
	}
	message, err := api.Order(req)
	if err != nil {
		return
	}
	fmt.Println(message)
}
func GetOrdersPending() {
	acc := okx.OkxAccount{
		ApiKey:     apikey,
		ApiSecret:  sercet,
		Passphrase: passphrase,
		BaseURL:    "",
		Testnet:    true,
	}
	api := okx.NewOkxRestApi(&acc, "")
	req := models.OrdersPendingReq{}
	message, err := api.OrderPending(req)
	if err != nil {
		return
	}
	fmt.Println(message)
}
func GerOrdersHistory() {
	acc := okx.OkxAccount{
		ApiKey:     apikey,
		ApiSecret:  sercet,
		Passphrase: passphrase,
		BaseURL:    "",
		Testnet:    true,
	}
	api := okx.NewOkxRestApi(&acc, "")
	req := models.OrdersHistoryReq{
		InstType:   "",
		Uly:        "",
		InstFamily: "",
		InstId:     "",
		OrdType:    "",
		State:      "",
		Category:   "",
		After:      "",
		Before:     "",
		Begin:      "",
		End:        "",
		Limit:      "",
	}
	message, err := api.OrderHistory(req)
	if err != nil {
		return
	}
	fmt.Println(message)
}
func GetOrdersHistoryArchive() {
	acc := okx.OkxAccount{
		ApiKey:     apikey,
		ApiSecret:  sercet,
		Passphrase: passphrase,
		BaseURL:    "",
		Testnet:    true,
	}
	api := okx.NewOkxRestApi(&acc, "")
	req := models.OrdersHistoryArchiveReq{
		InstType:   okx.SWAP,
		Uly:        "",
		InstFamily: "",
		InstId:     "",
		OrdType:    "",
		State:      "",
		Category:   "",
		After:      "",
		Before:     "",
		Begin:      "",
		End:        "",
		Limit:      "",
	}
	message, err := api.OrderHistoryArchive(req)
	if err != nil {
		return
	}
	fmt.Println(message)
}
func GetFills() {
	acc := okx.OkxAccount{
		ApiKey:     apikey,
		ApiSecret:  sercet,
		Passphrase: passphrase,
		BaseURL:    "",
		Testnet:    true,
	}
	api := okx.NewOkxRestApi(&acc, "")
	req := models.FillsReq{
		InstType:   "",
		Uly:        "",
		InstFamily: "",
		InstId:     "",
		OrdType:    "",
		Category:   "",
		After:      "",
		Before:     "",
		Begin:      "",
		End:        "",
		Limit:      "",
	}
	message, err := api.Fills(req)
	if err != nil {
		return
	}
	fmt.Println(message)
}
func GetFillsHistory() {
	acc := okx.OkxAccount{
		ApiKey:     apikey,
		ApiSecret:  sercet,
		Passphrase: passphrase,
		BaseURL:    "",
		Testnet:    true,
	}
	api := okx.NewOkxRestApi(&acc, "")
	req := models.FillsHistoryReq{
		InstType:   okx.SWAP,
		Uly:        "",
		InstFamily: "",
		InstId:     "",
		OrdId:      "",
		After:      "",
		Before:     "",
		Begin:      "",
		End:        "",
		Limit:      "",
	}
	message, err := api.FillsHistory(req)
	if err != nil {
		return
	}
	fmt.Println(message)
}

// 资金
func Currencies() {
	acc := okx.OkxAccount{
		ApiKey:     apikey,
		ApiSecret:  sercet,
		Passphrase: passphrase,
		BaseURL:    "",
		Testnet:    true,
	}
	api := okx.NewOkxRestApi(&acc, "")
	req := models.CurrenciesReq{}
	message, err := api.Currencies(req)
	if err != nil {
		return
	}
	fmt.Println(message)
}
func Balances() {
	acc := okx.OkxAccount{
		ApiKey:     apikey,
		ApiSecret:  sercet,
		Passphrase: passphrase,
		BaseURL:    "",
		Testnet:    true,
	}
	api := okx.NewOkxRestApi(&acc, "")
	req := models.BalancesReq{}
	message, err := api.Balances(req)
	if err != nil {
		return
	}
	fmt.Println(message)
}
func NonTradableAssets() {
	acc := okx.OkxAccount{
		ApiKey:     apikey,
		ApiSecret:  sercet,
		Passphrase: passphrase,
		BaseURL:    "",
		Testnet:    true,
	}
	api := okx.NewOkxRestApi(&acc, "")
	req := models.NonTradableAssetsReq{}
	message, err := api.NonTradableAssets(req)
	if err != nil {
		return
	}
	fmt.Println(message)
}
func AssetValuation() {
	acc := okx.OkxAccount{
		ApiKey:     apikey,
		ApiSecret:  sercet,
		Passphrase: passphrase,
		BaseURL:    "",
		Testnet:    true,
	}
	api := okx.NewOkxRestApi(&acc, "")
	req := models.AssetValuationReq{}
	message, err := api.AssetValuation(req)
	if err != nil {
		return
	}
	fmt.Println(message)
}
func Transfer() {
	acc := okx.OkxAccount{
		ApiKey:     apikey,
		ApiSecret:  sercet,
		Passphrase: passphrase,
		BaseURL:    "",
		Testnet:    true,
	}
	api := okx.NewOkxRestApi(&acc, "")
	req := models.TransferReq{
		Ccy:         "USDT",
		Amt:         "1",
		From:        "18",
		To:          "6",
		SubAcct:     0,
		Type:        0,
		LoanTrans:   false,
		ClientId:    "",
		OmitPosRisk: false,
	}
	message, err := api.Transfer(req)
	if err != nil {
		return
	}
	fmt.Println(message)
}
func TransferState() {
	acc := okx.OkxAccount{
		ApiKey:     apikey,
		ApiSecret:  sercet,
		Passphrase: passphrase,
		BaseURL:    "",
		Testnet:    true,
	}
	api := okx.NewOkxRestApi(&acc, "")
	req := models.TransferStaterReq{
		TransId:  "248492529",
		ClientId: "",
		Type:     0,
	}
	message, err := api.TransferState(req)
	if err != nil {
		return
	}
	fmt.Println(message)
}

// 账户
func AccountBalance() {
	acc := okx.OkxAccount{
		ApiKey:     apikey,
		ApiSecret:  sercet,
		Passphrase: passphrase,
		BaseURL:    "",
		Testnet:    true,
	}
	api := okx.NewOkxRestApi(&acc, "")
	req := models.AccountBalanceReq{}
	message, err := api.AccountBalance(req)
	if err != nil {
		return
	}
	fmt.Println(message)
}
func AccountPosition() {
	acc := okx.OkxAccount{
		ApiKey:     apikey,
		ApiSecret:  sercet,
		Passphrase: passphrase,
		BaseURL:    "",
		Testnet:    true,
	}
	api := okx.NewOkxRestApi(&acc, "")
	req := models.AccountBalanceReq{}
	message, err := api.AccountBalance(req)
	if err != nil {
		return
	}
	fmt.Println(message)
}
func AccountPositionHistory() {
	acc := okx.OkxAccount{
		ApiKey:     apikey,
		ApiSecret:  sercet,
		Passphrase: passphrase,
		BaseURL:    "",
		Testnet:    true,
	}
	api := okx.NewOkxRestApi(&acc, "")
	req := models.AccountPositionRiskReq{}
	message, err := api.AccountPositionRisk(req)
	if err != nil {
		return
	}
	fmt.Println(message)
}
func AccountConfig() {
	acc := okx.OkxAccount{
		ApiKey:     apikey,
		ApiSecret:  sercet,
		Passphrase: passphrase,
		BaseURL:    "",
		Testnet:    true,
	}
	api := okx.NewOkxRestApi(&acc, "")
	message, err := api.AccountConfig()
	if err != nil {
		return
	}
	fmt.Println(message)
}
func SetPositionMode() {
	acc := okx.OkxAccount{
		ApiKey:     apikey,
		ApiSecret:  sercet,
		Passphrase: passphrase,
		BaseURL:    "",
		Testnet:    true,
	}
	api := okx.NewOkxRestApi(&acc, "")
	req := models.SetPositionModeReq{}
	message, err := api.SetPositionMode(req)
	if err != nil {
		return
	}
	fmt.Println(message)
}
func SetLeverage() {
	acc := okx.OkxAccount{
		ApiKey:     apikey,
		ApiSecret:  sercet,
		Passphrase: passphrase,
		BaseURL:    "",
		Testnet:    true,
	}
	api := okx.NewOkxRestApi(&acc, "")
	req := models.SetLeverageReq{
		InstId:  "",
		Ccy:     "",
		Lever:   "",
		MgnMode: "",
		PosSide: "",
	}
	message, err := api.SetLeverage(req)
	if err != nil {
		return
	}
	fmt.Println(message)
}
