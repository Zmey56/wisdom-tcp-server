package wisdom

import (
	"math/rand"
	"os"
)

type Wisdom interface {
	GetRandomQuote() string
}

type WisdomImpl struct{}

var quotes = []string{
	"Stay focused and keep shipping.",
	"Success is not final, failure is not fatal.",
	"Keep pushing your limits.",
}

func (w WisdomImpl) GetRandomQuote() string {
	if mockQuote := os.Getenv("MOCK_WISDOM_QUOTE"); mockQuote != "" {
		return mockQuote
	}
	return quotes[rand.Intn(len(quotes))]
}
