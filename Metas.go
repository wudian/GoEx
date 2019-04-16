package goex

import (
	"net/http"
)

type Order struct {
	Price,
	Amount,
	AvgPrice,
	DealAmount,
	Fee float64
	OrderID2  string
	OrderID   int
	OrderTime int
	Status    TradeStatus
	Currency  CurrencyPair
	Side      TradeSide
}

type Trade struct {
	Tid    int64        `json:"tid"`
	Type   TradeSide    `json:"type"`
	Amount float64      `json:"amount,string"`
	Price  float64      `json:"price,string"`
	Date   int64        `json:"date_ms"`
	Pair   CurrencyPair `json:"omitempty"`
}

type SubAccount struct {
	Currency Currency
	Amount,
	ForzenAmount,
	LoanAmount float64
}

type Account struct {
	Exchange    string
	Asset       float64 //总资产
	NetAsset    float64 //净资产
	SubAccounts map[Currency]SubAccount
}

type Ticker struct {
	Pair CurrencyPair `json:"omitempty"`
	Last float64      `json:"last,string"`
	Buy  float64      `json:"buy,string"`
	Sell float64      `json:"sell,string"`
	High float64      `json:"high,string"`
	Low  float64      `json:"low,string"`
	Vol  float64      `json:"vol,string"`
	Date int64       `json:"date"` // 单位:秒(second)
}

func NewTicker() *Ticker  {
	var tk Ticker
	tk.Last = 0
	tk.Buy = 0
	tk.Sell = 0
	tk.High =0
	tk.Low = 0
	tk.Vol=0
	return &tk
}

func (tk *Ticker)Add(ticker *Ticker)  {
	tk.Last += ticker.Last
	tk.Buy += ticker.Buy
	tk.Sell += ticker.Sell
	tk.High += ticker.High
	tk.Low += ticker.Low
	tk.Vol += ticker.Vol
}
//division
func (tk *Ticker)Div(n float64) *Ticker {
	tk.Last /= n
	tk.Buy /= n
	tk.Sell /= n
	tk.High /= n
	tk.Low /= n
	tk.Vol /= n
	return tk
}

//multiplication
func (tk *Ticker)Multi(n float64) *Ticker {
	tk.Last *= n
	tk.Buy *= n
	tk.Sell *= n
	tk.High *= n
	tk.Low *= n
	tk.Vol *= n
	return tk
}

func (tk *Ticker)Decimal() *Ticker {
	tk.Last = Decimal(tk.Last)
	tk.Buy = Decimal(tk.Buy)
	tk.Sell = Decimal(tk.Sell)
	tk.High = Decimal(tk.High)
	tk.Low = Decimal(tk.Low)
	tk.Vol = Decimal(tk.Vol)
	return tk
}

type FutureTicker struct {
	*Ticker
	ContractType string  `json:"omitempty"`
	ContractId   int     `json:"contractId"`
	LimitHigh    float64 `json:"limitHigh,string"`
	LimitLow     float64 `json:"limitLow,string"`
	HoldAmount   float64 `json:"hold_amount,string"`
	UnitAmount   float64 `json:"unitAmount,string"`
}

type DepthRecord struct {
	Price,
	Amount float64
}

type DepthRecords []DepthRecord

func (dr DepthRecords) Len() int {
	return len(dr)
}

func (dr DepthRecords) Swap(i, j int) {
	dr[i], dr[j] = dr[j], dr[i]
}

func (dr DepthRecords) Less(i, j int) bool {
	return dr[i].Price < dr[j].Price
}

type Depth struct {
	ContractType string //for future
	Pair         CurrencyPair
	//UTime        time.Time
	Date int64       // 单位:秒(second)
	AskList,
	BidList DepthRecords
}

func NewDepth() *Depth {
	var depth Depth
	for i:=0; i<10; i++ {
		var dr DepthRecord
		dr.Price = ToFloat64(0)
		dr.Amount = ToFloat64(0)
		depth.AskList = append(depth.AskList, dr)
		depth.BidList = append(depth.BidList, dr)
	}
	return &depth
}

type APIConfig struct {
	HttpClient    *http.Client
	Endpoint      string
	ApiKey        string
	ApiSecretKey  string
	ApiPassphrase string //for okex.com v3 api
	ClientId      string //for bitstamp.net , huobi.pro

	Lever int //杠杆倍数 , for future
}

type Kline struct {
	Pair      CurrencyPair
	Timestamp int64
	Open,
	Close,
	High,
	Low,
	Vol float64
}

type FutureKline struct {
	*Kline
	Vol2 float64 //个数
}

type FutureSubAccount struct {
	Currency      Currency
	AccountRights float64 //账户权益
	KeepDeposit   float64 //保证金
	ProfitReal    float64 //已实现盈亏
	ProfitUnreal  float64
	RiskRate      float64 //保证金率
}

type FutureAccount struct {
	FutureSubAccounts map[Currency]FutureSubAccount
}

type FutureOrder struct {
	OrderID2     string //请尽量用这个字段替代OrderID字段
	Price        float64
	Amount       float64
	AvgPrice     float64
	DealAmount   float64
	OrderID      int64
	OrderTime    int64
	Status       TradeStatus
	Currency     CurrencyPair
	OType        int     //1：开多 2：开空 3：平多 4： 平空
	LeverRate    int     //倍数
	Fee          float64 //手续费
	ContractName string
}

type FuturePosition struct {
	BuyAmount      float64
	BuyAvailable   float64
	BuyPriceAvg    float64
	BuyPriceCost   float64
	BuyProfitReal  float64
	CreateDate     int64
	LeverRate      int
	SellAmount     float64
	SellAvailable  float64
	SellPriceAvg   float64
	SellPriceCost  float64
	SellProfitReal float64
	Symbol         CurrencyPair //btc_usd:比特币,ltc_usd:莱特币
	ContractType   string
	ContractId     int64
	ForceLiquPrice float64 //预估爆仓价
}
