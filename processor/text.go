package processor

import (
	"regexp"
	"sort"
	"strings"
)

// cleanText removes from text all special characters except white spaces and converts it to lowercase.
func (processor textProcessor) cleanText(text string) string {

	// convert to lowercase
	cleanedText := strings.ToLower(text)

	// replace special characters with spaces
	symbols := regexp.MustCompile(`[^\p{L}\s]`)
	cleanedText = symbols.ReplaceAllString(cleanedText, "")

	// remove leading, trailing and large spaces
	spaces := regexp.MustCompile(`\s+`)
	cleanedText = strings.TrimSpace(cleanedText)
	cleanedText = spaces.ReplaceAllString(cleanedText, " ")

	return cleanedText
}

// getWords returns a slice with every word of a text, in lowercase. It could have repeated words.
func (processor textProcessor) getWords(text string) []string {
	cleannedText := processor.cleanText(text)
	if cleannedText == "" {
		return nil
	}
	words := strings.Split(cleannedText, " ")
	return words
}

// isStopWord checks if a word is a stopword
func (processor textProcessor) isStopword(word string) bool {
	_, ok := processor.stopwords[word]
	return ok
}

// isStopWord checks if a word is a custom stopword
func (processor textProcessor) isCustomStopword(word string) bool {
	_, ok := processor.customStopwords[word]
	return ok
}

// GetWordsSet returns a sorted slice of strings containing every different word of a text, in lowercase.
func (processor textProcessor) GetWordsSet(text string) []string {
	words := processor.getWords(text)
	if words == nil {
		return nil
	}
	set := make([]string, 0, len(words))
	setOcurr := processor.GetWordsSetWithOcurrence(text)
	if len(setOcurr) == 0 {
		return nil
	}
	for word := range setOcurr {
		set = append(set, word)
	}
	sort.Strings(set)
	return set
}

// GetWordsSetWithOcurrence returns a map where the key is a word that appears in the text and the value is the number of times it appears.
func (processor textProcessor) GetWordsSetWithOcurrence(text string) map[string]int {

	words := processor.getWords(text)
	if words == nil {
		return nil
	}

	set := make(map[string]int, len(words))
	if processor.filter {
		if !processor.onlyCustom {
			for _, word := range words {
				if processor.isCustomStopword(word) || processor.isStopword(word) {
					continue
				}
				set[word] = set[word] + 1
			}
		} else {
			for _, word := range words {
				if processor.isCustomStopword(word) {
					continue
				}
				set[word] = set[word] + 1
			}
		}
	} else {
		for _, el := range words {
			set[el] = set[el] + 1
		}
	}
	if len(set) == 0 {
		return nil
	}
	return set
}
