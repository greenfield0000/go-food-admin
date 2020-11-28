package database

import (
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

// OpenDB - Открытие бд
func OpenDB() (db *sqlx.DB, err error) {
	dataBaseConf := "host=127.0.0.1 user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " port=" + dbPort + " sslmode=" + sslMode + " TimeZone=" + dbTimeZone
	return sqlx.Connect("postgres", dataBaseConf)
}

// startAutoMigrate - Function with start automigrate by struct
func StartAutoMigrate() {
	db, err := OpenDB()
	if err != nil {
		log.Fatal("startAutoMigrate is error = ", err.Error())
		return
	}
	// Запуск миграции
	db.MustExec(Schema)
	if err != nil {
		log.Fatal("startAutoMigrate is error = ", err.Error())
		return
	}
}
