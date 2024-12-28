package hw03frequencyanalysis

import (
	"regexp"
	"sort"

	"github.com/samber/lo"
)

var reg = "\\s+"

type elem struct {
	word  string
	count int
}

func Top10(input string) []string {
	re := regexp.MustCompile(reg)
	words := re.Split(input, -1)
	dict := make(map[string]int)
	for _, word := range words {
		if word == "" {
			continue
		}

		dict[word]++
	}

	elems := make([]elem, 0)
	for word, count := range dict {
		elems = append(elems, elem{word, count})
	}

	sort.Slice(elems, func(i, j int) bool {
		if elems[i].count == elems[j].count {
			return elems[i].word < elems[j].word
		}
		return elems[i].count > elems[j].count
	})

	top := lo.Map(elems, func(item elem, _ int) string {
		return item.word
	})
	if len(top) >= 10 {
		return top[:10]
	}
	return top
}
