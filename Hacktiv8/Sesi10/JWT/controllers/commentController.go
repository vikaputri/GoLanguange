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

func CreateComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	Comment := models.Comment{}
	Photo := models.Photo{}
	userID := uint(userData["id"].(float64))
	Comment.UserID = userID

	if err := c.ShouldBindJSON(&Comment); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := db.Model(&Photo).Where("id = ?", Comment.PhotoID).First(&Photo).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": "Photo ID not found",
		})
		return
	}

	err1 := db.Debug().Create(&Comment).Error

	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err1.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         Comment.ID,
		"message":    Comment.Message,
		"photo_id":   Comment.PhotoID,
		"user_id":    Comment.UserID,
		"created_at": Comment.CreatedAt,
	})
}

func GetAllComment(c *gin.Context) {
	db := database.GetDB()
	Comments := []models.Comment{}
	err := db.Find(&Comments).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	db.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("ID", "Email", "Username")
	}).Preload("Photo").Find(&Comments)

	c.JSON(http.StatusOK, Comments)
}

func CommentUpdate(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	Comment := models.Comment{}
	commentId, _ := strconv.Atoi(c.Param("commentId"))
	userID := uint(userData["id"].(float64))
	Comment.UserID = userID
	Comment.ID = uint(commentId)

	if err := c.ShouldBindJSON(&Comment); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := db.Model(&Comment).Where("id = ?", commentId).Updates(models.Comment{
		Message: Comment.Message,
	}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	db.Model(&Comment).Where("id = ?", Comment.ID).First(&Comment)

	c.JSON(http.StatusOK, gin.H{
		"id":         Comment.ID,
		"message":    Comment.Message,
		"photo_id":   Comment.PhotoID,
		"user_id":    Comment.UserID,
		"updated_at": Comment.UpdatedAt,
	})
}

func CommentDelete(c *gin.Context) {
	db := database.GetDB()
	Comment := models.Comment{}
	commentId, _ := strconv.Atoi(c.Param("commentId"))
	Comment.ID = uint(commentId)
	err := db.Where("id = ?", Comment.ID).Delete(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Your comment has been successfully deleted",
	})

}
