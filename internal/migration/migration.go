package customer

import (
	pn "github.com/munaja/pnm-exam-jti/internal/model/phone-number"
	"github.com/munaja/pnm-exam-jti/internal/model/provider"
	"github.com/munaja/pnm-exam-jti/internal/model/user"
)

func GetModelList() (data []interface{}) {
	tableList := []interface{}{
		&pn.PhoneNumber{},
		&provider.Provider{},
		&user.User{},
	}
	data = append(data, tableList...)

	return data
}
