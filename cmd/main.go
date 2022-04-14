package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/category/handler"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/category/repository"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/category/service"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/user/handler"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/user/repository"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/user/service"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/product/handler"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/product/repository"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/product/service"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/cart/handler"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/cart/repository"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/cart/service"
	"github.com/hkaya15/PicusSecurity/Final_Project/pkg/base/config"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/base/db"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/base/graceful"
	logger "github.com/hkaya15/PicusSecurity/Final_Project/pkg/base/log"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/base/middleware"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Loadconfig Failed: %v", err)
	}

	logger.NewLogger(cfg)
	defer logger.Close()

	errload := godotenv.Load("../.env")
	if errload != nil {
		zap.L().Fatal("Error loading .env file")
	}

	// It is possible to integrate different db technologies
	base := DBBase{DbType: &POSTGRES{}}
	db, err := base.DbType.Create(cfg)
	if err != nil {
		zap.L().Fatal("DB cannot init", zap.Error(err))
	}

	gin.SetMode(gin.ReleaseMode)
	g := gin.Default()
	LoggerMiddleware(g)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.ServerConfig.Port),
		Handler:      g,
		ReadTimeout:  time.Duration(cfg.ServerConfig.ReadTimeoutSecs * int64(time.Second)),
		WriteTimeout: time.Duration(cfg.ServerConfig.WriteTimeoutSecs * int64(time.Second)),
	}

	getUp(g, db, cfg)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Fatal("main.listen and serve: ", zap.Error(err))
		}
	}()
	zap.L().Debug(os.Getenv("START_SERVER"))
	ShutdownGin(srv, time.Duration(cfg.ServerConfig.TimeoutSecs*int64(time.Second)))

}

func getUp(g *gin.Engine, db *gorm.DB, cfg *config.Config) {
	rootRouter := g.Group(cfg.ServerConfig.RoutePrefix)
	userRooter := rootRouter.Group("/user")
	categoryRooter := rootRouter.Group("/category")
	productRooter := rootRouter.Group("/product")
	cartRooter := rootRouter.Group("/cart")

	userRepo := NewUserRepository(db)
	userService := NewUserService(userRepo)
	NewUserHandler(userRooter, userService, cfg)

	categoryRepo := NewCategoryRepository(db)
	categoryService := NewCategoryService(categoryRepo)
	NewCategoryHandler(categoryRooter, categoryService, cfg)

	productRepo := NewProductRepository(db)
	productService := NewProductService(productRepo)
	NewProductHandler(productRooter, productService, cfg)

	cartRepo := NewCartRepository(db)
	cartService := NewCartService(cartRepo, productRepo)
	NewCartHandler(cartRooter, cartService, cfg)

}
