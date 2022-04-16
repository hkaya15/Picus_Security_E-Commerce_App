package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hkaya15/PicusSecurity/Final_Project/pkg/base/config"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/base/errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type StatusHandler struct {
	cfg *config.Config
	db  *gorm.DB
}

func NewStatusHandler(r *gin.RouterGroup, cfg *config.Config, database *gorm.DB) {
	h := &StatusHandler{cfg: cfg, db: database}
	r.GET("/", h.checkstatus)
}

func (s *StatusHandler) checkstatus(c *gin.Context) {
	db, err := s.db.DB()
	if err != nil {
		zap.L().Fatal("Cannot get sql database instance", zap.Error(err))
		c.JSON(ErrorResponse(err))
		return
	}
	if err := db.Ping(); err != nil {
		zap.L().Fatal("Cannot ping database", zap.Error(err))
		c.JSON(ErrorResponse(err))
		return
	}
	zap.L().Info("DB HEALTH CHECK OK :",zap.String("DB:",fmt.Sprintf("Postgres: %v",db.Stats())) )
	c.JSON(http.StatusOK, false)
}
