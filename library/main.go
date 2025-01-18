package main

import (
	"fmt"
	"library/internal"
)

func main() {
	library := internal.InitRandomLibrary(3, 3)
	library.PrintLibrary()
	fmt.Print("---------------------\n")

	library.Sort()
	library.PrintLibrary()
	fmt.Print("---------------------\n")

	books := []internal.Book{
		{
			Author: "Bob",
			Title:  "123",
			Genre:  "Fantasy",
		},
		{
			Author: "Alise",
			Title:  "213",
			Genre:  "Fantasy",
		},
		{
			Author: "Bob",
			Title:  "321",
			Genre:  "Adventure",
		},
		{
			Author: "George",
			Title:  "123",
			Genre:  "Roman",
		},
	}

	testLibrary := internal.NewLibrary(books, 2)
	fmt.Print("------Тест------\n")
	testLibrary.PrintLibrary()
	fmt.Print("---------------------\n")

	testLibrary.Shuffle()
	fmt.Print("------После раскидывания------\n")
	testLibrary.PrintLibrary()
	fmt.Print("---------------------\n")

	testLibrary.Sort()
	fmt.Print("------После сортировки------\n")
	testLibrary.PrintLibrary()
	fmt.Print("---------------------\n")
}
