package huobi

import (
	"github.com/wudian/GoEx"
	"github.com/stretchr/testify/assert"
	"net"
	"net/http"
	"net/url"
	"testing"
	"time"
)

var httpProxyClient = &http.Client{
	Transport: &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return &url.URL{
				Scheme: "socks5",
				Host:   "172.16.226.24:8998"}, nil
		},
		Dial: (&net.Dialer{
			Timeout: 10 * time.Second,
		}).Dial,
	},
	Timeout: 10 * time.Second,
}

var (
	apikey    = "478c732a-4f0cd03e-e779076d-0a87b"
	secretkey = "e1b9b5c7-5a9bff84-e72dc5eb-f5dd9"
)

//
var hbpro = NewHuoBiProSpot(http.DefaultClient, apikey, secretkey)

func TestHuobiPro_GetTicker(t *testing.T) {
	//return
	ticker, err := hbpro.GetTicker(goex.BTC_USDT)
	assert.Nil(t, err)
	t.Log(ticker)
}

func TestHuobiPro_GetDepth(t *testing.T) {
	dep, err := hbpro.GetDepth(10, goex.ETH_USDT)
	assert.Nil(t, err)
	t.Log(dep)
	//t.Log(dep.BidList)
}

func TestHuobiPro_GetAccountInfo(t *testing.T) {
	return
	info, err := hbpro.GetAccountInfo("point")
	assert.Nil(t, err)
	t.Log(info)
}

//获取点卡剩余
func TestHuoBiPro_GetPoint(t *testing.T) {
	return
	point := NewHuoBiProPoint(httpProxyClient, apikey, secretkey)
	acc, _ := point.GetAccount()
	t.Log(acc.SubAccounts[HBPOINT])
}

//获取现货资产信息
func TestHuobiPro_GetAccount(t *testing.T) {
	return
	acc, err := hbpro.GetAccount()
	assert.Nil(t, err)
	t.Log(acc.SubAccounts)
}

func TestHuobiPro_LimitBuy(t *testing.T) {
	return
	ord, err := hbpro.LimitBuy("", "0.09122", goex.BCC_BTC)
	assert.Nil(t, err)
	t.Log(ord)
}

func TestHuobiPro_LimitSell(t *testing.T) {
	return
	ord, err := hbpro.LimitSell("1", "0.212", goex.BCC_BTC)
	assert.Nil(t, err)
	t.Log(ord)
}

func TestHuobiPro_MarketSell(t *testing.T) {
	return
	ord, err := hbpro.MarketSell("0.1738", "0.212", goex.BCC_BTC)
	assert.Nil(t, err)
	t.Log(ord)
}

func TestHuobiPro_MarketBuy(t *testing.T) {
	return
	ord, err := hbpro.MarketBuy("0.02", "", goex.BCC_BTC)
	assert.Nil(t, err)
	t.Log(ord)
}

func TestHuobiPro_GetUnfinishOrders(t *testing.T) {
	return
	ords, err := hbpro.GetUnfinishOrders(goex.ETC_USDT)
	assert.Nil(t, err)
	t.Log(ords)
}

func TestHuobiPro_CancelOrder(t *testing.T) {
	return
	r, err := hbpro.CancelOrder("600329873", goex.ETH_USDT)
	assert.Nil(t, err)
	t.Log(r)
	t.Log(err)
}

func TestHuobiPro_GetOneOrder(t *testing.T) {
	return
	ord, err := hbpro.GetOneOrder("1116237737", goex.LTC_BTC)
	assert.Nil(t, err)
	t.Log(ord)
}

func TestHuobiPro_GetOrderHistorys(t *testing.T) {
	return
	ords, err := hbpro.GetOrderHistorys(goex.NewCurrencyPair2("HT_USDT"), 1, 3)
	t.Log(err)
	t.Log(ords)
}

//func TestHuobiPro_GetDepthWithWs(t *testing.T) {
//	return
//	hbpro.GetDepthWithWs(goex.BTC_USDT, func(dep *goex.Depth) {
//		log.Println("%+v", *dep)
//	})
//	time.Sleep(time.Minute)
//}
//
//func TestHuobiPro_GetTickerWithWs(t *testing.T) {
//	return
//	hbpro.GetTickerWithWs(goex.BTC_USDT, func(ticker *goex.Ticker) {
//		log.Println("%+v", *ticker)
//	})
//	time.Sleep(time.Minute)
//}
//
//func TestHuobiPro_GetKLineWithWs(t *testing.T) {
//	return
//	hbpro.GetKLineWithWs(goex.BTC_USDT, goex.KLINE_PERIOD_60MIN, func(kline *goex.Kline) {
//		log.Println("%+v", *kline)
//	})
//	time.Sleep(time.Minute)
//}

func TestHuobiPro_GetCurrenciesList(t *testing.T) {
	return
	hbpro.GetCurrenciesList()
}

func TestHuobiPro_GetCurrenciesPrecision(t *testing.T) {
	return
	t.Log(hbpro.GetCurrenciesPrecision())
}
