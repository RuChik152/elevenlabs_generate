package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

func parseData(path string) {

	var wg sync.WaitGroup
	var list map[string]string

	file, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Err: ", err)
		return
	}

	err = json.Unmarshal(file, &list)
	if err != nil {
		fmt.Println("Err: ", err)
	}

	//fmt.Println(list)

	semaphore := make(chan struct{}, 2)
	for i, v := range list {
		//fmt.Printf("Index: %s - Value: %s\n", i, v)
		wg.Add(1)
		semaphore <- struct{}{}

		go func(v string, i string) {
			defer wg.Done()
			defer func() { <-semaphore }()
			runGenerate(v, i)
		}(v, i)
	}

	wg.Wait()
}
