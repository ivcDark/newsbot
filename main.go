package main

import (
	"github.com/ivcDark/newsbot/cmd"
	_ "github.com/mattn/go-sqlite3" // <-- обязательно, чтобы зарегистрировать драйвер
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

func main() {
	http.Handle("/metrics", promhttp.Handler())
	log.Println("Сервис запущен на :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

	cmd.Execute()
}
