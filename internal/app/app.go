package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mant1COREX/pet-project/configs"
	"github.com/mant1COREX/pet-project/internal/handlers"
	"github.com/mant1COREX/pet-project/internal/repository"
	"github.com/mant1COREX/pet-project/internal/service"
	"github.com/mant1COREX/pet-project/pkg/postgres"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Run() {
	//logrus
	logrus.SetFormatter(new(logrus.JSONFormatter))

	//Configs
	if err := configs.InitConfig(); err != nil {
		logrus.Fatalf("error initialization configs: %s", err.Error())
	}

	//.env
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env vars: %s", err.Error())
	}

	//DB
	pool, err := postgres.NewPG(context.Background(), postgres.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("failed init db: %s", err.Error())
	}

	defer pool.Close()

	// Repositories
	repos := repository.NewRepository(pool)

	// Service
	services := service.NewService(repos)

	// Handlers
	handlers := handlers.NewHandler(services)

	logrus.Print("5")
	//HTTP server
	app := fiber.New(fiber.Config{
		Prefork:       false,   // включаем предварительное форкование для увеличения производительности на многоядерных процессорах
		ServerHeader:  "Fiber", // добавляем заголовок для идентификации сервера
		CaseSensitive: true,    // включаем чувствительность к регистру в URL
		StrictRouting: true,    // включаем строгую маршрутизацию
	})

	logrus.Print("6")
	handlers.InitRoutes(app)

	logrus.Print("7")
	logrus.Fatal(app.Listen(":" + viper.GetString("port")))

	logrus.Print("8")
	logrus.Print("Pet project test task API started")

	//gracefull shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // Перехватываем SIGINT и SIGTERM
	<-quit                                               // Ждем сигнала

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := app.ShutdownWithContext(ctx); err != nil {
		logrus.Fatalf("Ошибка при остановке сервера: %v", err)
	}

	logrus.Println("Server stopping")

}
