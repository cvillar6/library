package book

func DeleteBook(book *Book, library *[]Book) {
	for bookIndex, bookValue := range *library {
		if bookValue.ID == book.ID {
			*library = append((*library)[:bookIndex], (*library)[bookIndex+1:]...)
		}
	}
}
