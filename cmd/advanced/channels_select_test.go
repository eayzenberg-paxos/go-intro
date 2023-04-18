package advanced

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestBasicSelect(t *testing.T) {

	/* The select statement waits until one of the communication operations is ready.
	If multiple channels are ready at the same time, one is chosen at random. The
	selected case then executes, and the program continues to the next statement
	after the select block.

	In the example, we have four cases in the select block:

	- The first two cases wait for a message to be received on channel1 and channel2,
	respectively. When a message is received on one of these channels, the
	corresponding case executes.
	- The third case waits for a message to be received on channel3. When a message is
	received, it is stored in the message variable and the case executes.
	- The fourth case sends a message on channel4. If the channel is not ready to
	receive the message, the select statement will block until it is. */

	channel1 := make(chan string)
	channel2 := make(chan string)
	channel3 := make(chan string)
	channel4 := make(chan string)

	go func() {
		channel3 <- "hello"
	}()

	var message string

	select {
	case <-channel1:
		// Handle a message from channel1
	case <-channel2:
		// Handle a message from channel2
	case message = <-channel3:
		// Handle a message from channel3
	case channel4 <- message:
		// Send a message on channel4
	}

	assert.Equal(t, "hello", message)
}

func TestSelectDefault(t *testing.T) {

	/* The default case executes if none of the other cases are ready. This is useful
	for handling the case where no communication is ready, or for implementing a
	timeout mechanism. */

	channel1 := make(chan string) // We'll never put anything in channel1
	var message string

	select {

	case message = <-channel1:

	default:
		message = "it was defaulted"
	}

	assert.Equal(t, "it was defaulted", message)
}

func TestSelectInFor(t *testing.T) {

	/* In this example, we define a sendMessages function that sends messages on two
	channels, c1 and c2, with a short delay between each message.

	We then use a for loop with a select statement to wait for messages on both
	channels. When a message is received on one of the channels, the corresponding
	case in the select statement executes and prints the expected text */

	sendMessages := func(c1, c2 chan string) {
		for i := 0; i < 5; i++ {
			c1 <- fmt.Sprintf("Message %d", i)
			time.Sleep(time.Millisecond * 100)
		}
		close(c1)
		for i := 0; i < 5; i++ {
			c2 <- fmt.Sprintf("Message %d", i)
			time.Sleep(time.Millisecond * 100)
		}
		close(c2)
	}

	c1 := make(chan string)
	c2 := make(chan string)

	go sendMessages(c1, c2)

	openC1, openC2 := true, true
	for {
		select {

		/* Since we are also evaluating if the channels are open, both first branches
		are eligible to run. Go chooses pseudo randomly in this case and eventually both
		branches are executed */

		case msg1, ok := <-c1:
			openC1 = ok
			if ok {
				fmt.Println("Received message from c1:", msg1)
			} else {
				fmt.Println("c1 channel closed")
			}
		case msg2, ok := <-c2:
			openC2 = ok
			if ok {
				fmt.Println("Received message from c2:", msg2)
			} else {
				fmt.Println("c2 channel closed")
			}
		default:
			t.Failed() // Never executed
		}

		// Exit loop when both channels are closed
		if !openC1 && !openC2 {
			break
		}
	}
}

func TestSelectDefaultWithSleep(t *testing.T) {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)

	ticked := 0
	boomed := 0

	for boomed == 0 {
		select {
		case <-tick:
			fmt.Println("tick.")
			ticked++
		case <-boom:
			fmt.Println("BOOM!")
			boomed++
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}

	assert.Equal(t, 4, ticked)
	assert.Equal(t, 1, boomed)
}
