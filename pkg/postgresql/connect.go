package connect

import (
	"database/sql"
	"fmt"
	"webPractice1/pkg/errorPrinter"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
	_ "github.com/spf13/viper"
)

type Config struct {
	DB Postgres
}

type Postgres struct {
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
	Sslmode  string
}

func PostgresqlConnect() *sql.DB {
	if err := godotenv.Load("../.env"); err != nil {
		errorPrinter.PrintCallerFunctionName(err)
	}
	cfg := new(Config)
	if err := envconfig.Process("db", &cfg.DB); err != nil {
		errorPrinter.PrintCallerFunctionName(err)
	}
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.Dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		errorPrinter.PrintCallerFunctionName(err)
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		errorPrinter.PrintCallerFunctionName(err)
		return nil
	}

	fmt.Println("Successfully connected!")
	return db
}
