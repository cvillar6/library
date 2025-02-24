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
	library := []book.Book{
		{
			ID:          1,
			Title:       "The Alchemist",
			Description: "A wonderful book",
			Author:      "Paulo Coelho",
			Review:      5,
		},
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Menu:")
		fmt.Println("1. Create a new book")
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
			fmt.Println("Creating a new book...")

			// Create a new book
			newBook := book.Book{
				ID:          1,
				Title:       "The Alchemist",
				Description: "A wonderful book",
				Author:      "Paulo Coelho",
				Review:      5,
			}

			createdBook, err := book.CreateBook(&newBook, &library)

			if err != nil {
				fmt.Println(err)
			} else {
				library = append(library, createdBook)
				fmt.Println("Book created successfully!")
				fmt.Println("")
				fmt.Println("Library:")
				for _, book := range library {
					fmt.Println(book)
				}
			}

		default:
			fmt.Println("Invalid option. Please choose a valid option.")

		}
	}

}
