package main

import "fmt"

func main() {
	words := []string{"go", "python", "go", "java", "go", "python"}
	fmt.Println(CountWords(words))
}

func CountWords(words []string) map[string]int { //подсчитываем в мапе кол-во слов из массива строк
	m := make(map[string]int)
	for _, word := range words {
		m[word]++
	}
	return m
}
