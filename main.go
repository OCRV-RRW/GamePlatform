package main

import (
	"context"
	"fmt"
	"path/filepath"
	"time"

	"log/slog"
	"os"

	"gameplatform/internal/DTO"
	"gameplatform/internal/api/v1/auth"
	"gameplatform/internal/api/v1/game"
	"gameplatform/internal/api/v1/user"
	"gameplatform/internal/config"
	"gameplatform/internal/view"

	"gameplatform/internal/dbconn"
	"gameplatform/internal/html"
	flogger "gameplatform/internal/logger"
	"gameplatform/internal/middleware"
	"gameplatform/internal/repository"
	"gameplatform/internal/utils"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})
	fh := flogger.FiberHandler{Handler: handler}
	logger := slog.New(fh)
	slog.SetDefault(logger)

	conf, err := config.Load()
	if err != nil {
		logger.Error(fmt.Sprintf("Couldn't load config: %v", err))
		os.Exit(1)
	}

	ctx := context.Background()
	database_conn := dbconn.NewDatabaseConnection(conf.DatabaseURL, ctx)
	defer database_conn.CloseConnection(ctx)
	redis_conn := dbconn.NewRedisConnection(ctx, conf.RedisHost, conf.RedisPort)
	defer redis_conn.Close()
	minio_conn := dbconn.NewMinioConnection(conf)

	user_repository := repository.NewUserRepository(&database_conn, &repository.UserConverterImpl{})
	gameRepository := repository.NewGameRepository(&database_conn, &repository.GameConverterImpl{})

	smtp := utils.SMTP{
		EmailFrom: conf.EmailFrom,
		User:      conf.SMTPUser,
		Pass:      conf.SMTPPass,
		Host:      conf.SMTPHost,
		Port:      conf.SMTPPort,
	}

	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	root, err := findRootDir()
	engine := html.New(filepath.Join(root, "ui/html"), conf.Debug)
	engine.Load()

	app := fiber.New(fiber.Config{
		Views: &engine,
	})

	micro := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowHeaders: "Origin,Content-Type,Accept,Content-Length,Accept-Language," +
			"Accept-Encoding,Connection,Access-Control-Allow-Origin,Authorization",
		AllowOrigins:     "http://localhost:3000,http://localhost:8000,https://*.ocrv-game.ru,https://ocrv-game.ru",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	var staticCacheDuration time.Duration
	if conf.Debug {
		staticCacheDuration = time.Millisecond * 50
	} else {
		staticCacheDuration = time.Minute * 3
	}

	app.Static("/static", "ui/static", fiber.Static{
		CacheDuration: staticCacheDuration,
	})

	swagger_conf := swagger.Config{
		BasePath: "/",
		FilePath: "./docs/swagger.json",
		Path:     "swagger",
		Title:    "Swagger API Docs",
	}

	slog.Info(fmt.Sprintf("Cookie Secure: %v", conf.CookieSecure))

	app.Use(swagger.New(swagger_conf))
	loggerMiddelware := flogger.NewLoggerMiddelware(conf.Debug)
	app.Use(loggerMiddelware.Handle)

	viewHandler := view.NewViewHanlder(gameRepository, &view.ViewConverterImpl{})
	view.AddRoutes(app, &viewHandler)

	app.Mount("/api/v1", micro)
	userMiddleware := middleware.NewUserMiddleware(user_repository, conf)

	authHandler := auth.NewAuthHandler(user_repository, &redis_conn, &smtp, conf)
	auth.AddRoutes(micro, &authHandler, &userMiddleware)

	convImpl := &DTO.UserConverterImpl{}
	userHandler := user.NewUserHandler(conf, &smtp, user_repository, &redis_conn, convImpl)
	user.AddRoutes(micro, userHandler, &userMiddleware)

	gameConvImp := &DTO.GameConverterImpl{}
	gameHandler := game.NewGameHandler(conf, gameRepository, &redis_conn, &minio_conn, gameConvImp)
	game.AddRoutes(micro, gameHandler, &userMiddleware)

	micro.Get("/", func(c *fiber.Ctx) error {
		slog.InfoContext(c.Context(), "Hellooooo")
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"message": "Hello",
		})
	})

	micro.Get("/healthchecker", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"message": "good",
		})
	})

	micro.All("*", func(c *fiber.Ctx) error {
		path := c.Path()
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "fail",
			"message": fmt.Sprintf("Path: %v does not exists on this server", path),
		})
	})

	app.Listen(conf.Host)
}

func findRootDir() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			println(dir)
			return dir, nil
		}

		parentDir := filepath.Dir(dir)
		if parentDir == dir {
			return "", os.ErrNotExist
		}
		dir = parentDir
	}
}
