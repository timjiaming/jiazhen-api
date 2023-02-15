package Exchange

import (
	"fmt"
	"jiazhen-api/okx"
	"jiazhen-api/okx/models"
)

func GetFundingrateHistory(instId, before, after, limit string) {
	acc := okx.OkxAccount{
		Testnet: true,
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
		return
	}
	for k, v := range message {
		fmt.Println(k, v)
	}
}
