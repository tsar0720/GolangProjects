package main

import "fmt"

//Структура представляющая подписчика включает в себя другую структуру
type Subscriber struct {
	Name        string
	Rate        float64
	Active      bool
	HomeAddress Address
}

//Структура представляющая адрес
type Address struct {
	Street   string
	City     string
	Building string
}

// printInfo печатает информацию о подписчике
func printInfo(s Subscriber) {
	fmt.Println("Name: ", s.Name, "\nRate:", s.Rate, "\nActive: ", s.Active)
	fmt.Println("Home Adress: ", s.HomeAddress)
}

//defaultSubscriber принимает 1 аргумент (имя) и возвращает структуру подписчика с дефолтовыми значениями
func defaultSubscriber(name string) Subscriber {
	var s Subscriber
	s.Name = name
	s.Rate = 11.1
	s.Active = true
	s.HomeAddress.City = "London"

	return s
}

// applyDiscount получает 1 аргумент указатель на подписчика и меняет значение переменной в структуре
func applyDiscount(s *Subscriber) {
	s.Rate = 4.99
}

func main() {

	//инициализация структуры адреса
	address := Address{
		Street:   "Polevaya",
		City:     "Samara",
		Building: "1",
	}

	//Создаем дефолтового подписчика
	subscriber1 := defaultSubscriber("Igor")

	//Присваиваем структуру которую создали ранее в структуру  подписчика
	subscriber1.HomeAddress = address

	//Меняем значение структуры передав ее ссылку
	applyDiscount(&subscriber1)

	//Печатаем структуру
	printInfo(subscriber1)
}
