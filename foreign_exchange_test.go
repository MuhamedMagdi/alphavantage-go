//go:build integration
// +build integration

package alphaVantage

import (
	"os"
	"testing"

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
	res, _ := c.GetCryptoIntraDay(options)
	assert.Equal(t, res.MetaData.DigitalCurrencyCode, "ETH")
}

func TestGetCryptoDaily(t *testing.T) {
	c := NewClient(os.Getenv("ALPHA_Vantage_API_KEY"))
	options := &CryptoDailyOptions{
		Symbol: "ETH",
		Market: "USD",
	}
	res, _ := c.GetCryptoDaily(options)
	assert.Equal(t, res.MetaData.DigitalCurrencyCode, "ETH")
}
