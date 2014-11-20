package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	sayHello()
	os.Exit(0)
}

func sayHello() {
	c := make(chan string)
	go func() {
		fmt.Println(<-c)
		time.Sleep(200 * time.Millisecond)
		c <- "SO LONG!"
	}()
	c <- "OH HAI ʕ•͡ᴥ•ʔ"
	fmt.Println(<-c)
}
