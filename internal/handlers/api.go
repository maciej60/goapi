package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/maciej60/goapi/internal/middleware"
)

func Handler(r *chi.Mux) {
	r.Use(chimiddle.StripSlashes)	// Global middleware

	r.Route("/", func(router chi.Router) {
		router.Post("/login", Login)
	})

	r.Route("/account", func(router chi.Router) {
		router.Use(middleware.Authorization) 
		router.Post("/coins", GetCoinBalance)
	})
}