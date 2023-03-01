package main

import "flag"

func main()  {
	// token = flags.Get(token)

	// tgClient = telegram.New(token)

	// fetcher = fetcher.New()
	// processor = processor.New()

	// consumer.Start(fetcher, processor)
}

func mustToken() string {
	token := flag.String(
		"token-bot-token", 
		"", 
		"token for access")

	flag.Parse()

	if *token == "" {
		log.Fatal()
	}
}