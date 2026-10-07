package main

import (
	"fmt"
	"os"

	"library-app/config"
	"library-app/library"
	"library-app/notifications"
)

func main() {
	// === Разбор аргументов командной строки ===
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

		default:
			fmt.Printf("Неизвестная команда: %s\n", command)
			fmt.Println("Доступные команды: save <файл>, load <файл>")
			os.Exit(1)
		}
	}
	demo()
}

func handleSave(filename string) {
	lib := library.New()

	// Наполним демо-данными (в реальном проекте — загрузили бы из хранилища).
	if _, err := lib.AddBook("Война и мир", "Л. Толстой", 1869); err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	if _, err := lib.AddBook("Преступление и наказание", "Ф. Достоевский", 1866); err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	if err := lib.SaveToCSV(filename); err != nil {
		fmt.Println("Ошибка сохранения:", err)
		return
	}
	fmt.Printf("Сохранено книг: %d → %s\n", len(lib.GetAllBooks()), filename)
}

func handleLoad(filename string) {
	lib := library.New()
	if err := lib.LoadFromCSV(filename); err != nil {
		fmt.Println("Ошибка загрузки:", err)
		return
	}
	fmt.Printf("Загружено книг: %d из %s\n", len(lib.GetAllBooks()), filename)
	for _, b := range lib.GetAllBooks() {
		fmt.Println(" -", b)
	}
}

func demo() {
	// === config ===
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

	// === library ===
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

	// === Сценарий 1: успешная выдача ===
	fmt.Println("\n[1] Успешная выдача книги")
	if err := myLibrary.IssueBookToReader(book1.ID, reader1.ID); err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Printf("Книга %s выдана читателю %s\n", book1, reader1)
	}

	// === Сценарий 2: повторная выдача ===
	fmt.Println("\n[2] Попытка выдать уже выданную книгу")
	if err := myLibrary.IssueBookToReader(book1.ID, reader2.ID); err != nil {
		fmt.Println("Ожидаемая ошибка:", err)
	} else {
		fmt.Println("Неожиданно: книга выдана повторно")
	}

	// === Сценарий 3: несуществующий читатель ===
	fmt.Println("\n[3] Попытка выдать книгу несуществующему читателю")
	if err := myLibrary.IssueBookToReader(book2.ID, 999); err != nil {
		fmt.Println("Ожидаемая ошибка:", err)
	} else {
		fmt.Println("Неожиданно: книга выдана")
	}

	// === Сценарий 4: успешный возврат ===
	fmt.Println("\n[4] Успешный возврат книги")
	if err := myLibrary.ReturnBook(book1.ID); err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Printf("Книга %s возвращена в библиотеку\n", book1)
	}

	// === Сценарий 5: повторный возврат ===
	fmt.Println("\n[5] Попытка вернуть книгу, которая уже в библиотеке")
	if err := myLibrary.ReturnBook(book1.ID); err != nil {
		fmt.Println("Ожидаемая ошибка:", err)
	} else {
		fmt.Println("Неожиданно: книга возвращена повторно")
	}

	// === notifications ===
	fmt.Println("\n--- Уведомления ---")
	var emailNotifier notifications.Notifier = &notifications.EmailNotifier{Email: "ivan@example.com"}
	if err := emailNotifier.Notify(reader1.String(), "Книга выдана"); err != nil {
		fmt.Println("Ошибка уведомления:", err)
	}

	var smsNotifier notifications.Notifier = &notifications.SMSNotifier{Phone: "+7-900-000-00-00"}
	if err := smsNotifier.Notify(reader2.String(), "Напоминание о возврате"); err != nil {
		fmt.Println("Ошибка уведомления:", err)
	}

	// === domain ===
	fmt.Println("\n--- Каталог ---")
	for _, b := range myLibrary.GetAllBooks() {
		fmt.Println(" -", b)
	}

	reader2.Deactivate()
	if err := myLibrary.IssueBookToReader(book2.ID, reader2.ID); err != nil {
		fmt.Println("Ожидаемая ошибка:", err)
	}
}
