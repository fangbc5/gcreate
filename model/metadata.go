package model

type Metadata struct {
	Id         int `gorm:"primary_key:true"`
	Name       string
	Code       string
	Comment    string
	Meta       string
}
