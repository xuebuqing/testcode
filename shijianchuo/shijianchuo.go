package main

import (
	"fmt"
	"time"
)

func main() {

	TimeStamp()
}
func TimeStamp() {
	now := time.Now().Unix()
	tim := now / 100
	fmt.Println(tim)
}
