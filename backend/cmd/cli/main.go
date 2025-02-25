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
			bookReview, _ := strconv.Atoi(review)

			bookID := len(library) + 1

			newBook := book.Book{
				ID:          bookID,
				Title:       bookTitle,
				Description: bookDescription,
				Author:      bookAuthor,
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

		default:
			fmt.Println("Invalid option. Please choose a valid option.")
		}

		fmt.Println("")
	}

}
