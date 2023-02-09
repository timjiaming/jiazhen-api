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
		client.SetBaseURL(HTTPBaseURL)
	} else {
		client.SetBaseURL(endpoint)
	}
	api := OkxRestApi{
		Account: acc,
		client:  client,
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
	}

	if body != "" {
		request.SetBody(body)
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
	if responseObj.Code == 0 {
		return responseObj.Data, nil
	}
	return nil, fmt.Errorf("invalid response code %d, %s", responseObj.Code, responseObj.Msg)
}
