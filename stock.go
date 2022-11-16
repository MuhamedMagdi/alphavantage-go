package alphaVantage

import (
	"fmt"
)

func (c *Client) GetIntraDay(options *IntraDayOptions) (*IntraDayList, error) {
	const function = "TIME_SERIES_INTRADAY"
	var symbol, interval string
	adjusted := "true"
	outputSize := "compact"

	if options != nil {
		symbol = options.Symbol
		interval = options.Interval
		if options.Adjusted != "" {
			adjusted = options.Adjusted
		}
		if options.OutputSize != "" {
			outputSize = options.OutputSize
		}
	}

	res := IntraDayList{}
	query := fmt.Sprintf("%s/query?function=%v&symbol=%v&interval=%v&adjusted=%v&outputsize=%v&apikey=%v", c.baseURL, function, symbol, interval, adjusted, outputSize, c.apiKey)

	if err := c.processRequest(query, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *Client) GetDaily(options *DailyOptions) (*DailyList, error) {
	const function = "TIME_SERIES_DAILY"
	var symbol string
	outputSize := "compact"

	if options != nil {
		symbol = options.Symbol
		if options.OutputSize != "" {
			outputSize = options.OutputSize
		}
	}

	res := DailyList{}
	query := fmt.Sprintf("%s/query?function=%v&symbol=%v&outputsize=%v&apikey=%v", c.baseURL, function, symbol, outputSize, c.apiKey)

	if err := c.processRequest(query, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *Client) GetDailyAdjusted(options *DailyAdjustedOptions) (*DailyAdjustedList, error) {
	const function = "TIME_SERIES_DAILY_ADJUSTED"
	var symbol string
	outputSize := "compact"

	if options != nil {
		symbol = options.Symbol
		if options.OutputSize != "" {
			outputSize = options.OutputSize
		}
	}

	res := DailyAdjustedList{}
	query := fmt.Sprintf("%s/query?function=%v&symbol=%v&outputsize=%v&apikey=%v", c.baseURL, function, symbol, outputSize, c.apiKey)

	if err := c.processRequest(query, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *Client) GetWeekly(options *WeeklyOptions) (*WeeklyList, error) {
	const function = "TIME_SERIES_WEEKLY"
	var symbol string

	if options != nil {
		symbol = options.Symbol
	}

	res := WeeklyList{}
	query := fmt.Sprintf("%s/query?function=%v&symbol=%v&apikey=%v", c.baseURL, function, symbol, c.apiKey)

	if err := c.processRequest(query, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *Client) GetWeeklyAdjusted(options *WeeklyAdjustedOptions) (*WeeklyAdjustedList, error) {
	const function = "TIME_SERIES_WEEKLY_ADJUSTED"
	var symbol string

	if options != nil {
		symbol = options.Symbol
	}

	res := WeeklyAdjustedList{}
	query := fmt.Sprintf("%s/query?function=%v&symbol=%v&apikey=%v", c.baseURL, function, symbol, c.apiKey)

	if err := c.processRequest(query, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *Client) GetMonthly(options *MonthlyOptions) (*MonthlyList, error) {
	const function = "TIME_SERIES_MONTHLY"
	var symbol string

	if options != nil {
		symbol = options.Symbol
	}

	res := MonthlyList{}
	query := fmt.Sprintf("%s/query?function=%v&symbol=%v&apikey=%v", c.baseURL, function, symbol, c.apiKey)

	if err := c.processRequest(query, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *Client) GetMonthlyAdjusted(options *MonthlyAdjustedOptions) (*MonthlyAdjustedList, error) {
	const function = "TIME_SERIES_MONTHLY_ADJUSTED"
	var symbol string

	if options != nil {
		symbol = options.Symbol
	}

	res := MonthlyAdjustedList{}
	query := fmt.Sprintf("%s/query?function=%v&symbol=%v&apikey=%v", c.baseURL, function, symbol, c.apiKey)

	if err := c.processRequest(query, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *Client) GetQuote(options *QuoteOptions) (*QuoteList, error) {
	const function = "GLOBAL_QUOTE"
	var symbol string

	if options != nil {
		symbol = options.Symbol
	}

	res := QuoteList{}
	query := fmt.Sprintf("%s/query?function=%v&symbol=%v&apikey=%v", c.baseURL, function, symbol, c.apiKey)

	if err := c.processRequest(query, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
