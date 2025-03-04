package main

import (
	"flag"
	"fmt"
	"time"
)

var timeoutDurartion = flag.Duration("timeout.duration", 5, "HTTP timeout in second")

func main() {
	flag.Parse()
	fmt.Println("Hello...!")
	t1 := time.Now()
	time.Sleep(*timeoutDurartion)
	fmt.Println("Time taken ", time.Since(t1))
}
