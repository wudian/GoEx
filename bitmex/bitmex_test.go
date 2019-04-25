package bitmex

import (
	"github.com/wudian/GoEx"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"net/url"
	"net"
	"time"
)

var httpProxyClient = &http.Client{
	Transport: &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return &url.URL{
				Scheme: "socks5",
				Host:   "127.0.0.1:55307"}, nil
		},
		Dial: (&net.Dialer{
			Timeout: 10 * time.Second,
		}).Dial,
	},
	Timeout: 10 * time.Second,
}
var mex = New(httpProxyClient, "", "")

func TestBitmex_GetTicker(t *testing.T) {
	//return
	ticker, err := mex.GetTicker(goex.BTC_USDT)
	assert.Nil(t, err)
	t.Log(ticker)
}

func TestBitmex_GetDepth(t *testing.T) {
	return
	dep, err := mex.GetDepth(2, goex.NewCurrencyPair(goex.XBT, goex.USD))
	assert.Nil(t, err)
	t.Log(dep)
}
