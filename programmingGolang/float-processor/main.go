package main

import (
	"fmt"
	"math"
)

func main() {
	var num float64
	fmt.Scan(&num)

	if num <= 0 {
		// Число <= 0: выводим с 2 знаками после запятой
		fmt.Printf("число %.2f не подходит", num)
	} else if num > 10000 {
		// Число > 10000: экспоненциальное представление
		fmt.Printf("%e", num)
	} else {
		// 0 < число <= 10000: возводим в квадрат и обрезаем до 4 знаков
		squared := num * num
		// Обрезаем дробную часть (не округляем!)
		truncated := math.Trunc(squared*10000) / 10000
		fmt.Printf("%.4f", truncated)
	}
}
