package model

type User struct {
	ID       int       `sql:"AUTO_INCREMENT"`
	Email    string
	Password []byte
	Projects []Project `gorm:"many2many:user_projects;"`
}

type Project struct {
	ID   int       `sql:"AUTO_INCREMENT"`
	Name string
	Code string
}