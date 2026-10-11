package storage

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"library-app/domain"
)

type storageData struct {
	Books   []*domain.Book   `json:"books"`
	Readers []*domain.Reader `json:"readers"`
}

func SaveLibraryToJSON(filePath string, books []*domain.Book, readers []*domain.Reader) error {
	data := storageData{
		Books:   books,
		Readers: readers,
	}

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("ошибка сериализации JSON: %w", err)
	}

	return os.WriteFile(filePath, jsonData, 0644)
}

func LoadLibraryFromJSON(filePath string) ([]*domain.Book, []*domain.Reader, error) {
	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, nil, fmt.Errorf("не удалось прочитать файл '%s': %w", filePath, err)
	}

	var data storageData

	if err := json.Unmarshal(jsonData, &data); err != nil {
		return nil, nil, fmt.Errorf("ошибка десериализации JSON: %w", err)
	}

	return data.Books, data.Readers, nil
}

func SaveBooksToCSV(filename string, books []*domain.Book) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("не удалось создать файл '%s': %w", filename, err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{"ID", "Title", "Author", "Year", "IsIssued", "ReaderID"}
	if err := writer.Write(header); err != nil {
		return fmt.Errorf("ошибка записи заголовка: %w", err)
	}

	for _, b := range books {
		record := []string{
			strconv.Itoa(b.ID),
			b.Title,
			b.Author,
			strconv.Itoa(b.Year),
			strconv.FormatBool(b.IsIssued),
			strconv.Itoa(b.ReaderID),
		}
		if err := writer.Write(record); err != nil {
			return fmt.Errorf("ошибка записи книги '%s': %w", b.Title, err)
		}
	}

	if err := writer.Error(); err != nil {
		return fmt.Errorf("ошибка при сохранении CSV: %w", err)
	}
	return nil
}

func LoadBooksFromCSV(filename string) ([]*domain.Book, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("не удалось открыть файл '%s': %w", filename, err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения CSV: %w", err)
	}

	if len(rows) < 2 {
		return nil, fmt.Errorf("файл '%s' пуст или содержит только заголовок", filename)
	}

	books := make([]*domain.Book, 0, len(rows)-1)
	for i, row := range rows[1:] {
		if len(row) != 6 {
			return nil, fmt.Errorf("строка %d: ожидалось 6 полей, получено %d", i+2, len(row))
		}

		id, err := strconv.Atoi(row[0])
		if err != nil {
			return nil, fmt.Errorf("строка %d: неверный ID '%s': %w", i+2, row[0], err)
		}
		year, err := strconv.Atoi(row[3])
		if err != nil {
			return nil, fmt.Errorf("строка %d: неверный год '%s': %w", i+2, row[3], err)
		}
		isIssued, err := strconv.ParseBool(row[4])
		if err != nil {
			return nil, fmt.Errorf("строка %d: неверный IsIssued '%s': %w", i+2, row[4], err)
		}
		readerID, err := strconv.Atoi(row[5])
		if err != nil {
			return nil, fmt.Errorf("строка %d: неверный ReaderID '%s': %w", i+2, row[5], err)
		}

		books = append(books, &domain.Book{
			ID:       id,
			Title:    row[1],
			Author:   row[2],
			Year:     year,
			IsIssued: isIssued,
			ReaderID: readerID,
		})
	}

	return books, nil
}
