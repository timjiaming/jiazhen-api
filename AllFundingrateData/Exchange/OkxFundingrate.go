package Exchange

import (
	"github.com/timjiaming/jiazhen-api/AllFundingrateData/ORM"
	"github.com/timjiaming/jiazhen-api/okx"
	"github.com/timjiaming/jiazhen-api/okx/models"
	"strconv"
	"time"
)

func GetFundingrateHistory(instId, before, after, limit string) ([]ORM.FundingrateData, error) {
	acc := okx.OkxAccount{
		Testnet: false,
	}
	api := okx.NewOkxRestApi(&acc, "")
	req := models.GetFundingrateHistoryReq{
		InstId: instId,
		Before: before,
		After:  after,
		Limit:  limit,
	}
	message, err := api.GetFundingrateHistory(req)
	if err != nil {
		return nil, err
	}
	var dataList []ORM.FundingrateData
	for _, v := range message {
		f, _ := strconv.ParseFloat(v.FundingRate, 64)
		t, _ := strconv.ParseInt(v.FundingTime, 10, 64)
		tm := time.Unix(0, t*int64(time.Millisecond))
		//utc8time := tm.Format("2006-01-02 15:04:05")
		funding := ORM.FundingrateData{
			//Id:          instId + v.FundingTime,
			Symbol:      instId,
			Fundingrate: f,
			T:           v.FundingTime,
			Exchange:    "OKX",
			UTC8TIME:    tm,
		}
		dataList = append(dataList, funding)
	}
	return dataList, nil
}

func GetInstruments() ([]string, error) {
	var symbolList []string
	acc := okx.OkxAccount{
		Testnet: false,
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
		return nil, err
	}
	for _, v := range message {
		symbolList = append(symbolList, v.InstId)
	}
	return symbolList, nil
}
