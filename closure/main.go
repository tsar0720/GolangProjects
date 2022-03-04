package main

import (
	"fmt"
)

func totalPrice(initPrice int) func(int) int {
	sum := initPrice // инициализируется
	return func(x int) int {
		sum += x //увеличивается
		return sum
	}
}

func main() {
	//замыкание - функция возвращает другую функцию которая
	//использует параметры функции родителя
	//и изменяет состояние функции родителя
	orderPrice := totalPrice(1) // 1 вернет функцию orderPrice становится функцией
	fmt.Println(orderPrice(1))  // 2
	fmt.Println(orderPrice(1))  // 3
}
