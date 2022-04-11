package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/user/handler"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/user/repository"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/user/service"
	"github.com/hkaya15/PicusSecurity/Final_Project/pkg/base/config"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/base/db"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/base/graceful"
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

	gin.SetMode(gin.ReleaseMode)
	g := gin.Default()

	g.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// your custom format
		message := fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
		zap.L().Info(message)
		return message
	}))

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.ServerConfig.Port),
		Handler:      g,
		ReadTimeout:  time.Duration(cfg.ServerConfig.ReadTimeoutSecs * int64(time.Second)),
		WriteTimeout: time.Duration(cfg.ServerConfig.WriteTimeoutSecs * int64(time.Second)),
	}

	rootRouter := g.Group(cfg.ServerConfig.RoutePrefix)
	authRooter := rootRouter.Group("/user")

	userRepo := NewUserRepository(db)
	userService := NewUserService(userRepo)
	NewUserHandler(authRooter, userService, cfg)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Fatal("main.listenandserve: ", zap.Error(err))
		}
	}()
	zap.L().Debug("Server started")
	ShutdownGin(srv, time.Duration(cfg.ServerConfig.TimeoutSecs*int64(time.Second)))

}
