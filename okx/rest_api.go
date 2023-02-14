package okx

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"jiazhen-api/common"
	"jiazhen-api/okx/models"
	"net/url"
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
func apiPublicTemplateWithQuery(api *OkxRestApi, method string, path string, query map[string]string, response any) error {
	data, err := api.SendPublicRequest(method, path, query, "")
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &response)
}

// 获取所有产品行情信息
func (api *OkxRestApi) GetTickers(req models.GetTickersReq) (*models.GetTickersData, error) {
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
		return &response, err
	}
	return nil, err
}

// 获取单个产品行情信息
func (api *OkxRestApi) GetTicker(req models.GetTickerReq) (*models.GetTickerData, error) {
	query := make(map[string]string)
	if req.InstId != "" {
		query["instId"] = req.InstId
	}
	response := models.GetTickerData{}
	err := apiPublicTemplateWithQuery(api, resty.MethodGet, PathTicker, query, &response)
	if err == nil {
		return &response, err
	}
	return nil, err
}

// 获取指数行情
func (api *OkxRestApi) GetIndexTickers(req models.GetIndexTickersReq) (*models.GetIndexTickersData, error) {
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
		return &response, err
	}
	return nil, err
}

// 获取产品深度
func (api *OkxRestApi) GetBooks(req models.GetBooksReq) (*models.GetBooksData, error) {
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
		return &response, err
	}
	return nil, err
}

// 获取产品轻量深度
func (api *OkxRestApi) GetBooksLite(req models.GetBookLiteReq) (*models.GetBookLiteData, error) {
	query := make(map[string]string)
	if req.InstId != "" {
		query["instId"] = req.InstId
	}
	response := models.GetBookLiteData{}
	err := apiPublicTemplateWithQuery(api, resty.MethodGet, PathBookLite, query, &response)
	if err == nil {
		return &response, err
	}
	return nil, err
}

// 获取交易产品K线数据
func (api *OkxRestApi) GetCandles(req models.GetCandlesReq) (*models.GetCandlesData, error) {
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
		return &response, err
	}
	return nil, err
}
func (api *OkxRestApi) GetHistoryCandles(req models.GetHistoryCandlesReq) (*models.GetHistoryCandlesData, error) {
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
		return &response, err
	}
	return nil, err
}
func (api *OkxRestApi) GetMarketPriceCandles(req models.GetMarketPriceCandlesReq) (*models.GetMarketPriceCandelsData, error) {
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
		return &response, nil
	}
	return nil, err
}
func (api *OkxRestApi) GetHistoryMarketPriceCandles(req models.GetHistoryMarketPriceReq) (*models.GetHistoryMarketPriceCandlesData, error) {
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
		return &response, nil
	}
	return nil, err
}
func (api *OkxRestApi) GetTrades(req models.GetTradesReq) (*models.GetTradesData, error) {
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
		return &response, nil
	}
	return nil, err
}

func (api *OkxRestApi) GetHistoryTrades(req models.GetHistoryTradesReq) (*models.GetHistoryData, error) {
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
		return &response, nil
	}
	return nil, err
}
func (api *OkxRestApi) GetPlatform24Volume() (*models.GetPlatform24VolumeData, error) {
	query := make(map[string]string)
	response := models.GetPlatform24VolumeData{}
	err := apiPublicTemplateWithQuery(api, resty.MethodGet, PathPlatform24Volume, query, &response)
	if err == nil {
		return &response, nil
	}
	return nil, err
}
func (api *OkxRestApi) GetInstruments(req models.GetInstrumentsReq) (*models.GetInstrumentsData, error) {
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
		return &response, nil
	}
	return nil, err
}

// 获取交割和行权记录
func (api *OkxRestApi) GetDeliveryExerciseHistory(req models.GetDeliveryExerciseHistoryReq) (*models.GetDeliveryExerciseHistoryData, error) {
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
		return &response, nil
	}
	return nil, err
}

// 获取持仓总量
func (api *OkxRestApi) GetOpenInterest(req models.GetOpenInterestReq) (*models.GetOpenInterestData, error) {
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
		return &response, nil
	}
	return nil, err
}

// 获取永续合约当前资金费率
func (api *OkxRestApi) GetFundingrate(req models.GetFundingrateReq) (*models.GetFundingrateData, error) {
	query := make(map[string]string)
	if req.InstId != "" {
		query["instId"] = req.InstId
	}
	response := models.GetFundingrateData{}
	err := apiPublicTemplateWithQuery(api, resty.MethodGet, PathFundingrate, query, &response)
	if err == nil {
		return &response, nil
	}
	return nil, err
}

// 获取永续合约历史资金费率
func (api *OkxRestApi) GetFundingrateHistory(req models.GetFundingrateHistoryReq) (*models.GetFundingrateHistoryData, error) {
	query := make(map[string]string)
	if req.InstId != "" {
		query["instId"] = req.InstId
	}
	if req.Before != "" {
		query["before"] = req.InstId
	}
	if req.After != "" {
		query["after"] = req.InstId
	}
	if req.Limit != "" {
		query["limit"] = req.InstId
	}
	response := models.GetFundingrateHistoryData{}
	err := apiPublicTemplateWithQuery(api, resty.MethodGet, PathFundingrateHistory, query, &response)
	if err == nil {
		return &response, nil
	}
	return nil, err
}

// 获取限价
func (api *OkxRestApi) GetPriceLimit(req models.GetPriceLimitReq) (*models.GetPriceLimitData, error) {
	query := make(map[string]string)
	if req.InstId != "" {
		query["instId"] = req.InstId
	}
	response := models.GetPriceLimitData{}
	err := apiPublicTemplateWithQuery(api, resty.MethodGet, PathPriceLimit, query, &response)
	if err == nil {
		return &response, nil
	}
	return nil, err
}
func (api *OkxRestApi) GetEstimatedPrice(req models.GetEstimatedPriceReq) (*models.GetEstimatedPriceData, error) {
	query := make(map[string]string)
	if req.InstId != "" {
		query["instId"] = req.InstId
	}
	response := models.GetEstimatedPriceData{}
	err := apiPublicTemplateWithQuery(api, resty.MethodGet, PathPriceLimit, query, &response)
	if err == nil {
		return &response, nil
	}
	return nil, err
}
