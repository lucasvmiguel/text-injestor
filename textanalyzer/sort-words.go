package textanalyzer

import "sort"

type pair struct {
	Key   string
	Value int
}

type pairList []pair

func (p pairList) Len() int           { return len(p) }
func (p pairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p pairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func (c *Client) sortMostUsedWords() pairList {
	plWordsMap := make(pairList, c.wordsMap.Count())
	index := 0

	for key, value := range c.wordsMap.Items() {
		strValue, _ := value.(int)
		plWordsMap[index] = pair{key, strValue}
		index++
	}

	sort.Sort(sort.Reverse(plWordsMap))
	return plWordsMap
}
