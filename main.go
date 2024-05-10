package main

// func main() {
// 	// call display function with go routine
// 	go display(1)
// 	time.Sleep(1 * time.Second)

// 	// call display function without go routine
// 	display(2)
// }

// func display(number int) {
// 	// Display the number
// 	fmt.Println(number)
// }

// func main() {
// 	fmt.Println("Start")

// 	go func() {
// 		fmt.Println("Middle")
// 	}()

// 	time.Sleep(1 * time.Second)
// 	fmt.Println("End")
// }

// func main() {
// 	fmt.Println("Start")
// 	// Create a channel
// 	ch := make(chan int)

// 	// Send data to channel
// 	go func() {
// 		ch <- 10
// 	}()

// 	// Receive data from channel
// 	value := <-ch
// 	fmt.Println("Value: ", value)

// 	fmt.Println("End")
// }

// func main() {
// 	fmt.Println("Start")
// 	// Create a channel
// 	ch := make(chan int)

// 	// Send data to channel
// 	go func() {
// 		ch <- 10
// 		close(ch)
// 	}()

// 	// Receive data from channel
// 	for {
// 		value, ok := <-ch
// 		if !ok {
// 			fmt.Println("Channel closed")
// 			break
// 		}

// 		fmt.Println("Value: ", value)
// 	}

// 	fmt.Println("End")
// }

// func main() {
// 	fmt.Println("Start")
// 	// Create a channel
// 	ch := make(chan int)

// 	// Send data to channel
// 	go func() {
// 		for i := 0; i < 5; i++ {
// 			ch <- i
// 		}
// 		close(ch)
// 	}()

// 	// Receive data from channel
// 	for value := range ch {
// 		fmt.Println("Value: ", value)
// 	}

// 	fmt.Println("End")
// }

// func main() {
// 	fmt.Println("Start")
// 	// Create a channel
// 	ch := make(chan int, 3)

// 	// Send data to channel
// 	ch <- 10
// 	ch <- 20
// 	ch <- 30

// 	// Receive data from channel
// 	fmt.Println("Length of channel: ", len(ch))
// 	fmt.Println("Capacity of channel: ", cap(ch))

// 	fmt.Println("End")
// }

// func main() {
// 	fmt.Println("Start")
// 	// Create a channel
// 	ch1 := make(chan string)
// 	ch2 := make(chan string)

// 	// Send data to channel
// 	go func() {
// 		time.Sleep(1 * time.Second)
// 		ch1 <- "From channel 1"
// 	}()

// 	go func() {
// 		time.Sleep(2 * time.Second)
// 		ch2 <- "From channel 2"
// 	}()

// 	// Receive data from channel
// 	select {
// 	case value := <-ch1:
// 		fmt.Println("Value from channel 1: ", value)
// 	case value := <-ch2:
// 		fmt.Println("Value from channel 2: ", value)
// 	}

// 	fmt.Println("End")
// }

// func main() {
// 	// creating channel
// 	mychannel := make(chan int)
// 	select {
// 	case <-mychannel:

// 	default:
// 		fmt.Println("Not found")
// 	}
// }

// func main() {
// 	// creating channel
// 	ch1 := make(chan string)
// 	ch2 := make(chan string)

// 	// send data to channel
// 	go func() {
// 		for i := 0; i < 3; i++ {
// 			ch1 <- "From channel 1"
// 		}
// 	}()

// 	go func() {
// 		ch2 <- "From channel 2"
// 	}()

// 	select {
// 	case value := <-ch1:
// 		fmt.Println("Value from channel 1: ", value)
// 	case value := <-ch2:
// 		fmt.Println("Value from channel 2: ", value)
// 	}
// }

// func main() {
// 	fmt.Println("Start")
// 	// Create a WaitGroup
// 	var wg sync.WaitGroup

// 	// Add 1 to the WaitGroup
// 	wg.Add(1)

// 	// call display function with go routine
// 	go func() {
// 		// call Done function on WaitGroup
// 		defer wg.Done()
// 		fmt.Println(1)
// 	}()

// 	// Wait until the WaitGroup counter goes to 0
// 	wg.Wait()

// 	fmt.Println("End")
// }
