package todo

import (
	"context"
)

// A service aqui que criamos vai conter a logica, ela é um objet (struct) pois mais para frente iremos injetar a dependencia
// nela, calma que vai fazer sentido em outro artigo
type Service struct {
}

func NewService() Service {
	return Service{}
}

// Dessa forma aque fizemos, basta que o método da service chame a função da model para que list todos os todos.
func (s Service) ListTodos(ctx context.Context) []Todo {
	return GetAllTodos(ctx)
}
