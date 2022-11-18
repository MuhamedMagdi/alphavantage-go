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
