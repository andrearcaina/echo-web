package routes

import (
	"net/http"

	"github.com/andrearcaina/echo-web/internal/api/http/handler"
	"github.com/andrearcaina/echo-web/pkg/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func InitRouter() chi.Router {
	// the underlying chi router is just a http.NewServeMux
	r := chi.NewRouter()
	// chi middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// the root route
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		utils.SendJSON(w, http.StatusOK, map[string]string{"message": "Hello World!"})
	})

	// testing custom handler
	handler := handler.CustomHandler{}

	r.Route("/api/v1", func(r chi.Router) {
		r.Mount("/handle", handler.InitRoutes())
	})

	return r
}
