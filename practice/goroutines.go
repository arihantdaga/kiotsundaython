package main

import (
	"time"
)

// Running goroutines on a shared variable does not ensure atomicity.
func main() {
	count := 0
	for i := 0; i < 1000; i++ {
		go func() {
			count++
			// fmt.Println(count)
		}()
	}
	time.Sleep(time.Second * 2)
	println(count) // this count will not be 1000 neccesarily. it can be 950/990 etc.
}
