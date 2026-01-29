package main

import "fmt"

func main() {
	var worker []Worker
	em := Employee{Person{"test", 12}, 1000}
	m := Manager{Person{"m", 33}, 2000, 10}
	worker = append(worker, em)
	worker = append(worker, m)
	for _, w := range worker {
		fmt.Println(w.GetInfo())
		fmt.Println(w.GetSalary())
	}
} //Код дня: 1074

type Worker interface {
	GetSalary() float64
	GetInfo() string
}

type Person struct {
	Name string
	Age  int
}
type Employee struct {
	Person
	Salary float64
}

func (e Employee) GetSalary() float64 {
	return e.Salary
}
func (e Employee) GetInfo() string {
	return e.Name
}

type Manager struct {
	Person
	Salary float64
	Bonus  float64
}

func (m Manager) GetSalary() float64 {
	return m.Salary + m.Bonus
}
func (m Manager) GetInfo() string {
	return m.Name
}
