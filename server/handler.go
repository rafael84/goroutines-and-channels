package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"regexp"
	"strings"

	"github.com/rafael84/goroutines-and-channels/api"
)

const (
	FilesDir  = "./files"
	FilesPath = "/files/"

	serverBaseURL = "http://localhost:8080/"
)

var (
	datePattern = regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)
)

func downloadURL(date string, filename string) string {
	return path.Join(serverBaseURL, FilesPath, date, filename)
}

func DailyExportsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := r.URL.Query()
	exportDate := params.Get("export-date")

	if !datePattern.MatchString(exportDate) {
		api.WriteErrorResponse(w, 400, api.ErrCodeInvalidParam, fmt.Sprintf(`export-date "%s" is not valid`, exportDate))
		return
	}

	files, err := ioutil.ReadDir(path.Join(FilesDir, exportDate))
	if err != nil {
		log.Printf("could not list files: %v\n", err)
	}

	response := &api.DailyExportsResponse{
		Addresses:   []string{},
		SnakesCount: []string{},
	}

	for _, file := range files {
		filename := file.Name()
		url := downloadURL(exportDate, filename)

		if strings.HasPrefix(filename, "addresses") {
			response.Addresses = append(response.Addresses, url)
		}

		if strings.HasPrefix(filename, "snakes_count") {
			response.SnakesCount = append(response.SnakesCount, url)
		}
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		api.WriteErrorResponse(w, 500, api.ErrCodeGeneric, fmt.Sprintf("could not encode response: %s", err.Error()))
	}
}
