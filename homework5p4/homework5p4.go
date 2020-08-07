package main

import (
	"io/ioutil"
	"log"
	"os"
)

// Программа копирует файлы, которые существуют в той же директории
// что и сама программа.
// Принцип работы следующий:
// для копирования файла - запустите программу homeworks5p4, указав после
// названия программы, через пробел, название файла, который необходимо скопировать
// и (так же через пробел) название нового файла в который скопируются данные.
// примеры запуска:
// (1)go run homeworks5p4.go somefilename.extension newfilename.extension
// если программа существует в бинарнике:
// (2)./homeworks5p4 somefilename.extension newfilename.extension
// Для копирования самой программы homeworks5p4.go, сначала необходимо выполнить
// go build homeworks5p4.go, затем можно копировать программу, выполнив в терминале:
// ./homeworks5p4 homeworks5p4.go somefile.extension

func getPath() string {
	pathFile, err := os.Getwd()
	if err != nil {
		log.Fatal()
	}
	return pathFile
}

func searchFileNamesInDir(pathFile string) (files []os.FileInfo) {
	files, err := ioutil.ReadDir(pathFile)
	if err != nil {
		log.Fatal()
	}
	return files
}

func compareAndCreateNewFile(files []os.FileInfo, flagInfo []string, pathFile string) {
	for _, file := range files {
		if file.Name() == flagInfo[1] {
			data, _ := ioutil.ReadFile(file.Name())
			createNewFile(pathFile, flagInfo, data)
		}
	}
}

func createNewFile(pathFile string, flagInfo []string, data []byte) {
	newFileCreate, err := os.Create(pathFile + "/" + flagInfo[2])
	if err != nil {
		log.Fatal()
	}
	err = ioutil.WriteFile(newFileCreate.Name(), data, 0666)
	if err != nil {
		log.Fatal()
	}
	defer newFileCreate.Close()
}

func main() {
	flagInfo := os.Args
	pathFile := getPath()
	files := searchFileNamesInDir(pathFile)
	compareAndCreateNewFile(files, flagInfo, pathFile)
}
