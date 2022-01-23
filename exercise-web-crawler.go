package main

import (
	"fmt"
	"sync"
)


type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

type Cache struct {
	v   map[string]int
	mu sync.Mutex
}

func (c Cache) hasURL(url string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, ok := c.v[url]
	
	if ok == false { return false
	} else { return true }
}

func (c Cache) addURLToCache(url string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.v[url] = 1
}


// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, wg *sync.WaitGroup, c *Cache ) {
	defer wg.Done()
	if depth <= 0 {
		return
	}
	if c.hasURL(url) {
		return
	}
	
	c.addURLToCache(url)
	body, urls, err := fetcher.Fetch(url)
	
	if err != nil {
		fmt.Println(err)
		return
	}
	
	fmt.Printf("found: %s %q\n", url, body)
	
	for _, u := range urls {
		wg.Add(1)
		
		go Crawl(u, depth-1, fetcher, wg, c)
	}
}

func main() {
	c := Cache{v: make(map[string]int)}
	var wg sync.WaitGroup
	
	wg.Add(1)
	go Crawl("http://golang.org/", 4, fetcher, &wg, &c)
	wg.Wait()
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}