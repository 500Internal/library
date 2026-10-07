package library

import (
	"fmt"

	"library-app/domain"
	"library-app/storage"
)

type Library struct {
	books   []*domain.Book
	readers []*domain.Reader

	lastBookID   int
	lastReaderID int
}

// New — фабричная функция Возвращает пустую библиотеку с не-nil слайсами
func New() *Library {
	return &Library{
		books:   []*domain.Book{},
		readers: []*domain.Reader{},
	}
}

func (lib *Library) AddBook(title, author string, year int) (*domain.Book, error) {
	if err := validateBook(title, author); err != nil {
		return nil, err
	}
	for _, b := range lib.books {
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
	lib.books = append(lib.books, book)
	return book, nil
}

func (lib *Library) AddReader(firstName, lastName string) (*domain.Reader, error) {
	for _, r := range lib.readers {
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
	lib.readers = append(lib.readers, reader)
	return reader, nil
}

func (lib *Library) FindBookByID(id int) (*domain.Book, error) {
	for _, book := range lib.books {
		if book.ID == id {
			return book, nil
		}
	}
	return nil, fmt.Errorf("книга с ID %d не найдена в библиотеке", id)
}

func (lib *Library) FindReaderByID(id int) (*domain.Reader, error) {
	for _, reader := range lib.readers {
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
	return lib.books
}

// SaveToCSV — обертка над storage.SaveBooksToCSV
func (lib *Library) SaveToCSV(filename string) error {
	return storage.SaveBooksToCSV(filename, lib.books)
}

// LoadFromCSV — читает книги из файла и заменяет ими текущий список
func (lib *Library) LoadFromCSV(filename string) error {
	books, err := storage.LoadBooksFromCSV(filename)
	if err != nil {
		return err
	}
	lib.books = books

	maxID := 0
	for _, b := range books {
		if b.ID > maxID {
			maxID = b.ID
		}
	}
	lib.lastBookID = maxID
	return nil
}

// validateBook — вспомогательная функция, используется только внутри пакета
func validateBook(title, author string) error {
	if title == "" {
		return fmt.Errorf("название книги не может быть пустым")
	}
	if author == "" {
		return fmt.Errorf("автор книги не может быть пустым")
	}
	return nil
}
