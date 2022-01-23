package main

import (
	"fmt"
	watcher "github.com/buraksekili/bak"
	"time"
)

func main() {
	w := watcher.New(watcher.Conf{File: "test.txt", Duration: 1 * time.Second})
	ch := w.Watch()
	for range ch {
		fmt.Println("change detected")
	}
}
