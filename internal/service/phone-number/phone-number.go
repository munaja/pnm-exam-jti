package phonenumber

import (
	"errors"
	"math/rand"
	"strconv"
	"time"

	sc "github.com/jinzhu/copier"
	dg "github.com/karincake/apem/databasegorm"
	l "github.com/karincake/apem/lang"
	gh "github.com/karincake/getuk"
	td "github.com/karincake/tempe/data"
	te "github.com/karincake/tempe/error"
	"gorm.io/gorm"

	m "github.com/munaja/pnm-exam-jti/internal/model/phone-number"
	sh "github.com/munaja/pnm-exam-jti/pkg/servicehelper"
)

const source = "phone-number"

func Create(input m.CreateDto) (*m.PhoneNumber, error) {
	data := &m.PhoneNumber{}
	if err := sc.Copy(data, input); err != nil {
		return nil, te.XErrors{"struct": te.XError{Code: "copy-fail", Message: l.I.Msg("data-copy-fail")}}
	}

	toInt, _ := strconv.Atoi(data.Number)
	data.OddStatus = toInt%2 != 0

	if err := dg.I.Create(&data).Error; err != nil {
		ed := sh.Event{
			Feature: "phone-number",
			Action:  "create-data",
			Source:  source,
			Status:  "failed",
			ECode:   "data-create-fail",
			EDetail: err.Error(),
		}
		return nil, sh.SetError(ed, nil)
	}

	return data, nil
}

func Update(id int, input m.UpdateDto) (*m.PhoneNumber, error) {
	data := &m.PhoneNumber{}

	if err := dg.I.Where("Id = ?", id).First(data).Error; err != nil {
		return nil, returnFetchError(err)
	}
	if err := sc.Copy(data, input); err != nil {
		return nil, te.XErrors{"struct": te.XError{Code: "copy-fail", Message: l.I.Msg("data-copy-fail")}}
	}

	toInt, _ := strconv.Atoi(data.Number)
	data.OddStatus = toInt%2 != 0

	if err := dg.I.Save(&data).Error; err != nil {
		ed := sh.Event{
			Feature: "phone-number",
			Action:  "create-data",
			Source:  source,
			Status:  "failed",
			ECode:   "data-create-fail",
			EDetail: err.Error(),
		}
		return nil, sh.SetError(ed, nil)
	}

	return data, nil
}

func Delete(id int, dto m.DeleteDto) (*string, error) {
	data := &m.PhoneNumber{}

	// if err := dg.I.Where("Id = ? AND User_Id = ?", id, dto.User_Id).First(data).Error; err != nil {
	if err := dg.I.Where("Id = ?", id).First(data).Error; err != nil {
		return nil, returnFetchError(err)
	}
	if err := dg.I.Delete(data).Error; err != nil {
		return nil, err
	}

	msg := l.I.Msg("data-delete-success")
	return &msg, nil
}

func GetList(input m.FilterListDto) (*td.Data, error) {
	data := []m.PhoneNumberForList{}
	pagination := gh.Pagination{
		Page:     input.Page,
		PageSize: input.PageSize,
	}
	count := int64(0)

	err := dg.I.
		Model(&m.PhoneNumber{}).
		Select("phonenumber.Id, phonenumber.Number, phonenumber.User_Id, provider.Name Provider_Name").
		Joins("JOIN provider ON phonenumber.Provider_Id=provider.Id").
		Scopes(gh.Filter(input)).
		Count(&count).
		Scopes(gh.Paginate(input, &pagination)).
		Find(&data).
		Error
	if err != nil {
		return nil, sh.SetError(sh.Event{
			Feature: "phone-number",
			Action:  "get-list",
			Source:  source,
			Status:  "failed",
			ECode:   "data-getList-fail",
			EDetail: err.Error(),
		}, nil)

	}

	return &td.Data{
		Meta: td.II{
			"page":     pagination.Page,
			"pageSize": pagination.PageSize,
			"count":    count,
		},
		Data: data,
	}, nil
}

func GetDetail(input m.FilterDetailDto) (*m.PhoneNumber, error) {
	data := &m.PhoneNumber{}

	err := dg.I.
		Model(&m.PhoneNumber{}).
		Scopes(gh.Filter(input)).
		Find(&data).
		Error
	if err != nil {
		return nil, sh.SetError(sh.Event{
			Feature: "phone-number",
			Action:  "get-list",
			Source:  source,
			Status:  "failed",
			ECode:   "data-getList-fail",
			EDetail: err.Error(),
		}, nil)

	}

	return data, nil
}

func GenRandom(input m.GenRandomDto) (string, error) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	data := &m.PhoneNumber{}
	for i := 0; i < input.Count; i++ {
		mandatory := (strconv.Itoa(1+r1.Intn(9999999)) + "0000000")[0:9]
		optional := strconv.Itoa(r1.Intn(9999))[0:1]
		provider_id := r1.Intn(4)
		toInt, _ := strconv.Atoi(data.Number)
		data = &m.PhoneNumber{
			Number:      "08" + mandatory + optional,
			Provider_Id: provider_id,
			User_Id:     input.User_Id,
			OddStatus:   toInt%2 != 0,
		}
		if err := dg.I.Create(&data).Error; err != nil {
			ed := sh.Event{
				Feature: "phone-number",
				Action:  "create-data",
				Source:  source,
				Status:  "failed",
				ECode:   "data-create-fail",
				EDetail: err.Error(),
			}
			return "", sh.SetError(ed, nil)
		}
	}
	return "data berhasil di random", nil
}

func returnFetchError(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return te.XError{Code: "data-notFound", Message: l.I.Msg("data-notFound")}
	} else {
		return err
	}
}
