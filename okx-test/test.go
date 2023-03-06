package main

import (
	"errors"
	"fmt"
	"github.com/timjiaming/jiazhen-api/okx"
	"github.com/timjiaming/jiazhen-api/okx/models"
	"gopkg.in/yaml.v3"
	"os"
	"strconv"
	"time"
)

type Config struct {
	User     user     `json:"user"`
	Spot     spot     `yaml:"spot"`
	CoinSwap coinSwap `yaml:"coinSwap"`
}
type user struct {
	Apikey        string
	Sercetkey     string
	Passphrase    string
	Interval      int64
	Total         int64
	Intervalprice int64
}

type spot struct {
	Symbol string
	Sz     string
	Side   string
}

type coinSwap struct {
	Symbol string
	Sz     string
	Side   string
}

var config Config

func readYaml() (Config, error) {
	dataBttes, err := os.ReadFile("/Users/tim/GolandProjects/jiazhen-api/okx-test/user.yaml")
	if err != nil {
		fmt.Println("read yaml err", err)
		return config, err
	}
	err = yaml.Unmarshal(dataBttes, &config)
	if err != nil {
		fmt.Println("yaml unmarshal err", err)
		return config, err
	}
	return config, nil

}
func main() {
	var total int64
	var err error
	config, err = readYaml()
	if err != nil {
		return
	}
	total = 0
	for {
		if total >= config.User.Total {
			break
		} else {
			err = placeSpot()
			if err != nil {
				return
			}
			err = placeSwap()
			if err != nil {
				return
			}
			total = total + config.User.Intervalprice
			fmt.Println("total:", total)
			time.Sleep(time.Duration(config.User.Interval) * time.Second)
		}
	}
}
func placeSpot() error {
	side := ""
	size := ""
	acc := okx.OkxAccount{
		ApiKey:     config.User.Apikey,
		ApiSecret:  config.User.Sercetkey,
		Passphrase: config.User.Passphrase,
		BaseURL:    "",
		Testnet:    false,
	}
	api := okx.NewOkxRestApi(&acc, "")
	if config.Spot.Symbol == "" || config.Spot.Sz == "" {
		return errors.New("symbol or sz empty")
	}
	if config.Spot.Side == okx.Buy {
		side = okx.Buy
		size = config.Spot.Sz //USDT数量
	} else if config.Spot.Side == okx.Sell {
		side = okx.Sell
		price, err := GetPrice()
		if err != nil {
			return err
		}
		sz, _ := strconv.ParseFloat(config.Spot.Sz, 64)
		size = strconv.FormatFloat(sz/price, 'f', 3, 64) //币的数量
		//size = 价格/数量
	}
	if side == "" {
		return errors.New("side is err")
	}
	req := models.PlaceOrderReq{
		InstId:  config.Spot.Symbol,
		TdMode:  okx.Cross,
		Side:    side,
		OrdType: okx.Market,
		Sz:      size,
	}
	message, err := api.PlaceOrder(req)
	if err != nil {
		fmt.Println("place spot err")
		fmt.Println(err)
		return err
	}
	fmt.Println(message)
	return nil
}
func placeSwap() error {
	side := ""
	size := config.CoinSwap.Sz
	acc := okx.OkxAccount{
		ApiKey:     config.User.Apikey,
		ApiSecret:  config.User.Sercetkey,
		Passphrase: config.User.Passphrase,
		BaseURL:    "",
		Testnet:    false,
	}
	api := okx.NewOkxRestApi(&acc, "")
	if config.CoinSwap.Side == okx.Buy {
		side = okx.Buy
	} else if config.CoinSwap.Side == okx.Sell {
		side = okx.Sell
	}
	if side == "" {
		return errors.New("side is err")
	}
	if config.CoinSwap.Symbol == "" || config.CoinSwap.Sz == "" {
		return errors.New("symbol or sz empty")
	}
	req := models.PlaceOrderReq{
		InstId:  config.CoinSwap.Symbol,
		TdMode:  okx.Cross,
		Side:    side,
		OrdType: okx.Market,
		Sz:      size,
	}
	message, err := api.PlaceOrder(req)
	if err != nil {
		fmt.Println("place coinswap err")
		fmt.Println(err)
		return err
	}
	fmt.Println(message)
	return nil
}
func GetPrice() (float64, error) {
	acc := okx.OkxAccount{
		Testnet: false,
	}
	api := okx.NewOkxRestApi(&acc, "")
	req := models.GetBookLiteReq{InstId: config.Spot.Symbol}
	message, err := api.GetBooksLite(req)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	fmt.Println(message[0].Asks[0][0])
	price, err := strconv.ParseFloat(message[0].Asks[0][0], 64)
	if err != nil {
		return 0, err
	}
	return price, nil
}
