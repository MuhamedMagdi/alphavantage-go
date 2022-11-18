package alphaVantage

type CurrencyExchangeRateOptions struct {
	FromCurrency string `json:"from_currency"`
	ToCurrency   string `json:"to_currency"`
}

type CryptoIntrDayOptions struct {
	Symbol     string `json:"symbol"`
	Market     string `json:"market"`
	Interval   string `json:"interval"`
	OutputSize string `json:"outputsize,omitempty"`
}

type CryptoDailyOptions struct {
	Symbol string `json:"symbol"`
	Market string `json:"market"`
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

type CryptoIntraDayList struct {
	MetaData              CryptoIntraDayMetaData        `json:"Meta Data"`
	TimeSeriesCrypto1Min  map[string]CryptoIntraDayOHLC `json:"Time Series Crypto (1min),omitempty"`
	TimeSeriesCrypto5Min  map[string]CryptoIntraDayOHLC `json:"Time Series Crypto (5min),omitempty"`
	TimeSeriesCrypto15Min map[string]CryptoIntraDayOHLC `json:"Time Series Crypto (15min),omitempty"`
	TimeSeriesCrypto30Min map[string]CryptoIntraDayOHLC `json:"Time Series Crypto (30min),omitempty"`
	TimeSeriesCrypto60Min map[string]CryptoIntraDayOHLC `json:"Time Series Crypto (60min),omitempty"`
}

type CryptoIntraDayMetaData struct {
	Information         string `json:"1. Information"`
	DigitalCurrencyCode string `json:"2. Digital Currency Code"`
	DigitalCurrencyName string `json:"3. Digital Currency Name"`
	MarketCode          string `json:"4. Market Code"`
	MarketName          string `json:"5. Market Name"`
	LastRefreshed       string `json:"6. Last Refreshed"`
	Interval            string `json:"7. Interval"`
	OutputSize          string `json:"8. Output Size"`
	TimeZone            string `json:"9. Time Zone"`
}

type CryptoIntraDayOHLC struct {
	Open   string `json:"1. open"`
	High   string `json:"2. high"`
	Low    string `json:"3. low"`
	Close  string `json:"4. close"`
	Volume int    `json:"5. volume"`
}

type CryptoDailyList struct {
	MetaData                       CryptoDailyMetaData          `json:"Meta Data"`
	TimeSeriesDigitalCurrencyDaily map[string]map[string]string `json:"Time Series (Digital Currency Daily)"`
	// FIXME: update the type				  ^
}

type CryptoDailyMetaData struct {
	Information         string `json:"1. Information"`
	DigitalCurrencyCode string `json:"2. Digital Currency Code"`
	DigitalCurrencyName string `json:"3. Digital Currency Name"`
	MarketCode          string `json:"4. Market Code"`
	MarketName          string `json:"5. Market Name"`
	LastRefreshed       string `json:"6. Last Refreshed"`
	TimeZone            string `json:"7. Time Zone"`
}
