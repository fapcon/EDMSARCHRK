package router

import (
	"github.com/go-chi/chi"
	cnt "user/internal/controller"
)

func Route(cnt *cnt.HandleUser) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/api/profile", cnt.Profile)
	r.Get("/api/list", cnt.List)
	return r
}
