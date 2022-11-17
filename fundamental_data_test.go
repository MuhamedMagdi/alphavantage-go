//go:build integration
// +build integration

package alphaVantage

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCompanyOverview(t *testing.T) {
	c := NewClient(os.Getenv("ALPHA_Vantage_API_KEY"))
	options := &CompanyOverviewOptions{
		Symbol: "IBM",
	}
	res, _ := c.GetCompanyOverview(options)
	assert.Equal(t, res.Symbol, "IBM")
}

func TestGetIncomeStatement(t *testing.T) {
	c := NewClient(os.Getenv("ALPHA_Vantage_API_KEY"))
	options := &IncomeStatementOptions{
		Symbol: "IBM",
	}
	res, _ := c.GetIncomeStatement(options)
	assert.Equal(t, res.Symbol, "IBM")
}
