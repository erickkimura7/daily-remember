package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	httphandler "github.com/erickkimura7/daily-remember/api/http"
	"github.com/erickkimura7/daily-remember/config"
	"github.com/erickkimura7/daily-remember/jobscheduler"
	notificationevent "github.com/erickkimura7/daily-remember/notificationEvent"
	mockrepository "github.com/erickkimura7/daily-remember/repository/mock"
	"github.com/gin-gonic/gin"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	config, err := config.NewConfig("./")

	if err != nil {
		log.Fatal(err)
	}

	repo, err := mockrepository.NewMockRepository()

	if err != nil {
		log.Fatal(err)
	}

	notificacaoService := notificationevent.NewNotificationService(repo)

	handler := httphandler.NewHandler(notificacaoService)

	r := gin.Default()

	r.GET("/", handler.GetAllNotification)
	r.POST("/", handler.AddNotification)
	r.DELETE("/:id", handler.RemoveNotification)

	go jobscheduler.PoolJob()

	errs := make(chan error, 2)

	go func() {
		log.Printf("Listening on port :%s\n", config.Server.Port)
		err := r.Run(fmt.Sprintf(":%s", config.Server.Port))
		errs <- err
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	fmt.Printf("Terminated %s", <-errs)
}
