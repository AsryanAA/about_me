package main

import (
	docs "about_me/api"
	"about_me/internal/config"
	"about_me/internal/http-server/handlers"
	"about_me/internal/http-server/view"
	"about_me/internal/lib/logger/sl"
	"about_me/internal/storage/sqlite"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	// TODO: init config: cleanenv
	cfg := config.MustLoad()

	// TODO: init logger: slog
	log := setupLogger(cfg.Env)
	log.Info("starting about-me", slog.String("env", cfg.Env))
	log.Debug("debug messages")

	/*
		id, err := storage.AddWorkPlace("ЧОУ ООШ Максимовой 'Улыбка'", "01.10.2015", "29.02.2016",
			"Учитель математики и информатики")
		fmt.Println(id, err)

		id, err = storage.AddWorkPlace("ООО Центр подготовки 'Супер'", "01.11.2015", "30.04.2016",
			"Репетитор математики и информатики")
		fmt.Println(id, err)
	*/

	// TODO: init storage: SQLite
	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}

	// TODO: init router: gin
	router := gin.Default()

	api := router.Group("api")
	{
		v1 := api.Group("v1")
		{
			workPlace := v1.Group("work_place")
			{
				workPlace.POST("create", handlers.CreateWorkPlace(storage))
				workPlace.GET("read", handlers.ReadWorkPlaces(storage))
				workPlace.PATCH("update", handlers.UpdateWorkPlace(storage))
			}
		}
	}

	docs.SwaggerInfo.Title = "Автодокументация к приложению Онлайн резюме"
	docs.SwaggerInfo.BasePath = "api/v1"

	// use ginSwagger middleware to serve the API docs
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.LoadHTMLGlob("templates/*.html")
	router.Static("/view/css", "templates/css")

	pages := router.Group("view")
	{
		pages.GET("/home_page", view.ReadWorkPlaces(storage))
	}

	log.Info("starting server", slog.String("on address & port", fmt.Sprintf("%s:%d", cfg.Address, cfg.Port)))

	// TODO: run server
	err = router.Run(fmt.Sprintf("%s:%d", cfg.Address, cfg.Port))
	if err != nil {
		log.Error("can not start server", sl.Err(err))
		os.Exit(1)
	}

	log.Error("server stopped")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}
