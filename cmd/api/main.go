package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type healthCheckResponse struct {
	Status      string `json:"status"`
	ProjectName string `json:"project_name"`
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	response := healthCheckResponse{
		Status:      "available",
		ProjectName: "Library API",
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(jsonData); err != nil {
		log.Printf("ошибка записи ответа: %v", err)
	}
}

func main() {
	http.HandleFunc("/healthcheck", healthCheckHandler)

	addr := ":8080"
	fmt.Println("Сервер запущен на http://localhost" + addr)
	fmt.Println("Проверка: http://localhost" + addr + "/healthcheck")

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
