package main

import (
	"log"
	"sync"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	createGeneralDir()
	getJson()
}

func main() {
	runner()
}

func runner() {

	var wg sync.WaitGroup

	semaphore := make(chan struct{}, 1)
	for i, v := range langList {
		wg.Add(1)
		semaphore <- struct{}{}
		go func(v string, i string) {
			defer wg.Done()
			defer func() { <-semaphore }()
			parseData(v, i)
		}(v, i)
	}

	wg.Wait()
}
