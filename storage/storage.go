package storage

import (
	"encoding/json"
	"fmt"
	"os"

	"library-app/domain"
	"library-app/library"
)

type storageData struct {
	Books   []*domain.Book   `json:"books"`
	Readers []*domain.Reader `json:"readers"`
}

func SaveLibraryToJSON(filePath string, lib *library.Library) error {
	data := storageData{
		Books:   lib.Books,
		Readers: lib.Readers,
	}

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("ошибка сериализации JSON: %w", err)
	}

	return os.WriteFile(filePath, jsonData, 0644)
}

func LoadLibraryFromJSON(filePath string) (*library.Library, error) {
	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var data storageData
	if err := json.Unmarshal(jsonData, &data); err != nil {
		return nil, fmt.Errorf("ошибка десериализации JSON: %w", err)
	}

	lib := library.New()
	lib.Books = data.Books
	lib.Readers = data.Readers
	lib.UpdateIDs()

	return lib, nil
}