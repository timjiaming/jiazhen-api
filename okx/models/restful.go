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

type PlaceOrderReq struct {
	InstId          string `json:"instId"`
	TdMode          string `json:"tdMode"`
	Ccy             string `json:"ccy"`
	ClOrdId         string `json:"clOrdId"`
	Tag             string `json:"tag"`
	Side            string `json:"side"`
	PosSide         string `json:"posSide"`
	OrdType         string `json:"ordType"`
	Sz              string `json:"sz"`
	Px              string `json:"px"`
	ReduceOnly      string `json:"reduceOnly"`
	TgtCcy          string `json:"tgtCcy"`
	BanAmend        string `json:"banAmend"`
	TpTriggerPx     string `json:"tpTriggerPx"`
	TpOrdPx         string `json:"tpOrdPx"`
	SlTriggerPx     string `json:"slTriggerPx"`
	SlOrdPx         string `json:"slOrdPx"`
	TpTriggerPxType string `json:"tpTriggerPxType"`
	SlTriggerPxType string `json:"slTriggerPxType"`
	QuickMgnType    string `json:"quickMgnType"`
}
type GetPlaceOrderData []placeOrderData
type placeOrderData struct {
	ClOrdId string `json:"clOrdId"`
	OrdId   string `json:"ordId"`
	Tag     string `json:"tag"`
	SCode   string `json:"sCode"`
	SMsg    string `json:"sMsg"`
}
type PlaceBatchOrdersReq []PlaceOrderReq
type GetPlacBatheData []placeOrderData

type CancelOrderReq struct {
	InstId  string `json:"instId"`
	OrdId   string `json:"ordId"`
	ClOrdId string `json:"clOrdId"`
}
type GetCancelOrderData []cancelOrder
type cancelOrder struct {
	ClOrdId string `json:"clOrdId"`
	OrdId   string `json:"ordId"`
	SCode   string `json:"sCode"`
	SMsg    string `json:"sMsg"`
}
type CancelBatchOrdersReq []CancelOrderReq
type GetCancelBatchOrdersData []cancelOrder
type AmendOrderReq struct {
	InstId    string `json:"instId"`
	CxlOnFail string `json:"cxlOnFail"`
	OrdId     string `json:"ordId"`
	ClOrdId   string `json:"clOrdId"`
	ReqId     string `json:"reqId"`
	NewSz     string `json:"newSz"`
	NewPx     string `json:"newPx"`
}
type GetAmendOrderData []amendOrder
type amendOrder struct {
	ClOrdId string `json:"clOrdId"`
	OrdId   string `json:"ordId"`
	ReqId   string `json:"reqId"`
	SCode   string `json:"sCode"`
	SMsg    string `json:"sMsg"`
}
type AmendBatchOrdersReq []AmendOrderReq
type GetAmendBatchOrderData []amendOrder

type ClosePositionReq struct {
	InstId  string `json:"instId"`
	PosSide string `json:"posSide"`
	MgnMode string `json:"mgnMode"`
	Ccy     string `json:"ccy"`
	AutoCxl string `json:"autoCxl"`
	ClOrdId string `json:"clOrdId"`
	Tag     string `json:"tag"`
}
type GetClosePosition []closePosition
type closePosition struct {
	InstId  string `json:"instId"`
	PosSide string `json:"posSide"`
}
type OrderReq struct {
	InstId  string `json:"instId"`
	OrdId   string `json:"ordId"`
	ClOrdId string `json:"clOrdId"`
}
type GetOrderData []order
type order struct {
	InstType           string `json:"instType"`
	InstId             string `json:"instId"`
	Ccy                string `json:"ccy"`
	OrdId              string `json:"ordId"`
	ClOrdId            string `json:"clOrdId"`
	Tag                string `json:"tag"`
	Px                 string `json:"px"`
	Sz                 string `json:"sz"`
	Pnl                string `json:"pnl"`
	OrdType            string `json:"ordType"`
	Side               string `json:"side"`
	PosSide            string `json:"posSide"`
	TdMode             string `json:"tdMode"`
	AccFillSz          string `json:"accFillSz"`
	FillPx             string `json:"fillPx"`
	TradeId            string `json:"tradeId"`
	FillSz             string `json:"fillSz"`
	FillTime           string `json:"fillTime"`
	Source             string `json:"source"`
	State              string `json:"state"`
	AvgPx              string `json:"avgPx"`
	Lever              string `json:"lever"`
	TpTriggerPx        string `json:"tpTriggerPx"`
	TpTriggerPxType    string `json:"tpTriggerPxType"`
	TpOrdPx            string `json:"tpOrdPx"`
	SlTriggerPx        string `json:"slTriggerPx"`
	SlTriggerPxType    string `json:"slTriggerPxType"`
	SlOrdPx            string `json:"slOrdPx"`
	FeeCcy             string `json:"feeCcy"`
	Fee                string `json:"fee"`
	RebateCcy          string `json:"rebateCcy"`
	Rebate             string `json:"rebate"`
	TgtCcy             string `json:"tgtCcy"`
	Category           string `json:"category"`
	ReduceOnly         string `json:"reduceOnly"`
	CancelSource       string `json:"cancelSource"`
	CancelSourceReason string `json:"cancelSourceReason"`
	QuickMgnType       string `json:"quickMgnType"`
	UTime              string `json:"uTime"`
	CTime              string `json:"cTime"`
}
type OrdersPendingReq struct {
	InstType   string `json:"instType"`
	Uly        string `json:"uly"`
	InstFamily string `json:"instFamily"`
	InstId     string `json:"instId"`
	OrdType    string `json:"ordType"`
	State      string `json:"state"`
	After      string `json:"after"`
	Before     string `json:"before"`
	Limit      string `json:"limit"`
}
type GetOrdersPendingData []ordersPending
type ordersPending struct {
	AccFillSz       string `json:"accFillSz"`
	AvgPx           string `json:"avgPx"`
	CTime           string `json:"cTime"`
	Category        string `json:"category"`
	Ccy             string `json:"ccy"`
	ClOrdId         string `json:"clOrdId"`
	Fee             string `json:"fee"`
	FeeCcy          string `json:"feeCcy"`
	FillPx          string `json:"fillPx"`
	FillSz          string `json:"fillSz"`
	FillTime        string `json:"fillTime"`
	InstId          string `json:"instId"`
	InstType        string `json:"instType"`
	Lever           string `json:"lever"`
	OrdId           string `json:"ordId"`
	OrdType         string `json:"ordType"`
	Pnl             string `json:"pnl"`
	PosSide         string `json:"posSide"`
	Px              string `json:"px"`
	Rebate          string `json:"rebate"`
	RebateCcy       string `json:"rebateCcy"`
	Side            string `json:"side"`
	SlOrdPx         string `json:"slOrdPx"`
	SlTriggerPx     string `json:"slTriggerPx"`
	SlTriggerPxType string `json:"slTriggerPxType"`
	Source          string `json:"source"`
	State           string `json:"state"`
	Sz              string `json:"sz"`
	Tag             string `json:"tag"`
	TdMode          string `json:"tdMode"`
	TgtCcy          string `json:"tgtCcy"`
	TpOrdPx         string `json:"tpOrdPx"`
	TpTriggerPx     string `json:"tpTriggerPx"`
	TpTriggerPxType string `json:"tpTriggerPxType"`
	TradeId         string `json:"tradeId"`
	ReduceOnly      string `json:"reduceOnly"`
	QuickMgnType    string `json:"quickMgnType"`
	UTime           string `json:"uTime"`
}

type OrdersHistoryReq struct {
	InstType   string `json:"instType"`
	Uly        string `json:"uly"`
	InstFamily string `json:"instFamily"`
	InstId     string `json:"instId"`
	OrdType    string `json:"ordType"`
	State      string `json:"state"`
	Category   string `json:"category"`
	After      string `json:"after"`
	Before     string `json:"before"`
	Begin      string `json:"begin"`
	End        string `json:"end"`
	Limit      string `json:"limit"`
}
type GetOrderHistoryData []orderHistory
type orderHistory struct {
	InstType           string `json:"instType"`
	InstId             string `json:"instId"`
	Ccy                string `json:"ccy"`
	OrdId              string `json:"ordId"`
	ClOrdId            string `json:"clOrdId"`
	Tag                string `json:"tag"`
	Px                 string `json:"px"`
	Sz                 string `json:"sz"`
	OrdType            string `json:"ordType"`
	Side               string `json:"side"`
	PosSide            string `json:"posSide"`
	TdMode             string `json:"tdMode"`
	AccFillSz          string `json:"accFillSz"`
	FillPx             string `json:"fillPx"`
	TradeId            string `json:"tradeId"`
	FillSz             string `json:"fillSz"`
	FillTime           string `json:"fillTime"`
	Source             string `json:"source"`
	State              string `json:"state"`
	AvgPx              string `json:"avgPx"`
	Lever              string `json:"lever"`
	TpTriggerPx        string `json:"tpTriggerPx"`
	TpTriggerPxType    string `json:"tpTriggerPxType"`
	TpOrdPx            string `json:"tpOrdPx"`
	SlTriggerPx        string `json:"slTriggerPx"`
	SlTriggerPxType    string `json:"slTriggerPxType"`
	SlOrdPx            string `json:"slOrdPx"`
	FeeCcy             string `json:"feeCcy"`
	Fee                string `json:"fee"`
	RebateCcy          string `json:"rebateCcy"`
	Rebate             string `json:"rebate"`
	TgtCcy             string `json:"tgtCcy"`
	Pnl                string `json:"pnl"`
	Category           string `json:"category"`
	ReduceOnly         string `json:"reduceOnly"`
	CancelSource       string `json:"cancelSource"`
	CancelSourceReason string `json:"cancelSourceReason"`
	UTime              string `json:"uTime"`
	CTime              string `json:"cTime"`
}

type OrdersHistoryArchiveReq struct {
	InstType   string `json:"instType"`
	Uly        string `json:"uly"`
	InstFamily string `json:"instFamily"`
	InstId     string `json:"instId"`
	OrdType    string `json:"ordType"`
	State      string `json:"state"`
	Category   string `json:"category"`
	After      string `json:"after"`
	Before     string `json:"before"`
	Begin      string `json:"begin"`
	End        string `json:"end"`
	Limit      string `json:"limit"`
}
type GetOrdersHistoryArchiveData []ordersHistoryArchive
type ordersHistoryArchive struct {
	InstType           string `json:"instType"`
	InstId             string `json:"instId"`
	Ccy                string `json:"ccy"`
	OrdId              string `json:"ordId"`
	ClOrdId            string `json:"clOrdId"`
	Tag                string `json:"tag"`
	Px                 string `json:"px"`
	Sz                 string `json:"sz"`
	OrdType            string `json:"ordType"`
	Side               string `json:"side"`
	PosSide            string `json:"posSide"`
	TdMode             string `json:"tdMode"`
	AccFillSz          string `json:"accFillSz"`
	FillPx             string `json:"fillPx"`
	TradeId            string `json:"tradeId"`
	FillSz             string `json:"fillSz"`
	FillTime           string `json:"fillTime"`
	Source             string `json:"source"`
	State              string `json:"state"`
	AvgPx              string `json:"avgPx"`
	Lever              string `json:"lever"`
	TpTriggerPx        string `json:"tpTriggerPx"`
	TpTriggerPxType    string `json:"tpTriggerPxType"`
	TpOrdPx            string `json:"tpOrdPx"`
	SlTriggerPx        string `json:"slTriggerPx"`
	SlTriggerPxType    string `json:"slTriggerPxType"`
	SlOrdPx            string `json:"slOrdPx"`
	FeeCcy             string `json:"feeCcy"`
	Fee                string `json:"fee"`
	RebateCcy          string `json:"rebateCcy"`
	Rebate             string `json:"rebate"`
	TgtCcy             string `json:"tgtCcy"`
	Pnl                string `json:"pnl"`
	Category           string `json:"category"`
	ReduceOnly         string `json:"reduceOnly"`
	CancelSource       string `json:"cancelSource"`
	CancelSourceReason string `json:"cancelSourceReason"`
	UTime              string `json:"uTime"`
	CTime              string `json:"cTime"`
}
type FillsReq struct {
	InstType   string `json:"instType"`
	Uly        string `json:"uly"`
	InstFamily string `json:"instFamily"`
	InstId     string `json:"instId"`
	OrdType    string `json:"ordType"`
	Category   string `json:"category"`
	After      string `json:"after"`
	Before     string `json:"before"`
	Begin      string `json:"begin"`
	End        string `json:"end"`
	Limit      string `json:"limit"`
}
type GetFillsData []fills
type fills struct {
	InstType   string `json:"instType"`
	Uly        string `json:"uly"`
	InstFamily string `json:"instFamily"`
	InstId     string `json:"instId"`
	OrdId      string `json:"ordId"`
	After      string `json:"after"`
	Before     string `json:"before"`
	Begin      string `json:"begin"`
	End        string `json:"end"`
	Limit      string `json:"limit"`
}
type FillsHistoryReq struct {
	InstType   string `json:"instType"`
	Uly        string `json:"uly"`
	InstFamily string `json:"instFamily"`
	InstId     string `json:"instId"`
	OrdId      string `json:"ordId"`
	After      string `json:"after"`
	Before     string `json:"before"`
	Begin      string `json:"begin"`
	End        string `json:"end"`
	Limit      string `json:"limit"`
}
type GetFillsHistoryData []fillsHistory
type fillsHistory struct {
	InstType string `json:"instType"`
	InstId   string `json:"instId"`
	TradeId  string `json:"tradeId"`
	OrdId    string `json:"ordId"`
	ClOrdId  string `json:"clOrdId"`
	BillId   string `json:"billId"`
	Tag      string `json:"tag"`
	FillPx   string `json:"fillPx"`
	FillSz   string `json:"fillSz"`
	Side     string `json:"side"`
	PosSide  string `json:"posSide"`
	ExecType string `json:"execType"`
	FeeCcy   string `json:"feeCcy"`
	Fee      string `json:"fee"`
	Ts       string `json:"ts"`
	FillTime string `json:"fillTime"`
}

type CurrenciesReq struct {
	Ccy string `json:"ccy"`
}
type GetCurrenciesData []currencies
type currencies struct {
	CanDep               bool   `json:"canDep"`
	CanInternal          bool   `json:"canInternal"`
	CanWd                bool   `json:"canWd"`
	Ccy                  string `json:"ccy"`
	Chain                string `json:"chain"`
	DepQuotaFixed        string `json:"depQuotaFixed"`
	DepQuoteDailyLayer2  string `json:"depQuoteDailyLayer2"`
	LogoLink             string `json:"logoLink"`
	MainNet              bool   `json:"mainNet"`
	MaxFee               string `json:"maxFee"`
	MaxWd                string `json:"maxWd"`
	MinDep               string `json:"minDep"`
	MinDepArrivalConfirm string `json:"minDepArrivalConfirm"`
	MinFee               string `json:"minFee"`
	MinWd                string `json:"minWd"`
	MinWdUnlockConfirm   string `json:"minWdUnlockConfirm"`
	Name                 string `json:"name"`
	NeedTag              bool   `json:"needTag"`
	UsedDepQuotaFixed    string `json:"usedDepQuotaFixed"`
	UsedWdQuota          string `json:"usedWdQuota"`
	WdQuota              string `json:"wdQuota"`
	WdTickSz             string `json:"wdTickSz"`
}
type BalancesReq struct {
	Ccy string `json:"ccy"`
}
type GetBalancesData []balancesData
type balancesData struct {
	AvailBal  string `json:"availBal"`
	Bal       string `json:"bal"`
	Ccy       string `json:"ccy"`
	FrozenBal string `json:"frozenBal"`
}
type NonTradableAssetsReq struct {
	Ccy string `json:"ccy"`
}
type NonTradableAssetsData []nonTradableAssets
type nonTradableAssets struct {
	Bal      string `json:"bal"`
	CanWd    bool   `json:"canWd"`
	Ccy      string `json:"ccy"`
	Chain    string `json:"chain"`
	CtAddr   string `json:"ctAddr"`
	Fee      string `json:"fee"`
	LogoLink string `json:"logoLink"`
	MinWd    string `json:"minWd"`
	Name     string `json:"name"`
	NeedTag  bool   `json:"needTag"`
	WdAll    bool   `json:"wdAll"`
	WdTickSz string `json:"wdTickSz"`
}
type AssetValuationReq struct {
	Ccy string `json:"ccy"`
}
type GetAssetValuationData []assetValuationData
type assetValuationData struct {
	Details  details `json:"details"`
	TotalBal string  `json:"totalBal"`
	Ts       string  `json:"ts"`
}
type details struct {
	Classic string `json:"classic"`
	Earn    string `json:"earn"`
	Funding string `json:"funding"`
	Trading string `json:"trading"`
}
type TransferReq struct {
	Ccy         string `json:"ccy"`
	Amt         string `json:"amt"`
	From        string `json:"from"`
	To          string `json:"to"`
	SubAcct     int    `json:"subAcct"`
	Type        int    `json:"type"`
	LoanTrans   bool   `json:"loanTrans"`
	ClientId    string `json:"clientId"`
	OmitPosRisk bool   `json:"omitPosRisk"`
}
type GetTransferData []transfer
type transfer struct {
	TransId  string `json:"transId"`
	Ccy      string `json:"ccy"`
	ClientId string `json:"clientId"`
	From     string `json:"from"`
	Amt      string `json:"amt"`
	To       string `json:"to"`
}
type TransferStaterReq struct {
	TransId  string `json:"transId"`
	ClientId string `json:"clientId"`
	Type     int    `json:"type"`
}
type GetTransferStaterData []transferStater
type transferStater struct {
	Amt      string `json:"amt"`
	Ccy      string `json:"ccy"`
	ClientId string `json:"clientId"`
	From     string `json:"from"`
	InstId   string `json:"instId"`
	State    string `json:"state"`
	SubAcct  string `json:"subAcct"`
	To       string `json:"to"`
	ToInstId string `json:"toInstId"`
	TransId  string `json:"transId"`
	Type     string `json:"type"`
}
type GetBillsData []bills
type bills struct {
	BillId   string `json:"billId"`
	Ccy      string `json:"ccy"`
	ClientId string `json:"clientId"`
	BalChg   string `json:"balChg"`
	Bal      string `json:"bal"`
	Type     string `json:"type"`
	Ts       string `json:"ts"`
}
type AccountBalanceReq struct {
	Ccy string `json:"ccy"`
}
type GetAccountBalanceData []accountBalance
type accountBalance struct {
	AdjEq       string           `json:"adjEq"`
	Imr         string           `json:"imr"`
	IsoEq       string           `json:"isoEq"`
	MgnRatio    string           `json:"mgnRatio"`
	Mmr         string           `json:"mmr"`
	NotionalUsd string           `json:"notionalUsd"`
	OrdFroz     string           `json:"ordFroz"`
	TotalEq     string           `json:"totalEq"`
	UTime       string           `json:"uTime"`
	Details     []accountDetails `json:"details"`
}
type accountDetails struct {
	AvailBal      string `json:"availBal"`
	AvailEq       string `json:"availEq"`
	CashBal       string `json:"cashBal"`
	Ccy           string `json:"ccy"`
	CrossLiab     string `json:"crossLiab"`
	DisEq         string `json:"disEq"`
	Eq            string `json:"eq"`
	EqUsd         string `json:"eqUsd"`
	FrozenBal     string `json:"frozenBal"`
	Interest      string `json:"interest"`
	IsoEq         string `json:"isoEq"`
	IsoLiab       string `json:"isoLiab"`
	IsoUpl        string `json:"isoUpl"`
	Liab          string `json:"liab"`
	MaxLoan       string `json:"maxLoan"`
	MgnRatio      string `json:"mgnRatio"`
	NotionalLever string `json:"notionalLever"`
	OrdFrozen     string `json:"ordFrozen"`
	Twap          string `json:"twap"`
	UTime         string `json:"uTime"`
	Upl           string `json:"upl"`
	UplLiab       string `json:"uplLiab"`
	StgyEq        string `json:"stgyEq"`
	SpotInUseAmt  string `json:"spotInUseAmt"`
}
type AccountPosition struct {
	InstType string `json:"instType"`
	InstId   string `json:"instId"`
	PosId    string `json:"posId"`
}
type GetAccountPositionData []accountPosition
type accountPosition struct {
	Adl            string           `json:"adl"`
	AvailPos       string           `json:"availPos"`
	AvgPx          string           `json:"avgPx"`
	CTime          string           `json:"cTime"`
	Ccy            string           `json:"ccy"`
	DeltaBS        string           `json:"deltaBS"`
	DeltaPA        string           `json:"deltaPA"`
	GammaBS        string           `json:"gammaBS"`
	GammaPA        string           `json:"gammaPA"`
	Imr            string           `json:"imr"`
	InstId         string           `json:"instId"`
	InstType       string           `json:"instType"`
	Interest       string           `json:"interest"`
	Last           string           `json:"last"`
	UsdPx          string           `json:"usdPx"`
	Lever          string           `json:"lever"`
	Liab           string           `json:"liab"`
	LiabCcy        string           `json:"liabCcy"`
	LiqPx          string           `json:"liqPx"`
	MarkPx         string           `json:"markPx"`
	Margin         string           `json:"margin"`
	MgnMode        string           `json:"mgnMode"`
	MgnRatio       string           `json:"mgnRatio"`
	Mmr            string           `json:"mmr"`
	NotionalUsd    string           `json:"notionalUsd"`
	OptVal         string           `json:"optVal"`
	PTime          string           `json:"pTime"`
	Pos            string           `json:"pos"`
	PosCcy         string           `json:"posCcy"`
	PosId          string           `json:"posId"`
	PosSide        string           `json:"posSide"`
	SpotInUseAmt   string           `json:"spotInUseAmt"`
	SpotInUseCcy   string           `json:"spotInUseCcy"`
	ThetaBS        string           `json:"thetaBS"`
	ThetaPA        string           `json:"thetaPA"`
	TradeId        string           `json:"tradeId"`
	BizRefId       string           `json:"bizRefId"`
	BizRefType     string           `json:"bizRefType"`
	QuoteBal       string           `json:"quoteBal"`
	BaseBal        string           `json:"baseBal"`
	BaseBorrowed   string           `json:"baseBorrowed"`
	BaseInterest   string           `json:"baseInterest"`
	QuoteBorrowed  string           `json:"quoteBorrowed"`
	QuoteInterest  string           `json:"quoteInterest"`
	UTime          string           `json:"uTime"`
	Upl            string           `json:"upl"`
	UplRatio       string           `json:"uplRatio"`
	VegaBS         string           `json:"vegaBS"`
	VegaPA         string           `json:"vegaPA"`
	CloseOrderAlgo []closeOrderAlgo `json:"closeOrderAlgo"`
}
type closeOrderAlgo struct {
	AlgoId          string `json:"algoId"`
	SlTriggerPx     string `json:"slTriggerPx"`
	SlTriggerPxType string `json:"slTriggerPxType"`
	TpTriggerPx     string `json:"tpTriggerPx"`
	TpTriggerPxType string `json:"tpTriggerPxType"`
	CloseFraction   string `json:"closeFraction"`
}
type PositionHistoryReq struct {
	InstType string `json:"instType"`
	InstId   string `json:"instId"`
	MgnMode  string `json:"mgnMode"`
	Type     string `json:"type"`
	PosId    string `json:"posId"`
	After    string `json:"after"`
	Before   string `json:"before"`
	Limit    string `json:"limit"`
}
type GetPositionHistoryData []positionHistory
type positionHistory struct {
	CTime         string `json:"cTime"`
	Ccy           string `json:"ccy"`
	CloseAvgPx    string `json:"closeAvgPx"`
	CloseTotalPos string `json:"closeTotalPos"`
	InstId        string `json:"instId"`
	InstType      string `json:"instType"`
	Lever         string `json:"lever"`
	MgnMode       string `json:"mgnMode"`
	OpenAvgPx     string `json:"openAvgPx"`
	OpenMaxPos    string `json:"openMaxPos"`
	Pnl           string `json:"pnl"`
	PnlRatio      string `json:"pnlRatio"`
	PosId         string `json:"posId"`
	Direction     string `json:"direction"`
	TriggerPx     string `json:"triggerPx"`
	Type          string `json:"type"`
	UTime         string `json:"uTime"`
	Uly           string `json:"uly"`
}
type AccountPositionRiskReq struct {
	InstType string `json:"instType"`
}
type GetAccountPositionRiskData []accountPositionRisk
type accountPositionRisk struct {
	AdjEq   string  `json:"adjEq"`
	BalData balData `json:"balData"`
	PosData posData `json:"posData"`
}
type balData struct {
	Ccy   string `json:"ccy"`
	DisEq string `json:"disEq"`
	Eq    string `json:"eq"`
}
type posData struct {
	BaseBal     string `json:"baseBal"`
	Ccy         string `json:"ccy"`
	InstId      string `json:"instId"`
	InstType    string `json:"instType"`
	MgnMode     string `json:"mgnMode"`
	NotionalCcy string `json:"notionalCcy"`
	NotionalUsd string `json:"notionalUsd"`
	Pos         string `json:"pos"`
	PosCcy      string `json:"posCcy"`
	PosId       string `json:"posId"`
	PosSide     string `json:"posSide"`
	QuoteBal    string `json:"quoteBal"`
}
type GetaccountConfigData []accountConfig
type accountConfig struct {
	AcctLv         string `json:"acctLv"`
	AutoLoan       bool   `json:"autoLoan"`
	CtIsoMode      string `json:"ctIsoMode"`
	GreeksType     string `json:"greeksType"`
	Level          string `json:"level"`
	LevelTmp       string `json:"levelTmp"`
	MgnIsoMode     string `json:"mgnIsoMode"`
	PosMode        string `json:"posMode"`
	SpotOffsetType string `json:"spotOffsetType"`
	Uid            string `json:"uid"`
	Label          string `json:"label"`
	OpAuth         string `json:"opAuth"`
	Ip             string `json:"ip"`
}
type SetPositionModeReq struct {
	PosMode string `json:"posMode"`
}
type GetSetPositionModeData []setPositionMode
type setPositionMode struct {
	PosMode string `json:"posMode"`
}
type SetLeverageReq struct {
	InstId  string `json:"instId"`
	Ccy     string `json:"ccy"`
	Lever   string `json:"lever"`
	MgnMode string `json:"mgnMode"`
	PosSide string `json:"posSide"`
}
type GetSetLeverageData []setLeverage
type setLeverage struct {
	Lever   string `json:"lever"`
	MgnMode string `json:"mgnMode"`
	InstId  string `json:"instId"`
	PosSide string `json:"posSide"`
}
