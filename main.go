package main

import (
	"Bookshop/src/server"
	"Bookshop/src/services"
	"fmt"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"

	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type config struct {
	Port       string `env:"PORT" envDefault:"3000"`
	AppVersion string `env:"APP_VERSION" envDefault:"v0.0.0"`
	POSTGRES   struct {
		Port     string `env:"POSTGRES_PORT" envDefault:"5432"`
		Host     string `env:"POSTGRES_HOST" envDefault:"localhost"`
		User     string `env:"POSTGRES_USER" envDefault:"backend"`
		Password string `env:"POSTGRES_PASSWORD" envDefault:"1234567890"`
		Database string `env:"POSTGRES_DATABASE" envDefault:"bookshop"`
		SslMode  string `env:"POSTGRES_SSL_MODE" envDefault:"disable"`
	}
}

func main() {
	cfg := initEnvironment()
	db := initDB(cfg)
	initFiberServer(cfg, db)
}

func initFiberServer(cfg config, db services.GormPg) {
	fiberSv := server.New(&server.ServerConfig{
		AppVersion:    cfg.AppVersion,
		ListenAddress: fmt.Sprintf(":%s", cfg.Port),
	}, db)

	fiberSv.Start()
}
func initEnvironment() config {
	// load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// set config
	// ให้ตัวแปรกับenv match กัน
	var cfg config
	err = env.Parse(&cfg)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	return cfg
}

func initDB(cfg config) services.GormPg {
	// connect database
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.POSTGRES.Host,
		cfg.POSTGRES.Port,
		cfg.POSTGRES.User,
		cfg.POSTGRES.Password,
		cfg.POSTGRES.Database,
		cfg.POSTGRES.SslMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	bookPg := services.SetUpPosgresql(db)
	return bookPg
}
