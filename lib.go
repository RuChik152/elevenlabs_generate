package main

import (
	"log"
	"os"
)

var langList map[string]string

func getJson() {
	files, err := os.ReadDir("./lang")
	if err != nil {
		log.Fatal("lang file not found")
	}

	for i := 0; i < len(files); i++ {
		name := getName(files[i].Name())
		
		
	}

}

func getName(fileName string) string {
	
}
