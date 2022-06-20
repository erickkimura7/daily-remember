package httphandler

import (
	"log"

	nt "github.com/erickkimura7/daily-remember/notificationEvent"
	"github.com/gin-gonic/gin"
)

type NotificationHandler interface {
	GetAllNotification(*gin.Context)
	AddNotification(*gin.Context)
}

type handler struct {
	notificationService nt.NotificationService
}

func NewHandler(notificationService nt.NotificationService) NotificationHandler {
	return &handler{notificationService}
}

func (h *handler) GetAllNotification(c *gin.Context) {
	allNotifications, err := h.notificationService.ListAllNotifications()

	if err != nil {
		log.Fatalln(err)
	}

	c.JSON(200, allNotifications)
}

func (h *handler) AddNotification(c *gin.Context) {

	if err := h.notificationService.AddNotification(&nt.Event{}); err != nil {
		log.Fatalln(err)
	}

	c.JSON(201, gin.H{})
}
