package todo

import (
	"context"
	"errors"
)

// Todo é a nossa estrutura de tarefas a fazer, por hora ela vai ter apenas ID inteiro, titulo e se está completa
type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// atualizamos agora o slice para o map, desta forma podemos ter um melhor controle da nossa base temporária
// de tal forma que podemos validar, por exemplo, se estamos tentando inserir um todo de ID n, porém o ID n já
// existe na base, algo similar ao que aconteceria em uma Base de Dados real.
var todos = map[int]Todo{
	1: {
		ID:        1,
		Title:     "Meu primeiro todo!",
		Completed: true,
	},
	2: {
		ID:        2,
		Title:     "Da minha primeira api Go!",
		Completed: false,
	},
}

// com a visão de pacotes do Go, podemos criar uma função sem precisar de um objeto (struct) para ter o método
// entretanto não é considerado uma boa prática na visão de POO, mas quis trazer que também é possível, diferente da service e do
// handler.
func GetAllTodos(ctx context.Context) []Todo {
	// Aqui atualizamos a nossa função para que possamos converter o mapa em um slice com todos os Todos
	result := make([]Todo, len(todos))
	for k, v := range todos {
		result[k-1] = v
	}
	return result
}

// agora temos uma função para adicionar o Todo na nossa base, e ainda vamos prevenir que acidentalmente
// se salva um todo novo em cima de um todo já existente
func AddTodo(ctx context.Context, newTodo Todo) error {
	// toda vez que vamos acessar um item (valor) dentro de um map em go, temos dois campos que são retornados
	// o primeiro campo é o valor em si, mas aqui tem uma pegadinha, e é por isso que vem o segundo campo to tipo bool (booleano)
	// esse segundo retorno, que chamei de 'ok' vai indicar se item existe no mapa (true) ou não (false).
	_, ok := todos[newTodo.ID]
	if ok {
		return errors.New("todo already exists")
	}
	todos[newTodo.ID] = newTodo
	return nil
}

// agora temos uma função para atualizar o Todo na nossa base, e ainda vamos prevenir que acidentalmente
// se tente atualizar um todo novo em cima de um todo que não existe
func UpdateTodo(ctx context.Context, todo Todo) error {
	_, ok := todos[todo.ID]
	if !ok {
		return errors.New("todo not found")
	}
	todos[todo.ID] = todo
	return nil
}

// e agora temos uma função para deletar o Todo na nossa base, e ainda vamos prevenir que acidentalmente
// se tente deletar um todo novo em cima de um todo que não existe
func DeleteTodo(ctx context.Context, id int) error {
	_, ok := todos[id]
	if !ok {
		return errors.New("todo not found")
	}
	delete(todos, id)
	return nil
}
