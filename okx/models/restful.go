package models

import "encoding/json"

type BaseResponse struct {
	Code string          `json:"code"`
	Msg  string          `json:"msg"`
	Data json.RawMessage `json:"data"`
}
type GetTickersReq struct {
	InstType   string `json:"instType"`
	Uly        string `json:"uly"`
	InstFamily string `json:"instFamily"`
}
type GetTickersData []TickersListData

type TickersListData struct {
	InstType  string `json:"instType"`
	InstId    string `json:"instId"`
	Last      string `json:"last"`
	LastSz    string `json:"lastSz"`
	AskPx     string `json:"askPx"`
	AskSz     string `json:"askSz"`
	BidPx     string `json:"bidPx"`
	BidSz     string `json:"bidSz"`
	Open24H   string `json:"open24h"`
	High24H   string `json:"high24h"`
	Low24H    string `json:"low24h"`
	VolCcy24H string `json:"volCcy24h"`
	Vol24H    string `json:"vol24h"`
	SodUtc0   string `json:"sodUtc0"`
	SodUtc8   string `json:"sodUtc8"`
	Ts        string `json:"ts"`
}
type GetTickerReq struct {
	InstId string `json:"instId"`
}
type GetTickerData []tickerData

type tickerData struct {
	InstType  string `json:"instType"`
	InstId    string `json:"instId"`
	Last      string `json:"last"`
	LastSz    string `json:"lastSz"`
	AskPx     string `json:"askPx"`
	AskSz     string `json:"askSz"`
	BidPx     string `json:"bidPx"`
	BidSz     string `json:"bidSz"`
	Open24H   string `json:"open24h"`
	High24H   string `json:"high24h"`
	Low24H    string `json:"low24h"`
	VolCcy24H string `json:"volCcy24h"`
	Vol24H    string `json:"vol24h"`
	Ts        string `json:"ts"`
	SodUtc0   string `json:"sodUtc0"`
	SodUtc8   string `json:"sodUtc8"`
}
type GetIndexTickersReq struct {
	QuoteCcy string `json:"quoteCcy"`
	InstId   string `json:"instId"`
}

type GetIndexTickersData []indexTickersData

type indexTickersData struct {
	InstId  string `json:"instId"`
	IdxPx   string `json:"idxPx"`
	High24H string `json:"high24h"`
	SodUtc0 string `json:"sodUtc0"`
	Open24H string `json:"open24h"`
	Low24H  string `json:"low24h"`
	SodUtc8 string `json:"sodUtc8"`
	Ts      string `json:"ts"`
}

type GetBooksReq struct {
	InstId string `json:"instId"`
	SZ     string `json:"sz"`
}
type GetBooksData []books
type books struct {
	Asks [][]string `json:"asks"`
	Bids [][]string `json:"bids"`
	Ts   string     `json:"ts"`
}

// 获取产品轻量深度
type GetBookLiteReq struct {
	InstId string `json:"instId"`
}
type GetBookLiteData []books

// 获取交易产品K线数据
type GetCandlesReq struct {
	InstId string `json:"instId"`
	Bar    string `json:"bar"`
	After  string `json:"after"`
	Before string `json:"before"`
	Limit  string `json:"limit"`
}

// 获取交易产品K线数据
type GetCandlesData [][]string

// 获取交易产品历史K线数据
type GetHistoryCandlesReq struct {
	InstId string `json:"instId"`
	After  string `json:"after"`
	Before string `json:"before"`
	Bar    string `json:"bar"`
	Limit  string `json:"limit"`
}

type GetHistoryCandlesData [][]string

// 获取标记价格K线数据
type GetMarketPriceCandlesReq struct {
	InstId string `json:"instId"`
	After  string `json:"after"`
	Before string `json:"before"`
	Bar    string `json:"bar"`
	Limit  string `json:"limit"`
}

type GetMarketPriceCandelsData [][]string

// 获取标记价格历史K线数据
type GetHistoryMarketPriceReq struct {
	InstId string `json:"instId"`
	After  string `json:"after"`
	Before string `json:"before"`
	Bar    string `json:"bar"`
	Limit  string `json:"limit"`
}

type GetHistoryMarketPriceCandlesData [][]string

// 获取交易产品公共成交数据
type GetTradesReq struct {
	InstId string `json:"instId"`
	Limit  string `json:"limit"`
}
type GetTradesData []TradesData

type TradesData struct {
	InstId  string `json:"instId"`
	Side    string `json:"side"`
	Sz      string `json:"sz"`
	Px      string `json:"px"`
	TradeId string `json:"tradeId"`
	Ts      string `json:"ts"`
}

// 获取交易产品公共历史成交数据
type GetHistoryTradesReq struct {
	InstId string `json:"instId"`
	After  string `json:"after"`
	Before string `json:"before"`
	Bar    string `json:"bar"`
	Limit  string `json:"limit"`
}
type GetHistoryData []HistoryTradesData
type HistoryTradesData struct {
	InstId  string `json:"instId"`
	Side    string `json:"side"`
	Sz      string `json:"sz"`
	Px      string `json:"px"`
	TradeId string `json:"tradeId"`
	Ts      string `json:"ts"`
}

// 获取平台24小时总成交量
type GetPlatform24Volume struct {
}
type GetPlatform24VolumeData []platform24VolumeData

type platform24VolumeData struct {
	VolCny      string `json:"volCny"`
	BlockVolUsd string `json:"blockVolUsd"`
	VolUsd      string `json:"volUsd"`
	BlockVolCny string `json:"blockVolCny"`
	Ts          string `json:"ts"`
}
type GetInstrumentsReq struct {
	InstType   string `json:"instType"`
	Uly        string `json:"uly"`
	InstFamily string `json:"instFamily"`
	InstId     string `json:"instId"`
}
type GetInstrumentsData []instrumentsData
type instrumentsData struct {
	Alias        string `json:"alias"`
	BaseCcy      string `json:"baseCcy"`
	Category     string `json:"category"`
	CtMult       string `json:"ctMult"`
	CtType       string `json:"ctType"`
	CtVal        string `json:"ctVal"`
	CtValCcy     string `json:"ctValCcy"`
	ExpTime      string `json:"expTime"`
	InstFamily   string `json:"instFamily"`
	InstId       string `json:"instId"`
	InstType     string `json:"instType"`
	Lever        string `json:"lever"`
	ListTime     string `json:"listTime"`
	LotSz        string `json:"lotSz"`
	MaxIcebergSz string `json:"maxIcebergSz"`
	MaxLmtSz     string `json:"maxLmtSz"`
	MaxMktSz     string `json:"maxMktSz"`
	MaxStopSz    string `json:"maxStopSz"`
	MaxTriggerSz string `json:"maxTriggerSz"`
	MaxTwapSz    string `json:"maxTwapSz"`
	MinSz        string `json:"minSz"`
	OptType      string `json:"optType"`
	QuoteCcy     string `json:"quoteCcy"`
	SettleCcy    string `json:"settleCcy"`
	State        string `json:"state"`
	Stk          string `json:"stk"`
	TickSz       string `json:"tickSz"`
	Uly          string `json:"uly"`
}

type GetDeliveryExerciseHistoryReq struct {
	InstType   string `json:"instType"`
	Uly        string `json:"uly"`
	InstFamily string `json:"instFamily"`
	After      string `json:"after"`
	Before     string `json:"before"`
	Limit      string `json:"limit"`
}

type GetDeliveryExerciseHistoryData []DeliveryExerciseHistory
type DeliveryExerciseHistory struct {
	Ts      string                           `json:"ts"`
	Details []DeliveryExerciseHistoryDetails `json:"details"`
}
type DeliveryExerciseHistoryDetails struct {
	Type   string `json:"type"`
	InstId string `json:"instId"`
	Px     string `json:"px"`
}

type GetOpenInterestReq struct {
	InstType   string `json:"instType"`
	Uly        string `json:"uly"`
	InstFamily string `json:"instFamily"`
	Limit      string `json:"limit"`
}

type GetOpenInterestData []openInterestData

type openInterestData struct {
	InstType string `json:"instType"`
	InstId   string `json:"instId"`
	Oi       string `json:"oi"`
	OiCcy    string `json:"oiCcy"`
	Ts       string `json:"ts"`
}

type GetFundingrateReq struct {
	InstId string `json:"instId"`
}
type GetFundingrateData []fundingrate
type fundingrate struct {
	FundingRate     string `json:"fundingRate"`
	FundingTime     string `json:"fundingTime"`
	InstId          string `json:"instId"`
	InstType        string `json:"instType"`
	NextFundingRate string `json:"nextFundingRate"`
	NextFundingTime string `json:"nextFundingTime"`
}

type GetFundingrateHistoryReq struct {
	InstId string `json:"instId"`
	Before string `json:"before"`
	After  string `json:"after"`
	Limit  string `json:"limit"`
}
type GetFundingrateHistoryData []fundingrateHistory
type fundingrateHistory struct {
	InstType     string `json:"instType"`
	InstId       string `json:"instId"`
	FundingRate  string `json:"fundingRate"`
	RealizedRate string `json:"realizedRate"`
	FundingTime  string `json:"fundingTime"`
}

type GetPriceLimitReq struct {
	InstId string `json:"instId"`
}
type GetPriceLimitData []priceLimitData
type priceLimitData struct {
	InstType string `json:"instType"`
	InstId   string `json:"instId"`
	BuyLmt   string `json:"buyLmt"`
	SellLmt  string `json:"sellLmt"`
	Ts       string `json:"ts"`
}
type GetEstimatedPriceReq struct {
	InstId string `json:"instId"`
}
type GetEstimatedPriceData []estimatedPriceData

type estimatedPriceData struct {
	InstType string `json:"instType"`
	InstId   string `json:"instId"`
	SettlePx string `json:"settlePx"`
	Ts       string `json:"ts"`
}
