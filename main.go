package main

import (
	"fmt"
	"os"

	"library-app/config"
	"library-app/library"
	"library-app/notifications"
	"library-app/storage"
)

func main() {
	if len(os.Args) >= 2 {
		command := os.Args[1]

		switch command {
		case "save":
			if len(os.Args) < 3 {
				fmt.Println("Использование: save <имя_файла.csv>")
				os.Exit(1)
			}
			handleSave(os.Args[2])
			return

		case "load":
			if len(os.Args) < 3 {
				fmt.Println("Использование: load <имя_файла.csv>")
				os.Exit(1)
			}
			handleLoad(os.Args[2])
			return

		case "savejson":
			if len(os.Args) < 3 {
				fmt.Println("Использование: savejson <имя_файла.json>")
				os.Exit(1)
			}
			handleSaveJSON(os.Args[2])
			return

		case "loadjson":
			if len(os.Args) < 3 {
				fmt.Println("Использование: loadjson <имя_файла.json>")
				os.Exit(1)
			}
			handleLoadJSON(os.Args[2])
			return

		default:
			fmt.Printf("Неизвестная команда: %s\n", command)
			fmt.Println("Доступные команды: save <файл>, load <файл>, savejson <файл>, loadjson <файл>")
			os.Exit(1)
		}
	}

	demo()
}

func handleSave(filename string) {
	lib := library.New()

	if _, err := lib.AddBook("Война и мир", "Л. Толстой", 1869); err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	if _, err := lib.AddBook("Преступление и наказание", "Ф. Достоевский", 1866); err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	if err := storage.SaveBooksToCSV(filename, lib.Books); err != nil {
		fmt.Println("Ошибка сохранения:", err)
		return
	}
	fmt.Printf("Сохранено книг: %d → %s\n", len(lib.Books), filename)
}

func handleLoad(filename string) {
	books, err := storage.LoadBooksFromCSV(filename)
	if err != nil {
		fmt.Println("Ошибка загрузки:", err)
		return
	}
	fmt.Printf("Загружено книг: %d из %s\n", len(books), filename)
	for _, b := range books {
		fmt.Println(" -", b)
	}
}

func handleSaveJSON(filename string) {
	lib := library.New()

	if _, err := lib.AddBook("Война и мир", "Л. Толстой", 1869); err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	if _, err := lib.AddReader("Иван", "Иванов"); err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	if err := storage.SaveLibraryToJSON(filename, lib.Books, lib.Readers); err != nil {
		fmt.Println("Ошибка сохранения JSON:", err)
		return
	}
	fmt.Printf("Сохранено: книг %d, читателей %d → %s\n", len(lib.Books), len(lib.Readers), filename)
}

func handleLoadJSON(filename string) {
	books, readers, err := storage.LoadLibraryFromJSON(filename)
	if err != nil {
		fmt.Println("Ошибка загрузки JSON:", err)
		return
	}

	lib := library.New()
	lib.Books = books
	lib.Readers = readers
	lib.UpdateIDs()

	fmt.Printf("Загружено: книг %d, читателей %d из %s\n", len(lib.Books), len(lib.Readers), filename)
	for _, b := range lib.Books {
		fmt.Println(" -", b)
	}
}

func demo() {
	cfgWithPort := map[string]string{"PORT": "8080"}
	cfgWithoutPort := map[string]string{"HOST": "localhost"}

	if port, err := config.GetPort(cfgWithPort); err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Порт:", port)
	}

	if port, err := config.GetPort(cfgWithoutPort); err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Порт:", port)
	}

	fmt.Println("--- Библиотека ---")

	myLibrary := library.New()

	book1, err := myLibrary.AddBook("Война и мир", "Л. Толстой", 1869)
	if err != nil {
		fmt.Println("Ошибка добавления книги:", err)
		return
	}
	fmt.Println("Добавлена книга:", book1)

	book2, err := myLibrary.AddBook("Преступление и наказание", "Ф. Достоевский", 1866)
	if err != nil {
		fmt.Println("Ошибка добавления книги:", err)
		return
	}
	fmt.Println("Добавлена книга:", book2)

	if _, err := myLibrary.AddBook("Война и мир", "Л. Толстой", 1869); err != nil {
		fmt.Println("Ожидаемая ошибка:", err)
	}

	reader1, err := myLibrary.AddReader("Иван", "Иванов")
	if err != nil {
		fmt.Println("Ошибка добавления читателя:", err)
		return
	}
	fmt.Println("Добавлен читатель:", reader1)

	reader2, err := myLibrary.AddReader("Мария", "Петрова")
	if err != nil {
		fmt.Println("Ошибка добавления читателя:", err)
		return
	}
	fmt.Println("Добавлен читатель:", reader2)

	fmt.Println("--- Сценарии использования ---")

	fmt.Println("\n[1] Успешная выдача книги")
	if err := myLibrary.IssueBookToReader(book1.ID, reader1.ID); err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Printf("Книга %s выдана читателю %s\n", book1, reader1)
	}

	fmt.Println("\n[2] Попытка выдать уже выданную книгу")
	if err := myLibrary.IssueBookToReader(book1.ID, reader2.ID); err != nil {
		fmt.Println("Ожидаемая ошибка:", err)
	} else {
		fmt.Println("Неожиданно: книга выдана повторно")
	}

	fmt.Println("\n[3] Попытка выдать книгу несуществующему читателю")
	if err := myLibrary.IssueBookToReader(book2.ID, 999); err != nil {
		fmt.Println("Ожидаемая ошибка:", err)
	} else {
		fmt.Println("Неожиданно: книга выдана")
	}

	fmt.Println("\n[4] Успешный возврат книги")
	if err := myLibrary.ReturnBook(book1.ID); err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Printf("Книга %s возвращена в библиотеку\n", book1)
	}

	fmt.Println("\n[5] Попытка вернуть книгу, которая уже в библиотеке")
	if err := myLibrary.ReturnBook(book1.ID); err != nil {
		fmt.Println("Ожидаемая ошибка:", err)
	} else {
		fmt.Println("Неожиданно: книга возвращена повторно")
	}

	fmt.Println("\n--- Уведомления ---")
	var emailNotifier notifications.Notifier = &notifications.EmailNotifier{Email: "ivan@example.com"}
	if err := emailNotifier.Notify(reader1.String(), "Книга выдана"); err != nil {
		fmt.Println("Ошибка уведомления:", err)
	}

	var smsNotifier notifications.Notifier = &notifications.SMSNotifier{Phone: "+7-900-000-00-00"}
	if err := smsNotifier.Notify(reader2.String(), "Напоминание о возврате"); err != nil {
		fmt.Println("Ошибка уведомления:", err)
	}

	fmt.Println("\n--- Каталог ---")
	for _, b := range myLibrary.GetAllBooks() {
		fmt.Println(" -", b)
	}

	reader2.Deactivate()
	if err := myLibrary.IssueBookToReader(book2.ID, reader2.ID); err != nil {
		fmt.Println("Ожидаемая ошибка:", err)
	}

	fmt.Println("\n--- JSON: сохранение ---")
	if err := storage.SaveLibraryToJSON("library.json", myLibrary.Books, myLibrary.Readers); err != nil {
		fmt.Println("Ошибка сохранения JSON:", err)
	} else {
		fmt.Println("Библиотека сохранена в library.json")
	}

	fmt.Println("\n--- JSON: загрузка ---")
	loadedBooks, loadedReaders, err := storage.LoadLibraryFromJSON("library.json")
	if err != nil {
		fmt.Println("Ошибка загрузки JSON:", err)
	} else {
		fmt.Println("Загружено книг:", len(loadedBooks))
		fmt.Println("Загружено читателей:", len(loadedReaders))
	}
}
