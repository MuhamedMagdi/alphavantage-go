package alphaVantage

type NewsAndSentimentOptions struct {
	Tickers  string `json:"tickers,omitempty"`
	Topics   string `json:"topics,omitempty"`
	TimeFrom string `json:"time_from,omitempty"`
	TimeTo   string `json:"time_to,omitempty"`
	Sort     string `json:"sort,omitempty"`
	Limit    int    `json:"limit,omitempty"`
}

type WinningPortfoliosOptions struct {
	Season string `json:"season,omitempty"`
}

type NewsAndSentimentList struct {
	Items                    string `json:"items"`
	SentimentScoreDefinition string `json:"sentiment_score_definition"`
	RelevanceScoreDefinition string `json:"relevance_score_definition"`
	Feed                     []Feed `json:"feed"`
}

type Topics struct {
	Topic          string `json:"topic"`
	RelevanceScore string `json:"relevance_score"`
}

type TickerSentiment struct {
	Ticker               string `json:"ticker"`
	RelevanceScore       string `json:"relevance_score"`
	TickerSentimentScore string `json:"ticker_sentiment_score"`
	TickerSentimentLabel string `json:"ticker_sentiment_label"`
}

type Feed struct {
	Title                 string            `json:"title"`
	URL                   string            `json:"url"`
	TimePublished         string            `json:"time_published"`
	Authors               []string          `json:"authors"`
	Summary               string            `json:"summary"`
	BannerImage           string            `json:"banner_image"`
	Source                string            `json:"source"`
	CategoryWithinSource  string            `json:"category_within_source"`
	SourceDomain          string            `json:"source_domain"`
	Topics                []Topics          `json:"topics"`
	OverallSentimentScore float64           `json:"overall_sentiment_score"`
	OverallSentimentLabel string            `json:"overall_sentiment_label"`
	TickerSentiment       []TickerSentiment `json:"ticker_sentiment"`
}

type WinningPortfoliosList struct {
	Season  string    `json:"season"`
	Ranking []Ranking `json:"ranking"`
}

type Portfolio struct {
	Symbol string `json:"symbol"`
	Shares string `json:"shares"`
}

type Ranking struct {
	Rank             string      `json:"rank"`
	Portfolio        []Portfolio `json:"portfolio"`
	MeasurementStart string      `json:"measurement_start"`
	StartValueUsd    string      `json:"start_value_usd"`
	MeasurementEnd   string      `json:"measurement_end"`
	EndValueUsd      string      `json:"end_value_usd"`
	PercentGain      string      `json:"percent_gain"`
}
