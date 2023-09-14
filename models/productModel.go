package models

import "gorm.io/gorm"

type Product struct{
	gorm.Model
	
	ID string `gorm:"not null" json:"firstName"`
	Name  string `gorm:"not null" json:"lastName"`
	Price  string `gorm:"not null" json:"lastName"`
	
}