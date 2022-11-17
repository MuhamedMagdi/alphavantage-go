package alphaVantage

import "fmt"

func (c *Client) GetCompanyOverview(options *CompanyOverviewOptions) (*CompanyOverviewList, error) {
	const function = "OVERVIEW"
	var symbol string
	if options != nil {
		symbol = options.Symbol
	}

	res := CompanyOverviewList{}
	query := fmt.Sprintf("%s/query?function=%s&symbol=%s&apikey=%s", c.baseURL, function, symbol, c.apiKey)

	if err := c.processRequest(query, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *Client) GetIncomeStatement(options *IncomeStatementOptions) (*IncomeStatementList, error) {
	const function = "INCOME_STATEMENT"
	var symbol string
	if options != nil {
		symbol = options.Symbol
	}

	res := IncomeStatementList{}
	query := fmt.Sprintf("%s/query?function=%s&symbol=%s&apikey=%s", c.baseURL, function, symbol, c.apiKey)

	if err := c.processRequest(query, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *Client) GetBalanceSheet(options *BalanceSheetOptions) (*BalanceSheetList, error) {
	const function = "BALANCE_SHEET"
	var symbol string
	if options != nil {
		symbol = options.Symbol
	}

	res := BalanceSheetList{}
	query := fmt.Sprintf("%s/query?function=%s&symbol=%s&apikey=%s", c.baseURL, function, symbol, c.apiKey)

	if err := c.processRequest(query, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *Client) GetCashFlow(options *CashFlowOptions) (*CashFlowList, error) {
	const function = "CASH_FLOW"
	var symbol string
	if options != nil {
		symbol = options.Symbol
	}

	res := CashFlowList{}
	query := fmt.Sprintf("%s/query?function=%s&symbol=%s&apikey=%s", c.baseURL, function, symbol, c.apiKey)

	if err := c.processRequest(query, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *Client) GetEarnings(options *EarningsOptions) (*EarningsList, error) {
	const function = "EARNINGS"
	var symbol string
	if options != nil {
		symbol = options.Symbol
	}

	res := EarningsList{}
	query := fmt.Sprintf("%s/query?function=%s&symbol=%s&apikey=%s", c.baseURL, function, symbol, c.apiKey)

	if err := c.processRequest(query, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
