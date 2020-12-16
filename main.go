package main

import (
	"fmt"
	"github.com/cui-bo/keypass/Config"
	"github.com/cui-bo/keypass/Models"
	"github.com/cui-bo/keypass/Routes"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/jinzhu/gorm"

	_ "github.com/cui-bo/keypass/docs"
)

var err error

func main() {
	r := gin.Default()

	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}

	defer Config.DB.Close()

	Config.DB.AutoMigrate(&Models.User{})

	// Route
	r = Routes.SetupRouter()

	// Swagger
	url := ginSwagger.URL("http://localhost:9090/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.Run(":9090")
}
