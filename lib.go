package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"regexp"
	"strings"
)

var langList map[string]string

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
			log.Println("Соотвествует", name)

			log.Printf("%q\n", reName.FindString(name))
			log.Println(strings.ReplaceAll(reName.FindString(name), "_", ""))

			createDir(strings.ReplaceAll(reName.FindString(name), "_", ""))

		} else {
			log.Println("Не соотвествует", name)
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

	// path := path.Join(".", "error", nameDir)
	// if err := os.Mkdir(path, 0750); err != nil {
	// 	log.Fatalln(err)
	// }

	for i := 0; i < len(list); i++ {
		path := path.Join(".", list[i], nameDir)
		if err := os.Mkdir(path, 0750); err != nil {
			log.Println(err)
		}
	}
}
