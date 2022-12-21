package main

import (
	"log"
	"net/http"
	"os"
	"time"
	"todo/auth"
	"todo/todo"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/time/rate"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// set vairable for package main
var (
	buildCommit = "dev"
	buildTime   = time.Now().String()
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

	router.GET("/limit", limitedHandler)
	router.GET("/token", auth.AccessToken(signature))

	// #ldflag get value from git
	// if build with ldflag >> value come from git
	router.GET("/commit", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"buildCommit": buildCommit,
			"buildTime":   buildTime,
		})
	})

	// middleware
	protected := router.Group("", auth.Protect(signature))

	// assign middleware to route
	todosHandler := todo.NewTodoHandler(db)
	protected.POST("/todos", todosHandler.NewTask)

	router.Run()
}

// Test limit req
var limiter = rate.NewLimiter(5, 5)
func limitedHandler(ctx *gin.Context) {
	if !limiter.Allow() {
		ctx.AbortWithStatus(http.StatusTooManyRequests)
		return
	}

	ctx.JSON(200, gin.H{
		"message": "work fine & not over limit",
	})

}
