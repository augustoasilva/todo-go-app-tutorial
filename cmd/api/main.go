package main

import (
	"fmt"
	"net/http"

	"github.com/augustoasilva/todo-go-app-tutorial/internal/todo"
)

func main() {
	service := todo.NewService()
	handler := todo.NewHandler(service)

	http.HandleFunc("/todos", handler.ListTodos)
	// Aqui depois vamos adicionar nossar outras rotas

	fmt.Println("Servidor iniciado em local host na porta 8080")
	_ = http.ListenAndServe(":8080", nil)
}
