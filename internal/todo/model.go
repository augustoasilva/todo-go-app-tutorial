package todo

import "context"

// Todo é a nossa estrutura de tarefas a fazer, por hora ela vai ter apenas ID inteiro, titulo e se está completa
type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// esse slice é para representar nossa base de forma simplificada, em outro artigo iremos conectar com uma base dados
// mas por hora vamos deixar assim para ficar mais simples
var todos = []Todo{
	{
		ID:        1,
		Title:     "Meu primeiro todo!",
		Completed: true,
	},
	{
		ID:        2,
		Title:     "Da minha primeira api Go!",
		Completed: false,
	},
}

// com a visão de pacotes do Go, podemos criar uma função sem precisar de um objeto (struct) para ter o método
// não é considerado uma boa prática na visão de POO, mas quis trazer que também é possível, diferente da service e do
// handler
func GetAllTodos(ctx context.Context) []Todo {
	return todos
}
