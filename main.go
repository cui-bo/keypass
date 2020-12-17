package main

import (
	"fmt"
	"github.com/cui-bo/keypass/config"
	"github.com/cui-bo/keypass/routes"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "github.com/cui-bo/keypass/docs"
)

type Conf struct {
	MySQL string
	Mode  string
	Port  string
}

var conf Conf
var err error

// Function executed before main
func init() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		panic(fmt.Errorf("Fatal err config file: %s \n", err))
	}
	conf.MySQL = viper.GetString("mysql.dsn")
	conf.Mode = viper.GetString("mode")
	conf.Port = viper.GetString("port")
}

// @title Swagger For Keypass API
// @version 1.0
// @description This is an API for creating hash in order to create keypass.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /v1
func main() {
	r := gin.Default()

	// config
	fmt.Println("run app in mode:", conf.Mode)

	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}

	defer config.DB.Close()

	r = routes.SetupUserRouter(r)

	// Swagger
	url := ginSwagger.URL("http://localhost:9090/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.Run(":9090")
}
