package dev

import (
	"sort"
	"strings"
)

func findAnagrams(words []string) map[string][]string {
	anagramSets := make(map[string][]string)
	sortedWords := make([]string, len(words))
	for i := 0; i < len(words); i++ {
		sortedWords[i] = sortString(words[i])
	}
	for _, word := range words {
		isExist := false
		word = strings.ToLower(word)

		for k, _ := range anagramSets {
			if sortString(word) == sortString(k) {
				anagramSets[k] = append(anagramSets[k], word)
				isExist = true
				break
			}
		}
		if !isExist {
			anagramSets[word] = []string{word}
		}

	}

	for key, set := range anagramSets {
		if len(set) <= 1 {
			delete(anagramSets, key)
		} else {
			sort.Strings(set)
			anagramSets[key] = set
		}
	}

	return anagramSets
}

func sortString(s string) string {
	sChars := strings.Split(s, "")
	sort.Strings(sChars)
	return strings.Join(sChars, "")
}
