package storage

import "library-app/domain"

// Storage — in-memory хранилище книг и читателей
type Storage struct {
	books   []*domain.Book
	readers []*domain.Reader
}

// NewStorage — фабрика, возвращает хранилище с не-nil слайсами
func NewStorage() *Storage {
	return &Storage{
		books:   []*domain.Book{},
		readers: []*domain.Reader{},
	}
}

// AddBook добавляет книгу в хранилище
func (s *Storage) AddBook(b *domain.Book) {
	s.books = append(s.books, b)
}

// AddReader добавляет читателя в хранилище
func (s *Storage) AddReader(r *domain.Reader) {
	s.readers = append(s.readers, r)
}

// Books возвращает все книги
func (s *Storage) Books() []*domain.Book {
	return s.books
}

// Readers возвращает всех читателей
func (s *Storage) Readers() []*domain.Reader {
	return s.readers
}

// countBooks — вспомогательная функция для внутреннего использования
func (s *Storage) countBooks() int {
	return len(s.books)
}

// countReaders — аналогично
func (s *Storage) countReaders() int {
	return len(s.readers)
}
