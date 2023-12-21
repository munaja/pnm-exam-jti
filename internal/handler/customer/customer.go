package customer

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/munaja/pnm-exam-jti/internal/handler/all-basic-common/auth"
	er "github.com/munaja/pnm-exam-jti/internal/handler/all-basic-common/errors"
	"github.com/munaja/pnm-exam-jti/internal/handler/all-basic-common/home"
	pnm "github.com/munaja/pnm-exam-jti/internal/handler/customer/phone-number"
	"github.com/munaja/pnm-exam-jti/internal/handler/customer/provider"
	ma "github.com/munaja/pnm-exam-jti/internal/middleware/auth"
)

func SetRoutes() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.NotFound(er.NotFoundResponse)
	r.MethodNotAllowed(er.MethodNotAllowedResponse)

	r.Get("/", home.Index)

	r.Route("/auth", func(r chi.Router) {
		r.Post("/login-via-google", auth.LoginViaGoogle)
		r.Get("/logout", auth.Logout)
	})

	r.Route("/phone-number", func(r chi.Router) {
		r.Use(ma.GuardMW)
		r.Post("/", pnm.Create)
		r.Patch("/{id}", pnm.Update)
		r.Delete("/{id}", pnm.Delete)
		r.Get("/", pnm.GetList)
		r.Get("/{id}", pnm.GetDetail)
		r.Get("/gen-random", pnm.GenRandom)
	})

	r.Route("/provider", func(r chi.Router) {
		r.Get("/", provider.GetList)
	})

	return r
}
