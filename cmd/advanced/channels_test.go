package advanced

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChannels(t *testing.T) {

	/* In this example, we create a simple channel ch that can hold integers. We then
	start a new goroutine that sends the value 42 on the channel. Finally, we
	receive the value from the channel and print it to the console. */

	ch := make(chan int)

	go func() { // creates a non OS level thread called 'goroutine'
		ch <- 42
	}()

	assert.Equal(t, 42, <-ch)

	close(ch) // prevent new usages in channel

	_, ok := <-ch // However we can still check if it's closed. Careful, we are still reading
	if !ok {
		fmt.Println("Channel is closed")
	} else {
		fmt.Println("Channel is open")
		t.Failed()
	}
}
