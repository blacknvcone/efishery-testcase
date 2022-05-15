package main

import (
	"log"
	"os"
	"time"

	"github.com/blacknvcone/efishery-testcase/svc-fetch/common/config"
	"github.com/blacknvcone/efishery-testcase/svc-fetch/common/logger"
	fetchHttp "github.com/blacknvcone/efishery-testcase/svc-fetch/fetch/delivery/http"
	repoFetch "github.com/blacknvcone/efishery-testcase/svc-fetch/fetch/repository"
	ucFetch "github.com/blacknvcone/efishery-testcase/svc-fetch/fetch/usecase"
	iamHttp "github.com/blacknvcone/efishery-testcase/svc-fetch/iam/delivery/http"
	ucIam "github.com/blacknvcone/efishery-testcase/svc-fetch/iam/usecase"
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

	//Init IAM Module
	iamUc := ucIam.NewIAMUseCase(contextTimeout)

	//Init Fetch Module
	fetchRepo := repoFetch.NewFetchRepository()
	fetchUc := ucFetch.NewFetchUseCase(fetchRepo, contextTimeout)

	//Registering All Route Module
	fetchHttp.NewV1FetchHandler(router, iamUc, fetchUc, os.Getenv("JWT_SECRET"), logger)
	iamHttp.NewV1IamHandler(router, iamUc, os.Getenv("JWT_SECRET"), logger)

	router.Run(":" + os.Getenv("SERVER_PORT"))

}
