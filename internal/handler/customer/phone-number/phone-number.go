package phonenumber

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	hh "github.com/karincake/apem/handlerhelper"

	"github.com/munaja/pnm-exam-jti/internal/helper/jwt"
	m "github.com/munaja/pnm-exam-jti/internal/model/phone-number"
	s "github.com/munaja/pnm-exam-jti/internal/service/phone-number"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var input m.CreateDto
	if hh.ValidateStructByIOR(w, r.Body, &input) == false {
		return
	}
<<<<<<< HEAD

	input.User_Id = r.Context().Value("authInfo").(*jwt.AuthInfo).User_Id // should be safe being validated by auth
=======
>>>>>>> main
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

	input.User_Id = r.Context().Value("authInfo").(*jwt.AuthInfo).User_Id // should be safe being validated by auth
	res, err := s.Update(id, input)
	hh.DataResponse(w, res, err)

}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateInt(w, "id", chi.URLParam(r, "id"))
	if id <= 0 {
		return
	}

	input := m.DeleteDto{}
	input.User_Id = r.Context().Value("authInfo").(*jwt.AuthInfo).User_Id // should be safe being validated by auth
	res, err := s.Delete(id, input)

	hh.DataResponse(w, res, err)
}

func GetList(w http.ResponseWriter, r *http.Request) {
	x := ""
	input := m.FilterListDto{OddStatus_Opt: &x}
	if hh.ValidateStructByURL(w, *r.URL, &input) == false {
		return
	}
	if *input.OddStatus_Opt == "" {
		*input.OddStatus_Opt = "eq"
	}

	input.User_Id = r.Context().Value("authInfo").(*jwt.AuthInfo).User_Id // should be safe being validated by auth
	input.PageSize = 100
	res, err := s.GetList(input)
	hh.DataResponse(w, res, err)
}

func GetDetail(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateInt(w, "id", chi.URLParam(r, "id"))
	if id <= 0 {
		return
	}

	input := m.FilterDetailDto{}
	input.User_Id = r.Context().Value("authInfo").(*jwt.AuthInfo).User_Id // should be safe being validated by auth
	input.Id = id
	res, err := s.GetDetail(input)
	hh.DataResponse(w, res, err)
}

func GenRandom(w http.ResponseWriter, r *http.Request) {
	input := m.GenRandomDto{
		User_Id: r.Context().Value("authInfo").(*jwt.AuthInfo).User_Id,
		Count:   25,
	}
	res, err := s.GenRandom(input)
	hh.DataResponse(w, res, err)
}
