package phonenumber

import (
	"errors"

	sc "github.com/jinzhu/copier"
	dg "github.com/karincake/apem/databasegorm"
	l "github.com/karincake/apem/lang"
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

	if err := dg.I.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&data).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
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

	if err := dg.I.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(&data).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
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

func returnFetchError(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return te.XError{Code: "data-notFound", Message: l.I.Msg("data-notFound")}
	} else {
		return err
	}
}
