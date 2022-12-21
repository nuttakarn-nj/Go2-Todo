package main

import (
	"log"
	"os"
	"todo/auth"
	"todo/todo"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// load env
	errEnv := godotenv.Load("local.env")
	if errEnv != nil {
		log.Printf("Please consider env %s", errEnv)
	}
	signature := []byte(os.Getenv("SIGNATURE"))

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("Database connection failed")
	}

	db.AutoMigrate(&todo.Todo{})

	// init router
	router := gin.Default()
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/token", auth.AccessToken(signature))

	// middleware
	protected := router.Group("", auth.Protect(signature))

	// assign middleware to route
	todosHandler := todo.NewTodoHandler(db)
	protected.POST("/todos", todosHandler.NewTask)

	router.Run()
}
