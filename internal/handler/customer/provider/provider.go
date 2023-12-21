package provider

import (
	"net/http"

	hh "github.com/karincake/apem/handlerhelper"

	m "github.com/munaja/pnm-exam-jti/internal/model/provider"
	s "github.com/munaja/pnm-exam-jti/internal/service/provider"
)

func GetList(w http.ResponseWriter, r *http.Request) {
	var input m.FilterDto
	if hh.ValidateStructByURL(w, *r.URL, &input) == false {
		return
	}

	res, err := s.GetList(input)
	hh.DataResponse(w, res, err)
}
