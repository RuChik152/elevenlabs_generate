package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

func parseData(path string, lang string) {

	var wg sync.WaitGroup
	var list map[string]string

	file, err := os.ReadFile(filepath.Join(".", "lang", path))
	if err != nil {
		fmt.Println("Err: ", err)
		return
	}

	err = json.Unmarshal(file, &list)
	if err != nil {
		fmt.Println("Err: ", err)
	}

	semaphore := make(chan struct{}, 2)
	for i, v := range list {

		wg.Add(1)
		semaphore <- struct{}{}

		go func(v string, i string) {
			defer wg.Done()
			defer func() { <-semaphore }()
			runGenerate(v, i, lang)
		}(v, i)
	}

	wg.Wait()
}
