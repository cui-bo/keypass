package models

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"
)

type User struct {
	Id           uint      `json:"id"`
	Uuid         string    `json:"uuid"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	Phone        string    `json:"phone"`
	Address      string    `json:"address"`
	CreationDate time.Time `json:"creation_date"`
}

type UserLogin struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (u *User) TableName() string {
	return "user"
}

func (u *User) UnmarshalJSON(b []byte) error {
	aux := struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Phone    string `json:"phone"`
		Address  string `json:"address"`
	}{}
	if err := json.Unmarshal(b, &aux); err != nil {
		return err
	}

	u.Name = aux.Name
	u.Email = aux.Email
	u.Password = aux.Password
	u.Phone = aux.Phone
	u.Address = aux.Address

	return nil
}

func (u User) MarshalJSON() ([]byte, error) {
	aux := struct {
		Id           uint      `json:"id"`
		Uuid         string    `json:"uuid"`
		Name         string    `json:"name"`
		Email        string    `json:"email"`
		Password     string    `json:"password"`
		Phone        string    `json:"phone"`
		Address      string    `json:"address"`
		CreationDate time.Time `json:"creation_date"`
	}{
		Id:           u.Id,
		Uuid:         u.Uuid,
		Name:         u.Name,
		Email:        u.Email,
		Password:     u.Password,
		Phone:        u.Phone,
		Address:      u.Address,
		CreationDate: u.CreationDate,
	}
	return json.Marshal(aux)
}

func Hash(clear string) (hashed string) {
	h := sha256.New()
	h.Write([]byte(clear))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func (u *UserLogin) UnmarshalJSON(b []byte) error {
	aux := struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}{}

	if err := json.Unmarshal(b, &aux); err != nil {
		return err
	}

	u.Login = aux.Login
	u.Password = Hash(aux.Password)

	return nil
}

func (u UserLogin) MarshalJSON() ([]byte, error) {
	return json.Marshal(nil)
}
