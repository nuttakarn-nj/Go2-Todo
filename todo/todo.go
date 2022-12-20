package todo

import (
	"net/http"
	"strings"
	"todo/auth"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Todo struct {
	// Tag json name that title (front end send title)
	// Tag `json:"xxxxx"`
	Title string `json:"title"` // Export Title
	gorm.Model
}

// Todo = data
// Todo.Tablename()
func (Todo) TableName() string {
	return "todos"
}

// TodoHandler.db
type TodoHandler struct {
	db *gorm.DB
}

func NewTodoHandler(db *gorm.DB) *TodoHandler {
	return &TodoHandler{db: db}
}

// t.NewTask()
func (t *TodoHandler) NewTask(c *gin.Context) {
	// middleware
	// check token
	authorization := c.Request.Header.Get("Authorization")
	token := strings.TrimPrefix(authorization, "Bearer ")

	authErr := auth.Protect(token)

	if authErr != nil {
		// stop to go next middleware
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var todo Todo
	err := c.ShouldBindJSON(&todo) // To bind a request body into a type

	// Error
	if err != nil {
		// make JSON
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(), // response
		})

		return
	}

	// insert data
	result := t.db.Create(&todo)

	// Error
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	// Success
	c.JSON(http.StatusCreated, gin.H{
		"ID": todo.Model.ID, // ID from Model
	})
}
