package controllers

import (
	"middleware/database"
	"middleware/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreatePhoto(c *gin.Context) {
	db := database.GetDB()
	UserData := c.MustGet("userData").(jwt.MapClaims)
	//contentType := helpers.GetContentType(c)
	userID := uint(UserData["id"].(float64))
	Photo := models.Photo{}
	Photo.UserID = userID

	//if contentType == appJSON {
	//	c.ShouldBindJSON(&Photo)
	//} else {
	//	c.ShouldBindJSON(&Photo)
	//}

	if err := c.ShouldBindJSON(&Photo); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := db.Create(&Photo).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"caption":    Photo.Caption,
		"created_at": Photo.CreatedAt,
		"id":         Photo.ID,
		"photo_url":  Photo.PhotoUrl,
		"title":      Photo.Title,
		"user_id":    Photo.UserID,
	})
}

func GetAllPhoto(c *gin.Context) {
	db := database.GetDB()
	Photos := []models.Photo{}
	err := db.Find(&Photos).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	db.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("Email", "Username")
	}).Find(&Photos)

	c.JSON(http.StatusOK, Photos)
}

func PhotoUpdate(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	Photo := models.Photo{}
	photoId, _ := strconv.Atoi(c.Param("photoId"))
	userID := uint(userData["id"].(float64))
	Photo.UserID = userID
	Photo.ID = uint(photoId)

	if err := c.ShouldBindJSON(&Photo); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := db.Model(&Photo).Where("id = ?", photoId).Updates(models.Photo{
		Title:    Photo.Title,
		Caption:  Photo.Caption,
		PhotoUrl: Photo.PhotoUrl,
	}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":        Photo.ID,
		"title":     Photo.Title,
		"caption":   Photo.Caption,
		"photo_url": Photo.PhotoUrl,
		"user_id":   Photo.UserID,
		"update_at": Photo.UpdatedAt,
	})
}

func PhotoDelete(c *gin.Context) {
	db := database.GetDB()
	Photo := models.Photo{}
	photoId, _ := strconv.Atoi(c.Param("photoId"))
	Photo.ID = uint(photoId)

	err := db.Where("id = ?", photoId).Delete(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your Photo Has been Successfully Deleted",
	})
}
