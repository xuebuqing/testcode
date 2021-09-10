package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	for i := 0; i < 100; i++ {
		time.Sleep(time.Second)
		fmt.Println(time.Since(start))
	}
}
