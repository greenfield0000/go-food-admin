package database

import (
	"log"

	"github.com/greenfield0000/go-food/microservices/go-food-admin/database/migration"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	dbHost     = "central-db"
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
	dataBaseConf := "host=" + dbHost + " user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " port=" + dbPort + " sslmode=" + sslMode + " TimeZone=" + dbTimeZone
	connect, err := sqlx.Connect("postgres", dataBaseConf)

	DatabaseHolder.Db = connect
	//DatabaseHolder.Db.SetMaxIdleConns(10)
	//DatabaseHolder.Db.SetMaxOpenConns(100 )

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
