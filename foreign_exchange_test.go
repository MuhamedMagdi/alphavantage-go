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
