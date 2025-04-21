package main

import "sync"

type post struct {
	views int
	mu    sync.Mutex
}

func (p *post) inc(wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
		p.mu.Unlock()
	}()

	p.mu.Lock()
	p.views += 1
}

func main() {
	var wg sync.WaitGroup

	myPost := post{
		views: 0,
	}

	for range 100 {
		wg.Add(1)
		go myPost.inc(&wg)
	}

	wg.Wait()
	println(myPost.views)
}
