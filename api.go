package main

import (
	"log"

	"github.com/joho/godotenv"
)

//  = map[string]string{
// 	//"CHN":  "resource_map_CHN.json",
// 	//"ENG":  "resource_map_ENG.json",
// 	//"FR":   "resource_map_FR.json",
// 	//"ESP":  "resource_map_ESP.json",
// 	//"ARAB": "resource_map_ARAB.json",
// 	// "CHN":  "new_com_CHN.json",
// 	// "ENG":  "new_com_ENG.json",
// 	// "FR":   "new_com_FR.json",
// 	// "ESP":  "new_com_ESP.json",
// 	// "ARAB": "new_com_ARAB.json",
// }

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	createGeneralDir()
	getJson()

}

func main() {
	// runner()
}

// func runner() {

// 	var wg sync.WaitGroup

// 	semaphore := make(chan struct{}, 1)
// 	for i, v := range langList {
// 		wg.Add(1)
// 		semaphore <- struct{}{}
// 		go func(v string, i string) {
// 			defer wg.Done()
// 			defer func() { <-semaphore }()
// 			parseData(v, i)
// 		}(v, i)
// 	}

// 	wg.Wait()
// }
