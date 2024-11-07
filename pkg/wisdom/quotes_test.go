package wisdom

import "testing"

func TestGetRandomQuote(t *testing.T) {
	quote := (WisdomImpl{}).GetRandomQuote()
	if quote == "" {
		t.Error("Expected a non-empty quote")
	}
}
