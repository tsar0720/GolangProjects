package main

import "fmt"

type Point struct {
	X int
	Y int
}

func main() {
	//инициализируем мапу {} - инициализация
	pointsMap := map[string]Point{}

	//Инициализация с помощью явного указания
	pointsMap["word"] = Point{
		X: 1,
		Y: 2,
	}
	fmt.Println(pointsMap["word"])
	fmt.Println(pointsMap["word"])

	//аналог инициализации мапы через make()
	otherPointsMap := make(map[string]Point)

	// можно проинициализировать  Point{1, 2} (вроде позиционных аргументов)
	otherPointsMap["hello"] = Point{1, 2}
	fmt.Println(otherPointsMap["hello"])

	//инициализация сразу ключом и значением
	someMap := map[string]Point{
		"b": {
			X: 1,
			Y: 2,
		},
	}

	fmt.Println(someMap)

	//объявили но не проинициализировали мапу
	var oneMorePointsMap map[string]Point

	//проверяем на nil и инициализируем
	if oneMorePointsMap == nil {
		oneMorePointsMap = map[string]Point{
			"a": {12, 21},
		}
	}
	fmt.Println(oneMorePointsMap["a"])
	fmt.Println(oneMorePointsMap["a"].X)
	fmt.Println(oneMorePointsMap["a"].Y)
	oneMorePointsMap["a"].method()

	//Проверка того что в мапе есть значение
	value, ok := oneMorePointsMap["a"]
	if ok {
		fmt.Println(value)
		fmt.Println("Такой ключ есть в мапе")
	}
}

func (p Point) method() {
	fmt.Println("call Point method")
	fmt.Println(p)
}
