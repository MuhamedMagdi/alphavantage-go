package alphaVantage

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

type Client struct {
	apiKey     string
	baseURL    string
	HTTPClient *http.Client
}

type errorResponse struct {
	ErrorMessage string `json:"Error Message"`
}

func NewClient(apiKey string) *Client {
	return &Client{
		apiKey:  apiKey,
		baseURL: "https://www.alphavantage.co",
		HTTPClient: &http.Client{
			Timeout: 2 * time.Minute,
		},
	}
}

func (c *Client) sendRequest(req *http.Request, v interface{}) error {
	req.Header.Set("Accept", "application/json; charset=utf-8")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	var successBody bytes.Buffer
	errorBody := io.TeeReader(res.Body, &successBody)

	var errRes errorResponse
	if err = json.NewDecoder(errorBody).Decode(&errRes); err == nil {
		if len(errRes.ErrorMessage) != 0 {
			return errors.New(errRes.ErrorMessage)
		}
	}
	fullResponse := v
	if err = json.NewDecoder(&successBody).Decode(&fullResponse); err != nil {
		return err
	}

	return nil
}
