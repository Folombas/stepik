package main

import (
	"fmt"
	"strconv"
)

// FirstDigitLoop - находит первую цифру числа математическим способом
func FirstDigitLoop(n int) int {
	if n == 0 {
		return 0
	}

	// Делим число на 10 пока оно больше или равно 10
	for n >= 10 {
		n = n / 10
	}
	return n
}

// FirstDigitString - находит первую цифру через преобразование в строку
func FirstDigitString(n int) int {
	if n == 0 {
		return 0
	}

	str := strconv.Itoa(n)
	firstChar := string(str[0])
	digit, _ := strconv.Atoi(firstChar)
	return digit
}

// FirstDigitLogarithm - находит первую цифру с использованием логарифма
func FirstDigitLogarithm(n int) int {
	if n == 0 {
		return 0
	}

	// Находим количество цифр в числе
	power := 1
	temp := n
	for temp >= 10 {
		temp /= 10
		power *= 10
	}

	// Первая цифра = n / 10^(количество_цифр-1)
	return n / power
}

func main() {
	var n int

	fmt.Println("Введите натуральное число (не более 10000):")
	_, err := fmt.Scan(&n)

	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

	// Проверка на отрицательные числа
	if n < 0 {
		fmt.Println("Число должно быть неотрицательным!")
		return
	}

	// Проверка на превышение лимита (хотя по условию до 10000, но алгоритмы работают для любых)
	if n > 10000 {
		fmt.Println("Внимание: число превышает 10000, но алгоритм продолжает работу")
	}

	// Используем математический способ (основной)
	firstDigit := FirstDigitLoop(n)

	fmt.Printf("\n=== РЕЗУЛЬТАТЫ ===\n")
	fmt.Printf("Введенное число: %d\n", n)
	fmt.Printf("Первая цифра (математический способ): %d\n", firstDigit)

	// Демонстрация других способов для сравнения
	fmt.Printf("\n=== СРАВНЕНИЕ СПОСОБОВ ===\n")
	fmt.Printf("Через строку: %d\n", FirstDigitString(n))
	fmt.Printf("С помощью логарифма: %d\n", FirstDigitLogarithm(n))

	// Проверка на корректность всех методов
	if FirstDigitLoop(n) == FirstDigitString(n) && FirstDigitLoop(n) == FirstDigitLogarithm(n) {
		fmt.Println("\n✓ Все методы дали одинаковый результат!")
	}

	// Дополнительная информация
	fmt.Printf("\n=== ИНФОРМАЦИЯ ===\n")
	fmt.Printf("Количество цифр в числе: %d\n", countDigits(n))

	if n >= 10 {
		lastDigit := n % 10
		fmt.Printf("Последняя цифра: %d\n", lastDigit)
		fmt.Printf("Сумма первой и последней цифры: %d\n", firstDigit+lastDigit)
	}
}

// countDigits - подсчитывает количество цифр в числе
func countDigits(n int) int {
	if n == 0 {
		return 1
	}

	count := 0
	for n != 0 {
		n /= 10
		count++
	}
	return count
}
