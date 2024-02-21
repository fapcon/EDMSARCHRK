package router

import (
	cnt "geo/internal/controller"
	middleware "geo/internal/middleware"
	"github.com/go-chi/chi"
	"gitlab.com/ptflp/gopubsub/queue"
)

func Route(cnt *cnt.HandleGeo, queuer queue.MessageQueuer) *chi.Mux {
	r := chi.NewRouter()
	limit := 5
	limiter := middleware.NewLimit(limit, queuer)
	r.Use(limiter.Limiter)
	r.Post("/api/address/search", cnt.SearchHandle)
	r.Post("/api/address/search", cnt.GeocodeHandle)

	return r
}
