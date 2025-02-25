package book

import "fmt"

func ReadBook(library *[]Book) {
	for _, book := range *library {
		fmt.Println(book)
	}
}
