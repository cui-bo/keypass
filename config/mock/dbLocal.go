package mock

import (
	"github.com/cui-bo/keypass/models"
	uuid "github.com/satori/go.uuid"
	"time"
)

type mockDB struct {
	users map[uint]*models.User
}

func New() *mockDB {
	var db mockDB
	db.users = make(map[uint]*models.User)
	return &db
}

func (db *mockDB) CreateUser(u *models.User) (*models.User, error) {
	u.Uuid = uuid.NewV4().String()
	u.CreationDate = time.Now()
	db.users[u.Id] = u
	return u, nil
}

func (db *mockDB) GetUser(id uint) (*models.User, error) {

	return db.users[id], nil
}

func (db *mockDB) DeleteUser(id uint) (*models.User, error) {
	u, err := db.GetUser(id)
	if err != nil {
		return nil, err
	}
	delete(db.users, id)
	return u, nil
}

func (db *mockDB) UpdateUser(id uint, payload *models.Payload) (*models.User, error) {
	u, err := db.GetUser(id)
	if err != nil {
		return nil, err
	}

	payload.ToString(&u.Name, "name")
	payload.ToString(&u.Email, "email")
	payload.ToString(&u.Phone, "phone")
	payload.ToString(&u.Address, "address")

	return u, nil
}

func (db *mockDB) GetAllUser() ([]*models.User, error) {
	us := make([]*models.User, len(db.users))
	for _, u := range db.users {
		us = append(us, u)
	}
	return us, nil
}
