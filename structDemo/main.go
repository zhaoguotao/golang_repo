package main

import "fmt"

type Book struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	var book1 Book
	var book2 Book
	book1.author = "garrett"
	book1.title = "c base"
	book1.subject = "enjoy yourself"
	book1.book_id = 2344
	book2.author = "zhaoguotao"
	fmt.Println(book1)
	fmt.Println(book2)
	fmt.Println(book2.author)
}
