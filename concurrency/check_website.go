package concurrency

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	// channel containing result type
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			// concurrent writes can be sent to channel safely
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		// sequencially extract (receive) stuff inside the channel
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}
