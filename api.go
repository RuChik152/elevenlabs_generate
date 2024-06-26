package main

import (
	"log"
	"sync"

	"github.com/joho/godotenv"
)

var langList map[string]string = map[string]string{
	"CHN":  "LNG_M_CHN.json",
	"ENG":  "LNG_M_ENG.json",
	"FR":   "LNG_M_FR.json",
	"ESP":  "LNG_M_ESP.json",
	"ARAB": "LNG_M_ARAB.json",
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	runner()
	//parseData("LNG_M_CHN.json")
	//runGenerate("要打开这扇门，您需要在KargaVR.com激活您朋友的傀儡 <color=orange>在您完成此教程之前</color>。\r\n\r\n在这扇门后面有一个礼物送给你们两个：傀儡的动力包。\r\n\r\n剩余级别: ", "ACTV_CAVE_LVL")
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
