package handler

import "github.com/gin-gonic/gin"

type XIXI map[string]interface{}

func Home(c *gin.Context) {
	c.HTML(200, "index.html", XIXI{
		"Title":   "Welcome to My Website",
		"User":    "Guest",
		"Items":   []string{"Item 1", "Item 2", "Item 3"},
		"Year":    2025,
		"Message": "Hello, World!",
	})
}
