package http

import (
	"net/http"

	"github.com/blacknvcone/efishery-testcase/svc-fetch/common/logger"
	"github.com/blacknvcone/efishery-testcase/svc-fetch/domain"
	"github.com/gin-gonic/gin"
)

type v1FetchHandler struct {
	IAMUseCase domain.IIamUseCase
	log        logger.LogInfoFormat
	jwtSecret  string
}

func NewV1FetchHandler(router *gin.Engine, iamUseCase domain.IIamUseCase, jwtSecret string, logger logger.LogInfoFormat) {

	h := &v1FetchHandler{
		IAMUseCase: iamUseCase,
		log:        logger,
		jwtSecret:  jwtSecret,
	}

	//Implement JWT Middleware Here
	v1 := router.Group("/v1")
	v1.Use(h.IAMUseCase.AuthorizationHTTP(jwtSecret))
	{
		v1.GET("/ping", h.Ping)
	}

}

func (v *v1FetchHandler) Ping(g *gin.Context) {
	g.JSON(http.StatusOK, "OK")
}
