package processor

import (
	"fmt"
	"strings"

	"github.com/JhonSalgado/text-processor/processor/stopwords"
)

// empty is an empty struct that occupies zero bytes of storage.
// It's perfect when you don't need the value part of a key-value map.
type empty struct{}

// Filter is a struct to store the user's filtering preferences.
// OnlyCustom is a bool to indicate if you want to filter just your custom stopwords.
// Langs is a list that indicates which set of stopwords among those included in this package should be filtered,
// if this field is empty and OnlyCustom is false, all available languages will be considered.
// CustomStopWords is a list of words you want to filter in addition to those included in this package.
type Filter struct {
	OnlyCustom      bool
	Langs           []string
	CustomStopwords []string
}

// textProcessor is a struct that contains filter specifications and methods for processing text.
type textProcessor struct {
	filter          bool
	onlyCustom      bool
	stopwords       map[string]empty
	customStopwords map[string]empty
}

// GetTextProcessor returns a text processor that does not filter stopwords.
func GetTextProcessor() textProcessor {
	return textProcessor{filter: false}
}

// GetTextProcessor returns a text processor that filters stopwords.
func GetTextProcessorWithStopWordsFilter(filter Filter) (textProcessor, error) {
	processor := textProcessor{
		filter:     true,
		onlyCustom: filter.OnlyCustom,
	}
	if !filter.OnlyCustom {
		processor.stopwords = make(map[string]empty)
		err := processor.storeStopwordsByLanguage(filter.Langs)
		if err != nil {
			return processor, err
		}
	}
	if len(filter.CustomStopwords) > 0 {
		processor.customStopwords = make(map[string]empty, len(filter.CustomStopwords))
		processor.storeCustomStopwords(filter.CustomStopwords)
	}
	return processor, nil
}

// loadStopWordsFromLang loads the stopwords from a single language file.
func (processor textProcessor) loadStopWordsFromLang(lang string, words []string) {
	for _, word := range words {
		processor.stopwords[word] = empty{}
	}
}

// storeStopwordsByLanguage stores in the text processor stopwords for each specified language.
// If there is none specified, it stores them all.
func (processor textProcessor) storeStopwordsByLanguage(langs []string) error {
	var err error = nil
	if len(langs) > 0 {
		for _, lang := range langs {
			words, ok := stopwords.Stopwords[lang]
			if !ok {
				err = fmt.Errorf("language code '%s' not supported", lang)
				break
			}
			processor.loadStopWordsFromLang(lang, words)
		}
	} else {
		for lang, words := range stopwords.Stopwords {
			processor.loadStopWordsFromLang(lang, words)
		}
	}
	return err
}

// storeCustomStopwords stores the custom stopwords in the text processor.
func (processor textProcessor) storeCustomStopwords(stopWords []string) {
	for _, word := range stopWords {
		processor.customStopwords[strings.ToLower(word)] = empty{}
	}
}
