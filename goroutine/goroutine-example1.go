package goroutine

import (
	"fmt"
	"time"
)

func bufferedChannel() {
	charChannel := make(chan string, 3)
	chars := []string{"a", "b", "c"}

	for _, s := range chars {
		// select {
		// case charChannel <- s:
		// }
		charChannel <- s

	}

	close(charChannel)

	for result := range charChannel {
		fmt.Println(result)
	}

}

func RunGoroutineExample1() {
	bufferedChannel()
	time.Sleep(time.Second * 2)
	fmt.Println("hi")
}
