package httphandler

import (
	nt "github.com/erickkimura7/daily-remember/notificationEvent"
	"github.com/gin-gonic/gin"
)

type NotificationHandler interface {
	GetAllNotification(*gin.Context)
	AddNotification(*gin.Context)
	RemoveNotification(*gin.Context)
}

type handler struct {
	notificationService nt.NotificationService
}

type AddNotificationRequest struct {
}

func NewHandler(notificationService nt.NotificationService) NotificationHandler {
	return &handler{notificationService}
}

func (h *handler) GetAllNotification(c *gin.Context) {
	allNotifications, err := h.notificationService.ListAllNotifications()

	if err != nil {
		c.JSON(500, gin.H{"msg": err})
		return
	}

	c.JSON(200, allNotifications)
}

func (h *handler) AddNotification(c *gin.Context) {

	if err := h.notificationService.AddNotification(&nt.Event{}); err != nil {
		c.JSON(500, gin.H{"msg": err})
		return
	}

	c.JSON(201, gin.H{})
}

type RemoveNotificationRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

func (h *handler) RemoveNotification(c *gin.Context) {

	var request RemoveNotificationRequest

	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}

	if err := h.notificationService.RemoveNotification(request.ID); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}

	c.JSON(204, gin.H{})
}
