package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Yershuaq/library/library" // Замените YourName на ваше имя
)

func main() {
	lib := library.NewLibrary()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1. Add Book")
		fmt.Println("2. Borrow Book")
		fmt.Println("3. Return Book")
		fmt.Println("4. List Books")
		fmt.Println("5. Exit")
		fmt.Print("Enter choice: ")

		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			var id, title, author string
			fmt.Print("Enter book ID: ")
			fmt.Scan(&id)
			fmt.Print("Enter book title: ")
			scanner.Scan()
			title = scanner.Text()
			fmt.Print("Enter book author: ")
			scanner.Scan()
			author = scanner.Text()
			lib.AddBook(library.Book{ID: id, Title: title, Author: author})
		case "2":
			fmt.Print("Enter book ID to borrow: ")
			scanner.Scan()
			lib.BorrowBook(scanner.Text())
		case "3":
			fmt.Print("Enter book ID to return: ")
			scanner.Scan()
			lib.ReturnBook(scanner.Text())
		case "4":
			lib.ListBooks()
		case "5":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice, try again.")
		}
	}
}
