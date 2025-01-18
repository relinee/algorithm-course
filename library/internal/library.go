package internal

import (
	"fmt"
	"math/rand/v2"
)

type Library struct {
	Shelfs       []Shelf
	MaxShelfSize int
}

type Shelf struct {
	Books []Book
}

func (l *Library) PrintLibrary() {
	for i, shelf := range l.Shelfs {
		fmt.Printf("Полка %v:\n", i+1)
		for _, book := range shelf.Books {
			fmt.Println(book.ToString() + "\n")
		}
	}
}

func (l *Library) BooksCount() int32 {
	cnt := 0
	for _, shelf := range l.Shelfs {
		cnt += len(shelf.Books)
	}
	return int32(cnt)
}

func (l *Library) Sort() {
	allBooks := make([]Book, 0, len(l.Shelfs))
	for _, shelf := range l.Shelfs {
		allBooks = append(allBooks, shelf.Books...)
	}

	allBooks = mergeSort(allBooks)

	shelfs := make([]Shelf, 0)
	curShelf := newShelf(l.MaxShelfSize)
	currentGenre := ""

	for _, book := range allBooks {
		if currentGenre == "" {
			currentGenre = book.Genre
		}

		if len(curShelf.Books) == l.MaxShelfSize {
			shelfs = append(shelfs, curShelf)
			currentGenre = book.Genre
			curShelf = newShelf(l.MaxShelfSize)
			curShelf.Books = append(curShelf.Books, book)
			continue
		}

		if book.Genre == currentGenre {
			curShelf.Books = append(curShelf.Books, book)
			continue
		}

		//  Новый жанр на новую полку
		shelfs = append(shelfs, curShelf)
		currentGenre = book.Genre
		curShelf = newShelf(l.MaxShelfSize)
		curShelf.Books = append(curShelf.Books, book)
	}

	if len(curShelf.Books) > 0 {
		shelfs = append(shelfs, curShelf)
	}

	l.Shelfs = shelfs
}

func (l *Library) Shuffle() {
	allBooks := make([]Book, 0, l.BooksCount())
	for _, shelf := range l.Shelfs {
		allBooks = append(allBooks, shelf.Books...)
	}

	rand.Shuffle(len(allBooks), func(i, j int) {
		allBooks[i], allBooks[j] = allBooks[j], allBooks[i]
	})

	l.Shelfs = NewLibrary(allBooks, l.MaxShelfSize).Shelfs
}

func NewLibrary(books []Book, maxShelfSize int) Library {
	shelfs := make([]Shelf, 0)
	curShelf := newShelf(maxShelfSize)

	for _, book := range books {
		curShelf.Books = append(curShelf.Books, book)

		if len(curShelf.Books) == maxShelfSize {
			shelfs = append(shelfs, curShelf)
			curShelf = newShelf(maxShelfSize)
		}
	}

	if len(curShelf.Books) > 0 {
		shelfs = append(shelfs, curShelf)
	}

	return Library{shelfs, maxShelfSize}
}

func InitRandomLibrary(booksCnt, maxShelfSize int) Library {
	books := make([]Book, 0, booksCnt)
	for i := 0; i < booksCnt; i++ {
		books = append(books, InitRandomBook())
	}
	return NewLibrary(books, maxShelfSize)
}

func newShelf(maxShelfSize int) Shelf {
	return Shelf{make([]Book, 0, maxShelfSize)}
}

func compareBooks(book1, book2 Book) bool {
	if book1.Genre != book2.Genre {
		return book1.Genre < book2.Genre
	}
	if book1.Author != book2.Author {
		return book1.Author < book2.Author
	}
	return book1.Title < book2.Title
}

func mergeSort(books []Book) []Book {
	if len(books) <= 1 {
		return books
	}

	mid := len(books) / 2
	left := mergeSort(books[:mid])
	right := mergeSort(books[mid:])

	return merge(left, right)
}

func merge(left, right []Book) []Book {
	result := make([]Book, 0, len(left)+len(right))
	l, r := 0, 0

	for l < len(left) && r < len(right) {
		if compareBooks(left[l], right[r]) {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}

	result = append(result, left[l:]...)
	result = append(result, right[r:]...)

	return result
}
