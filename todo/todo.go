package todo

import (
	"net/http"

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
func (Todo) Tablename() string {
	return "todos"
}

type TodoHandler struct {
	db *gorm.DB
}

func NewTodoHandler(db *gorm.DB) *TodoHandler {
	return &TodoHandler{db: db}
}

// t.NewTask()
func (t *TodoHandler) NewTask(c *gin.Context) {
	var todo Todo
	err := c.ShouldBindJSON(todo) // To bind a request body into a type

	if err != nil {
		// make JSON
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(), // response
		})

		return
	}
}
