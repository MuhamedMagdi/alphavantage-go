package alphaVantage

type CurrencyExchangeRateOptions struct {
	FromCurrency string `json:"from_currency"`
	ToCurrency   string `json:"to_currency"`
}

type FXIntrDayOptions struct {
	FromSymbol string `json:"from_symbol"`
	ToSymbol   string `json:"to_symbol"`
	Interval   string `json:"interval"`
	OutputSize string `json:"outputsize,omitempty"`
}

type FXDailyOptions struct {
	FromSymbol string `json:"from_symbol"`
	ToSymbol   string `json:"to_symbol"`
	OutputSize string `json:"outputsize,omitempty"`
}

type CurrencyExchangeRateList struct {
	RealtimeCurrencyExchangeRate struct {
		FromCurrencyCode string `json:"1. From_Currency Code"`
		FromCurrencyName string `json:"2. From_Currency Name"`
		ToCurrencyCode   string `json:"3. To_Currency Code"`
		ToCurrencyName   string `json:"4. To_Currency Name"`
		ExchangeRate     string `json:"5. Exchange Rate"`
		LastRefreshed    string `json:"6. Last Refreshed"`
		TimeZone         string `json:"7. Time Zone"`
		BidPrice         string `json:"8. Bid Price"`
		AskPrice         string `json:"9. Ask Price"`
	} `json:"Realtime Currency Exchange Rate"`
}

type FXIntraDayList struct {
	MetaData              FXIntraDayMetaData        `json:"Meta Data"`
	TimeSeriesCrypto1Min  map[string]FXIntraDayOHLC `json:"Time Series FX (1min),omitempty"`
	TimeSeriesCrypto5Min  map[string]FXIntraDayOHLC `json:"Time Series FX (5min),omitempty"`
	TimeSeriesCrypto15Min map[string]FXIntraDayOHLC `json:"Time Series FX (15min),omitempty"`
	TimeSeriesCrypto30Min map[string]FXIntraDayOHLC `json:"Time Series FX (30min),omitempty"`
	TimeSeriesCrypto60Min map[string]FXIntraDayOHLC `json:"Time Series FX (60min),omitempty"`
}

type FXIntraDayMetaData struct {
	Information    string `json:"1. Information"`
	FromSymbol     string `json:"2. From Symbol"`
	ToSymbol       string `json:"3. To Symbol"`
	LastRefreshed  string `json:"4. Last Refreshed"`
	Interval       string `json:"5. Interval"`
	OutputSize     string `json:"6. Output Size"`
	TimeZone       string `json:"7. Time Zone"`
}

type FXIntraDayOHLC struct {
	Open   string `json:"1. open"`
	High   string `json:"2. high"`
	Low    string `json:"3. low"`
	Close  string `json:"4. close"`
}

type FXDailyList struct {
	MetaData          FXDailyMetaData        `json:"Meta Data"`
	TimeSeriesFXDaily map[string]FXDailyOHLC `json:"Time Series FX (Daily)"`
}

type FXDailyMetaData struct {
	Information   string `json:"1. Information"`
	FromSymbol    string `json:"2. From Symbol"`
	ToSymbol      string `json:"3. To Symbol"`
	OutputSize    string `json:"4. Output Size"`
	LastRefreshed string `json:"5. Last Refreshed"`
	TimeZone      string `json:"6. Time Zone"`
}

type FXDailyOHLC FXIntraDayOHLC
