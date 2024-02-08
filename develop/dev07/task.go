package dev

import "time"

func or(channels ...<-chan interface{}) <-chan interface{} {
	orChannel := make(chan interface{})

	go func() {
		defer close(orChannel)

		done := make(chan struct{})

		for _, ch := range channels {
			go func(c <-chan interface{}) {
				select {
				case <-c:
					close(done)
				case <-done:
					return
				}
			}(ch)
		}

		<-done
	}()

	return orChannel
}

func sig(after time.Duration) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
	}()
	return c
}
