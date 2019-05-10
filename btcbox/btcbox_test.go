package btcbox

import (
	"github.com/stretchr/testify/assert"
	"github.com/wudian/GoEx"
	"net/http"
	"testing"
)

var btcbox = New(http.DefaultClient, "", "")

func TestBtcBox_GetTicker(t *testing.T) {
	ticker, err := btcbox.GetTicker(goex.BTC_JPY)
	assert.Nil(t, err)
	t.Log(ticker)
}

func TestBtcBox_GetDepth(t *testing.T) {
	dep, err := btcbox.GetDepth(5, goex.ETH_JPY)
	assert.Nil(t, err)
	t.Log(dep.BidList)
}

func TestBtcBox_CancelOrder(t *testing.T) {

}
