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
)
