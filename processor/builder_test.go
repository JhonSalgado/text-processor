package processor

import "testing"

func TestGetProcessor(t *testing.T) {
	got := GetTextProcessor()
	if got.filter != false {
		t.Fatalf("Expected processor filter to be: %t\n Got: %t", false, got.filter)
	}
	if len(got.stopwords) > 0 {
		t.Fatalf("Expected stopwords to be: 0\n Got: %d", len(got.stopwords))
	}
	if len(got.customStopwords) > 0 {
		t.Fatalf("Expected custom stopwords to be: 0\n Got: %d", len(got.customStopwords))
	}
}

func TestGetTextProcessorWithStopWordsFilterCustomStopWords(t *testing.T) {
	filter := Filter{
		OnlyCustom:      true,
		CustomStopwords: []string{"hola", "hello"},
		Langs:           []string{"en"},
	}
	proc, err := GetTextProcessorWithStopWordsFilter(filter)
	if err != nil {
		t.Fatalf("Error found: %s", err)
	}
	if proc.filter != true {
		t.Fatalf("Expected processor filter to be: %t\n Got: %t", true, proc.filter)
	}
	if proc.onlyCustom != true {
		t.Fatalf("Expected processor only custom to be: %t\n Got: %t", true, proc.onlyCustom)
	}
	if len(proc.stopwords) > 0 {
		t.Fatalf("Expected stopwords to be: %d\n Got: %d", 0, len(proc.stopwords))
	}
	if len(proc.customStopwords) != 2 {
		t.Fatalf("Expected custom stopwords to be: %d\n Got: %d", 2, len(proc.customStopwords))
	}
}

func TestGetTextProcessorWithStopWordsFilterBasicStopWords(t *testing.T) {
	filter := Filter{
		OnlyCustom:      false,
		CustomStopwords: []string{},
		Langs:           []string{"ar"},
	}
	proc, err := GetTextProcessorWithStopWordsFilter(filter)
	if err != nil {
		t.Fatalf("Error found: %s", err)
	}
	if proc.filter != true {
		t.Fatalf("Expected processor filter to be: %t\n Got: %t", true, proc.filter)
	}
	if proc.onlyCustom != false {
		t.Fatalf("Expected processor only custom to be: %t\n Got: %t", false, proc.onlyCustom)
	}
	if len(proc.stopwords) < 160 {
		t.Fatalf("Expected stopwords to be more than: %d\n Got: %d", 160, len(proc.stopwords))
	}
	if len(proc.customStopwords) > 0 {
		t.Fatalf("Expected custom stopwords to be: %d\n Got: %d", 0, len(proc.customStopwords))
	}
}

func TestGetTextProcessorWithStopWordsFilterAllBasicStopwords(t *testing.T) {
	filter := Filter{
		OnlyCustom:      false,
		CustomStopwords: []string{},
		Langs:           []string{},
	}
	proc, err := GetTextProcessorWithStopWordsFilter(filter)
	if err != nil {
		t.Fatalf("Error found: %s", err)
	}
	if proc.filter != true {
		t.Fatalf("Expected processor filter to be: %t\n Got: %t", true, proc.filter)
	}
	if proc.onlyCustom != false {
		t.Fatalf("Expected processor only custom to be: %t\n Got: %t", false, proc.onlyCustom)
	}
	if len(proc.stopwords) < 8500 {
		t.Fatalf("Expected stopwords to be more than: %d\n Got: %d", 8500, len(proc.stopwords))
	}
	if len(proc.customStopwords) > 0 {
		t.Fatalf("Expected custom stopwords to be: %d\n Got: %d", 0, len(proc.customStopwords))
	}
}

func TestGetTextProcessorWithStopWordsFilterBadLang(t *testing.T) {
	badCode := "zz"
	filter := Filter{
		OnlyCustom:      false,
		CustomStopwords: []string{},
		Langs:           []string{badCode},
	}
	_, err := GetTextProcessorWithStopWordsFilter(filter)
	if err == nil {
		t.Fatalf("Expected error to be 'language code '%s' not supported', got nil", badCode)
	}
}
