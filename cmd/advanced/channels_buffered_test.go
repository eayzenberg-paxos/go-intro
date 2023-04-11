package advanced

import (
	"fmt"
	"testing"
	"time"
)

func TestChannelsBuffered(t *testing.T) {

	/* In this example, we create a buffered channel ch that can hold up to 2 integers.
	We then send the values 1 and 2 on the channel. Finally, we receive the values
	from the channel and print them to the console.

	Note that in this example, we do not need to use a separate goroutine to send
	values on the channel, because the channel is buffered and can hold multiple
	values at once. If the channel were not buffered, we would need to use a
	separate goroutine to send values on the channel in order to prevent blocking */

	var ch chan int = make(chan int, 2) // notice type 'chan int' can be buffered or unbuffered
	ch <- 1
	ch <- 2
	close(ch) // closing the channel will prevent further writing but we can still read

	fmt.Println(fmt.Sprintf("The received values are %v and %v", <-ch, <-ch))
}

func TestChannelBufferedWithRange(t *testing.T) {

	var ch = make(chan int, 2) // notice type 'chan int' can be buffered or unbuffered
	ch <- 1
	ch <- 2

	/* In Go, the range keyword can be used with a channel to iterate over the
	values sent on the channel. When used with a buffered channel, range ch will
	iterate over the buffered values in the channel, as well as any future values
	that are sent on the channel until it is closed. */

	go func() {
		for val := range ch { // however range operator is blocking
			fmt.Println(val)
		}
	}()

	ch <- 3 // oops, missed one. The prev loop will take care of it

	time.Sleep(10 * time.Millisecond)
}
