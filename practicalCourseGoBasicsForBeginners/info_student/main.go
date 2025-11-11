package main

import "fmt"

// Определяем структуру Person
type Person struct {
	Name string
	Age  int
}

// Определяем структуру Student с встроенным полем Pesron
type Student struct {
	Person    // Встроенное поле
	StudentID int
}

func main() {
	// Создаём экземпляр структуры Student
	student := Student{
		Person: Person{
			Name: "Гоша Гошников",
			Age:  37,
		},
		StudentID: 101,
	}

	// Выводим информацию о студенте с прямым доступом к полям
	fmt.Println("Имя студента:", student.Name)
	fmt.Println("Возраст студента:", student.Age)
	fmt.Println("ID студента:", student.StudentID)
}
