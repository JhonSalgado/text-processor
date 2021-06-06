# [WIP] text-processor
This package helps you to obtain the set of unique words present in a text and includes the possibility to filter stopwords from various languages, obtained from https://github.com/Alir3z4/stop-words. It also allows you to add your own words to filter.

It is similar to tokenize if you take the words as your tokens, but because it is focused on working with many languages it does not have advanced techniques such as Tagging Part of Speech (PoS).

The processor is not intended to be robust, that is, words that are spelled differently are considered different words (words in singular and plural, misspelled, etc.), except for the words with capital letters (it is case insentsitive).

Before obtaining the words, a pre-processing is carried out to clean the text and remove all non-word characters.

## Methods
This package has just two methods to process a text:
### GetWordsSet
Receives a text string and returns a slice with all the unique words in it sorted alphabetically. The complexity of this method is O(n log n), because of the sorting, with n being the number of words in the text.
### GetWordsSetWithOcurrence
Receives a text string and returns a value-key map where the keys are the words and the values are the number of times they occurred in the text. The complexity of this method is O(n), where n is the number of words in the text.

Complexity is maintained even when using filters, because stopwords are stored in haspmaps, so checking if a word is a stopword takes constant time O(1). If this is done for n words, the total time is just O(n).

## Usage
In order to use the methods mentioned above you need to create a text processor.
### Without filters
You need to create a text processor following the example below:
```
package main

import (
	"fmt"
	
    // import the package
	"github.com/JhonSalgado/text-processor/processor"
)

func main() {
    // create the processor
	textProcessor := processor.GetTextProcessor()
	text := "Hey! hey!! hey!!! how are you today?"

	// get words
	wordsSlice := textProcessor.GetWordsSet(text)
	fmt.Println("slice:", wordsSlice)

	// get words with ocurrence
	wordsMap := textProcessor.GetWordsSetWithOcurrence(text)
	fmt.Println("map:", wordsMap)
}
```
Output:
```
slice: [are hey how today you]
map: map[are:1 hey:3 how:1 today:1 you:1]
```
### With stopwords filter
Before creating a text processor you need to create a filter, with the Filter struct provided by the package:
```
filter := processor.Filter{
	OnlyCustom:      false,
	CustomStopwords: []string{"bye"},
	Langs:           []string{"en"},
}
```
`OnlyCustom` is a bool to indicate if you want to filter just your custom stopwords. If it is false, the stopwords of the languages present in the `Langs` list that are included in this package will also be filtered. If `Langs` is empty and `OnlyCustom`  is false, all available languages will be considered.
In `Langs` you only have to enter the ISO 639-1 codes of the languages, not their full name.
`CustomStopWords` is a list of the words you want to filter in replacement or in addition to those included in this package

After creating the filter you should use it as shown in the example below:
```
func main() {
	filter := processor.Filter{
		OnlyCustom:      false,
		CustomStopwords: []string{"bye"},
		Langs:           []string{"en"},
	}
	textProcessor, err := processor.GetTextProcessorWithStopWordsFilter(filter)
	if err != nil {
		log.Fatal(err)
	}
	text := "Hi, how are you today? I'm super fine. Bye!"
	wordsSet := textProcessor.GetWordsSetWithOcurrence(text)
	fmt.Println(wordsSet)
}
```
Output:
```map[fine:1 super:1]```

Explanation:
The method returns only those words because in the filter we indicated that we wanted to filter the english stopwords included in the package, and precisely the words that did not appear are classified as stopwords in the english language, we also filtered 'bye' because it was present in the custom stopwords.

## Available languages:

You can check the list of available languages and the list of stopwords for each one in its source repository https://github.com/Alir3z4/stop-words

Here is the list of ISO 639-1 codes accepted by the filter and their respective language:

* **ar**: arabic
* **bg**: bulgarian
* **ca**: catalan
* **cs**: czech
* **da**: danish
* **nl**: dutch
* **en**: english
* **fi**: finnish
* **fr**: french
* **de**: german
* **gu**: gujarati
* **he**: hebrew
* **hi**: hindi
* **hu**: hungarian
* **id**: indonesian
* **ms**: malaysian
* **it**: italian
* **nb**: norwegian
* **pl**: polish
* **pt**: portuguese
* **ro**: romanian
* **ru**: russian
* **sk**: slovak
* **es**: spanish
* **sv**: swedish
* **tr**: turkish
* **uk**: ukrainian
* **vi**: vietnamese