package services

import (
	"errors"
	"fmt"
	"github.com/cui-bo/keypass/config"
	"github.com/cui-bo/keypass/models"
	"strings"
)

func GetAllUsers(users *[]models.User) (err error) {
	if err = config.DB.Find(users).Error; err != nil {
		return err
	}
	return nil
}

func CreateUser(user *models.User) (err error) {
	if err = config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func GetUserById(user *models.User, id string) (err error) {
	if err = config.DB.Where("id=?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUser(user *models.User, id string) (err error) {
	if err = config.DB.Save(user).Error; err != nil {
		return err
	}
	return nil
}

func DeleteUser(user *models.User, id string) (err error) {
	if err = config.DB.Where("id=?", id).Delete(user).Error; err != nil {
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

func ValidatePayload(user *models.User) []error {
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

//
//// @Description check if the login pass are correct and gives backe a JWT value
//// @Accept json
//// @Produce json
//// @Param user body model.UserLogin true "Add a User"
//// @Success 200 {object} string nil
//// @Failure 400 {string} string nil
//// @Router /login [post]
//func (su *ServiceUser) LoginUser(ctx *gin.Context) {
//	var payload models.UserLogin
//	err := ctx.BindJSON(&payload)
//	if err != nil {
//		log.Println("/users bad request", err.Error())
//		ctx.JSON(http.StatusBadRequest, nil)
//		return
//	}
//
//	u2, err := su.DB.GetUserByEmail(payload.Login)
//	if err != nil {
//		log.Println("/users not found", err.Error())
//		ctx.JSON(http.StatusNotFound, nil)
//		return
//	}
//
//	if u2.Password != payload.Password {
//		log.Println("/users not authorized")
//		ctx.JSON(http.StatusUnauthorized, nil)
//		return
//	}
//
//	jwtValue, err := middleware.NewJWT(u2.Uuid, u2.Name)
//	if err != nil {
//		log.Println("/users not internal server error", err)
//		ctx.JSON(http.StatusInternalServerError, nil)
//		return
//	}
//	ctx.JSON(http.StatusOK, gin.H{"jwt": jwtValue})
//}
