package todo

import (
	"context"
)

// A service aqui que criamos vai conter a logica, ela é um objet (struct) pois mais para frente iremos injetar a dependencia
// nela, calma que vai fazer sentido em outro artigo. Também adiciono aqui neste artigo que por hora não temos validações
// de regra de negócio, ainda, mas em outro artigo já vamos adicionar algumas.
type Service struct {
}

func NewService() Service {
	return Service{}
}

// Dessa forma aque fizemos, basta que o método da service chame a função da model para que list todos os todos.
func (s Service) ListTodos(ctx context.Context) []Todo {
	return GetAllTodos(ctx)
}

// igualmente para inserir um novo todos
func (s Service) InsertTodo(ctx context.Context, todo Todo) error {
	return AddTodo(ctx, todo)
}

// igualmente para atualizar um todos
func (s Service) UpdateTodo(ctx context.Context, todo Todo) error {
	return UpdateTodo(ctx, todo)
}

// e para deletar um todos
func (s Service) DeleteTodo(ctx context.Context, id int) error {
	return DeleteTodo(ctx, id)
}
