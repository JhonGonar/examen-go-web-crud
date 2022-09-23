package models

import "gorm.io/gorm"

type Odontologo struct {
	gorm.Model
	Nombre    string `gorm:"type:varchar(30);not null"`
	Apellido  string `gorm:"type:varchar(50);not null"`
	Matricula string `gorm:"type:varchar(10);not null;unique_index"`
}
