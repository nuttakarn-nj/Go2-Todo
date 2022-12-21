package main

import (
	"todo/auth"
	"todo/todo"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
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
	signature := []byte("==mySignature==")

	router.GET("/token", auth.AccessToken(signature))

	// middleware
	protected := router.Group("", auth.Protect(signature))

	// assign middleware to route
	todosHandler := todo.NewTodoHandler(db)
	protected.POST("/todos", todosHandler.NewTask)

	router.Run()
}
