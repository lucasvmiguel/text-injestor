package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"

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

	taClient, err := textanalyzer.New(string(body))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "500 internal server error")
		return
	}

	statsResponse := statsResponse{
		Stats: *buildStatsResponse(taClient),
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

// buildStatsResponse build stats response splitting different tasks in differents goroutines
// for bigger texts the results will be better
// but for small texts the results will be equal or worst
// you can check with the benchmark that I jave created on the stats_test.go
func buildStatsResponse(taClient textanalyzer.Client) *stats {
	statsObj := &stats{
		Words:      taClient.NumberOfWords(),
		Characters: taClient.NumberOfChars(),
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		statsObj.Lines = taClient.NumberOfLines()
		wg.Done()
	}()
	go func() {
		statsObj.Top5Words = taClient.FiveMostUsedWords()
		wg.Done()
	}()

	wg.Wait()

	return statsObj
}
