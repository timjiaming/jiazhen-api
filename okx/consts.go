package okx

const (
	HTTPBaseURL                   = "https://www.okx.com"
	HTTPBaseURLAws                = "https://aws.okx.com"
	PathTickers                   = "/api/v5/market/tickers"
	PathTicker                    = "/api/v5/market/ticker"
	PathIndexTickers              = "/api/v5/market/index-tickers"
	PathBooks                     = "/api/v5/market/books"
	PathBookLite                  = "/api/v5/market/books-lite"
	PathCandles                   = "/api/v5/market/candles"
	PathHistoryCandles            = "/api/v5/market/history-candles"
	PathMarketPriceCandles        = "/api/v5/market/mark-price-candles"
	PathHistoryMarketPriceCandles = "/api/v5/market/history-mark-price-candles"
	PathTrades                    = "/api/v5/market/trades"
	PathHistoryTrades             = "/api/v5/market/history-trades"
	PathPlatform24Volume          = "/api/v5/market/platform-24-volume"

	PathInstruments             = "/api/v5/public/instruments"
	PathDeliveryExerciseHistory = "/api/v5/public/delivery-exercise-history"
	PathOpenInterest            = "/api/v5/public/open-interest"
	PathFundingrate             = "/api/v5/public/funding-rate"
	PathFundingrateHistory      = "/api/v5/public/funding-rate-history"
	PathPriceLimit              = "/api/v5/public/price-limit"
	PathPlace                   = "/api/v5/trade/order"
	PathBatchOrders             = "/api/v5/trade/batch-orders"
	PathCancelOrder             = "/api/v5/trade/cancel-order"
	PathCancelBatchOrders       = "/api/v5/trade/cancel-batch-orders"
	PathAmendOrder              = "/api/v5/trade/amend-order"
	PathAmendBatchOrders        = "/api/v5/trade/amend-batch-orders"
	PathClosePosition           = "/api/v5/trade/close-position"
	PathOrder                   = "/api/v5/trade/order"
	PathOrdersPending           = "/api/v5/trade/orders-pending"
	PathOrdersHistory           = "/api/v5/trade/orders-history"
	PathOrdersHistoryArchive    = "/api/v5/trade/orders-history-archive"
	PathFills                   = "/api/v5/trade/fills"
	PathFillsHistory            = "/api/v5/trade/fills-history"
	PathCurrencies              = "/api/v5/asset/currencies"
	PathBalances                = "/api/v5/asset/balances"
	PathNonTradableAssets       = "/api/v5/asset/non-tradable-assets"
	PathAssetValuation          = "/api/v5/asset/asset-valuation"
	PathTransfer                = "/api/v5/asset/transfer"
	PathTransferState           = "/api/v5/asset/transfer-state"
	PathAccountBalance          = "/api/v5/account/balance"
	PathAccountPosition         = "/api/v5/account/positions"
	PathAccountPositionRisk     = "/api/v5/account/account-position-risk"
	PathAccountConfig           = "/api/v5/account/config"
	PathSetPositionMode         = "/api/v5/account/set-position-mode"
	PathSetLeverage             = "/api/v5/account/set-leverage"
)
const (
	URLRealWSPublic  = "wss://ws.okx.com:8443/ws/v5/public"                // 实盘公共频道
	URLRealWSPrivate = "wss://ws.okx.com:8443/ws/v5/private"               // 实盘私有频道
	URLMockWSPublic  = "wss://ws.okx.com:8443/ws/v5/public?brokerId=9999"  // 模拟盘公共频道
	URLMockWSPrivate = "wss://ws.okx.com:8443/ws/v5/private?brokerId=9999" // 模拟盘私有频道

	URLAwsWSPublic  = "wss://wsaws.okx.com:8443/ws/v5/public"
	URLAwsWSPrivate = "wss://wsaws.okx.com:8443/ws/v5/private"
)
const (
	channelBooks   = "books"
	channelTickers = "tickers"
	channelTrades  = "trades"
	channelBbo     = "bbo-tbt"
	subscribe      = "subscribe"
)
const (
	Isolated          = "isolated"          //逐仓
	Cross             = "cross"             //全仓
	Cash              = "cash"              //非保证金
	Buy               = "buy"               //买
	Sell              = "sell"              //卖
	Market            = "market"            //市价单
	Limit             = "limit"             //限价单
	Postonly          = "post_only"         //只做maker单
	Fok               = "fok"               //全部成交或立即取消
	Ioc               = "ioc"               //立即成交并取消剩余
	Optimal_limit_ioc = "optimal_limit_ioc" //市价委托立即成交并取消剩余（仅适用交割、永续）

	SPOT    = "SPOT"    //币币
	MARGIN  = "MARGIN"  //币币杠杆
	SWAP    = "SWAP"    // 永续合约
	FUTURES = "FUTURES" // 交割合约
	OPTION  = "OPTION"  // 期权
)
