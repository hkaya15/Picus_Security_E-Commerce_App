package db

import (
	"github.com/hkaya15/PicusSecurity/Final_Project/pkg/base/config"
	"gorm.io/gorm"
)

type DBSelector interface {
	Create(config *config.Config) (*gorm.DB, error)
}

type DBBase struct {
	DbType DBSelector
}

