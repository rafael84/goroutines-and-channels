package api

import (
	"time"
)

// Interface defines the third-party-server APIs
type Interface interface {

	// DailyExports returns the list of event files for a given date
	DailyExports(date time.Time) (*DailyExportsResponse, error)
}
