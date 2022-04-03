package main

func main() {
	message := make(chan int)
	// for loop
	for i := 0; i < 100; i++ {
		go func(j int) {

			// time.Sleep(time.Second * 2)
			message <- j
		}(i)
	}
	for i := 0; i < 100; i++ {
		println(<-message) // but these messages wont be in same order like 1to 100.
	}

	// time.Sleep(time.Second)
}
