package Controllers

import (
	"github.com/cui-bo/keypass/Models"
	"github.com/cui-bo/keypass/Services"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	uuid "github.com/satori/go.uuid"
	"log"
	"net/http"
	"time"
)

// @Description get all users
// @Accept json
// @Produce json
// @Param uuid path string true "Some ID"
// @Success 200 {object} Models.User "ok"
// @Router /users [get]
func GetUsers(ctx *gin.Context) {
	var users []Models.User

	err := Services.GetAllUsers(&users)
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
	} else {
		ctx.JSON(http.StatusOK, users)
	}
}

// @Description create a User from the payload.
// @Accept json
// @Produce json
// @Param user body Models.User true "Add a User"
// @Success 200 {object} Models.User
// @Failure 400 {string} string nil
// @Router /user [post]
func CreateUser(ctx *gin.Context) {
	var user Models.User
	// uuid
	u := uuid.NewV4().String()
	user.Uuid = u
	// date
	d := time.Now()
	user.CreationDate = d

	ctx.BindJSON(&user)

	errs := Services.ValidatePayload(&user)
	if len(errs) != 0 {
		log.Println("/users bad request", errs)
		ctx.JSON(http.StatusBadRequest, errs)
		return
	}

	err := Services.CreateUser(&user)
	if err != nil {
		log.Println("/users bad request", err.Error())
		ctx.AbortWithStatus(http.StatusNotFound)
	} else {
		ctx.JSON(http.StatusOK, user)
	}
}

// @Description get user by id
// @Accept json
// @Produce json
// @Param id path string "Some ID"
// @Success 200 {object} Models.User "ok"
// @Failure 400 {string} string "We need ID!!"
// @Failure 404 {string} string "Can not find ID"
// @Router /user/{id} [get]
func GetUserById(ctx *gin.Context) {
	userId := ctx.Params.ByName("id")
	var user Models.User
	err := Services.GetUserById(&user, userId)
	if err != nil {
		log.Println("/users bad request", err.Error())
		ctx.AbortWithStatus(http.StatusNotFound)
	} else {
		ctx.JSON(http.StatusOK, user)
	}
}

func UpdateUser(ctx *gin.Context) {
	userId := ctx.Params.ByName("id")
	var user Models.User

	// Check if user exist
	err := Services.GetUserById(&user, userId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, user)
	}

	ctx.BindJSON(&user)

	err = Services.UpdateUser(&user, userId)
	if err != nil {
		log.Println("/users bad request", err.Error())
		ctx.AbortWithStatus(http.StatusNotFound)
	} else {
		ctx.JSON(http.StatusOK, user)
	}
}

func DeleteUser(ctx *gin.Context) {
	userId := ctx.Params.ByName("id")
	var user Models.User

	// Check if user exist
	err := Services.GetUserById(&user, userId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, user)
	}

	ctx.BindJSON(&user)

	err = Services.DeleteUser(&user, userId)
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
	} else {
		ctx.JSON(http.StatusOK, gin.H{"id " + userId : " is deleted"})
	}

}