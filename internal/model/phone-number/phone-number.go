package phonenumber

import (
	sv "github.com/karincake/serabi"
)

type PhoneNumber struct {
	Id          int    `json:"id" gorm:"primaryKey;autoIncrement"`
	User_Id     int    `json:"user_id"`
	Number      string `json:"number" gorm:"size:30"`
	OddStatus   bool   `json:"oddStatus"`
	Provider_Id int    `json:"provider_id"`
}

type PhoneNumberForList struct {
	Id            int    `json:"id"`
	User_Id       int    `json:"user_id"`
	Number        string `json:"number"`
	OddStatus     bool   `json:"oddStatus"`
	Provider_Id   int    `json:"provider_id"`
	Provider_Name string `json:"provider_name"`
}

type CreateDto struct {
	User_Id     int    `json:"-"`
	Number      string `json:"number" validate:"required;minLength=10;maxLength=25;numeric;phoneNumberFormat"`
	Provider_Id int    `json:"provider_id" validate:"required"`
}

type UpdateDto struct {
	User_Id     int    `json:"-"`
	Number      string `json:"number" validate:"required;minLength=10;maxLength=25;numeric;phoneNumberFormat"`
	Provider_Id int    `json:"provider_id" validate:"required"`
}

type DeleteDto struct {
	User_Id int `json:"-"`
}

type FilterListDto struct {
	User_Id       int     `json:"-"`
	OddStatus     bool    `json:"oddStatus"`
	OddStatus_Opt *string `json:"oddStatus_opt"`
	Page          int     `json:"page"`
	PageSize      int     `json:"pagesize"`
}

type FilterDetailDto struct {
	Id      int `json:"id"`
	User_Id int `json:"-"`
}

type GenRandomDto struct {
	User_Id int `json:"-"`
	Count   int `json:"count"`
}

func init() {
	sv.AddTagForRegex("phoneNumberFormat", `^08[0-9]{8,14}$`, "should have valid phone number format: 08XXXXXXXX[XXXXXX], X is numeric, total length between 10-16")
}
