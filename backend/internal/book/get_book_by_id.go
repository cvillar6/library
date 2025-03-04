package book

func GetBookByID(bookID int, library *[]Book) *Book {
	book := Book{}

	for _, bookValue := range *library {
		if bookValue.ID == bookID {
			book = bookValue
			break
		}
	}

	return &book
}
