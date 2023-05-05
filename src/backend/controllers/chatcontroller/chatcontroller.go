package chatcontroller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"backend/model"

	"gorm.io/gorm"

	"backend/StringMatching"

	"github.com/gin-gonic/gin"
)

var AlgoType bool = true // true = KMP, false = BM

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

func TextInput(c *gin.Context) {
	type Message struct{
		content string
	}
	var input Message
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message" : err.Error()})
		return
	}
	var question []string
	dbRaw := model.DB.Raw("SELECT question FROM CHATS").Scan(&question)
	var answer []string
	dbAnswerRaw := model.DB.Raw("SELECT question FROM CHATS").Scan(&answer)
	if dbRaw.Error != nil || dbAnswerRaw.Error != nil{
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message" : dbRaw.Error.Error()})
		return
	}
	for _, str := range question {
		fmt.Println(str)
		if AlgoType {
			kmpRet := StringMatching.KMPMatch(str, input.content)
			fmt.Println(kmpRet)
			if len(kmpRet) > 0 {
				c.JSON(http.StatusOK, gin.H{"text": input.content, "reply" : str, "kmpRet" : kmpRet[0]})
				return
			}
		} else {
			bmRet := StringMatching.BMMatch(str, input.content)
			if bmRet != -1 {
				c.JSON(http.StatusOK, gin.H{"text": input.content, "reply" : str})
				return
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"text": input.content, "reply" : "Did not find an answer"})
				return
			}
		}
	}
	var chats []model.Chat
    abandonHope := model.DB.Where("question LIKE ?", "%"+input.content+"%").Find(&chats)
	if abandonHope.Error != nil {
		c.JSON(http.StatusOK, gin.H{"text": input.content, "reply" : "Did not find an answer"})
	} else {
		c.JSON(http.StatusOK, gin.H{"text": input.content, "reply" : abandonHope})
	}
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