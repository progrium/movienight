package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println(time.Now())
		<-time.After(2 * time.Second)
	}
}
