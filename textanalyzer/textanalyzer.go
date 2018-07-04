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
		return Client{}, errors.New("error empty text")
	}

	var wordsMap cmap.ConcurrentMap
	var totalWords int
	var err error

	wordsMap, totalWords, err = indexWordsMap(text)

	if err != nil {
		return Client{}, errors.Wrap(err, "error to index words map")
	}

	totalChars := utf8.RuneCountInString(removeInvalidChars(text))

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

// MostUsedWords return N most used words
func (c *Client) MostUsedWords(numberOfWords int) []string {
	mostUsedWords := c.sortMostUsedWords()
	var words []string

	for i, pairWord := range mostUsedWords {
		words = append(words, pairWord.Key)

		// break to get just N words
		if i == (numberOfWords - 1) {
			break
		}
	}

	return words
}

// indexWordsMap index all words of text in a map and count how many words
func indexWordsMap(text string) (cmap.ConcurrentMap, int, error) {
	var wordsMap cmap.ConcurrentMap
	var totalWords int

	cleanText, err := normalizeTextToIndex(text)
	if err != nil {
		return nil, 0, errors.Wrap(err, "error to clean map when it was indexing")
	}

	splittedWords := strings.Split(cleanText, " ")

	wordsMap, totalWords = buildWordsMapAndCount(splittedWords)

	return wordsMap, totalWords, nil
}

// normalizeTextToIndex remove symbols, put everything lowercase
func normalizeTextToIndex(text string) (string, error) {
	lowerText := strings.ToLower(text)
	wordsInSameLine := strings.Replace(lowerText, "\n", " ", -1)

	reg, err := regexp.Compile("[^a-zA-Z0-9 ]+")
	if err != nil {
		return "", errors.Wrap(err, "error to clean text when it was indexing")
	}
	cleanText := reg.ReplaceAllString(wordsInSameLine, " ")

	return cleanText, nil
}

// buildWordsMapAndCount puts all words in a map and count
// this function has two responsibilities to be more performatic
func buildWordsMapAndCount(words []string) (cmap.ConcurrentMap, int) {
	wordsCount := 0
	wordsMap := cmap.New()

	for _, word := range words {
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

		wordsCount = wordsCount + 1
	}

	return wordsMap, wordsCount
}

func removeInvalidChars(text string) string {
	return strings.NewReplacer(
		" ", "",
		"\t", "",
		"\n", "").Replace(text)
}
