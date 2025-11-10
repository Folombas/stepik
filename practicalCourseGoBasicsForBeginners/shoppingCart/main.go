package main

import "fmt"

func main() {
    // Создаем карту для корзины покупок
    shoppingCart := make(map[string]float64)

    // Добавляем товары с ценами
    shoppingCart["Молоко"] = 65.50
    shoppingCart["Хлеб"] = 40.00
    shoppingCart["Сыр"] = 150.75
    shoppingCart["Сок"] = 85.25

    // Вычисляем общую сумму с явным объявлением типа
    var total float64 = 0.0
    for _, price := range shoppingCart {
        total += price
    }

    // Выводим итоговую сумму с форматированием
    fmt.Printf("Итоговая сумма: %.2f руб.\n", total)
}
