package model

import (
	"github.com/jinzhu/gorm"
)

type (
	Project struct {
		gorm.Model
		Name     string
		Networks []Network
		Devices  []Device
	}

	Network struct {
		gorm.Model
		ProjectID uint
		Name      string
		Devices   []Device `gorm:"many2many:net_dev;"`
	}

	Device struct {
		gorm.Model
		ProjectID uint
		Name      string
		Networks  []Network `gorm:"many2many:net_dev;"`
	}
)
