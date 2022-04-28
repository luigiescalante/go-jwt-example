package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.jwt.example/httpd"
	"go.jwt.example/model"
	"log"
)

func main() {
	router := gin.Default()
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	model.InitDB()
	model.Migrate()
	router.GET("/", httpd.Index)
	router.GET("/_healthy", httpd.HealthyCheck)
	api := router.Group("/v1")
	{
		//users
		userApi := api.Group("/users")
		{
			userApi.GET("/", httpd.GetUser)
			userApi.GET("/id/:id", httpd.GetUserDetail)
			userApi.POST("/", httpd.NewUser)
			userApi.PUT("/id/:id", httpd.SaveUser)
			userApi.DELETE("/id/:id", httpd.DeleteUser)
		}
	}

	err = router.Run()
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
