package main

import "github.com/raugustin93/bible-cli/internal/bibleapi"

type config struct {
	bibleApiClient bibleapi.Client
}

func main() {
	cfg := config{
		bibleApiClient: bibleapi.NewClient(),
	}
	startRepl(&cfg)
}
