package book

import "fmt"

type Printable interface {
	PrintInfo()
}

func Print(p Printable) {
	p.PrintInfo()
}

type Book struct {
	Title  string
	Author string
	Pages  int
}

// Comoposicion, es como la herencia........
type TextBook struct {
	Book
	editorial string
	level     string
}

// Simula un constructor
func NewBook(title string, author string, pages int) *Book {
	return &Book{
		Title:  title,
		Author: author,
		Pages:  pages,
	}
}

// Recibe un puntero
func (b *Book) PrintInfo() {
	fmt.Printf("Title: %s\nAuthor: %s\nPages: %d\n", b.Title, b.Author, b.Pages)
}

//Si los atributos fueran privados, es decir que inicien con minusculas
/*
func (b *Book) SetTitle(title string) {
	b.title = title
}
*/

// Constructor de la composicion
func NewTextBook(title, author string, pages int, editorial, level string) *TextBook {
	return &TextBook{
		Book:      Book{title, author, pages},
		editorial: editorial,
		level:     level,
	}
}

func (b *TextBook) PrintInfo() {
	fmt.Printf("Title: %s\nAuthor: %s\nPages: %d\nEditorial: %s\nNivel: %s\n", b.Title, b.Author, b.Pages, b.editorial, b.level)
}
