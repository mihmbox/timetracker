package model

import (
	"app"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func GetConnection() (gorm.DB, error) {
	cfg := app.App.Config.Db
	return gorm.Open(cfg.Dialect, cfg.ConnectionString)
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
