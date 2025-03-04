package book

func UpdateBook(book *Book, library *[]Book) {
	for bookIndex, bookValue := range *library {
		if bookValue.ID == book.ID {
			(*library)[bookIndex] = *book
			break
		}
	}
}
