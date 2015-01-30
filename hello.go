package main

import (
	"fmt"
	"os"
	"time"

	"github.com/travis-repos/go-test-staging/a"
	"github.com/travis-repos/go-test-staging/a/b"
	"github.com/travis-repos/go-test-staging/a/b/c"
)

func main() {
	sayHello()
	fmt.Printf("%v\n", a.Apple)
	fmt.Printf("%v\n", b.Banana)
	fmt.Printf("%v\n", c.Carrot)
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
