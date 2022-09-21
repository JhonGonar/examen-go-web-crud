package models

import "gorm.io/gorm"

type Odontologo struct {
	gorm.Model
	Nombre    string `gorm:"type:varchar(30);not null"`
	Apellido  string `gorm:"ype:varchar(50);not null"`
	Matricula string `gorm:"ype:varchar(10);not null;unique_index"`
}
