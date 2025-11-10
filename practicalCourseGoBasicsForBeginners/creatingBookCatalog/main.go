package main

import "fmt"

// Определяем структуру Book
type Book struct {
	Title  string
	Author string
	Year   int
}

func main() {
	// Создаём слайс структур Book
	var catalog []Book

	// Добавляем книги в слайс
	catalog = append(catalog, Book{
		Title:  "Война и мир",
		Author: "Лев Толстой",
		Year:   1869,
	})

	catalog = append(catalog, Book{
		Title:  "Мастер и Маргарита",
		Author: "Михаил Булгаков",
		Year:   1967,
	})

	catalog = append(catalog, Book{
		Title:  "Преступление и наказание",
		Author: "Фёдор Достоевский",
		Year:   1866,
	})

	// Выводим заголовок каталога
	fmt.Println("Каталог книг:")

	//Используем цикл for range для вывода информации о каждой книге
	for _, book := range catalog {
		fmt.Printf("Название: '%', Автор: %s, Год: %d\n",
			book.Title, book.Author, book.Year)
	}
}
