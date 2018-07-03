package textanalyzer

import (
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/pkg/errors"

	cmap "github.com/orcaman/concurrent-map"
)

// Client struct to use package
type Client struct {
	totalWords int
	totalChars int
	wordsMap   cmap.ConcurrentMap
	text       string
}

// New receives the text and index
func New(text string) (Client, error) {
	if text == "" {
		return Client{}, errors.New("empty text")
	}

	wordsMap, totalWords, totalChars, err := indexWordsMap(text)
	if err != nil {
		return Client{}, errors.Wrap(err, "error to index words map")
	}

	return Client{
		text:       text,
		wordsMap:   wordsMap,
		totalWords: totalWords,
		totalChars: totalChars,
	}, nil
}

// NumberOfChars return the number of characters in the text
func (c *Client) NumberOfChars() int {
	return c.totalChars
}

// NumberOfLines return the number of lines in the text
func (c *Client) NumberOfLines() (total int) {
	splittedLines := strings.Split(c.text, "\n")
	for _, line := range splittedLines {
		if line != "" {
			total = total + 1
		}
	}

	return total
}

// NumberOfWords return the number of words in the text
func (c *Client) NumberOfWords() (total int) {
	return c.totalWords
}

// FiveMostUsedWords return five most used words
func (c *Client) FiveMostUsedWords() []string {
	mostUsedWords := c.sortMostUsedWords()

	if mostUsedWords.Len() >= 5 {
		return []string{
			mostUsedWords[0].Key,
			mostUsedWords[1].Key,
			mostUsedWords[2].Key,
			mostUsedWords[3].Key,
			mostUsedWords[4].Key,
		}
	}

	return []string{}
}

func indexWordsMap(text string) (cmap.ConcurrentMap, int, int, error) {
	wordsMap := cmap.New()
	totalWords := 0
	totalChars := 0

	totalChars = utf8.RuneCountInString(removeInvalidChars(text))

	// tranform break line into whitespace
	wordsInSameLine := strings.Replace(text, "\n", " ", 0)

	// clean text
	reg, err := regexp.Compile("[^a-zA-Z0-9 ]+")
	if err != nil {
		return nil, 0, 0, errors.Wrap(err, "error to clean map when it was indexing")
	}
	cleanText := reg.ReplaceAllString(wordsInSameLine, " ")

	// splitted words in array
	splittedWords := strings.Split(cleanText, " ")

	// put each word as a key in a map and count
	for _, word := range splittedWords {
		if word == "" {
			continue
		}

		numWords, ok := wordsMap.Get(word)

		countWords, _ := numWords.(int)
		if ok {
			countWords = countWords + 1
			wordsMap.Set(word, countWords)
		} else {
			wordsMap.Set(word, 1)
		}

		// count all words
		totalWords = totalWords + 1
	}

	return wordsMap, totalWords, totalChars, nil
}

func removeInvalidChars(text string) string {
	return strings.NewReplacer(
		" ", "",
		"\t", "",
		"\n", "").Replace(text)
}
