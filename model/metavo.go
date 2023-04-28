package model

type Metavo struct {
	Id         int `gorm:"primary_key:true"`
	Name       string
	Code       string
	Comment    string
	Meta       string
}