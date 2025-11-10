package main

import "fmt"

func main() {
	// Создаём слайс чисел от 1 до 5
	numbers := []int{1, 2, 3, 4, 5}

	fmt.Printf("Исходный слайс: %v\n", numbers)

	// Удаляем элемент с индексом 2 (число 3)
	numbers = append(numbers[:2], numbers[3:]...)

	fmt.Printf("Обновлённый слайс: %v\n", numbers)
}
