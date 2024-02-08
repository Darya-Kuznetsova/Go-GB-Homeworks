package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("Введите предложение")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	// Вставьте ваш код здесь

	text = strings.ToLower(text)
	fmt.Println(text)
	letterCounter := make(map[rune]int)
	var length int
	for _, v := range text {
		if unicode.IsLetter(v) {
			letterCounter[v]++
			length++
		}
	}

	for key, v := range letterCounter {
		fmt.Printf("%c: %v %0.2f \n", key, v, float64(v)/float64(length))
	}
}
