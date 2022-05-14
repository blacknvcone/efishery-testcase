package main

import (
	"log"
	"os"
	"time"

	"github.com/blacknvcone/efishery-testcase/svc-fetch/common/config"
	"github.com/blacknvcone/efishery-testcase/svc-fetch/common/logger"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	//Loading Env
	err := godotenv.Load(config.ProjectRootPath + "/.env")
	if err != nil {
		log.Fatal("Error loading .env file : ", err.Error())

	}

	//Initiate Logger
	logger, _ := logger.NewLogger()

	router := gin.New()
	contextTimeout := time.Duration(10 * time.Second)

	router.Run(":" + os.Getenv("SERVER_PORT"))

}
