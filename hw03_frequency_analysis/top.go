package hw03frequencyanalysis

import (
	"regexp"
	"sort"
)

var reg = "\\s+"

type wordInfo struct {
	word  string
	count int
}

func Top10(input string) []string {
	infos := calculateWords(input)

	sort.Slice(infos, func(i, j int) bool {
		if infos[i].count == infos[j].count {
			return infos[i].word < infos[j].word
		}
		return infos[i].count > infos[j].count
	})

	return getTop10Words(infos)
}

func getTop10Words(wordInfos []wordInfo) []string {
	result := make([]string, 0)
	for _, word := range wordInfos {
		result = append(result, word.word)
	}

	if len(result) >= 10 {
		return result[:10]
	}
	return result
}

func calculateWords(input string) []wordInfo {
	re := regexp.MustCompile(reg)
	words := re.Split(input, -1)
	dict := make(map[string]int)
	for _, word := range words {
		if word == "" {
			continue
		}

		dict[word]++
	}

	elems := make([]wordInfo, 0)
	for word, count := range dict {
		elems = append(elems, wordInfo{word, count})
	}

	return elems
}
