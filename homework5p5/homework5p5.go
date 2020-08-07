package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var count int = 0

func searchDirAndPrintAllInFiles(pathToFiles string, search []string) {
	// получаем список файлов относительно текущей директории
	files, err := ioutil.ReadDir(pathToFiles)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() == true {
			searchDirAndPrintAllInFiles(pathToFiles+"/"+file.Name(), search)
		} else {
			data, _ := ioutil.ReadFile(pathToFiles + "/" + file.Name())
			for i := 0; i < len(data); i++ {
				if string(data[i:i+len(search[1])]) == search[1] {
					// fmt.Printf("Содержание файла %v:\n%v\n\n", file.Name(), string(data))
					count++
				}
			}
		}
	}
}

func main() {
	search := os.Args
	fmt.Println("print args:", search[1])
	fmt.Println("print args:", len(search[1]))
	// путь текущей директории
	pathToFiles, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	searchDirAndPrintAllInFiles(pathToFiles, search)
	fmt.Println("Количество совпадений:", count)
}
