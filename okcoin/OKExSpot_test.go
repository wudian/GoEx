package okcoin

import (
	"github.com/wudian/GoEx"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var okexSpot = NewOKExSpot(http.DefaultClient, "06aa16ef-ec6e-425b-bbdb-145bee989842", "EF674EDE16BE81D7F2A2F83330BEDE36")

func TestOKExSpot_GetTicker(t *testing.T) {
	ticker, err := okexSpot.GetTicker(goex.ETC_BTC)
	assert.Nil(t, err)
	t.Log(ticker)
}

func TestOKExSpot_GetDepth(t *testing.T) {
	dep, err := okexSpot.GetDepth(2, goex.ETC_BTC)
	assert.Nil(t, err)
	t.Log(dep)
}

func TestOKExSpot_GetKlineRecords(t *testing.T) {
	klines, err := okexSpot.GetKlineRecords(goex.LTC_BTC, goex.KLINE_PERIOD_1MIN, 1000, -1)
	t.Log(err, klines)
}
