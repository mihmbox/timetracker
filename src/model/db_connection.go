package model
import (
	_ "github.com/lib/pq"
	"github.com/jinzhu/gorm"
)

const (
	CONNECTION_STRING = "host=localhost port=5432 user=postgres password=root dbname=TimeTrack sslmode=disable"
	DB_DIALECT = "postgres"
)

func GetConnection() (gorm.DB, error) {
	return gorm.Open(DB_DIALECT, CONNECTION_STRING)
}

func CreateDB() error {
	if db, err := GetConnection(); err == nil {
		db.CreateTable(&User{})
		db.CreateTable(&Project{})
		db.Close()
		return nil
	} else {
		return err
	}
}