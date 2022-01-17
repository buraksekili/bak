package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {

		fi, err := os.Stat("test.txt")
		if err != nil {
			log.Fatal(err)
		}
		prevTime := fi.ModTime()
		prevSize := fi.Size()

		for range time.Tick(1 * time.Second) {
			fi, err = os.Stat("test.txt")
			if err != nil {
				log.Fatal(err)
			}

			if fi.ModTime().After(prevTime) && fi.Size() != prevSize {
				fmt.Printf("new size %v, previously it was %v\n", fi.Size(), prevSize)
				prevTime = fi.ModTime()
				prevSize = fi.Size()
			} else {
				fmt.Println("no changes")
			}
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		sig := <-c
		fmt.Println("got signal:", sig)
		wg.Done()
	}()

	wg.Wait()
}
