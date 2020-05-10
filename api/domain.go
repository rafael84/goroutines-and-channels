package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// DailyExportsResponse is a successful response from the Daily Exports API
type DailyExportsResponse struct {
	Addresses   []string `json:"addresses"`
	SnakesCount []string `json:"snakes_count"`
}

const (
	ErrCodeGeneric = iota + 1000
	ErrCodeInvalidParam
)

// ErrorResponse defines a generic error response
type ErrorResponse struct {
	Code        int
	Description string
}

// WriteErrorResponse writes an error response to the response writer
func WriteErrorResponse(w http.ResponseWriter, status int, errCode int, errDescription string) {
	w.WriteHeader(status)
	log.Printf("error response %d: %s", errCode, errDescription)
	json.NewEncoder(w).Encode(&ErrorResponse{
		Code:        errCode,
		Description: errDescription,
	})
}

const (
	DateFormat = "2006-01-02"
)

func Yesterday() time.Time {
	return time.Now().AddDate(0, 0, -1)
}
