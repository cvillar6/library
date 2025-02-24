package book

import "errors"

func CreateBook(book *Book, library *[]Book) (Book, error) {
	for _, bookValue := range *library {
		if bookValue.ID == book.ID {
			return Book{}, errors.New("Book already exists")
		}
	}

	return *book, nil
}
