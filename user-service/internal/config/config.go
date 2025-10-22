package config

import (
	"auth-service/internal/api/handlers"
	"auth-service/internal/db/repositories"
	"flag"
	"os"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Config struct {
	Port string
	Env  string
	DB   struct {
		MONGODB_URI string
	}
}

func NewConfig() *Config {
	return &Config{}
}

func (cfg *Config) ParseFlags() error {
	flag.StringVar(&cfg.Port, "port", os.Getenv("PORT"), "API server port")
	flag.StringVar(&cfg.Env, "env", os.Getenv("ENV"), "Environment (development|staging|production)")

	flag.StringVar(&cfg.DB.MONGODB_URI, "mongodb_uri", os.Getenv("MONGODB_URI"), "MONGODB_URI")

	// flag.IntVar(&cfg.DB.MaxOpenConns, "db-max-open-conns", 25, "PostgreSQL max open connections")
	// flag.IntVar(&cfg.DB.MaxIdleConns, "db-max-idle-conns", 25, "PostgreSQL max idle connections")
	// flag.StringVar(&cfg.DB.MaxIdleTime, "db-max-idle-time", "15m", "PostgreSQL max connection idle time")

	// flag.Float64Var(&cfg.Limiter.RPS, "limiter-rps", 2, "Rate limiter maximum requests per second")
	// flag.IntVar(&cfg.Limiter.Burst, "limiter-burst", 4, "Rate limiter maximum burst")
	// flag.BoolVar(&cfg.Limiter.Enabled, "limiter-enabled", true, "Enable rate limiter")

	// flag.StringVar(&cfg.SMTP.Host, "smtp-host", os.Getenv("SMTP_HOST"), "SMTP host")
	// flag.IntVar(&cfg.SMTP.Port, "smtp-port", 25, os.Getenv("SMTP_PORT"))
	// flag.StringVar(&cfg.SMTP.Username, "smtp-username", os.Getenv("SMTP_USERNAME"), "SMTP username")
	// flag.StringVar(&cfg.SMTP.Password, "smtp-password", os.Getenv("SMTP_PASSWORD"), "SMTP password")
	// flag.StringVar(&cfg.SMTP.Sender, "smtp-sender", "openMovie <no-reply@test.user.net>", "SMTP sender")

	return nil
}

func (c *Config) InitializeHandlers(r *repositories.Repositories) *handlers.Handlers {
	return handlers.NewHandlers(r.UserRepository)
}

func (c *Config) InitializeRepositories(db *mongo.Database) *repositories.Repositories {
	return repositories.NewRepositories(db)
}
