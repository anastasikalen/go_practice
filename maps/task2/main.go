package main

import "fmt"

func main() {
	m := map[string]int{"go": 3, "python": 2, "java": 3}
	fmt.Println(InvertMap(m))
}
func InvertMap(m map[string]int) map[int][]string {
	inverted := make(map[int][]string, len(m))
	for k, v := range m {
		inverted[v] = append(inverted[v], k)
	}
	return inverted
}
