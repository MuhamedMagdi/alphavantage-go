package alphaVantage

import (
	"fmt"
)

type IntraDayOptions struct {
	Symbol     string `json:"symbol"`
	Interval   string `json:"interval"`
	Adjusted   string `json:"adjusted,omitempty"`
	OutputSize string `json:"outputsize,omitempty"`
}

type DailyOptions struct {
	Symbol     string `json:"symbol"`
	OutputSize string `json:"outputsize,omitempty"`
}

type DailyAdjustedOptions DailyOptions

type WeeklyOptions struct {
	Symbol string `json:"symbol"`
}

type WeeklyAdjustedOptions WeeklyOptions

type MonthlyOptions WeeklyOptions

type MonthlyAdjustedOptions WeeklyOptions

type IntraDayList struct {
	MetaData        IntraDayMetaData `json:"Meta Data"`
	TimeSeries1Min  map[string]OHLC  `json:"Time Series (1min),omitempty"`
	TimeSeries5Min  map[string]OHLC  `json:"Time Series (5min),omitempty"`
	TimeSeries15Min map[string]OHLC  `json:"Time Series (15min),omitempty"`
	TimeSeries30Min map[string]OHLC  `json:"Time Series (30min),omitempty"`
	TimeSeries60Min map[string]OHLC  `json:"Time Series (60min),omitempty"`
}

type DailyList struct {
	MetaData        DailyMetaData   `json:"Meta Data"`
	TimeSeriesDaily map[string]OHLC `json:"Time Series (Daily)"`
}

type DailyAdjustedList struct {
	MetaData        DailyMetaData           `json:"Meta Data"`
	TimeSeriesDaily map[string]OHLCAdjusted `json:"Time Series (Daily)"`
}

type WeeklyList struct {
	MetaData         WeeklyMetaData  `json:"Meta Data"`
	WeeklyTimeSeries map[string]OHLC `json:"Weekly Time Series"`
}

type WeeklyAdjustedList struct {
	MetaData         WeeklyMetaData          `json:"Meta Data"`
	WeeklyTimeSeries map[string]OHLCAdjusted `json:"Weekly Adjusted Time Series"`
}

type MonthlyList struct {
	MetaData          MonthlyMetaData `json:"Meta Data"`
	MonthlyTimeSeries map[string]OHLC `json:"Monthly Time Series"`
}

type MonthlyAdjustedList struct {
	MetaData          MonthlyAdjustedMetaData `json:"Meta Data"`
	MonthlyTimeSeries map[string]OHLCAdjusted `json:"Monthly Adjusted Time Series"`
}

type IntraDayMetaData struct {
	Information   string `json:"1. Information"`
	Symbol        string `json:"2. Symbol"`
	LastRefreshed string `json:"3. Last Refreshed"`
	Interval      string `json:"4. Interval"`
	OutputSize    string `json:"5. Output Size"`
	TimeZone      string `json:"6. Time Zone"`
}

type DailyMetaData struct {
	Information   string `json:"1. Information"`
	Symbol        string `json:"2. Symbol"`
	LastRefreshed string `json:"3. Last Refreshed"`
	OutputSize    string `json:"4. Output Size"`
	TimeZone      string `json:"5. Time Zone"`
}

type WeeklyMetaData struct {
	Information   string `json:"1. Information"`
	Symbol        string `json:"2. Symbol"`
	LastRefreshed string `json:"3. Last Refreshed"`
	TimeZone      string `json:"4. Time Zone"`
}

type MonthlyMetaData WeeklyMetaData

type MonthlyAdjustedMetaData WeeklyMetaData

type OHLC struct {
	Open   string `json:"1. open"`
	High   string `json:"2. high"`
	Low    string `json:"3. low"`
	Close  string `json:"4. close"`
	Volume string `json:"5. volume"`
}

type OHLCAdjusted struct {
	Open             string `json:"1. open"`
	High             string `json:"2. high"`
	Low              string `json:"3. low"`
	Close            string `json:"4. close"`
	AdjustedClose    string `json:"5. adjusted close"`
	Volume           string `json:"6. volume"`
	DividendAmount   string `json:"7. dividend amount"`
	SplitCoefficient string `json:"8. split coefficient,omitempty"`
}

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
	fmt.Println(res)
	return &res, nil
}
