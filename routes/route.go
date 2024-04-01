package routes

import (
	"learngo/handlers"
	"net/http"

	"github.com/go-chi/chi"
)

func Routes(router *chi.Mux) {
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome to Go lang"))
	})
	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("check"))
	})

	// group other routes with /api
	router.Route("/api", func(r chi.Router) {
		orgRoutes(r)
		departRoutes(r)
		roleRoutes(r)
		userRoutes(r)
		otpRoutes(r)
		AuthRoutes(r)
		LeadRoutes(r)
	})
}

func orgRoutes(r chi.Router) {
	r.Route("/org", func(r chi.Router) {
		r.Post("/", handlers.OrgCreateHandler)
		r.Get("/{id}", handlers.OrgHandler)
		r.Get("/", handlers.OrgListHandler)
		r.Put("/{id}", handlers.OrgUpdateHandler)
	})
}

func departRoutes(r chi.Router) {
	r.Route("/depart", func(r chi.Router) {
		r.Post("/", handlers.DepartCreateHandler)
		r.Get("/{id}", handlers.DepartHandler)
		r.Get("/", handlers.DepartListHandler)
		r.Put("/{id}", handlers.DepartUpdateHandler)
	})
}

func roleRoutes(r chi.Router) {
	r.Route("/role", func(r chi.Router) {
		r.Post("/", handlers.RoleCreateHandler)
		r.Get("/{id}", handlers.RoleHandler)
		r.Get("/", handlers.RoleListHandler)
		r.Put("/{id}", handlers.RoleUpdateHandler)
	})
}

func userRoutes(r chi.Router) {
	r.Route("/users", func(r chi.Router) {
		r.Post("/", handlers.UserCreateHandler)
		r.Get("/{id}", handlers.UserHandler)
		r.Get("/", handlers.UserListHandler)
		r.Put("/{id}", handlers.UserUpdateHandler)
	})
}

func otpRoutes(r chi.Router) {
	r.Route("/otp", func(r chi.Router) {
		r.Post("/", handlers.OrgCreateHandler)
		r.Get("/{id}", handlers.OtpHandler)
	})
}

func AuthRoutes(r chi.Router) {
	r.Route("/login", func(r chi.Router) {
		r.Post("/", handlers.LoginHandler)
	})
}

func LeadRoutes(r chi.Router) {
	r.Route("/lead", func(r chi.Router) {
		r.Post("/", handlers.LeadCreateHandler)
		r.Get("/{id}", handlers.LeadHandler)
		r.Get("/", handlers.LeadListHandler)
		r.Put("/{id}", handlers.LeadUpdateHandler)
	})
}
