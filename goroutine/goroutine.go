package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println(phrase)
	doneChan <- true
}
func slowGreeting(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hell0 ! ", phrase)
	doneChan <- true // data flow direction mark as an arrow like symbol
	close(doneChan)
}
func main() {
	// done := make([]chan bool, 4)
	done := make(chan bool) // creating a channel
	// dones[0] = make(chan bool)
	go greet("Hi Go!...", done)
	// dones[1] = make(chan bool)
	go slowGreeting("Hello !, will come after 3 seconds...", done)
	// dones[2] = make(chan bool)
	go greet("Hi Go! .... again am here !!", done)
	// fmt.Println(<-done)
	// for _, done := range dones {
	// 	<-done
	// }

	for doneChain := range done {
		fmt.Println(doneChain)
	}

}
