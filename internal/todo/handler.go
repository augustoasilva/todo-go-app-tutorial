package todo

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) Handler {
	return Handler{service: service}
}

// ListTodos lista todos todos que temos através da service injetada como dependencia
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
