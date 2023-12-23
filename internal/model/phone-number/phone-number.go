package phonenumber

type PhoneNumber struct {
	Id          int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Number      string `json:"number" gorm:"size:30"`
	Provider_Id int    `json:"provider_id"`
}

type CreateDto struct {
	Number      string `json:"number" validate:"required;minLength=10;maxLength=25"`
	Provider_Id int    `json:"provider_id" validate:"required"`
}

type UpdateDto struct {
	Number      string `json:"number" validate:"required;minLength=10;maxLength=25"`
	Provider_Id int    `json:"provider_id" validate:"required"`
}

type DeleteDto struct {
	User_Id *int `json:"-"`
}
