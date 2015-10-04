package model
import "time"

type User struct {
	ID       int `sql:"AUTO_INCREMENT"`
	Email    string
	Password []byte
	Created  time.Time `sql:"timestamp"`
	Network  string
	Projects []Project `gorm:"many2many:user_projects;"`
}

type Project struct {
	ID   int `sql:"AUTO_INCREMENT"`
	Name string
	Code string
}
