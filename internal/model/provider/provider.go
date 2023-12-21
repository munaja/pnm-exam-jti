package provider

type Provider struct {
	Id   int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Code string `json:"code" gorm:"size:10"`
	Name string `json:"name" gorm:"size:50"`
}
