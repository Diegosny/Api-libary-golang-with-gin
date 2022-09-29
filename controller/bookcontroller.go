package controller

import (
	"api/database"
	"api/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

func ShowBook(c *gin.Context) {
	id := c.Param("id")

	newId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has been integer",
		})

		return
	}

	db := database.GetDatabase()

	var book model.Book
	err = db.First(&book, newId).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Book not found " + err.Error(),
		})

		return
	}

	c.JSON(200, book)
}

func CreateBook(c *gin.Context) {
	db := database.GetDatabase()

	var book model.Book

	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot bind json: " + err.Error(),
		})
		return
	}

	err = db.Create(&book).Error
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Cannot create book: " + err.Error(),
		})
		return
	}

	c.JSON(200, book)
}

func ShowBooks(c *gin.Context) {
	db := database.GetDatabase()

	var books []model.Book

	err := db.Find(&books).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot list books: " + err.Error(),
		})
		return
	}

	c.JSON(200, books)
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	db := database.GetDatabase()

	var book model.Book

	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot bind json: " + err.Error(),
		})
		return
	}

	err = db.Where("id = ?", id).Updates(&book).Error
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Cannot update book: " + err.Error(),
		})
		return
	}

	c.JSON(200, book)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	newId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has been integer",
		})
	}

	db := database.GetDatabase()

	err = db.Delete(&model.Book{}, newId).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot delete book: " + err.Error(),
		})
		return
	}

	c.Status(204)
}
