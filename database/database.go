package database

import (
	"github.com/greenfield0000/go-food/microservices/go-food-admin/database/migration"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

const (
	dbName     = "central-db"
	dbUser     = "admin"
	dbPassword = "admin"
	dbPort     = "5432"
	sslMode    = "disable"
	dbTimeZone = "Europe/Moscow"
)

// databaseHolder databaseHolder connection database
type databaseHolder struct {
	Db *sqlx.DB
}

var DatabaseHolder = databaseHolder{}

// ConnectDB - Открытие бд
func ConnectDB() (databaseHolder, error) {
	dataBaseConf := "host=127.0.0.1 user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " port=" + dbPort + " sslmode=" + sslMode + " TimeZone=" + dbTimeZone
	connect, err := sqlx.Connect("postgres", dataBaseConf)

	DatabaseHolder.Db = connect
	return DatabaseHolder, err
}

// StartAutoMigrate - Function with start automigrate by struct
func StartAutoMigrate() {
	holder, err := ConnectDB()
	if err != nil {
		log.Fatal("startAutoMigrate is error = ", err.Error())
		return
	}
	// Запуск миграции
	holder.Db.MustExec(migration.Schema)
	if err != nil {
		log.Fatal("startAutoMigrate is error = ", err.Error())
		return
	}
}
