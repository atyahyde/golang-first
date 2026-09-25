package controllers

import (
	"net/http"

	"example.com/assesment-app/config"
	"example.com/assesment-app/models"
	"github.com/gin-gonic/gin"
)

func CreateEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	event.UserId = 1 //dummy

	if err := config.DB.Create(&event).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "Event created successfully",
		"data":    event,
	})
}

func GetEvents(context *gin.Context) {
	var events []models.Event

	if err := config.DB.Find(&events).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Events retrieved successfully",
		"data":    events,
	})
}
