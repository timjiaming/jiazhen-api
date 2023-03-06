package models

import "encoding/json"

type BaseTextMsg struct {
	Arg    MsgArg          `json:"arg"`
	Data   json.RawMessage `json:"data"`
	Event  string          `json:"event"`
	Action string          `json:"action"`
	Code   string          `json:"code"`
	ID     string          `json:"id"`
	Op     string          `json:"op"`
	Msg    string          `json:"msg"`
}
type MsgArg struct {
	Channel string `json:"channel"`
	InstID  string `json:"instId"`
	UID     string `json:"uid"`
}
type WsSubscribeReq struct {
	Op   string      `json:"op"`
	Args interface{} `json:"args"`
}
type Ticker struct {
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
