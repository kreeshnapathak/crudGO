package controllers

import (
	"crud/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//getting all the movies list
func GetMovies(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// movie := models.Movie{}
	var movies []models.Movie
	db.Find(&movies)

	c.JSON(http.StatusOK, gin.H{"data": movies})
}

// Find a movie by id
func GetMovie(c *gin.Context) { // Get model if exist
	var movie models.Movie
	db := c.MustGet("db").(*gorm.DB)
	idstr := c.Param("id")
	id, _ := strconv.Atoi(idstr)
	fmt.Println(db, "-----------------------", c.Param("idstr"))

	// if err := db.Where("id = ?", c.Param("id")).First(&movie).Error; err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	// 	return
	// }
	fmt.Println(db.First(&movie, id), "-----------------------", id)
	if err := db.First(&movie, id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": movie})
}

//creating a new movie
func CreateMovie(c *gin.Context) {
	// Validate input
	var input models.Movie
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Create task
	movie := models.Movie{MovieID: input.MovieID, MovieName: input.MovieName}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&movie)
	c.JSON(http.StatusOK, gin.H{"data": movie})
}
