package Models

import (
	"encoding/json"
	"github.com/cui-bo/keypass/Config"
	"time"
)

type User struct {
	Id      	 uint   `json:"id"`
	Uuid    	 string `json:"uuid"`
	Name    	 string `json:"name"`
	Email   	 string `json:"email"`
	Phone   	 string `json:"phone"`
	Address 	 string `json:"address"`
	CreationDate time.Time `json:"creation_date"`
}

func (u *User)TableName() string  {
	return "user"
}

func (u *User) UnmarshalJSON(b []byte) error {
	aux := struct {
		Name 		string `json:"name"`
		Email		string `json:"email"`
		Phone		string `json:"phone"`
		Address		string `json:"address"`
	}{}
	if err := json.Unmarshal(b, &aux); err != nil {
		return err
	}

	u.Name = aux.Name
	u.Email = aux.Email
	u.Phone = aux.Phone
	u.Address = aux.Address

	return nil
}

func (u User) MarshalJSON() ([]byte, error) {
	aux := struct {
		Id				uint		`json:"id"`
		Uuid			string		`json:"uuid"`
		Name			string		`json:"name"`
		Email			string		`json:"email"`
		Phone			string		`json:"phone"`
		Address			string		`json:"address"`
		CreationDate	time.Time `json:"creation_date"`
	}{
		Id:				u.Id,
		Uuid:           u.Uuid,
		Name:			u.Name,
		Email:			u.Email,
		Phone:			u.Phone,
		Address:        u.Address,
		CreationDate:	u.CreationDate,
	}
	return json.Marshal(aux)
}

func GetAllUsers(users *[]User) (err error) {
	if err = Config.DB.Find(users).Error; err != nil {
		return err
	}
	return nil
}

func CreateUser(user *User) (err error) {
	if err = Config.DB.Create(user).Error; err != nil {
		return err;
	}
	return nil
}

func GetUserById(user *User, id string) (err error) {
	if err = Config.DB.Where("id=?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUser(user *User, id string) (err error) {
	if err = Config.DB.Save(user).Error; err != nil {
		return err
	}
	return nil
}

func DeleteUser(user *User, id string) (err error) {
	if err = Config.DB.Where("id=?", id).Delete(user).Error; err != nil {
		return err
	}
	return nil
}