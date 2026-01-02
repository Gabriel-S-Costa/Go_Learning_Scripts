package main

import (
	"fmt"
	"time"
)

func say(s string, done chan string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
	done <- "Terminei"
}

func main() {
	// start := time.Now()
	// fmt.Println("Hello World!")
	// secs := time.Since(start).Seconds()
	// fmt.Printf("%.2fs \n", secs)
	done := make(chan string)
	go say("world", done)
	go say("hello", done)
	fmt.Println(<-done)
}
