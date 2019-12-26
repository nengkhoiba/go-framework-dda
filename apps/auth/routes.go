package auth

import (
	"github.com/go-chi/chi"
)

// Routes related to auth
func (a *App) Routes() chi.Router {
	r := chi.NewRouter()
	r.Post("/post", a.somefunction) //POST
	r.Get("/get", a.somefunction)   //GET

	r.Group(func(r chi.Router) {
		r.Use(a.Middleware.RequireAuthorization)
		r.Post("/post", a.somefunction) //POST AUTH
	})

	return r
}
