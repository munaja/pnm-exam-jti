package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/spf13/viper"

	dg "github.com/karincake/apem/databasegorm"
	l "github.com/karincake/apem/lang"
	ms "github.com/karincake/apem/memstorageredis"
	td "github.com/karincake/tempe/data"
	te "github.com/karincake/tempe/error"
	lh "github.com/munaja/pnm-exam-jti/pkg/langhelper"

	mu "github.com/munaja/pnm-exam-jti/internal/model/user"
)

//	type TokenDetails struct {
//		AccessToken  string
//		RefreshToken string
//		AccessUuid   string
//		RefreshUuid  string
//		AtExpires    int64
//		RtExpires    int64
//	}

// Generates token and store in redis at one place
// just return the error code
func GenTokenViaGoogle(input mu.LoginViaGoogleDto) (any, error) {
	req, err := http.NewRequest("GET", "https://www.googleapis.com/oauth2/v3/userinfo?access_token="+input.AccessToken, nil)
	if err != nil {
		return nil, te.XError{Code: "googleapi-request-error", Message: err.Error()}
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, te.XError{Code: "googleapi-request-error", Message: err.Error()}
	}
	userInfo := map[string]any{}
	err = json.NewDecoder(resp.Body).Decode(&userInfo)
	if err != nil {
		return nil, te.XError{Code: "googleapi-response-error", Message: err.Error()}
	}
	if _, ok := userInfo["error"]; ok {
		return nil, te.XError{Code: "googleapi-response-error", Message: userInfo["error_description"].(string)}
	}
	if _, ok := userInfo["email"]; !ok {
		return nil, te.XError{Code: "googleapi-response-error", Message: "no email available"}
	}

	// Get User
	email := userInfo["email"].(string)
	user := mu.User{}
	if errCode := getAndCheck(&user, mu.User{Email: email}); errCode != "" {
		if errCode == "record-not-found" {
			user.Email = email
			if err := dg.I.Save(&user).Error; err != nil {
				return nil, te.XErrors{"authentication": te.XError{Code: "user-create-fail", Message: "failed to create user account"}}
			}
		} else {
			return nil, te.XErrors{"authentication": te.XError{Code: errCode, Message: lh.ErrorMsgGen(errCode)}}
		}
	}

	// Access token prep
	id, err := uuid.NewRandom()
	if err != nil {
		panic(fmt.Sprintf(l.I.Msg("uuid-gen-fail"), err))
	}
	aUuid := id.String()

	// calculate
	durations := strings.Split(strings.ToLower(input.Duration), "-")
	duration := time.Hour * 24
	if len(durations) == 2 {
		val, err := strconv.Atoi(durations[0])
		if err == nil {
			if durations[1] == "m" {
				duration = time.Minute * time.Duration(val)
			} else if durations[1] == "h" {
				duration = time.Hour * time.Duration(val)
			} else if durations[1] == "d" {
				duration = time.Hour * 24 * time.Duration(val)
			}
		}
	}
	atExpires := time.Now().Add(duration).Unix()

	// key
	atSecretKey := viper.GetString("authConf.atSecretKey")

	// Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = user.Id
	atClaims["user_name"] = user.Name
	atClaims["user_email"] = user.Email
	atClaims["exp"] = atExpires
	atClaims["uuid"] = aUuid
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	ats, err := at.SignedString([]byte(atSecretKey))
	if err != nil {
		return nil, te.XErrors{"user": te.XError{Code: "token-sign-err", Message: lh.ErrorMsgGen("token-sign-err")}}
	}
	// Save to redis
	// now := time.Now()
	// atx := time.Unix(atExpires, 0) //converting Unix to UTC(to Time object)
	// err = ms.I.Set(aUuid, strconv.Itoa(user.Id), atx.Sub(now)).Err()
	// if err != nil {
	// 	panic(fmt.Sprintf(l.I.Msg("redis-store-fail"), err.Error()))
	// }

	// Oauth doesn't need fail count
	// user.FailedLoginAttemptCount = 0
	// user.LastSuccessLogin = time.Now()
	// user.LastAllowedLogin = time.Now()
	// dg.I.Save(&user)

	// Current data
	return td.II{
		"id":          strconv.Itoa(user.Id),
		"name":        user.Name,
		"email":       user.Email,
		"accessToken": ats,
	}, nil
}

func RevokeToken(uuid string) {
	ms.I.Del(uuid)
}
