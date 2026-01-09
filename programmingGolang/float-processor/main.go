package main

import (
	"fmt"
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
		// 0 < число <= 10000: возводим в квадрат
		squared := num * num
		// Выводим с 4 знаками после запятой
		fmt.Printf("%.4f", squared)
	}
}
