package main

import (
	"fmt"
	"net/http"

	"github.com/augustoasilva/todo-go-app-tutorial/internal/todo"
)

func main() {
	service := todo.NewService()
	handler := todo.NewHandler(service)

	// agora que temos mais rotas, e as rotas acabam se duplicando, para identificar qual rota combina com qual verbo http
	// e assim poder escolher o handler correto, é desta forma que fica usando o pacote padrão http do Go
	http.HandleFunc("GET /todos", handler.ListTodos)
	http.HandleFunc("POST /todos", handler.PostTodo)

	// e no caso de querermos passar um parametro ao final da url, que geralmente fazemos passando algum id, usamos um
	// wildcard na url, que é como chamamos o {} aqui. para este caso estamos usando {id}
	http.HandleFunc("PUT /todos/{id}", handler.PutTodo)
	http.HandleFunc("DELETE /todos/{id}", handler.DeleteTodo)

	fmt.Println("Servidor iniciado em local host na porta 8080")
	_ = http.ListenAndServe(":8080", nil)
}
