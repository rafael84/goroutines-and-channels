package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/rafael84/goroutines-and-channels/api"
)

type DailyExports struct {
	API api.Interface
}

func (de *DailyExports) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// get the export date parameter
	exportDate := api.Yesterday()
	dateParam := r.URL.Query().Get("export-date")
	if dateParam != "" {
		parsedDate, err := time.Parse(api.DateFormat, dateParam)
		if err != nil {
			api.WriteErrorResponse(w, 400, api.ErrCodeInvalidParam, fmt.Sprintf(`"%s" is not a valid date param`, dateParam))
		}
		exportDate = parsedDate
	}

	// get files from the given date
	response, err := de.API.DailyExports(exportDate)
	if err != nil {
		api.WriteErrorResponse(w, 500, api.ErrCodeGeneric, fmt.Sprintf("could not get response from server: %s", err.Error()))
	}

	// write the response
	json.NewEncoder(w).Encode(response)
}
