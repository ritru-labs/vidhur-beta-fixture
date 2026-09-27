package ids

import (
	"reflect"
	"testing"
)

func TestNormalizeWhitespaceAndStableOrder(t *testing.T) {
	input := []string{"  alpha ", "beta", "alpha", "\t", " beta\n", "gamma"}
	before := append([]string(nil), input...)
	want := []string{"alpha", "beta", "gamma"}
	if got := Normalize(input); !reflect.DeepEqual(got, want) {
		t.Fatalf("Normalize() = %q, want %q", got, want)
	}
	if !reflect.DeepEqual(input, before) {
		t.Fatal("input mutated")
	}
}

func TestNormalizeAlreadyCanonical(t *testing.T) {
	want := []string{"alpha", "beta"}
	if got := Normalize([]string{"alpha", "beta", "alpha", ""}); !reflect.DeepEqual(got, want) {
		t.Fatalf("Normalize() = %q, want %q", got, want)
	}
}
