<<<<<<< HEAD
package main

import (
	"bank"
	"bank/pkg/handler"
	"bank/pkg/repository"
	"bank/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error while initialization config : %s", err.Error())
	}
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(bank.Server)
	if err := srv.Run(viper.GetString("8080"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error while running server")
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
=======
package main

import (
	"bank"
	"bank/pkg/handler"
	"bank/pkg/repository"
	"bank/pkg/service"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func main() {
	if err := initConfig(); err != nil {
		logrus.Fatalf("error while initialization config : %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error while loading env param : %s", err.Error())

	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("error while db init : %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(bank.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error while running server")
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
>>>>>>> f48611f (first commit)
