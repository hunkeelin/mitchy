package main

import (
	"flag"
	"fmt"
	"time"
)

var (
	count = flag.Int("n", 5, "The \"n\" specify in the assignemnt")
)

func dowork(n int) {
	for i := 0; i < n; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}

func main() {
	flag.Parse()
	dowork(*count)
}
