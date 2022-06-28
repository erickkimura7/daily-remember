package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/erickkimura7/daily-remember/config"
	"github.com/gin-gonic/gin"
)

const (
	CONST_GO_PATH = "~/home/go"
)

func main() {
	config, err := config.NewConfig("./")

	if err != nil {
		panic(err)
	}

	r := gin.Default()

	r.GET("/", execScript)

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

func execScript(c *gin.Context) {
	home, _ := os.UserHomeDir()

	// if dir not exist create
	if _, err := os.Stat(getGoPath(home)); os.IsNotExist(err) {
		if err := os.MkdirAll(getGoPath(home), 0700); err != nil {
			responseError(c, err)

			return
		}
	}

	err := os.Chdir(getGoPath(home))

	if err != nil {
		responseError(c, err)
		return
	}

	if _, err := os.Stat(getProjectPath(home)); os.IsNotExist(err) {
		cmd := exec.Command("git", "clone", "https://github.com/erickkimura7/daily-remember.git")

		err := cmd.Run()

		if err != nil {
			responseError(c, err)
			return
		}
	}

	err = os.Chdir(getProjectPath(home))

	if err != nil {
		responseError(c, err)
		return
	}

	cmd := exec.Command("git", "pull")

	err = cmd.Run()

	if err != nil {
		responseError(c, err)
		return
	}

	cmd = exec.Command("make", "build-prod")

	err = cmd.Run()

	if err != nil {
		responseError(c, err)
		return
	}

	c.JSON(200, gin.H{
		"mensagem": "Pipeline executada com sucesso",
	})
}

func getProjectPath(home string) string {
	return filepath.Join(getGoPath(home), "daily-remember")
}

func getGoPath(home string) string {
	return filepath.Join(home, "go")
}

func responseError(c *gin.Context, err error) {
	log.Println(err)
	c.JSON(500, gin.H{
		"error": err,
	})
}
