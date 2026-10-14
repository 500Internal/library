package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"library-app/library"
	"library-app/storage"
)

func Run(lib *library.Library, dbPath string) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n=== Библиотека ===")
		fmt.Println("1. Добавить книгу")
		fmt.Println("2. Добавить читателя")
		fmt.Println("3. Выдать книгу")
		fmt.Println("4. Вернуть книгу")
		fmt.Println("5. Показать все книги")
		fmt.Println("6. Показать всех читателей")
		fmt.Println("0. Выход")
		fmt.Print("Выбор: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Неверный ввод, попробуйте снова.")
			continue
		}

		switch choice {
		case 1:
			fmt.Print("Название: ")
			title, _ := reader.ReadString('\n')
			fmt.Print("Автор: ")
			author, _ := reader.ReadString('\n')
			fmt.Print("Год: ")
			yearStr, _ := reader.ReadString('\n')
			year, err := strconv.Atoi(strings.TrimSpace(yearStr))
			if err != nil {
				fmt.Println("Неверный год.")
				continue
			}
			if _, err := lib.AddBook(strings.TrimSpace(title), strings.TrimSpace(author), year); err != nil {
				fmt.Println("Ошибка:", err)
			} else {
				fmt.Println("Книга добавлена.")
			}

		case 2:
			fmt.Print("Имя: ")
			first, _ := reader.ReadString('\n')
			fmt.Print("Фамилия: ")
			last, _ := reader.ReadString('\n')
			if _, err := lib.AddReader(strings.TrimSpace(first), strings.TrimSpace(last)); err != nil {
				fmt.Println("Ошибка:", err)
			} else {
				fmt.Println("Читатель добавлен.")
			}

		case 3:
			fmt.Print("ID книги: ")
			bStr, _ := reader.ReadString('\n')
			fmt.Print("ID читателя: ")
			rStr, _ := reader.ReadString('\n')
			bookID, _ := strconv.Atoi(strings.TrimSpace(bStr))
			readerID, _ := strconv.Atoi(strings.TrimSpace(rStr))
			if err := lib.IssueBookToReader(bookID, readerID); err != nil {
				fmt.Println("Ошибка:", err)
			} else {
				fmt.Println("Книга выдана.")
			}

		case 4:
			fmt.Print("ID книги: ")
			bStr, _ := reader.ReadString('\n')
			bookID, _ := strconv.Atoi(strings.TrimSpace(bStr))
			if err := lib.ReturnBook(bookID); err != nil {
				fmt.Println("Ошибка:", err)
			} else {
				fmt.Println("Книга возвращена.")
			}

		case 5:
			fmt.Println("--- Книги ---")
			for _, b := range lib.GetAllBooks() {
				fmt.Println(" -", b)
			}

		case 6:
			fmt.Println("--- Читатели ---")
			for _, r := range lib.GetAllReaders() {
				fmt.Println(" -", r)
			}

		case 0:
			fmt.Println("Сохранение данных и выход.")
			if err := storage.SaveLibraryToJSON(dbPath, lib); err != nil {
				fmt.Println("Произошла ошибка сохранения списка книг:", err)
			}
			return

		default:
			fmt.Println("Неизвестная команда.")
		}
	}
}
