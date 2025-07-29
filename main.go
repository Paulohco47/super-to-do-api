package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	_ "back-end-to-do-list/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title           To-Do List API
// @version         1.0
// @description     API para gerenciar tarefas (CRUD).
// @host            localhost:8080
// @BasePath        /

// @contact.name    Paulo Henrique
// @contact.email   contato@paulo.app.br

func main() {
	InitDB()

	http.HandleFunc("/tasks", tasksHandler)
	http.HandleFunc("/tasks/", taskHandler)

	// Rota para acessar a documentação Swagger
	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	fmt.Printf("[%s] ℹ️ Iniciando servidor na porta 8080...\n", time.Now().Format(time.RFC3339))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
