package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/lucasvmiguel/text-injestor/textanalyzer"
)

type statsResponse struct {
	Stats stats `json:"stats"`
}

type stats struct {
	Lines      int      `json:"lines"`
	Characters int      `json:"characters"`
	Words      int      `json:"words"`
	Top5Words  []string `json:"top_5_words"`
}

// Stats stats http handler - only POST method allowed
func Stats(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "404 page not found")
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, "422 unprocessable entity")
		return
	}

	taClient, err := textanalyzer.New(string(body), false)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "500 internal server error")
		return
	}

	statsResponse := statsResponse{
		Stats: stats{
			Lines:      taClient.NumberOfLines(),
			Characters: taClient.NumberOfChars(),
			Words:      taClient.NumberOfWords(),
			Top5Words:  taClient.FiveMostUsedWords(),
		},
	}

	data, err := json.Marshal(statsResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "500 internal server error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
