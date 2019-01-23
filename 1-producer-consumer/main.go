//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer szenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"fmt"
	"time"
)

func producer(out chan<- Tweet, stream Stream) {
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			close(out)
			return
		}
		out <- *tweet
	}
}

func consumer(in <-chan Tweet) {
	for t := range in {
		if t.IsTalkingAboutGo() {
			fmt.Println(t.Username, "\ttweets about golang")
		} else {
			fmt.Println(t.Username, "\tdoes not tweet about golang")
		}
	}
}

func main() {
	start := time.Now()

	stream := GetMockStream()

	ch := make(chan Tweet, 100)
	// Producer
	go producer(ch, stream)
	// Consumer
	consumer(ch)
	fmt.Printf("Process took %s\n", time.Since(start))
}
