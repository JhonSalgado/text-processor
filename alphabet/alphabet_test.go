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

func TestGetAlphabetLowercaseText(t *testing.T) {
	want := []string{"hello", "world"}
	got := GetAlphabet("hello world")
	compareSlices(want, got, t)
}

func TestGetAlphabetUppercaseText(t *testing.T) {
	want := []string{"hello", "world"}
	got := GetAlphabet("HELLO World")
	compareSlices(want, got, t)
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

func TestGetAlphabetWithOcurrenceNormalText(t *testing.T) {
	want := map[string]int{"hello": 2, "world": 1}
	got := GetAlphabetWithOcurrence("Hello hello world!")
	compareMaps(want, got, t)
}

func TestGetAlphabetWithOcurrenceEmptyText(t *testing.T) {
	got := GetAlphabetWithOcurrence("")
	if got != nil {
		t.Fatalf("Expected nil, got: %v", got)
	}
}
