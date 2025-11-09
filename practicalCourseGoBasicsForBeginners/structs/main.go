package main

import "fmt"

// Определяем структуру User
type User struct {
	Name string
	Age  int
}

func main() {
	// Создаём экземпляр структуры User
	user := User{
		Name: "Гоша",
		Age:  37,
	}

	// Выводим информацию о пользователе
	fmt.Println("Профиль пользователя:")
	fmt.Printf("Имя %s\n", user.Name)
	fmt.Printf("Возраст: %d\n", user.Age)
}
