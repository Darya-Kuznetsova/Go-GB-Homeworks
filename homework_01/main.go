package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Укажите полный путь до файла вторым аргументом")

	}

	filePth := os.Args[1]

	var fileName, fileExt string

	// Напишите код, который выведет следующее
	// filename: <name>
	// extension: <extension>

	// Подсказка. Возможно вам понадобится функция strings.LastIndex
	// Для проверки своего решения используйте функции filepath.Base() filepath.Ext(
	// ) Они могут помочь для проверки решения

	// ./work/home/gb.exe

	solidusIndex := strings.LastIndex(filePth, "/")
	fullFile := filePth[solidusIndex+1:]
	fileName = fullFile
	for _, v := range fullFile {
		if v == '.' {
			dotIndex := strings.IndexRune(fullFile, v)
			fileName = fullFile[:dotIndex]
			fileExt = fullFile[dotIndex+1:]
		}
	}

	fmt.Printf("filename: %s\n", fileName)
	fmt.Printf("extension: %s\n", fileExt)

}
