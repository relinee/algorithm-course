package internal

import (
	"fmt"

	"github.com/brianvoe/gofakeit/v6"
)

type Book struct {
	Author string
	Title  string
	Genre  string
}

func (b Book) ToString() string {
	return fmt.Sprintf("Название: \"%s\"\nЖанр: %s\nАвтор: %s", b.Title, b.Genre, b.Author)
}

func InitRandomBook() Book {
	randBook := gofakeit.Book()
	return Book{
		Author: randBook.Author,
		Title:  randBook.Title,
		Genre:  randBook.Genre,
	}
}
