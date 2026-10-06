package domain

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
	ID        int
	FirstName string
	LastName  string
	IsActive  bool
}

func (b *Book) String() string {
	return fmt.Sprintf("'%s' (%s, %d)", b.Title, b.Author, b.Year)
}

func (b *Book) IssueBook(readerID int) error {
	if b.IsIssued {
		return fmt.Errorf("книга '%s' уже выдана читателю с ID %d", b.Title, b.ReaderID)
	}
	b.IsIssued = true
	b.ReaderID = readerID
	return nil
}

func (b *Book) ReturnBook() error {
	if !b.IsIssued {
		return fmt.Errorf("книга '%s' и так в библиотеке", b.Title)
	}
	b.IsIssued = false
	b.ReaderID = 0
	return nil
}

func (r *Reader) String() string {
	return fmt.Sprintf("%s %s", r.FirstName, r.LastName)
}

func (r *Reader) Activate() {
	r.IsActive = true
}

func (r *Reader) Deactivate() {
	r.IsActive = false
}
