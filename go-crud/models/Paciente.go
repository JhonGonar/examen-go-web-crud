package models

import "gorm.io/gorm"

type Paciente struct {
	gorm.Model
	Nombre    string `gorm:"type:varchar(30);not null"`
	Apellido  string `gorm:"type:varchar(50);not null"`
	Domicilio string `gorm:"type:varchar(50);not null"`
	DNI       string `gorm:"type:varchar(10);not null;unique_index"`
	FechaAlta string `gorm:"not null"`
}
