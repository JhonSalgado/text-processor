package alphabet

import (
	"testing"
)

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

func TestCleanTextLowercase(t *testing.T) {
	want := "hello world"
	got := cleanText("hello world")
	if want != got {
		t.Fatalf("Different strings.\n Want: '%s'\n Got: '%s'", want, got)
	}
}

func TestCleanTextUppercase(t *testing.T) {
	want := "hello world"
	got := cleanText("HELLO World")
	if want != got {
		t.Fatalf("Different strings.\n Want: '%s'\n Got: '%s'", want, got)
	}
}

func TestCleanTextSpecialCharacters(t *testing.T) {
	want := "hello world"
	got := cleanText("  HEllo!?... [(World)]  \n,,, ")
	if want != got {
		t.Fatalf("Different strings.\n Want: '%s'\n Got: '%s'", want, got)
	}
}

func TestCleanTextEmpty(t *testing.T) {
	want := ""
	got := cleanText("")
	if want != got {
		t.Fatalf("Different strings.\n Want: '%s'\n Got: '%s'", want, got)
	}
}

func TestGetWordsSpecialCharacters(t *testing.T) {
	want := []string{"hello", "world", "this", "is", "a", "test"}
	got := GetWords("Hello, World!! (this/ is @a test.]?")
	compareSlices(want, got, t)
}

func TestGetWordsWordsRepeated(t *testing.T) {
	want := []string{"hello", "hello", "world", "worlds", "hello"}
	got := GetWords("Hello hello world worlds hellO")
	compareSlices(want, got, t)
}

func TestGetWordsEmptyText(t *testing.T) {
	got := GetWords("")
	if got != nil {
		t.Fatalf("Expected nil, got: %v", got)
	}
}

func TestGetAlphabetSpecialCharacters(t *testing.T) {
	want := []string{"hello", "world", "this", "is", "a", "test"}
	got := GetAlphabet("Hello, World!! (this/ is @a test.]?")
	compareSlices(want, got, t)
}

func TestGetAlphabetWordsRepeated(t *testing.T) {
	want := []string{"hello", "world", "worlds"}
	got := GetAlphabet("Hello hello world worlds")
	compareSlices(want, got, t)
}

func TestGetAlphabetEmptyText(t *testing.T) {
	got := GetAlphabet("")
	if got != nil {
		t.Fatalf("Expected nil, got: %v", got)
	}
}

func TestGetAlphabetWithOcurrenceSpecialCharacters(t *testing.T) {
	want := map[string]int{"hello": 1, "world": 1}
	got := GetAlphabetWithOcurrence("(Hello) + /woRld/!")
	compareMaps(want, got, t)
}

func TestGetAlphabetWithOcurrenceWordsRepeated(t *testing.T) {
	want := map[string]int{"hello": 3, "world": 1}
	got := GetAlphabetWithOcurrence("(Hello) + [hello] /woRld/! heLLo")
	compareMaps(want, got, t)
}

func TestGetAlphabetWithOcurrenceEmptyText(t *testing.T) {
	got := GetAlphabetWithOcurrence("")
	if got != nil {
		t.Fatalf("Expected nil, got: %v", got)
	}
}
