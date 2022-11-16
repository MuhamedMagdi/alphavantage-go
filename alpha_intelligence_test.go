//go:build integration
// +build integration

package alphaVantage

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNewsAndSentiment(t *testing.T) {
	c := NewClient(os.Getenv("ALPHA_Vantage_API_KEY"))
	options := &NewsAndSentimentOptions{
		Tickers: "IBM",
		Topics:  "technology",
		Sort:    "LATEST",
		Limit:   50,
	}
	res, _ := c.GetNewsAndSentiment(options)
	assert.Equal(t, res.Items, "50")
}

func TestGetWinningPortfolios(t *testing.T) {
	c := NewClient(os.Getenv("ALPHA_Vantage_API_KEY"))
	options := &WinningPortfoliosOptions{
		Season: "2022-10",
	}
	res, _ := c.GetWinningPortfolios(options)
	assert.Equal(t, res.Season, "2022-10")
}
