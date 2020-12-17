package config

import (
	"fmt"
	"github.com/cui-bo/keypass/models"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

type BDD interface {
	DBUser
}

type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     "localhost",
		Port:     3306,
		User:     "root",
		Password: "laposte",
		DBName:   "formation_go",
	}
	return &dbConfig
}

func DbURL(config *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.DBName,
	)
}

type DBUser interface {
	CreateUser(u *models.User) (*models.User, error)
	GetUser(uuid string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	DeleteUser(uuid string) (*models.User, error)
	UpdateUser(uuid string, payload *models.Payload) (*models.User, error)
	GetAllUser() ([]*models.User, error)
}
