package model

import "gorm.io/gorm"

type RoleCatalogTable struct {
	gorm.Model
	RoleId   string `gorm:"uniqueIndex"`
	RoleCode string
	RoleDesc string
	Lob      string
}

type RoleByLobTable struct {
	gorm.Model
	Lob     string `gorm:"unique"`
	RoleIds []string
}
