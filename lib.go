package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"regexp"
	"strings"
)

var langList = make(map[string]string, 0)

func getJson() {
	files, err := os.ReadDir("./lang")
	if err != nil {
		log.Fatal("lang file not found")
	}

	re := regexp.MustCompile(`^.*_.*_\.json$`)
	reName := regexp.MustCompile(`_.*_`)

	for i := 0; i < len(files); i++ {
		name := files[i].Name()

		if re.MatchString(name) {

			langName := strings.ReplaceAll(reName.FindString(name), "_", "")
			createDir(langName)

			log.Printf("file %s defined for %s language\n", langName, name)

			langList[langName] = name

		}
	}

}

func createGeneralDir() {
	_, errEr := os.Stat("error")
	if errEr != nil && os.IsNotExist(errEr) {
		if err := os.Mkdir("error", 0750); err != nil {
			fmt.Println(err)
		}
	}

	_, errOut := os.Stat("output")
	if errOut != nil && os.IsNotExist(errOut) {
		if err := os.Mkdir("output", 0750); err != nil {
			fmt.Println(err)
		}
	}
}

func createDir(nameDir string) {

	list := []string{
		"error",
		"output",
	}

	for i := 0; i < len(list); i++ {
		path := path.Join(".", list[i], nameDir)
		if err := os.Mkdir(path, 0750); err != nil {
			log.Println(err)
		}
	}
}
