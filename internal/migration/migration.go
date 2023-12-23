package customer

import (
	pn "github.com/munaja/pnm-exam-jti/internal/model/phone-number"
	"github.com/munaja/pnm-exam-jti/internal/model/provider"
)

func GetModelList() (data []interface{}) {
	tableList := []interface{}{
		&pn.PhoneNumber{},
		&provider.Provider{},
	}
	data = append(data, tableList...)

	return data
}
