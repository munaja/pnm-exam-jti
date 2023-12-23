package phonenumber

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	hh "github.com/karincake/apem/handlerhelper"

	m "github.com/munaja/pnm-exam-jti/internal/model/phone-number"
	s "github.com/munaja/pnm-exam-jti/internal/service/phone-number"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var input m.CreateDto
	if hh.ValidateStructByIOR(w, r.Body, &input) == false {
		return
	}
	res, err := s.Create(input)
	hh.DataResponse(w, res, err)
}

func Update(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateInt(w, "id", chi.URLParam(r, "id"))
	if id <= 0 {
		return
	}

	var input m.UpdateDto
	if hh.ValidateStructByIOR(w, r.Body, &input) == false {
		return
	}
	res, err := s.Update(id, input)
	hh.DataResponse(w, res, err)

}

func Delete(w http.ResponseWriter, r *http.Request) {
	// ctx := r.Context()
	// authInfo, ok := ctx.Value("authInfo").(*sau.AuthInfo)
	// if !ok {
	// 	hh.WriteJSON(w, http.StatusUnauthorized, nil, nil)
	// 	return
	// }

	id := hh.ValidateInt(w, "id", chi.URLParam(r, "id"))
	if id <= 0 {
		return
	}

	input := m.DeleteDto{}
	res, err := s.Delete(id, input)

	hh.DataResponse(w, res, err)
}

func GetList(w http.ResponseWriter, r *http.Request) {

}
