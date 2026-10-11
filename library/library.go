package library

import (
	"fmt"

	"library-app/domain"
)

type Library struct {
	Books   []*domain.Book
	Readers []*domain.Reader

	lastBookID   int
	lastReaderID int
}

func New() *Library {
	return &Library{
		Books:   []*domain.Book{},
		Readers: []*domain.Reader{},
	}
}

func (lib *Library) AddBook(title, author string, year int) (*domain.Book, error) {
	if err := validateBook(title, author); err != nil {
		return nil, err
	}
	for _, b := range lib.Books {
		if b.Title == title && b.Author == author {
			return nil, fmt.Errorf("книга '%s' автора '%s' уже есть в библиотеке", title, author)
		}
	}
	lib.lastBookID++
	book := &domain.Book{
		ID:     lib.lastBookID,
		Title:  title,
		Author: author,
		Year:   year,
	}
	lib.Books = append(lib.Books, book)
	return book, nil
}

func (lib *Library) AddReader(firstName, lastName string) (*domain.Reader, error) {
	for _, r := range lib.Readers {
		if r.FirstName == firstName && r.LastName == lastName {
			return nil, fmt.Errorf("читатель '%s %s' уже зарегистрирован", firstName, lastName)
		}
	}
	lib.lastReaderID++
	reader := &domain.Reader{
		ID:        lib.lastReaderID,
		FirstName: firstName,
		LastName:  lastName,
		IsActive:  true,
	}
	lib.Readers = append(lib.Readers, reader)
	return reader, nil
}

func (lib *Library) FindBookByID(id int) (*domain.Book, error) {
	for _, book := range lib.Books {
		if book.ID == id {
			return book, nil
		}
	}
	return nil, fmt.Errorf("книга с ID %d не найдена в библиотеке", id)
}

func (lib *Library) FindReaderByID(id int) (*domain.Reader, error) {
	for _, reader := range lib.Readers {
		if reader.ID == id {
			return reader, nil
		}
	}
	return nil, fmt.Errorf("читатель с ID %d не найден в библиотеке", id)
}

func (lib *Library) IssueBookToReader(bookID, readerID int) error {
	book, err := lib.FindBookByID(bookID)
	if err != nil {
		return err
	}
	reader, err := lib.FindReaderByID(readerID)
	if err != nil {
		return err
	}
	if !reader.IsActive {
		return fmt.Errorf("читатель '%s' деактивирован", reader)
	}
	return book.IssueBook(readerID)
}

func (lib *Library) ReturnBook(bookID int) error {
	book, err := lib.FindBookByID(bookID)
	if err != nil {
		return err
	}
	return book.ReturnBook()
}

func (lib *Library) GetAllBooks() []*domain.Book {
	return lib.Books
}

func (lib *Library) GetAllReaders() []*domain.Reader {
	return lib.Readers
}

// SetBooks / SetReaders — сеттеры для загрузки из storage.
func (lib *Library) SetBooks(books []*domain.Book) {
	lib.Books = books
	lib.UpdateIDs()
}

func (lib *Library) SetReaders(readers []*domain.Reader) {
	lib.Readers = readers
	lib.UpdateIDs()
}

// UpdateIDs — пересчитывает lastBookID и lastReaderID после загрузки.
func (lib *Library) UpdateIDs() {
	maxBookID := 0
	for _, book := range lib.Books {
		if book.ID > maxBookID {
			maxBookID = book.ID
		}
	}
	lib.lastBookID = maxBookID

	maxReaderID := 0
	for _, reader := range lib.Readers {
		if reader.ID > maxReaderID {
			maxReaderID = reader.ID
		}
	}
	lib.lastReaderID = maxReaderID
}

func validateBook(title, author string) error {
	if title == "" {
		return fmt.Errorf("название книги не может быть пустым")
	}
	if author == "" {
		return fmt.Errorf("автор книги не может быть пустым")
	}
	return nil
}