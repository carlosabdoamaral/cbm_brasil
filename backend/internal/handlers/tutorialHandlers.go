package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func CreateTutorial(c *gin.Context) {
	// code...
}

func GetAllTutorials(c *gin.Context) {
	// code...
}

// Must pass query: /path?id=1234
func GetTutorialById(c *gin.Context) {
	// shortcut for c.Request.URL.Query().Get("lastname")
	id := c.Query("id")
	fmt.Println(id)
}

// Must pass query: /path?id=1234
func UpdateTutorialById(c *gin.Context) {
	// code...
}

func DeleteTutorialById(c *gin.Context) {
	// code...
}

func SoftDeleteTutorialById(c *gin.Context) {
	// code...
}
