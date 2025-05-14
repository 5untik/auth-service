package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Ответ")
	})

	fmt.Println("Сервер запущен!")
	error1 := http.ListenAndServe(":8080", nil)
	if error1 != nil {
		log.Fatal("Произошла ошибка при запуске сервера:", error1)
	}
}
