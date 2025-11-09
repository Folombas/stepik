package main

import "fmt"

func main() {
	// Создаём слайс длиной 10
	numbers := make([]int, 10)

	// Заполняем числами от 1 до 10
	for i := 0; i < 10; i++ {
		numbers[i] = i + 1
	}

	// Выводим в требуемом формате
	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))
}
