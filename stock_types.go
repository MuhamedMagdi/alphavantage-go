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

type SearchOptions struct {
	Keywords string `json:"keywords"`
}

type IntraDayList struct {
	MetaData        IntraDayMetaData     `json:"Meta Data"`
	TimeSeries1Min  map[string]StockOHLC `json:"Time Series (1min),omitempty"`
	TimeSeries5Min  map[string]StockOHLC `json:"Time Series (5min),omitempty"`
	TimeSeries15Min map[string]StockOHLC `json:"Time Series (15min),omitempty"`
	TimeSeries30Min map[string]StockOHLC `json:"Time Series (30min),omitempty"`
	TimeSeries60Min map[string]StockOHLC `json:"Time Series (60min),omitempty"`
}

type DailyList struct {
	MetaData        DailyMetaData        `json:"Meta Data"`
	TimeSeriesDaily map[string]StockOHLC `json:"Time Series (Daily)"`
}

type DailyAdjustedList struct {
	MetaData        DailyMetaData                `json:"Meta Data"`
	TimeSeriesDaily map[string]StockOHLCAdjusted `json:"Time Series (Daily)"`
}

type WeeklyList struct {
	MetaData         WeeklyMetaData       `json:"Meta Data"`
	WeeklyTimeSeries map[string]StockOHLC `json:"Weekly Time Series"`
}

type WeeklyAdjustedList struct {
	MetaData         WeeklyMetaData               `json:"Meta Data"`
	WeeklyTimeSeries map[string]StockOHLCAdjusted `json:"Weekly Adjusted Time Series"`
}

type MonthlyList struct {
	MetaData          MonthlyMetaData      `json:"Meta Data"`
	MonthlyTimeSeries map[string]StockOHLC `json:"Monthly Time Series"`
}

type MonthlyAdjustedList struct {
	MetaData          MonthlyAdjustedMetaData      `json:"Meta Data"`
	MonthlyTimeSeries map[string]StockOHLCAdjusted `json:"Monthly Adjusted Time Series"`
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

type SearchList struct {
	BestMatches []struct {
		Symbol      string `json:"1. symbol"`
		Name        string `json:"2. name"`
		Type        string `json:"3. type"`
		Region      string `json:"4. region"`
		MarketOpen  string `json:"5. marketOpen"`
		MarketClose string `json:"6. marketClose"`
		Timezone    string `json:"7. timezone"`
		Currency    string `json:"8. currency"`
		MatchScore  string `json:"9. matchScore"`
	} `json:"bestMatches"`
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

type StockOHLC struct {
	Open   string `json:"1. open"`
	High   string `json:"2. high"`
	Low    string `json:"3. low"`
	Close  string `json:"4. close"`
	Volume string `json:"5. volume"`
}

type StockOHLCAdjusted struct {
	Open             string `json:"1. open"`
	High             string `json:"2. high"`
	Low              string `json:"3. low"`
	Close            string `json:"4. close"`
	AdjustedClose    string `json:"5. adjusted close"`
	Volume           string `json:"6. volume"`
	DividendAmount   string `json:"7. dividend amount"`
	SplitCoefficient string `json:"8. split coefficient,omitempty"`
}
