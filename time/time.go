package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hello world")

	now := time.Now().Format("03:04:05PM 01-02-06")

	fmt.Println(now)
}
