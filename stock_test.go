//go:build integration
// +build integration

package alphaVantage

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetIntraDay(t *testing.T) {
	c := NewClient(os.Getenv("ALPHA_Vantage_API_KEY"))
	options := &IntraDayOptions{
		Symbol:   "IBM",
		Interval: "5min",
	}
	res, _ := c.GetIntraDay(options)
	assert.Equal(t, res.MetaData.Symbol, "IBM")
}

func TestGetDaily(t *testing.T) {
	c := NewClient(os.Getenv("ALPHA_Vantage_API_KEY"))
	options := &DailyOptions{
		Symbol:     "IBM",
		OutputSize: "full",
	}
	res, _ := c.GetDaily(options)
	assert.Equal(t, res.MetaData.Symbol, "IBM")
}

func TestGetDailyAdjusted(t *testing.T) {
	c := NewClient(os.Getenv("ALPHA_Vantage_API_KEY"))
	options := &DailyAdjustedOptions{
		Symbol: "IBM",
	}
	res, _ := c.GetDailyAdjusted(options)
	assert.Equal(t, res.MetaData.Symbol, "IBM")
}

func TestGetWeekly(t *testing.T) {
	c := NewClient(os.Getenv("ALPHA_Vantage_API_KEY"))
	options := &WeeklyOptions{
		Symbol: "IBM",
	}
	res, _ := c.GetWeekly(options)
	assert.Equal(t, res.MetaData.Symbol, "IBM")
}

func TestGetWeeklyAdjusted(t *testing.T) {
	c := NewClient(os.Getenv("ALPHA_Vantage_API_KEY"))
	options := &WeeklyAdjustedOptions{
		Symbol: "IBM",
	}
	res, _ := c.GetWeeklyAdjusted(options)
	assert.Equal(t, res.MetaData.Symbol, "IBM")
}

func TestGetMonthly(t *testing.T) {
	c := NewClient(os.Getenv("ALPHA_Vantage_API_KEY"))
	options := &MonthlyOptions{
		Symbol: "IBM",
	}
	res, _ := c.GetMonthly(options)
	assert.Equal(t, res.MetaData.Symbol, "IBM")
}

func TestGetMonthlyAdjusted(t *testing.T) {
	c := NewClient(os.Getenv("ALPHA_Vantage_API_KEY"))
	options := &MonthlyAdjustedOptions{
		Symbol: "IBM",
	}
	res, _ := c.GetMonthlyAdjusted(options)
	assert.Equal(t, res.MetaData.Symbol, "IBM")
}

func TestGetQuote(t *testing.T) {
	c := NewClient(os.Getenv("ALPHA_Vantage_API_KEY"))
	options := &QuoteOptions{
		Symbol: "IBM",
	}
	res, _ := c.GetQuote(options)
	assert.Equal(t, res.GlobalQuote.Symbol, "IBM")
}

func TestSearch(t *testing.T) {
	c := NewClient(os.Getenv("ALPHA_Vantage_API_KEY"))
	options := &SearchOptions{
		Keywords: "IBM",
	}
	res, _ := c.Search(options)
	assert.Equal(t, res.BestMatches[0].Symbol, "IBM")
}
