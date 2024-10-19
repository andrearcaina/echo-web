package handler

import (
	"net/http"

	"github.com/andrearcaina/echo-web/pkg/utils"
	"github.com/go-chi/chi/v5"
)

type CustomHandler struct{}

// InitRoutes creates a handler with a custom router for certain tasks
// (this will be later changed for specificity)
func (h *CustomHandler) InitRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/test", h.Test)

	return r
}

func (h *CustomHandler) Test(w http.ResponseWriter, r *http.Request) {
	message := map[string]string{"message": "testing custom handler endpoint"}
	utils.SendJSON(w, http.StatusOK, message)
}
