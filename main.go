package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/wahyurudiyan/super-sharing/user-svc/config"
	"github.com/wahyurudiyan/super-sharing/user-svc/internal/controller"
	userRepo "github.com/wahyurudiyan/super-sharing/user-svc/internal/core/repository/users"
	userSvc "github.com/wahyurudiyan/super-sharing/user-svc/internal/core/services/users"
)

func main() {
	cfg := config.Get()
	db, err := sqlx.Connect("mysql", generateMysqlDatasource(cfg))
	if err != nil {
		logrus.Infoln(generateMysqlDatasource(cfg))
		logrus.Panicln(err)
		return
	}

	router := fiber.New()

	userRepo := userRepo.NewUserRepository(db)
	userSvc := userSvc.NewUserService(userRepo)

	controller := controller.NewController(router, userSvc)
	controller.Init()

	start(cfg, router)
}

func generateMysqlDatasource(cfg *config.Config) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", cfg.Mysql.Username, cfg.Mysql.Password, cfg.Mysql.Host, cfg.Mysql.Port, cfg.Mysql.Database)
}

func start(cfg *config.Config, router *fiber.App) {
	// Graceful Shutdown
	wait := make(chan os.Signal, 1)
	signal.Notify(wait, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithTimeout(context.Background(), cfg.App.Timeout)
	defer cancel()

	go func() {
		for {
			select {
			case <-ctx.Done():
				router.ShutdownWithContext(ctx)
				logrus.Infoln("service shutdown")
				return
			default:
				logrus.Infoln("service started")
				err := router.Listen(":" + cfg.App.Port)
				if err != nil {
					logrus.Panicln(err)
				}
			}
		}
	}()

	<-wait
}
