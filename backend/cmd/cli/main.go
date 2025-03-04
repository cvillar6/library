package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/cvillar6/library/internal/book"
)

func main() {
	fmt.Println("Welcome to the Bookstore!")
	fmt.Println("-------------------------")
	fmt.Println("")

	// Create a library
	library := []book.Book{}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Menu:")
		fmt.Println("1. Create a new book")
		fmt.Println("2. Read books")
		fmt.Println("3. Update book")
		fmt.Println("4. Delete book")
		fmt.Println("0. Exit")
		fmt.Println("")
		fmt.Print("Choose an option: ")

		option, _ := reader.ReadString('\n')
		option = strings.TrimSpace(option)
		optionNumber, err := strconv.Atoi(option)

		if err != nil {
			fmt.Println("Invalid option. Please enter a number.")
			fmt.Println("")
			continue
		}

		switch optionNumber {
		case 0:
			fmt.Println("Exiting the Bookstore. Goodbye!")
			return

		case 1:
			reader := bufio.NewReader(os.Stdin)

			fmt.Print("Enter book title: ")
			bookTitle, _ := reader.ReadString('\n')

			fmt.Print("Enter book description: ")
			bookDescription, _ := reader.ReadString('\n')

			fmt.Print("Enter book author: ")
			bookAuthor, _ := reader.ReadString('\n')

			fmt.Print("Enter book review: ")
			review, _ := reader.ReadString('\n')
			bookReview, _ := strconv.Atoi(strings.TrimSpace(review))

			bookID := len(library) + 1

			newBook := book.Book{
				ID:          bookID,
				Title:       strings.TrimSpace(bookTitle),
				Description: strings.TrimSpace(bookDescription),
				Author:      strings.TrimSpace(bookAuthor),
				Review:      bookReview,
			}

			createdBook, err := book.CreateBook(&newBook, &library)

			if err != nil {
				fmt.Println(err)
			} else {
				library = append(library, createdBook)
				fmt.Println("Book created successfully!")
			}

		case 2:
			book.ReadBook(&library)

		case 3:
			reader := bufio.NewReader(os.Stdin)

			fmt.Print("Enter the book ID that you want to edit: ")
			ID, _ := reader.ReadString('\n')

			bookID, _ := strconv.Atoi(strings.TrimSpace(ID))

			selectedBook := book.GetBookByID(bookID, &library)

			fmt.Printf("Previous title %s, enter the new title: ", selectedBook.Title)
			newTitle, _ := reader.ReadString('\n')

			fmt.Printf("Previous description %s, enter the new description: ", selectedBook.Description)
			newDescription, _ := reader.ReadString('\n')

			fmt.Printf("Previous author %s, enter the new author: ", selectedBook.Author)
			newAuthor, _ := reader.ReadString('\n')

			fmt.Printf("Previous review %d, enter the new review: ", selectedBook.Review)
			review, _ := reader.ReadString('\n')
			newReview, _ := strconv.Atoi(strings.TrimSpace(review))

			if newTitle == "" {
				newTitle = selectedBook.Title
			}

			if newDescription == "" {
				newDescription = selectedBook.Description
			}

			if newAuthor == "" {
				newAuthor = selectedBook.Author
			}

			if newReview == 0 {
				newReview = selectedBook.Review
			}

			newBook := book.Book{
				ID:          bookID,
				Title:       strings.TrimSpace(newTitle),
				Description: strings.TrimSpace(newDescription),
				Author:      strings.TrimSpace(newAuthor),
				Review:      newReview,
			}

			book.UpdateBook(&newBook, &library)

		case 4:
			reader := bufio.NewReader(os.Stdin)

			fmt.Print("Enter the book ID that you want to delete: ")
			ID, _ := reader.ReadString('\n')

			bookID, _ := strconv.Atoi(strings.TrimSpace(ID))

			selectedBook := book.GetBookByID(bookID, &library)

			book.DeleteBook(selectedBook, &library)

		default:
			fmt.Println("Invalid option. Please choose a valid option.")
		}

		fmt.Println("")
	}

}
