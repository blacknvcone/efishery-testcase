package main

import (
	"log"
	"os"
	"time"

	"github.com/blacknvcone/efishery-testcase/svc-fetch/common/config"
	"github.com/blacknvcone/efishery-testcase/svc-fetch/common/logger"
	"github.com/blacknvcone/efishery-testcase/svc-fetch/fetch/delivery/http"
	"github.com/blacknvcone/efishery-testcase/svc-fetch/iam/usecase"
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

	//Init IAM
	iamUc := usecase.NewIAMUseCase(contextTimeout)

	//Registering All Route Module
	http.NewV1FetchHandler(router, iamUc, os.Getenv("JWT_SECRET"), logger)

	// iamRepo := _iamRepo.NewMysqlIAMRepository(db)
	// iamUseCase := _iamUseCase.NewIAMUseCase(iamRepo, contextTimeout)
	// _iamDelivery.NewIAMHandler(router, iamUseCase, logger)

	// prodRepo := _productRepo.NewMysqlProductRepository(db)
	// prodUC := _productUseCase.NewProductUseCase(prodRepo, contextTimeout)
	// _productDelivery.NewProductHandler(router, prodUC, iamUseCase, logger)

	router.Run(":" + os.Getenv("SERVER_PORT"))

}
