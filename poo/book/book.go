package book

import "fmt"

type Pirntable interface {
	PrintInfo()
}

func Print(p Pirntable) {
	p.PrintInfo()
}

type Book struct {
	title  string
	author string
	pages  int
}

func NewBook(title string, author string, pages int) *Book {
	return &Book{title, author, pages}
}

func (b *Book) PrintInfo() {
	fmt.Printf("Title: %s by %s pages: %d", b.title, b.author, b.pages)
}

func (b *Book) SetTitle(title string) {
	b.title = title
}

func (b *Book) GetTitle() string {
	return b.title
}

func (b *Book) SetAuthor(author string) {
	b.author = author
}

func (b *Book) GetAuthor() string {
	return b.author
}

func (b *Book) SetPages(pages int) {
	b.pages = pages
}

func (b *Book) GetPages() int {
	return b.pages
}

type TextBook struct {
	Book
	edition int
	level   string
}

func NewTextBook(title string, author string, pages int, edition int, level string) *TextBook {
	return &TextBook{Book{title, author, pages}, edition, level}
}

func (t *TextBook) PrintInfo() {
	fmt.Printf("Title: %s by %s pages: %d edition: %d level: %s", t.title, t.author, t.pages, t.edition, t.level)
}
