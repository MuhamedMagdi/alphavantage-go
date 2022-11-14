//go:build integration
// +build integration

package alphaVantage

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
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
