package processor

import (
	"sort"
	"testing"
)

var processor textProcessor
var processorCustomStopwords textProcessor
var processorBasicStopwords textProcessor
var processorAllStopwords textProcessor

func init() {
	// Simple processor
	processor = GetTextProcessor()

	// Processor with custom stopwords filter
	filterCustomStopwords := Filter{
		OnlyCustom:      true,
		CustomStopwords: []string{"hello", "bye"},
	}
	processorCustomStopwords, _ = GetTextProcessorWithStopWordsFilter(filterCustomStopwords)

	// Processor with basic stopwords filter
	filterBasicStopwords := Filter{
		OnlyCustom: false,
		Langs:      []string{"en"},
	}
	processorBasicStopwords, _ = GetTextProcessorWithStopWordsFilter(filterBasicStopwords)

	// Processor with basic and custom stopwords filter
	filterAllStopwords := Filter{
		OnlyCustom:      false,
		CustomStopwords: []string{"eat"},
	}
	processorAllStopwords, _ = GetTextProcessorWithStopWordsFilter(filterAllStopwords)
}

func compareSlices(want []string, got []string, t *testing.T) {
	lenWant := len(want)
	lenGot := len(got)
	if lenWant != lenGot {
		t.Fatalf("Slices have different lenght. Want: %v, got: %v.", want, got)
	}
	for i := 0; i < lenGot; i++ {
		if want[i] != got[i] {
			t.Fatalf("Difference found at position %d. Want: %s, got: %s.", i, want[i], got[i])
		}
	}
}

func compareMaps(want map[string]int, got map[string]int, t *testing.T) {
	lenWant := len(want)
	lenGot := len(got)
	if lenWant != lenGot {
		t.Fatalf("Slices have different lenght. Want: %v, got: %v.", want, got)
	}
	for key, value := range want {
		if value != got[key] {
			t.Fatalf("Different values for word '%s'. Want: %d, got: %d.", key, value, got[key])
		}
	}
}

// ====== Common methods =======
func TestCleanTextLowercase(t *testing.T) {
	want := "hello world"
	got := processor.CleanText("hello world")
	if want != got {
		t.Fatalf("Different strings.\n Want: '%s'\n Got: '%s'", want, got)
	}
}

func TestCleanTextUppercase(t *testing.T) {
	want := "hello world"
	got := processor.CleanText("HELLO World")
	if want != got {
		t.Fatalf("Different strings.\n Want: '%s'\n Got: '%s'", want, got)
	}
}

func TestCleanTextSpecialCharacters(t *testing.T) {
	want := "hello world"
	got := processor.CleanText("  HEllo!?... [(World)]  \n,,, ")
	if want != got {
		t.Fatalf("Different strings.\n Want: '%s'\n Got: '%s'", want, got)
	}
}

func TestCleanTextUnicodeCharacters(t *testing.T) {
	want := "yo nací en españa ты в россии"
	got := processor.CleanText("¡¡¡yo nací en españa!!! ты в россии")
	if want != got {
		t.Fatalf("Different strings.\n Want: '%s'\n Got: '%s'", want, got)
	}
}

func TestCleanTextEmpty(t *testing.T) {
	want := ""
	got := processor.CleanText("")
	if want != got {
		t.Fatalf("Different strings.\n Want: '%s'\n Got: '%s'", want, got)
	}
}

func TestGetWords(t *testing.T) {
	want := []string{"hello", "world", "this", "is", "a", "test"}
	got := processor.getWords("Hello, World!! (this/ is @a test.]?")
	compareSlices(want, got, t)
}

func TestGetWordsWordsRepeated(t *testing.T) {
	want := []string{"hello", "hello", "world", "worlds", "hello"}
	got := processor.getWords("Hello hello world worlds hellO")
	compareSlices(want, got, t)
}

func TestGetWordsEmptyText(t *testing.T) {
	got := processor.getWords("")
	if got != nil {
		t.Fatalf("Expected nil, got: %v", got)
	}
}

// ====== No filter =======
func TestGetWordsSetSpecialCharacters(t *testing.T) {
	want := []string{"a", "hello", "is", "test", "this", "world"}
	got := processor.GetWordsSet("Hello, World!! (this/ is @a test.]?")
	sort.Strings(got)
	compareSlices(want, got, t)
}

func TestGetWordsSetWordsRepeated(t *testing.T) {
	want := []string{"hello", "world", "worlds"}
	got := processor.GetWordsSet("Hello hello world worlds")
	sort.Strings(got)
	compareSlices(want, got, t)
}

func TestGetWordsSetEmptyText(t *testing.T) {
	got := processor.GetWordsSet("")
	if got != nil {
		t.Fatalf("Expected nil, got: %v", got)
	}
}

func TestGetWordsSetWithOcurrenceSpecialCharacters(t *testing.T) {
	want := map[string]int{"hello": 1, "world": 1}
	got := processor.GetWordsSetWithOcurrence("(Hello) + /woRld/!")
	compareMaps(want, got, t)
}

func TestGetWordsSetWithOcurrenceWordsRepeated(t *testing.T) {
	want := map[string]int{"hello": 3, "world": 1}
	got := processor.GetWordsSetWithOcurrence("(Hello) + [hello] /woRld/! heLLo")
	compareMaps(want, got, t)
}

func TestGetWordsSetWithOcurrenceEmptyText(t *testing.T) {
	got := processor.GetWordsSetWithOcurrence("")
	if got != nil {
		t.Fatalf("Expected nil, got: %v", got)
	}
}

// ====== Filter custom stopwords =======
func TestGetWordsCustomStopwordsFilter(t *testing.T) {
	want := []string{"a", "is", "test", "this", "world"}
	got := processorCustomStopwords.GetWordsSet("Hello, World!! (this/ is @a test.]? bye")
	sort.Strings(got)
	compareSlices(want, got, t)
}

func TestGetWordsCustomStopwordsFilterJustStopwords(t *testing.T) {
	got := processorCustomStopwords.GetWordsSet("-Hello +hellO -bye +bye")
	if got != nil {
		t.Fatalf("Expected nil, got: %v", got)
	}
}

func TestGetWordsSetWithOcurrenceCustomStopwordsFilter(t *testing.T) {
	want := map[string]int{"world": 1, "hi": 1}
	got := processorCustomStopwords.GetWordsSetWithOcurrence("hi, (Hello) + /woRld/! bye")
	compareMaps(want, got, t)
}

func TestGetWordsSetWithOcurrenceCustomStopwordsFilterJustStopwords(t *testing.T) {
	got := processorCustomStopwords.GetWordsSetWithOcurrence("-Hello +hellO -bye +bye")
	if got != nil {
		t.Fatalf("Expected nil, got: %v", got)
	}
}

// ====== Filter basic stopwords =======
func TestGetWordsBasicStopwordsFilter(t *testing.T) {
	want := []string{"earthquake"}
	got := processorBasicStopwords.GetWordsSet("Hello!! (this/ is @an earthquake.]?")
	sort.Strings(got)
	compareSlices(want, got, t)
}

func TestGetWordsBasicStopwordsFilterJustStopwords(t *testing.T) {
	got := processorBasicStopwords.GetWordsSet("-Hello this is me")
	if got != nil {
		t.Fatalf("Expected nil, got: %v", got)
	}
}

func TestGetWordsSetWithOcurrenceBasicStopwordsFilter(t *testing.T) {
	want := map[string]int{"pizza": 2, "eat": 1}
	got := processorBasicStopwords.GetWordsSetWithOcurrence("I like pizza, I want to eat pizza right now")
	compareMaps(want, got, t)
}

func TestGetWordsSetWithOcurrenceBasicStopwordsFilterJustStopwords(t *testing.T) {
	got := processorBasicStopwords.GetWordsSetWithOcurrence("-Hello this is me")
	if got != nil {
		t.Fatalf("Expected nil, got: %v", got)
	}
}

// ====== Filter all stopwords =======
func TestGetWordsAllStopwordsFilter(t *testing.T) {
	want := []string{"earthquake", "terremoto"}
	got := processorAllStopwords.GetWordsSet("Hello!! (this/ is @an earthquake o un terremoto.]?")
	sort.Strings(got)
	compareSlices(want, got, t)
}

func TestGetWordsAllStopwordsFilterJustStopwords(t *testing.T) {
	got := processorAllStopwords.GetWordsSet("Hello, erais, of, del")
	if got != nil {
		t.Fatalf("Expected nil, got: %v", got)
	}
}

func TestGetWordsSetWithOcurrenceAllStopwordsFilter(t *testing.T) {
	want := map[string]int{"pizza": 2, "chico": 1}
	got := processorAllStopwords.GetWordsSetWithOcurrence("I want to eat pizza right now. Aquel chico tiene pizza")
	compareMaps(want, got, t)
}

func TestGetWordsSetWithOcurrenceAllStopwordsFilterJustStopwords(t *testing.T) {
	got := processorAllStopwords.GetWordsSetWithOcurrence("Hello, erais, of, del")
	if got != nil {
		t.Fatalf("Expected nil, got: %v", got)
	}
}
