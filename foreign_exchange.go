package alphaVantage

import "fmt"

func (c *Client) GetCurrencyExchangeRate(options *CurrencyExchangeRateOptions) (*CurrencyExchangeRateList, error) {
	const function = "CURRENCY_EXCHANGE_RATE"
	var fromCurrency, toCurrency string
	if options != nil {
		fromCurrency = options.FromCurrency
		toCurrency = options.ToCurrency
	}

	res := CurrencyExchangeRateList{}
	query := fmt.Sprintf("%s/query?function=%s&from_currency=%s&to_currency=%s&apikey=%s", c.baseURL, function, fromCurrency, toCurrency, c.apiKey)

	if err := c.processRequest(query, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *Client) GetFXIntraDay(options *FXIntrDayOptions) (*FXIntraDayList, error) {
	const function = "FX_INTRADAY"
	var fromSymbol, toSymbol, interval, outputSize string
	if options != nil {
		fromSymbol = options.FromSymbol
        toSymbol = options.ToSymbol
		interval = options.Interval
		outputSize = options.OutputSize
	}

	res := FXIntraDayList{}
	query := fmt.Sprintf("%s/query?function=%s&from_symbol=%s&to_symbol=%s&interval=%s&outputsize=%s&apikey=%s", c.baseURL, function, fromSymbol, toSymbol, interval, outputSize, c.apiKey)

	if err := c.processRequest(query, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *Client) GetCryptoDaily(options *CryptoDailyOptions) (*CryptoDailyList, error) {
	const function = "DIGITAL_CURRENCY_DAILY"
	var symbol, market string
	if options != nil {
		symbol = options.Symbol
		market = options.Market
	}

	res := CryptoDailyList{}
	query := fmt.Sprintf("%s/query?function=%s&symbol=%s&market=%s&&apikey=%s", c.baseURL, function, symbol, market, c.apiKey)

	if err := c.processRequest(query, &res); err != nil {
		return nil, err
	}
	fmt.Println(res)
	return &res, nil
}
