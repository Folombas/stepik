package main

import (
	"fmt"
	"strconv"
)

// calculateAverage вфчисляет среднее арифметическое трёх чисел
func calculateAverage(a, b, c float64) float64 {
	return (a + b + c) / 3.0
}

func main() {
	var input1, input2, input3 string

	// Читаем три строки из stdin - без вывода приглашений!
	fmt.Scanln(&input1)
	fmt.Scanln(&input2)
	fmt.Scanln(&input3)

	// Преобразуем строки в числа
	num1, _ := strconv.ParseFloat(input1, 64)
	num2, _ := strconv.ParseFloat(input2, 64)
	num3, _ := strconv.ParseFloat(input3, 64)

	// Вычисляем среднее
	average := calculateAverage(num1, num2, num3)

	// Выводим только результат — как требует тест
	fmt.Printf("Среднее арифметическое: %.2f\n", average)
}
