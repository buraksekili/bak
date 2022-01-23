package main

import (
	"fmt"
	"time"

	"github.com/buraksekili/bak"
)

func main() {
	w := bak.New(bak.Conf{File: "test.txt", Duration: 1 * time.Second})
	ch := w.Watch()
	for range ch {
		fmt.Println("change detected")
	}
}
