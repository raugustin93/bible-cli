package main

type config struct {
	book string
}

func main() {
	cfg := config{
		book: "nv",
	}
	startRepl(&cfg)
}
