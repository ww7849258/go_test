package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hello Horus")
	now := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(now[:len("2006-01-02")])
	fmt.Println(now[len("2006-01-02 "):])
}
