package main

import (
	"fmt"
)

//коллбек изменение поведения внутри функции с помощью функции переданной как параметр
func doSomething(callback func(int, int) int, s string) int {
	result := callback(5, 1)
	fmt.Println(s)
	return result
}

func main() {
	sumCallback := func(n1, n2 int) int {
		return n1 + n2
	}
	result := doSomething(sumCallback, "plus")
	fmt.Println(result)

}
