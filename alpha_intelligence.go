package alphaVantage

import "fmt"

func (c *Client) GetNewsAndSentiment(options *NewsAndSentimentOptions) (*NewsAndSentimentList, error) {
	const function = "NEWS_SENTIMENT"
	var (
		tickers, topics, time_from, time_to, sort string
		limit                                     int
	)
	if options != nil {
		tickers = options.Tickers
		topics = options.Topics
		time_from = options.TimeFrom
		time_to = options.TimeTo
		sort = options.Sort
		limit = options.Limit
	}

	res := NewsAndSentimentList{}
	query := fmt.Sprintf("%s/query?function=%v&tickers=%s&topics=%s&time_from=%s&time_to=%s&sort=%s&limit=%d&apikey=%s", c.baseURL, function, tickers, topics, time_from, time_to, sort, limit, c.apiKey)

	if err := c.processRequest(query, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *Client) GetWinningPortfolios(options *WinningPortfoliosOptions) (*WinningPortfoliosList, error) {
	const function = "TOURNAMENT_PORTFOLIO"
	var season string
	if options != nil {
		season = options.Season
	}

	res := WinningPortfoliosList{}
	query := fmt.Sprintf("%s/query?function=%v&season=%s&apikey=%s", c.baseURL, function, season, c.apiKey)

	if err := c.processRequest(query, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
