package jwt

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	te "github.com/karincake/tempe/error"

	"github.com/golang-jwt/jwt"
	l "github.com/karincake/apem/lang"
	lh "github.com/munaja/pnm-exam-jti/pkg/langhelper"
	"github.com/spf13/viper"
)

func VerifyToken(r *http.Request, tokenType TokenType) (data *jwt.Token, errCode, errDetail string) {
	auth := r.Header.Get("Authorization")
	if auth == "" {
		return nil, "auth-missingHeader", ""
	}
	authArr := strings.Split(auth, " ")
	if len(authArr) == 2 {
		auth = authArr[1]
	}

	token, err := jwt.Parse(auth, func(token *jwt.Token) (any, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf(l.I.Msg("token-sign-unexcpeted"), token.Header["alg"])
		}
		if tokenType == AccessToken {
			return []byte(viper.GetString("authConf.atSecretKey")), nil
		} else {
			return []byte(viper.GetString("authConf.rtSecretKey")), nil
		}
	})
	if err != nil {
		return nil, "token-parse-fail", err.Error()
	}
	return token, "", ""
}

func ExtractToken(r *http.Request, tokenType TokenType) (*AuthInfo, error) {
	token, errCode, errDetail := VerifyToken(r, tokenType)
	if errCode != "" {
		return nil, te.XError{Code: errCode, Message: lh.ErrorMsgGen(errCode, errDetail)}
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid, ok := claims["uuid"].(string)
		if !ok {
			return nil, te.XError{Code: "token-invalid", Message: lh.ErrorMsgGen("token-invalid", "uuid not available")}
		}
		user_id, myErr := strconv.ParseInt(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if myErr != nil {
			return nil, te.XError{Code: "token-invalid", Message: lh.ErrorMsgGen("token-invalid", "uuid is not available")}
		}
		user_name := fmt.Sprintf("%v", claims["user_name"])
		user_email := fmt.Sprintf("%.f", claims["user_email"])
		return &AuthInfo{
			Uuid:       accessUuid,
			User_Id:    int(user_id),
			User_Name:  user_name,
			User_Email: user_email,
		}, nil
	}
	return nil, te.XError{Code: "token", Message: "token-invalid"}
}
