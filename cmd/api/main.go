package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type healthCheckResponse struct {
	Status      string `json:"status"`
	ProjectName string `json:"project_name"`
}

type infoResponse struct {
	Version    string `json:"version"`
	ServerTime string `json:"server_time"`
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

func infoHandler(w http.ResponseWriter, r *http.Request) {
	response := infoResponse{
		Version:    "1.0.0",
		ServerTime: time.Now().Format(time.RFC3339),
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

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	if _, err := w.Write([]byte("This page does not exist")); err != nil {
		log.Printf("ошибка записи ответа: %v", err)
	}
}

func main() {
	http.HandleFunc("/healthcheck", healthCheckHandler)
	http.HandleFunc("/api/v1/info", infoHandler)
	http.HandleFunc("/notfound", notFoundHandler)

	addr := ":8080"
	fmt.Println("Сервер запущен на http://localhost" + addr)
	fmt.Println("Проверка:   http://localhost" + addr + "/healthcheck")
	fmt.Println("Информация: http://localhost" + addr + "/api/v1/info")
	fmt.Println("404:        http://localhost" + addr + "/notfound")

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
