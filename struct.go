package main

import "fmt"

type Books struct {
	title string
	author string
	subject string
	book_id int
}

func main() {
	var Book1 Books
	var Book2 Books

	Book1.title = "Go Programming"
	Book1.author = "Mahesh Kumar"
	Book1.subject = "Go Programming Tutorial"
	Book1.book_id = 6495407

	Book2.title = "Telecom Billing"
	Book2.author = "Zara Ali"
	Book2.subject = "Telecom Billing Tutorial"
	Book2.book_id = 6495700

	fmt.Printf("Book 1 title: %s\n", Book1.title)
	fmt.Printf("Book 1 author: %s\n", Book1.author)
	fmt.Printf("Book 1 subject: %s\n", Book1.subject)
	fmt.Printf("Book 1 book_id: %d\n", Book1.book_id)
	fmt.Printf("\n")

	fmt.Printf("Book 2 title: %s\n", Book2.title)
	fmt.Printf("Book 2 author: %s\n", Book2.author)
	fmt.Printf("Book 2 subject: %s\n", Book2.subject)
	fmt.Printf("Book 2 book_id: %d\n", Book2.book_id)
}