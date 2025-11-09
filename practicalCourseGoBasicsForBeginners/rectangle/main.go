package main

import "fmt"

// Определяем структуру Rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

func main() {
	// Создаём экземпляр структуры с значениями 10.5 и 5.0
	rect := Rectangle{
		Width:  10.5,
		Height: 5.0,
	}

	// Вычисляем площадь и периметр
	area := rect.Width * rect.Height
	perimeter := 2 * (rect.Width + rect.Height)

	// Выводим результаты с форматированием до 2 знаков после запятой
	fmt.Println("Прямоугольник:")
	fmt.Printf("Ширина: %.2f\n", rect.Width)
	fmt.Printf("Высота: %.2f\n", rect.Height)
	fmt.Printf("Площадь: %.2f\n", area)
	fmt.Printf("Периметр: %.2f\n", perimeter)
}
