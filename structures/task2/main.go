package main

import "fmt"

func main() {
	c1 := Counter{
		1,
	}
	c1.Increment()
	fmt.Println(c1)
	c1.Increment()
	fmt.Println(c1)
}

type Counter struct {
	Value int
} //КОД ДНЯ 9426

func (c *Counter) Increment() {
	c.Value++
}
