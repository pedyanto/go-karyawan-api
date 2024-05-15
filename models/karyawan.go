package models

import (
	"gorm.io/gorm"
)

type Karyawan struct {
    gorm.Model
    Nama    string `gorm:"column:nama"`
    NoHP    string `gorm:"column:no_hp"`
    Email   string `gorm:"column:email"`
    Jabatan string `gorm:"column:jabatan"`
}

// Override the default table name
func (Karyawan) TableName() string {
    return "karyawan"
}
