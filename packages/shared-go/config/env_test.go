package config

import "testing"

func TestStringReturnsFallbackWhenUnset(t *testing.T) {
	t.Setenv("SHARED_GO_TEST_STRING", "")

	got := String("SHARED_GO_TEST_STRING", "fallback")
	if got != "fallback" {
		t.Fatalf("expected fallback, got %q", got)
	}
}

func TestStringListTrimsValues(t *testing.T) {
	t.Setenv("SHARED_GO_TEST_LIST", "alpha, beta ,,gamma")

	got := StringList("SHARED_GO_TEST_LIST", nil)
	want := []string{"alpha", "beta", "gamma"}
	if len(got) != len(want) {
		t.Fatalf("expected %d values, got %d", len(want), len(got))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("expected value %d to be %q, got %q", i, want[i], got[i])
		}
	}
}

func TestPositiveIntRejectsNonPositiveValue(t *testing.T) {
	t.Setenv("SHARED_GO_TEST_INT", "0")

	if _, err := PositiveInt("SHARED_GO_TEST_INT", 1); err == nil {
		t.Fatal("expected error for non-positive integer")
	}
}
