package routes

import (
	"crud/controllers"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	fmt.Println("Routes are set up")
	r.GET("/movies", controllers.GetMovies)
	r.GET("/movies/:id", controllers.GetMovie)
	r.POST("/movies", controllers.CreateMovie)
	// r.GET("/tasks/:id", controllers.FindTask)
	// r.PATCH("/tasks/:id", controllers.UpdateTask)
	// r.DELETE("tasks/:id", controllers.DeleteTask)
	return r
}
