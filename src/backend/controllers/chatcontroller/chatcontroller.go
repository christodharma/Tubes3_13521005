package chatcontroller

import (
	"encoding/json"
	"net/http"

	"backend/model"
	"gorm.io/gorm"
	
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	
	var chats []model.Chat

	model.DB.Find(&chats)

	c.JSON(http.StatusOK, gin.H{"chats" : chats})
}
func Show(c *gin.Context) {
	var chat model.Chat
	id := c.Param("id")

	if err := model.DB.First(&chat, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default: 
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message" : err.Error()})
		}
	}
	c.JSON(http.StatusOK, gin.H{"Chat": chat})
}
func Create(c *gin.Context) {
	var chat model.Chat
	if err := c.ShouldBindJSON(&chat); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message" : err.Error()})
		return
	}
	model.DB.Create(&chat)
	c.JSON(http.StatusOK, gin.H{"chat": chat})
}
func Update(c *gin.Context) {
	var chat model.Chat
	id := c.Param("id")

	if err := c.ShouldBindJSON(&chat); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if model.DB.Model(&chat).Where("id = ?", id).Updates(&chat).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "tidak dapat mengupdate chat"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diperbarui"})
}
func Delete(c *gin.Context) {
	var chat model.Chat

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.Id.Int64()
	if model.DB.Delete(&chat, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus chat"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}