package alphabet

import (
	"regexp"
	"strings"
)

// cleanText removes from text all special characters except white spaces and converts it to lowercase.
func cleanText(text string) string {

	// comvert to lowercase
	cleanedText := strings.ToLower(text)

	// replace special characters with spaces
	symbols := regexp.MustCompile(`[^\w]`)
	cleanedText = symbols.ReplaceAllString(cleanedText, " ")

	// remove leading, trailing and large spaces
	spaces := regexp.MustCompile(`\s+`)
	cleanedText = strings.TrimSpace(cleanedText)
	cleanedText = spaces.ReplaceAllString(cleanedText, " ")

	return cleanedText
}

// GetWords returns a slice with every word of a text, in lowercase. It could have repeated words.
func GetWords(text string) []string {
	cleannedText := cleanText(text)
	if cleannedText == "" {
		return nil
	}
	words := strings.Split(cleannedText, " ")
	return words
}

// GetAlphabet returns a slice of strings containing every different word of a text, in lowercase.
func GetAlphabet(text string) []string {
	words := GetWords(text)
	if words == nil {
		return nil
	}
	isRepeated := make(map[string]bool, len(words))
	alp := make([]string, 0, len(words))
	for _, val := range words {
		if !isRepeated[val] {
			isRepeated[val] = true
			alp = append(alp, val)
		}
	}
	return alp
}

// GetAlphabetWithOcurrence returns a map where the key is a word that appears in the text and the value is the number of times it appears.
func GetAlphabetWithOcurrence(text string) map[string]int {

	words := GetWords(text)
	if words == nil {
		return nil
	}

	alpOcurr := make(map[string]int, len(words))
	for _, el := range words {
		alpOcurr[el] = alpOcurr[el] + 1
	}
	return alpOcurr
}
