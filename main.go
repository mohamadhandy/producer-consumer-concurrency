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

func consumer(ch <-chan *Tweet, ch1 chan<- int) {
	for v := range ch {
		if v.IsTalkingAboutGo() {
			fmt.Println(v.Username, "\ttweets about golang")
		} else {
			fmt.Println(v.Username, "\tdoes not tweet about golang")
		}
	}
	ch1 <- -1
	fmt.Println(ch1)
}

func producer(ch chan<- *Tweet, stream Stream) {
	for {
		tweet, err := stream.Next()
		if err == nil {
			ch <- tweet
		} else {
			close(ch)
		}
	}
}

func main() {

	start := time.Now()
	stream := GetMockStream()

	pChannel := make(chan *Tweet)
	cChannel := make(chan int)

	go producer(pChannel, stream)
	go consumer(pChannel, cChannel)
	<-cChannel

	fmt.Printf("Process took %s\n", time.Since(start))

}
