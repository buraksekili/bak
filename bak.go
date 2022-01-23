package bak

import (
	"log"
	"os"
	"time"
)

type Conf struct {
	// File represents file name that Watcher watches changes on it.
	File     string
	Duration time.Duration
}

type Watcher struct {
	Config Conf
}

func New(config Conf) *Watcher {
	return &Watcher{config}
}

func (w *Watcher) Watch() chan struct{} {
	ch := make(chan struct{})
	go func() {
		fi, err := os.Stat(w.Config.File)
		if err != nil {
			log.Fatal(err)
		}
		prevTime := fi.ModTime()
		prevSize := fi.Size()

		for range time.Tick(w.Config.Duration) {
			fi, err = os.Stat(w.Config.File)
			if err != nil {
				log.Fatal(err)
			}

			if fi.ModTime().After(prevTime) && fi.Size() != prevSize {
				ch <- struct{}{}
				prevTime = fi.ModTime()
				prevSize = fi.Size()
			}
		}
	}()
	return ch
}
