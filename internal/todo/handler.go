package todo

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) Handler {
	return Handler{service: service}
}

// ListTodos lista todos os Todos que temos através da service injetada como dependencia
// de novo, calma que vai fazer sentido mais para frente.
func (h Handler) ListTodos(w http.ResponseWriter, r *http.Request) {
	// aqui ele chama a service que foi injetada como dependencia
	// (depois veremos melhor essa parte de injeção e como fazer de uma forma mais própria para isso e os seus benefícios)
	allTodos := h.service.ListTodos(r.Context())

	// com o resultado em mãos já podemos dizer que vamos retornar uma resposta do tipo json
	w.Header().Set("Content-Type", "application/json")
	// com status http 200 (OK)
	w.WriteHeader(http.StatusOK)
	// e por final pegamos nosso slice e convertemos em um json para ser respondido.
	_ = json.NewEncoder(w).Encode(allTodos)
}

// PostTodo insere o novo todos que enviamos através da service injetada como dependencia
func (h Handler) PostTodo(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	if errDecode := json.NewDecoder(r.Body).Decode(&todo); errDecode != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]string{"error": errDecode.Error()})
		return
	}

	if err := h.service.InsertTodo(r.Context(), todo); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// PostTodo atualiza o um todos que enviamos através da service injetada como dependencia
func (h Handler) PutTodo(w http.ResponseWriter, r *http.Request) {
	id, convErr := strconv.Atoi(strings.TrimSpace(r.PathValue("id")))
	if convErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]string{"error": convErr.Error()})
		return
	}

	var todo Todo
	if errDecode := json.NewDecoder(r.Body).Decode(&todo); errDecode != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]string{"error": errDecode.Error()})
		return
	}

	if id != todo.ID {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]string{"error": "invalid todo to update"})
		return
	}

	if err := h.service.UpdateTodo(r.Context(), todo); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
}

// PostTodo deleta o todos que enviamos através da service injetada como dependencia
func (h Handler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	id, convErr := strconv.Atoi(strings.TrimSpace(r.PathValue("id")))
	if convErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]string{"error": convErr.Error()})
		return
	}

	if err := h.service.DeleteTodo(r.Context(), id); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
}
