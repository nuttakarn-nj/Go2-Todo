package main

import (
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

	router := gin.Default()
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	todosHandler := todo.NewTodoHandler(db)
	router.POST("/todos", todosHandler.NewTask)

	router.Run()
}
