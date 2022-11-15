package alphaVantage

import (
	"fmt"
	"net/http"
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

type OHLC struct {
	Open   string `json:"1. open"`
	High   string `json:"2. high"`
	Low    string `json:"3. low"`
	Close  string `json:"4. close"`
	Volume string `json:"5. volume"`
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
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/query?function=%v&symbol=%v&interval=%v&adjusted=%v&outputsize=%v&apikey=%v", c.baseURL, function, symbol, interval, adjusted, outputSize, c.apiKey), nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	res := IntraDayList{}
	if err := c.sendRequest(req, &res); err != nil {
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

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/query?function=%v&symbol=%v&outputsize=%v&apikey=%v", c.baseURL, function, symbol, outputSize, c.apiKey), nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	res := DailyList{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
