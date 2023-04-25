package model

type Application struct {
	Id          int `gorm:"primary_key:true"`
	Name        string
	Code        string
	Comment     string
	Platform    string
	Port        string
	RegistryUrl string
	ConfigUrl   string
	Interfaces  []*Interface `gorm:"many2many:application_interface;"`
}
