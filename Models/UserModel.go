package Models

import (
	"encoding/json"
	"time"
)

type User struct {
	Id      	 uint   `json:"id"`
	Uuid    	 string `json:"uuid"`
	Name    	 string `json:"name"`
	Email   	 string `json:"email"`
	Password   	 string `json:"password"`
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
		Password  string `json:"password"`
		Phone		string `json:"phone"`
		Address		string `json:"address"`
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
		Id				uint		`json:"id"`
		Uuid			string		`json:"uuid"`
		Name			string		`json:"name"`
		Email			string		`json:"email"`
		Password		string		`json:"password"`
		Phone			string		`json:"phone"`
		Address			string		`json:"address"`
		CreationDate	time.Time 	`json:"creation_date"`
	}{
		Id:				u.Id,
		Uuid:           u.Uuid,
		Name:			u.Name,
		Email:			u.Email,
		Password:		u.Password,
		Phone:			u.Phone,
		Address:        u.Address,
		CreationDate:	u.CreationDate,
	}
	return json.Marshal(aux)
}
