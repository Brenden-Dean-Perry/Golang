package concurrnecy

import "fmt"

// Channels and how they work with goroutines
// Channels (1) Hold data, (2) Thread safe, (3) Listen when data is added or removed

// Channel with goroutine
func Channel_with_goroutine() {
	ch := make(chan int) // Create a channel of type int and unbuffered (synchronous).
	// Unbuffered channels require both a sender and a receiver to be ready at the same time.
	// This means that when you send data to an unbuffered channel, the sending goroutine will block until another goroutine is ready to receive the data, and vice versa.
	// Said another way, one one element can be in the channel at a time as it is passed directly from sender to receiver.
	go process(ch)
	for i := range ch {
		println("Received from channel:", i) // Receive data from channel
	}
}

// Channel with goroutine and buffered
func Channel_with_goroutine_buffered() {
	ch := make(chan int, 5) // Create a channel of type int and buffered (asynchronous) with a capacity of 5.
	// Buffered channels allow sending and receiving to occur independently.
	// The sender can send up to the buffer capacity without blocking, and the receiver can receive from the buffer without blocking as long as there are items in it.
	// This can help improve performance in certain scenarios where the producer and consumer operate at different speeds.
	go process(ch)
	for i := range ch {
		println("Received from channel:", i) // Receive data from channel
	}
}

func process(c chan int) {
	for i := 0; i < 5; i++ {
		c <- i
	}
	fmt.Println("Exiting process function")
	close(c) // Close the channel when done sending data.
	// This is important to avoid deadlocks.
	// Closing a channel notifies receivers that no more values will be sent on it.
}
