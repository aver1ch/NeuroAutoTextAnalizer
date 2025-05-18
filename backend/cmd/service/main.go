package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/kerilOvs/backend/internal/handlers"
	"github.com/kerilOvs/backend/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/kerilOvs/backend/internal/config"
	"github.com/kerilOvs/backend/pkg/logger"

	"github.com/kerilOvs/backend/internal/service"
	"github.com/kerilOvs/backend/internal/storage/minio"
	postgresstorage "github.com/kerilOvs/backend/internal/storage/postgres"
)

func main() {
	fmt.Println("Hi, i'm server:)")

	log := logger.Init("text", "debug")
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Error("Failed to read config:", slog.Any("error", err))
	}
	log.Info("Read config", slog.Any("config", cfg))

	log.Info("Connecting to db...")
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.Database.Host,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Dbname,
		strconv.Itoa(cfg.Database.Port),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error("Failed to connect to database:", slog.Any("error", err))
	}

	log.Info("Running migrations...")
	if err := db.AutoMigrate(
		&models.User{},
		&models.UserDoc{},
	); err != nil {
		log.Error("Failed to migrate database:", slog.Any("error", err))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Info("Connection to minio")
	minioClient, err := minio.New(ctx, cfg.Minio)
	if err != nil {
		log.Error("Failed to initialize MinIO client:", slog.Any("error", err))
	}

	userStorage := postgresstorage.NewUserPostgresStorage(db)
	userService := service.NewUserService(userStorage)

	docService := service.NewDocService(minioClient, cfg.Minio.Bucket)
	docHandler := handlers.NewDocHandler(userService, docService)

	e := echo.New()
	e.Use(handlers.Logging(log))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
			http.MethodPatch,
		},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAuthorization,
		},
	}))

	userHandler := handlers.NewUserHandler(userService)
	registerRoutes(e, userHandler, docHandler)

	serverAddr := ":" + strconv.Itoa(cfg.Server.Port)
	log.Info("Server started", slog.String("port", serverAddr))
	log.Error("Server stopped", slog.Any("error", e.Start(serverAddr)))
}

func registerRoutes(e *echo.Echo, userHandler *handlers.UserHandler, docHandler *handlers.DocHandler) {

	e.POST("/users", userHandler.CreateUser)                     // +
	e.DELETE("/users/:id", userHandler.DeleteUser)               // +
	e.GET("/users/:id", userHandler.GetUserById)                 // +
	e.PATCH("/users/:id/profile", userHandler.UpdateUserProfile) // +

	e.GET("/users/:id/docs", userHandler.GetUserDocs) // +
	//e.PUT("/users/:id/photos", userHandler.AddUserPhoto)
	e.DELETE("/users/:id/docs/:docId", userHandler.RemoveUserDoc) //

	// Фото маршруты
	e.POST("/users/:id/addphoto", docHandler.UploadDoc) //  по айди юзера добавляет доку
	e.GET("/photos/:id", docHandler.GetDoc)             //  по айди !doc! отдает доку
}
