package okx

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/timjiaming/jiazhen-api/common"
	"github.com/timjiaming/jiazhen-api/okx/models"
	"github.com/timjiaming/jiazhen-api/okx/utils"
	"net/url"
)

const (
	ACCEPT               = "ACCEPT"
	CONTENT_TYPE         = "Content-Type"
	OK_ACCESS_KEY        = "OK-ACCESS-KEY"
	OK_ACCESS_SIGN       = "OK-ACCESS-SIGN"
	OK_ACCESS_TIMESTAMP  = "OK-ACCESS-TIMESTAMP"
	OK_ACCESS_PASSPHRASE = "OK-ACCESS-PASSPHRASE"
	X_SIMULATE_TRADING   = "x-simulated-trading"

	APPLICATION_JSON      = "application/json"
	APPLICATION_JSON_UTF8 = "application/json; charset=UTF-8"
)

type OkxRestApi struct {
	Account *OkxAccount
	Client  *resty.Client
}

func NewOkxRestApi(acc *OkxAccount, endpoint string) *OkxRestApi {
	client := resty.NewWithClient(common.NewHTTPClient())
	if acc.BaseURL == "" {
		//acc.BaseURL = HTTPBaseURL
		client.SetBaseURL(HTTPBaseURL)
	}
	api := OkxRestApi{
		Account: acc,
		Client:  client,
	}
	//client.OnBeforeRequest()
	//client.OnAfterResponse()
	return &api
}
func (api *OkxRestApi) SendPublicRequest(method, endpoint string, query map[string]string, body string) (json.RawMessage, error) {
	request := api.Client.R()
	if query != nil {
		values := url.Values{}
		for k, v := range query {
			values.Add(k, v)
		}
		enc := values.Encode()
		if enc != "" {
			request.SetQueryString(enc)
		}
		if body != "" {
			request.SetBody(body)
		}
	}

	response, err := request.Execute(method, endpoint)
	if err != nil {
		return nil, err
	}
	code := response.StatusCode()
	responseBody := response.Body()
	if code < 200 || code >= 300 {
		return nil, fmt.Errorf("invalid HTTP status code %d, %s", code, string(responseBody))
	}
	responseObj := models.BaseResponse{}
	err = json.Unmarshal(responseBody, &responseObj)
	if err != nil {
		return nil, err
	}
	if responseObj.Code == "0" {
		return responseObj.Data, nil
	}
	return nil, fmt.Errorf("invalid response code %d, %s", responseObj.Code, responseObj.Msg)
}
func (api *OkxRestApi) SendRequest(method, endpoint string, query map[string]string, body string) (json.RawMessage, error) {
	signPath := endpoint
	req := api.Client.R()
	if query != nil {
		values := url.Values{}
		for k, v := range query {
			values.Add(k, v)
		}
		enc := values.Encode()
		if enc != "" {
			req.SetQueryString(enc)
			signPath = endpoint + "?" + enc
		}
	}
	ts, sign := api.Sign(method, signPath, body)
	headers := api.genHeaders(ts, sign)
	req.SetHeaders(headers)
	if body != "" {
		req.SetBody(body)
	}

	resp, err := req.Execute(method, endpoint)
	if err != nil {
		return nil, err
	}
	code := resp.StatusCode()
	responseBody := resp.Body()
	if code < 200 || code >= 300 { // 非2xx返回码认为是错误
		return nil, fmt.Errorf("invalid HTTP status code %d, %s", code, string(responseBody))
	}
	responseObj := models.BaseResponse{}
	err = json.Unmarshal(responseBody, &responseObj)
	if err != nil {
		return nil, err
	}
	if responseObj.Code == "0" { // HTTP返回码为200且应用层code为0，认为是成功响应
		return responseObj.Data, nil
	}
	return nil, fmt.Errorf("invalid response code %s, %s", responseObj.Code, responseObj.Msg)
}
func (api *OkxRestApi) genHeaders(timestamp string, sign string) map[string]string {
	headers := make(map[string]string)
	headers[ACCEPT] = APPLICATION_JSON
	headers[CONTENT_TYPE] = APPLICATION_JSON_UTF8
	headers[OK_ACCESS_KEY] = api.Account.ApiKey
	headers[OK_ACCESS_SIGN] = sign
	headers[OK_ACCESS_TIMESTAMP] = timestamp
	headers[OK_ACCESS_PASSPHRASE] = api.Account.Passphrase
	if api.Account.Testnet { // 模拟盘
		headers[X_SIMULATE_TRADING] = "1"
	}
	return headers
}
func apiPublicTemplateWithQuery(api *OkxRestApi, method string, path string, query map[string]string, response any) error {
	data, err := api.SendPublicRequest(method, path, query, "")
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &response)
}
func apiPriveTemplateWithQuery(api *OkxRestApi, method string, path string, query map[string]string, response any) error {
	data, err := api.SendRequest(method, path, query, "")
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &response)
}
func apiPrivateTemplateWithBody(api *OkxRestApi, method string, path string, req any, resp any) error {
	reqData, err := json.Marshal(req)
	if err != nil {
		return err
	}
	data, err := api.SendRequest(method, path, nil, string(reqData))
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return err
	}
	return nil
}
func (api *OkxRestApi) Sign(method string, uri string, body string) (string, string) {
	ts := utils.IsoTimestamp()
	preHash := fmt.Sprintf("%s%s%s%s", ts, method, uri, body)
	//fmt.Println("preHash:", preHash)
	p := []byte(preHash)
	h := hmac.New(sha256.New, []byte(api.Account.ApiSecret))
	h.Write(p)
	return ts, base64.StdEncoding.EncodeToString(h.Sum(nil))
}

// 获取所有产品行情信息
func (api *OkxRestApi) GetTickers(req models.GetTickersReq) (models.GetTickersData, error) {
	query := make(map[string]string)
	if req.InstType != "" {
		query["instType"] = req.InstType
	}
	if req.Uly != "" {
		query["url"] = req.Uly
	}
	if req.InstFamily != "" {
		query["instFamily"] = req.InstFamily
	}
	response := models.GetTickersData{}
	err := apiPublicTemplateWithQuery(api, resty.MethodGet, PathTickers, query, &response)
	if err == nil {
		return response, err
	}
	return nil, err
}

// 获取单个产品行情信息
func (api *OkxRestApi) GetTicker(req models.GetTickerReq) (models.GetTickerData, error) {
	query := make(map[string]string)
	if req.InstId != "" {
		query["instId"] = req.InstId
	}
	response := models.GetTickerData{}
	err := apiPublicTemplateWithQuery(api, resty.MethodGet, PathTicker, query, &response)
	if err == nil {
		return response, err
	}
	return nil, err
}

// 获取指数行情
func (api *OkxRestApi) GetIndexTickers(req models.GetIndexTickersReq) (models.GetIndexTickersData, error) {
	query := make(map[string]string)
	if req.QuoteCcy != "" {
		query["quoteCcy"] = req.QuoteCcy
	}
	if req.InstId != "" {
		query["instId"] = req.InstId
	}
	response := models.GetIndexTickersData{}
	err := apiPublicTemplateWithQuery(api, resty.MethodGet, PathIndexTickers, query, &response)
	if err == nil {
		return response, err
	}
	return nil, err
}

// 获取产品深度
func (api *OkxRestApi) GetBooks(req models.GetBooksReq) (models.GetBooksData, error) {
	query := make(map[string]string)
	if req.InstId != "" {
		query["instId"] = req.InstId
	}
	if req.SZ != "" {
		query["sz"] = req.SZ
	}
	response := models.GetBooksData{}
	err := apiPublicTemplateWithQuery(api, resty.MethodGet, PathBooks, query, &response)
	if err == nil {
		return response, err
	}
	return nil, err
}

// 获取产品轻量深度
func (api *OkxRestApi) GetBooksLite(req models.GetBookLiteReq) (models.GetBookLiteData, error) {
	query := make(map[string]string)
	if req.InstId != "" {
		query["instId"] = req.InstId
	}
	response := models.GetBookLiteData{}
	err := apiPublicTemplateWithQuery(api, resty.MethodGet, PathBookLite, query, &response)
	if err == nil {
		return response, err
	}
	return nil, err
}

// 获取交易产品K线数据
func (api *OkxRestApi) GetCandles(req models.GetCandlesReq) (models.GetCandlesData, error) {
	query := make(map[string]string)
	if req.InstId != "" {
		query["instId"] = req.InstId
	}
	if req.Bar != "" {
		query["bar"] = req.Bar
	}
	if req.After != "" {
		query["after"] = req.After
	}
	if req.Before != "" {
		query["before"] = req.Before
	}
	if req.Limit != "" {
		query["limit"] = req.Limit
	}
	response := models.GetCandlesData{}
	err := apiPublicTemplateWithQuery(api, resty.MethodGet, PathCandles, query, &response)
	if err == nil {
		return response, err
	}
	return nil, err
}
func (api *OkxRestApi) GetHistoryCandles(req models.GetHistoryCandlesReq) (models.GetHistoryCandlesData, error) {
	query := make(map[string]string)
	if req.InstId != "" {
		query["instId"] = req.InstId
	}
	if req.Bar != "" {
		query["bar"] = req.Bar
	}
	if req.After != "" {
		query["after"] = req.After
	}
	if req.Before != "" {
		query["before"] = req.Before
	}
	if req.Limit != "" {
		query["limit"] = req.Limit
	}
	response := models.GetHistoryCandlesData{}
	err := apiPublicTemplateWithQuery(api, resty.MethodGet, PathHistoryCandles, query, &response)
	if err == nil {
		return response, err
	}
	return nil, err
}
func (api *OkxRestApi) GetMarketPriceCandles(req models.GetMarketPriceCandlesReq) (models.GetMarketPriceCandelsData, error) {
	query := make(map[string]string)
	if req.InstId != "" {
		query["instId"] = req.InstId
	}
	if req.Bar != "" {
		query["bar"] = req.Bar
	}
	if req.After != "" {
		query["after"] = req.After
	}
	if req.Before != "" {
		query["before"] = req.Before
	}
	if req.Limit != "" {
		query["limit"] = req.Limit
	}
	response := models.GetMarketPriceCandelsData{}
	err := apiPublicTemplateWithQuery(api, resty.MethodGet, PathMarketPriceCandles, query, &response)
	if err == nil {
		return response, nil
	}
	return nil, err
}
func (api *OkxRestApi) GetHistoryMarketPriceCandles(req models.GetHistoryMarketPriceReq) (models.GetHistoryMarketPriceCandlesData, error) {
	query := make(map[string]string)
	if req.InstId != "" {
		query["instId"] = req.InstId
	}
	if req.Bar != "" {
		query["bar"] = req.Bar
	}
	if req.After != "" {
		query["after"] = req.After
	}
	if req.Before != "" {
		query["before"] = req.Before
	}
	if req.Limit != "" {
		query["limit"] = req.Limit
	}
	response := models.GetHistoryMarketPriceCandlesData{}
	err := apiPublicTemplateWithQuery(api, resty.MethodGet, PathHistoryMarketPriceCandles, query, &response)
	if err == nil {
		return response, nil
	}
	return nil, err
}
func (api *OkxRestApi) GetTrades(req models.GetTradesReq) (models.GetTradesData, error) {
	query := make(map[string]string)
	if req.InstId != "" {
		query["instId"] = req.InstId
	}
	if req.Limit != "" {
		query["limit"] = req.Limit
	}
	response := models.GetTradesData{}
	err := apiPublicTemplateWithQuery(api, resty.MethodGet, PathTrades, query, &response)
	if err == nil {
		return response, nil
	}
	return nil, err
}

func (api *OkxRestApi) GetHistoryTrades(req models.GetHistoryTradesReq) (models.GetHistoryData, error) {
	query := make(map[string]string)
	if req.InstId != "" {
		query["instId"] = req.InstId
	}
	if req.Bar != "" {
		query["bar"] = req.Bar
	}
	if req.After != "" {
		query["after"] = req.After
	}
	if req.Before != "" {
		query["before"] = req.Before
	}
	if req.Limit != "" {
		query["limit"] = req.Limit
	}
	response := models.GetHistoryData{}
	err := apiPublicTemplateWithQuery(api, resty.MethodGet, PathHistoryTrades, query, &response)
	if err == nil {
		return response, nil
	}
	return nil, err
}
func (api *OkxRestApi) GetPlatform24Volume() (models.GetPlatform24VolumeData, error) {
	query := make(map[string]string)
	response := models.GetPlatform24VolumeData{}
	err := apiPublicTemplateWithQuery(api, resty.MethodGet, PathPlatform24Volume, query, &response)
	if err == nil {
		return response, nil
	}
	return nil, err
}
func (api *OkxRestApi) GetInstruments(req models.GetInstrumentsReq) (models.GetInstrumentsData, error) {
	query := make(map[string]string)
	if req.InstType != "" {
		query["instType"] = req.InstType
	}
	if req.Uly != "" {
		query["uly"] = req.Uly
	}
	if req.InstFamily != "" {
		query["instFamily"] = req.InstFamily
	}
	if req.InstId != "" {
		query["instId"] = req.InstId
	}
	response := models.GetInstrumentsData{}
	err := apiPublicTemplateWithQuery(api, resty.MethodGet, PathInstruments, query, &response)
	if err == nil {
		return response, nil
	}
	return nil, err
}

// 获取交割和行权记录
func (api *OkxRestApi) GetDeliveryExerciseHistory(req models.GetDeliveryExerciseHistoryReq) (models.GetDeliveryExerciseHistoryData, error) {
	query := make(map[string]string)
	if req.InstType != "" {
		query["instType"] = req.InstType
	}
	if req.Uly != "" {
		query["uly"] = req.Uly
	}
	if req.InstFamily != "" {
		query["instFamily"] = req.InstFamily
	}
	if req.After != "" {
		query["after"] = req.After
	}
	if req.Before != "" {
		query["before"] = req.Before
	}
	if req.Limit != "" {
		query["limit"] = req.Limit
	}
	response := models.GetDeliveryExerciseHistoryData{}
	err := apiPublicTemplateWithQuery(api, resty.MethodGet, PathDeliveryExerciseHistory, query, &response)
	if err == nil {
		return response, nil
	}
	return nil, err
}

// 获取持仓总量
func (api *OkxRestApi) GetOpenInterest(req models.GetOpenInterestReq) (models.GetOpenInterestData, error) {
	query := make(map[string]string)
	if req.InstType != "" {
		query["instType"] = req.InstType
	}
	if req.Uly != "" {
		query["uly"] = req.Uly
	}
	if req.InstFamily != "" {
		query["instFamily"] = req.InstFamily
	}
	if req.InstType != "" {
		query["instId"] = req.InstType
	}
	response := models.GetOpenInterestData{}
	err := apiPublicTemplateWithQuery(api, resty.MethodGet, PathOpenInterest, query, &response)
	if err == nil {
		return response, nil
	}
	return nil, err
}

// 获取永续合约当前资金费率
func (api *OkxRestApi) GetFundingrate(req models.GetFundingrateReq) (models.GetFundingrateData, error) {
	query := make(map[string]string)
	if req.InstId != "" {
		query["instId"] = req.InstId
	}
	response := models.GetFundingrateData{}
	err := apiPublicTemplateWithQuery(api, resty.MethodGet, PathFundingrate, query, &response)
	if err == nil {
		return response, nil
	}
	return nil, err
}

// 获取永续合约历史资金费率
func (api *OkxRestApi) GetFundingrateHistory(req models.GetFundingrateHistoryReq) (models.GetFundingrateHistoryData, error) {
	query := make(map[string]string)
	if req.InstId != "" {
		query["instId"] = req.InstId
	}
	if req.Before != "" {
		query["before"] = req.Before
	}
	if req.After != "" {
		query["after"] = req.After
	}
	if req.Limit != "" {
		query["limit"] = req.Limit
	}
	response := models.GetFundingrateHistoryData{}
	err := apiPublicTemplateWithQuery(api, resty.MethodGet, PathFundingrateHistory, query, &response)
	if err == nil {
		return response, nil
	}
	return nil, err
}

func (api *OkxRestApi) GetPriceLimit(req models.GetPriceLimitReq) (models.GetPriceLimitData, error) {
	query := make(map[string]string)
	if req.InstId != "" {
		query["instId"] = req.InstId
	}
	response := models.GetPriceLimitData{}
	err := apiPublicTemplateWithQuery(api, resty.MethodGet, PathPriceLimit, query, &response)
	if err == nil {
		return response, nil
	}
	return nil, err
}
func (api *OkxRestApi) GetEstimatedPrice(req models.GetEstimatedPriceReq) (models.GetEstimatedPriceData, error) {
	query := make(map[string]string)
	if req.InstId != "" {
		query["instId"] = req.InstId
	}
	response := models.GetEstimatedPriceData{}
	err := apiPublicTemplateWithQuery(api, resty.MethodGet, PathPriceLimit, query, &response)
	if err == nil {
		return response, nil
	}
	return nil, err
}
func (api *OkxRestApi) PlaceOrder(req models.PlaceOrderReq) (models.GetPlaceOrderData, error) {
	resp := models.GetPlaceOrderData{}
	err := apiPrivateTemplateWithBody(api, resty.MethodPost, PathPlace, req, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (api *OkxRestApi) PlaceBatchOrder(req models.PlaceBatchOrdersReq) (models.GetPlacBatheData, error) {
	resp := models.GetPlacBatheData{}
	err := apiPrivateTemplateWithBody(api, resty.MethodPost, PathBatchOrders, req, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (api *OkxRestApi) CancelOrder(req models.CancelOrderReq) (models.GetCancelOrderData, error) {
	resp := models.GetCancelOrderData{}
	err := apiPrivateTemplateWithBody(api, resty.MethodPost, PathCancelOrder, req, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (api *OkxRestApi) CancelBatchOrder(req models.CancelBatchOrdersReq) (models.GetCancelBatchOrdersData, error) {
	resp := models.GetCancelBatchOrdersData{}
	err := apiPrivateTemplateWithBody(api, resty.MethodPost, PathCancelBatchOrders, req, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (api *OkxRestApi) AmendOrder(req models.AmendOrderReq) (models.GetAmendOrderData, error) {
	resp := models.GetAmendOrderData{}
	err := apiPrivateTemplateWithBody(api, resty.MethodPost, PathAmendOrder, req, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (api *OkxRestApi) AmendBatchOrder(req models.AmendBatchOrdersReq) (models.GetAmendBatchOrderData, error) {
	resp := models.GetAmendBatchOrderData{}
	err := apiPrivateTemplateWithBody(api, resty.MethodPost, PathAmendBatchOrders, req, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (api *OkxRestApi) ClosePosition(req models.ClosePositionReq) (models.GetClosePosition, error) {
	resp := models.GetClosePosition{}
	err := apiPrivateTemplateWithBody(api, resty.MethodPost, PathClosePosition, req, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (api *OkxRestApi) Order(req models.OrderReq) (models.GetOrderData, error) {
	query := make(map[string]string)
	query, _ = common.StructToStringMap(req)
	resp := models.GetOrderData{}
	err := apiPriveTemplateWithQuery(api, resty.MethodGet, PathOrder, query, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (api *OkxRestApi) OrderPending(req models.OrdersPendingReq) (models.GetOrdersPendingData, error) {
	query := make(map[string]string)
	query, _ = common.StructToStringMap(req)
	resp := models.GetOrdersPendingData{}
	err := apiPriveTemplateWithQuery(api, resty.MethodGet, PathOrdersPending, query, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (api *OkxRestApi) OrderHistory(req models.OrdersHistoryReq) (models.GetOrderHistoryData, error) {
	query := make(map[string]string)
	query, _ = common.StructToStringMap(req)
	resp := models.GetOrderHistoryData{}
	err := apiPriveTemplateWithQuery(api, resty.MethodGet, PathOrdersHistory, query, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (api *OkxRestApi) OrderHistoryArchive(req models.OrdersHistoryArchiveReq) (models.GetOrdersHistoryArchiveData, error) {
	query := make(map[string]string)
	query, _ = common.StructToStringMap(req)
	resp := models.GetOrdersHistoryArchiveData{}
	err := apiPriveTemplateWithQuery(api, resty.MethodGet, PathOrdersHistoryArchive, query, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (api *OkxRestApi) Fills(req models.FillsReq) (models.GetFillsData, error) {
	query := make(map[string]string)
	query, _ = common.StructToStringMap(req)
	resp := models.GetFillsData{}
	err := apiPriveTemplateWithQuery(api, resty.MethodGet, PathFills, query, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (api *OkxRestApi) FillsHistory(req models.FillsHistoryReq) (models.GetFillsHistoryData, error) {
	query := make(map[string]string)
	query, _ = common.StructToStringMap(req)
	resp := models.GetFillsHistoryData{}
	err := apiPriveTemplateWithQuery(api, resty.MethodGet, PathFillsHistory, query, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (api *OkxRestApi) Currencies(req models.CurrenciesReq) (models.GetCurrenciesData, error) {
	query := make(map[string]string)
	query, _ = common.StructToStringMap(req)
	resp := models.GetCurrenciesData{}
	err := apiPriveTemplateWithQuery(api, resty.MethodGet, PathCurrencies, query, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (api *OkxRestApi) Balances(req models.BalancesReq) (models.GetBalancesData, error) {
	query := make(map[string]string)
	query, _ = common.StructToStringMap(req)
	resp := models.GetBalancesData{}
	err := apiPriveTemplateWithQuery(api, resty.MethodGet, PathBalances, query, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (api *OkxRestApi) NonTradableAssets(req models.NonTradableAssetsReq) (models.NonTradableAssetsData, error) {
	query := make(map[string]string)
	query, _ = common.StructToStringMap(req)
	resp := models.NonTradableAssetsData{}
	err := apiPriveTemplateWithQuery(api, resty.MethodGet, PathNonTradableAssets, query, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (api *OkxRestApi) AssetValuation(req models.AssetValuationReq) (models.GetAssetValuationData, error) {
	query := make(map[string]string)
	query, _ = common.StructToStringMap(req)
	resp := models.GetAssetValuationData{}
	err := apiPriveTemplateWithQuery(api, resty.MethodGet, PathAssetValuation, query, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (api *OkxRestApi) Transfer(req models.TransferReq) (models.GetTransferData, error) {
	resp := models.GetTransferData{}
	err := apiPrivateTemplateWithBody(api, resty.MethodPost, PathTransfer, req, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (api *OkxRestApi) TransferState(req models.TransferStaterReq) (models.GetTransferData, error) {
	query := make(map[string]string)
	query, _ = common.StructToStringMap(req)
	resp := models.GetTransferData{}
	err := apiPriveTemplateWithQuery(api, resty.MethodGet, PathTransferState, query, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (api *OkxRestApi) AccountBalance(req models.AccountBalanceReq) (models.GetAccountBalanceData, error) {
	query := make(map[string]string)
	query, _ = common.StructToStringMap(req)
	response := models.GetAccountBalanceData{}
	err := apiPriveTemplateWithQuery(api, resty.MethodGet, PathAccountBalance, query, &response)
	if err == nil {
		return response, nil
	}
	return nil, err
}
func (api *OkxRestApi) AccountPositionRisk(req models.AccountPositionRiskReq) (models.GetAccountPositionRiskData, error) {
	query := make(map[string]string)
	query, _ = common.StructToStringMap(req)
	response := models.GetAccountPositionRiskData{}
	err := apiPriveTemplateWithQuery(api, resty.MethodGet, PathAccountPositionRisk, query, &response)
	if err == nil {
		return response, nil
	}
	return nil, err
}
func (api *OkxRestApi) AccountConfig() (models.GetaccountConfigData, error) {
	query := make(map[string]string)
	response := models.GetaccountConfigData{}
	err := apiPriveTemplateWithQuery(api, resty.MethodGet, PathAccountConfig, query, &response)
	if err == nil {
		return response, nil
	}
	return nil, err
}
func (api *OkxRestApi) SetPositionMode(req models.SetPositionModeReq) (models.GetSetPositionModeData, error) {
	resp := models.GetSetPositionModeData{}
	err := apiPrivateTemplateWithBody(api, resty.MethodPost, PathSetPositionMode, req, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (api *OkxRestApi) SetLeverage(req models.SetLeverageReq) (models.GetSetLeverageData, error) {
	resp := models.GetSetLeverageData{}
	err := apiPrivateTemplateWithBody(api, resty.MethodPost, PathSetLeverage, req, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
