package controllers

import (
	"assignment-3/database"
	"assignment-3/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RecordStatus(ctx *gin.Context) {
	db := database.GetDB()
	var el models.Element

	if err := ctx.ShouldBindJSON(&el); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	err := db.Debug().Create(&el).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data": el,
		"code": 200,
	})
}
