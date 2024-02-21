package router

import (
	cnt "auth/internal/controller"
	"github.com/go-chi/chi"
)

func Route(cnt *cnt.HandleAuth) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/api/register", cnt.Register)
	r.Get("/api/login", cnt.Login)
	return r
}
