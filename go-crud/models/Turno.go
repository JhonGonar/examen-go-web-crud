package models

import (
	"gorm.io/gorm"
)

type Turno struct {
	gorm.Model
	Fecha        string `gorm:"not null"`
	Hora         string `gorm:"not null"`
	Descripcion  string `gorm:"not null"`
	PacienteID   uint
	OdontologoID uint
}
