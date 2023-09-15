package main

import (
	"context"
	"log"
	"strconv"
	"time"
)

func main() {
	rootCtx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(rootCtx, 3*time.Second)
	defer cancel()
	done := make(chan bool)
	items := []string{}
	for i := 0; i < 10; i++ {
		items = append(items, strconv.FormatInt(int64(i), 10))
	}

	log.Println("[main] Start")
	go processData(timeoutCtx, done, items)

	select {
	case <-timeoutCtx.Done():
		log.Println("[main] " + timeoutCtx.Err().Error())
	case <-done:
		log.Println("[main] Success!")
	}
}

func processData(ctx context.Context, done chan<- bool, items []string) {
	for i := range items {
		select {
		case <-ctx.Done():
			log.Println("[processData] Abort processing data because context is canceled.")
			return
		default:
			log.Printf("[processData] Processing items[%d]\n", i)
			time.Sleep(1 * time.Second)
		}
	}

	done <- true
}
