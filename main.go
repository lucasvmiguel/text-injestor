package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"sort"
	"strings"
	"unicode/utf8"

	"github.com/davecgh/go-spew/spew"
)

func rankByWordCount(wordFrequencies map[string]int) PairList {
	pl := make(PairList, len(wordFrequencies))
	i := 0
	for k, v := range wordFrequencies {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "404 page not found")
		return
	}

	totalLines := 0
	totalWords := 0
	wordsMap := map[string]int{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, "422 unprocessable entity")
		return
	}

	strBody := string(body)

	// counting number of lines
	splittedLines := strings.Split(strBody, "\n")
	for _, line := range splittedLines {
		if line != "" {
			totalLines = totalLines + 1
		}
	}

	// tranform break line into whitespace
	wordsInSameLine := strings.Replace(strBody, "\n", " ", 0)

	// clean text
	reg, err := regexp.Compile("[^a-zA-Z0-9 ]+")
	if err != nil {
		log.Fatal(err)
	}
	cleanText := reg.ReplaceAllString(wordsInSameLine, " ")

	// splitted words in array
	splittedWords := strings.Split(cleanText, " ")

	// put each word as a key in a map and count
	for _, word := range splittedWords {
		if word == "" {
			continue
		}

		numWords, ok := wordsMap[word]
		if ok {
			wordsMap[word] = numWords + 1
		} else {
			wordsMap[word] = 1
		}

		totalWords = totalWords + 1
	}

	sortedPairList := rankByWordCount(wordsMap)

	spew.Dump(sortedPairList[0].Key)
	spew.Dump(sortedPairList[1].Key)
	spew.Dump(sortedPairList[2].Key)
	spew.Dump(sortedPairList[3].Key)
	spew.Dump(sortedPairList[4].Key)
	spew.Dump(totalLines)
	spew.Dump(totalWords)
	spew.Dump(utf8.RuneCountInString(strBody))

	fmt.Fprint(w, "done")
}

func main() {
	http.HandleFunc("/stats", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
