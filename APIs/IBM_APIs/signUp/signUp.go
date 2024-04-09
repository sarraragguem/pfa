package main

import (
	"signUp/controllers"
	"signUp/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadVariables()
	initializers.ConnectToDB()
	initializers.SyncDataBase()
}

func main() {
	router := gin.Default()
	router.GET("/getUsers", controllers.GetUsers)
	router.GET("/getUser/:email/:password", controllers.GetUserByUsernameAndPassword)
	router.GET("/login/:email/:password", controllers.Login)

	router.Run()
}
