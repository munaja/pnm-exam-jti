package provider

import (
	dg "github.com/karincake/apem/databasegorm"
	gh "github.com/karincake/getuk"
	td "github.com/karincake/tempe/data"

	m "github.com/munaja/pnm-exam-jti/internal/model/provider"
	sh "github.com/munaja/pnm-exam-jti/pkg/servicehelper"
)

const source = "phone-number"

func GetList(input m.FilterDto) (*td.Data, error) {
	var data []m.Provider
	var count int64

	var pagination gh.Pagination
	res := dg.I.
		Model(&m.Provider{}).
		Count(&count).
		Scopes(gh.Paginate(input, &pagination)).
		Find(&data)
	if res.Error != nil {
		return nil, sh.SetError(sh.Event{
			Feature: "phone-number",
			Action:  "create-data",
			Source:  source,
			Status:  "failed",
			ECode:   "data-create-fail",
			EDetail: res.Error.Error(),
		}, data)
	}

	return &td.Data{
		Meta: td.II{
			"count": count,
		},
		Data: data,
	}, nil
}
