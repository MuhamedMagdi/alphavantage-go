package alphaVantage

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

type QuoteOptions WeeklyOptions

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

type QuoteList struct {
	GlobalQuote struct {
		Symbol           string `json:"01. symbol"`
		Open             string `json:"02. open"`
		High             string `json:"03. high"`
		Low              string `json:"04. low"`
		Price            string `json:"05. price"`
		Volume           string `json:"06. volume"`
		LatestTradingDay string `json:"07. latest trading day"`
		PreviousClose    string `json:"08. previous close"`
		Change           string `json:"09. change"`
		ChangePercent    string `json:"10. change percent"`
	} `json:"Global Quote"`
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
