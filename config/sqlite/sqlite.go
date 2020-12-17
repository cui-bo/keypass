package sqlite

import (
	database "github.com/cui-bo/keypass/config"
	customErr "github.com/cui-bo/keypass/err"
	"github.com/cui-bo/keypass/models"
	uuid "github.com/satori/go.uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

var _ database.BDD = &SQLite{}

type SQLite struct {
	db *gorm.DB
}

func New(dbName string) *SQLite {
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.User{})
	//db.AutoMigrate(&models.Card{})

	return &SQLite{
		db: db,
	}
}

func (db *SQLite) SetDB(dbgorm *gorm.DB) {
	db.db = dbgorm
}

func (db *SQLite) CreateUser(u *models.User) (*models.User, error) {
	u.Uuid = uuid.NewV4().String()
	u.CreationDate = time.Now()
	db.db.Create(&u)
	return u, nil
}

func (db *SQLite) GetUser(id uint) (*models.User, error) {
	var u models.User
	db.db.Where("id = ?", id).First(&u)
	return &u, nil
}

func (db *SQLite) GetUserByEmail(email string) (*models.User, error) {
	var u models.User
	db.db.Where("email = ?", email).First(&u)
	if u.Id == 0 {
		return nil, customErr.NewErrNotFound("email"+email, nil)
	}
	return &u, nil
}

func (db *SQLite) DeleteUser(id uint) (*models.User, error) {
	var u models.User
	db.db.Where("id = ?", id).Delete(&u)
	return &u, nil
}

func (db *SQLite) UpdateUser(id uint, payload *models.Payload) (*models.User, error) {

	db.db.Model(&models.User{}).Where("id = ?", id).Updates(payload.Data)
	return db.GetUser(id)
}

func (db *SQLite) GetAllUser() ([]*models.User, error) {
	var us []*models.User
	db.db.Find(&us)
	return us, nil
}

