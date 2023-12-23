package auth

import (
	"context"
	"net/http"

	hh "github.com/karincake/apem/handlerhelper"
	td "github.com/karincake/tempe/data"

	"github.com/munaja/pnm-exam-jti/internal/helper/jwt"
	m "github.com/munaja/pnm-exam-jti/internal/model/user"
	s "github.com/munaja/pnm-exam-jti/internal/service/auth"
)

var Position m.Position

func LoginViaGoogle(w http.ResponseWriter, r *http.Request) {
	var input m.LoginViaGoogleDto
	if hh.ValidateStructByIOR(w, r.Body, &input) == false {
		return
	}

	input.Duration = "1-h"
	input.Position = Position
	boolTrue := true
	if Position == m.UPOperator {
		input.OptStatus = &boolTrue
	} else if Position == m.UPOwner {
		input.OwnerStatus = &boolTrue
	}
	res, err := s.GenTokenViaGoogle(input)
	if err != nil {
		hh.WriteJSON(w, http.StatusUnauthorized, td.II{"errors": err}, nil)
	} else {
		hh.DataResponse(w, res, nil)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	authInfo := context.Context.Value(r.Context(), "authInfo").(*jwt.AuthInfo)
	s.RevokeToken(authInfo.Uuid)
	hh.WriteJSON(w, http.StatusOK, td.IS{"message": "logged out"}, nil)
}
