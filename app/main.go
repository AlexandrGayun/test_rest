package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
	"test_task/app/rest"
	"test_task/business/service"
	"test_task/business/storage"
	"test_task/utils/config"
	"test_task/utils/logger"
)

func main() {
	l := logger.New()
	var cfg config.Config
	if err := godotenv.Load(".env"); err != nil {
		l.Fatal("can't load .env file", zap.Error(err))
	}
	if err := envconfig.Process("", &cfg); err != nil {
		l.Fatal("can't read OS env", zap.Error(err))
	}
	db, err := sql.Open("mysql", cfg.MysqlDsn)
	if err != nil {
		l.Fatal("can't establish database connection", zap.Error(err))
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		l.Fatal("database connection credentials are not valid", zap.Error(err))
	}
	queries := storage.New(db)
	st := storage.NewStorage(queries, l)
	s := service.New(l, st)
	restServer := rest.New(l, s, cfg)
	restServer.Start()
}
