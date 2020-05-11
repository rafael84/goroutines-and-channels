package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// HTTP implements the client Interface
type HTTP struct {
	BaseURL string
}

// NewHTTP creates a new HTTP client
func NewHTTP(baseURL string) *HTTP {
	return &HTTP{BaseURL: baseURL}
}

// DailyExports makes an HTTP request to the third-party-server and returns the parsed HTTP response
func (h *HTTP) DailyExports(date time.Time) (*DailyExportsResponse, error) {
	// format the request url
	params := url.Values{}
	params.Add("export-date", date.Format(DateFormat))
	requestURL := fmt.Sprintf("%s/daily-exports?%s", h.BaseURL, params.Encode())

	// send the request
	r, err := http.Post(requestURL, "", nil)
	if err != nil {
		return nil, fmt.Errorf("could not send request: %s", err.Error())
	}

	// decode the response
	var der DailyExportsResponse
	if err := json.NewDecoder(r.Body).Decode(&der); err != nil {
		return nil, fmt.Errorf("could not decode response body: %s", err.Error())
	}

	return &der, nil
}
