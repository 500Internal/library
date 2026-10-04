package main

import "fmt"

func main() {
	configWithPort := map[string]string{"PORT": "8080"}
	configWithoutPort := map[string]string{"HOST": "localhost"}

	if port, err := GetPortFromConfig(configWithPort); err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Порт:", port)
	}

	if port, err := GetPortFromConfig(configWithoutPort); err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Порт:", port)
	}

	fmt.Println("---Библиотека---")

	book := &Book{ID: 1, Title: "Война и мир", Author: "Л. Толстой", Year: 1869}
	reader := &Reader{ID: 1, Name: "Ярослав"}

	lib := &Library{
		Books:   []*Book{book},
		Readers: []*Reader{reader},
	}

	if err := lib.IssueBookToReader(1, 1); err != nil {
		fmt.Println("Ошибка выдачи:", err)
	} else {
		fmt.Println("Книга успешно выдана читателю")
	}

	if err := lib.ReturnBook(1); err != nil {
		fmt.Println("Ошибка возврата:", err)
	} else {
		fmt.Println("Книга успешно возвращена")
	}

	if err := lib.ReturnBook(1); err != nil {
		fmt.Println("Ожидаемая ошибка:", err)
	} else {
		fmt.Println("Книга успешно возвращена (неожиданно)")
	}
}
