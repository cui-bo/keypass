package Services

import (
	"errors"
	"fmt"
	"github.com/cui-bo/keypass/Config"
	"github.com/cui-bo/keypass/Models"
	"strings"
)

func GetAllUsers(users *[]Models.User) (err error) {
	if err = Config.DB.Find(users).Error; err != nil {
		return err
	}
	return nil
}

func CreateUser(user *Models.User) (err error) {
	if err = Config.DB.Create(user).Error; err != nil {
		return err;
	}
	return nil
}

func GetUserById(user *Models.User, id string) (err error) {
	if err = Config.DB.Where("id=?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUser(user *Models.User, id string) (err error) {
	if err = Config.DB.Save(user).Error; err != nil {
		return err
	}
	return nil
}

func DeleteUser(user *Models.User, id string) (err error) {
	if err = Config.DB.Where("id=?", id).Delete(user).Error; err != nil {
		return err
	}
	return nil
}

func validateSize(min, max int, value string) error {
	if len(value) < min || len(value) > max {
		if len(value) > max {
			return fmt.Errorf("got a value too long: %v", value[:max])
		}
		return fmt.Errorf("got a value too short: %v" + value)
	}
	return nil
}


func ValidatePayload(user *Models.User) []error {
	var errList []error

	err := validateSize(3, 200, user.Name)
	if err != nil {
		errList = append(errList, err)
	}

	err = validateSize(3, 320, user.Email)
	if err != nil {
		errList = append(errList, err)
	}
	if !strings.Contains(user.Email, "@") {
		errList = append(errList, errors.New("no @ found for the email"))
	}

	emailVals := strings.Split(user.Email, "@")
	if len(emailVals) != 2 {
		errList = append(errList, errors.New("email not valid"))
	}
	err = validateSize(1, 64, emailVals[1])
	if err != nil {
		errList = append(errList, err)
	}

	validateSize(3, 200, user.Password)
	if err != nil {
		errList = append(errList, err)
	}
	return errList
}
