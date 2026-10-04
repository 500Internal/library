package main

import "fmt"

type Book struct {
	ID       int
	Title    string
	Author   string
	Year     int
	IsIssued bool
	ReaderID int
}

type Reader struct {
	ID   int
	Name string
}

func (b *Book) ReturnBook() error {
	if !b.IsIssued {
		return fmt.Errorf("книга '%s' и так в библиотеке", b.Title)
	}
	b.IsIssued = false
	b.ReaderID = 0
	return nil
}

func (r *Reader) IssueBook(b *Book) {
	b.IsIssued = true
	b.ReaderID = r.ID
}
