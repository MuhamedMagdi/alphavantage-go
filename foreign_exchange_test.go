//go:build integration
// +build integration

package alphaVantage

import (
	"os"
	"testing"

	"fmt"
	"github.com/stretchr/testify/assert"
)

func TestGetCurrencyExchangeRate(t *testing.T) {
	c := NewClient(os.Getenv("ALPHA_Vantage_API_KEY"))
	options := &CurrencyExchangeRateOptions{
		FromCurrency: "BTC",
		ToCurrency:   "USD",
	}
	res, _ := c.GetCurrencyExchangeRate(options)
	assert.Equal(t, res.RealtimeCurrencyExchangeRate.FromCurrencyCode, "BTC")
}

func TestGetCryptoIntraDay(t *testing.T) {
	c := NewClient(os.Getenv("ALPHA_Vantage_API_KEY"))
	options := &CryptoIntrDayOptions{
		Symbol:     "ETH",
		Market:     "USD",
		Interval:   "5min",
		OutputSize: "full",
	}
	res, err := c.GetCryptoIntraDay(options)
	fmt.Println(err)
	assert.Equal(t, res.MetaData.DigitalCurrencyCode, "ETH")
}
