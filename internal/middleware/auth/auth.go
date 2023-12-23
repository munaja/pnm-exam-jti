package auth

import (
	"context"
	"net/http"

	te "github.com/karincake/tempe/error"
	// ms "github.com/karincake/apem/memstorageredis"

	hh "github.com/karincake/apem/handlerhelper"
	"github.com/munaja/pnm-exam-jti/internal/helper/jwt"
)

func GuardMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessDetail, err := jwt.ExtractToken(r, jwt.AccessToken)
		if err != nil {
			hh.WriteJSON(w, http.StatusUnauthorized, err.(te.XError), nil)
			return
		}

		// accessUuidRedis := ms.I.Get(accessUuid)
		// if accessUuidRedis.String() == "" {
		// 	return nil, te.XError{Code: "token-unidentified", Message: lh.ErrorMsgGen("token-unidentified")}
		// }

		ctx := context.WithValue(r.Context(), "authInfo", accessDetail)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
