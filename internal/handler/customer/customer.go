package customer

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	er "github.com/munaja/pnm-exam-jti/internal/handler/all-basic-common/errors"
	"github.com/munaja/pnm-exam-jti/internal/handler/all-basic-common/home"
	pnm "github.com/munaja/pnm-exam-jti/internal/handler/customer/phone-number"
)

func SetRoutes() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.NotFound(er.NotFoundResponse)
	r.MethodNotAllowed(er.MethodNotAllowedResponse)

	r.Get("/", home.Index)

	r.Route("/phone-number", func(r chi.Router) {
		r.Post("/", pnm.Create)
		r.Patch("/{id}", pnm.Update)
		r.Delete("/{id}", pnm.Delete)
		r.Get("/", pnm.GetList)
	})

	return r
}
