package controllers

import (
	"crud/models"
	"net/http"

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
	id := c.Param("id")
	if err := db.Where("id = ?", id).First(&movie).Error; err != nil {
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
	movie := models.Movie{ID: input.ID, MovieName: input.MovieName}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&movie)
	c.JSON(http.StatusOK, gin.H{"data": movie})
}

// Update a movie
func UpdateMovie(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var movie models.Movie
	id := c.Param("id")
	if err := db.Where("id = ?", id).First(&movie).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	// Validate input
	var input models.Movie
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Model(&movie).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": movie})
}

// Delete a movie
func DeleteMovie(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var movie models.Movie
	id := c.Param("id")
	if err := db.Where("id = ?", id).First(&movie).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	db.Delete(&movie)
	c.JSON(http.StatusOK, gin.H{"data": movie})
}
