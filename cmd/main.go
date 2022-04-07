package main

import (
	"log"

	"github.com/hkaya15/PicusSecurity/Final_Project/pkg/base/config"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/base/db"
	logger "github.com/hkaya15/PicusSecurity/Final_Project/pkg/base/log"
	"go.uber.org/zap"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Loadconfig Failed: %v", err)
	}

	logger.NewLogger(cfg)
	defer logger.Close()

	// It is possible to integrate different db technologies
	base := DBBase{DbType: &POSTGRES{}}
	db, err := base.DbType.Create(cfg)

	if err != nil {
		zap.L().Fatal("DB cannot init", zap.Error(err))
	}
	log.Println(db)
	
}
