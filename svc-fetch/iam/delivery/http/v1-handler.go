package http

import (
	"net/http"

	"github.com/blacknvcone/efishery-testcase/svc-fetch/common/logger"
	"github.com/blacknvcone/efishery-testcase/svc-fetch/domain"
	"github.com/gin-gonic/gin"
)

type v1IamHandler struct {
	IAMUseCase domain.IIamUseCase
	log        logger.LogInfoFormat
}

func NewV1IamHandler(router *gin.Engine, iamUseCase domain.IIamUseCase, logger logger.LogInfoFormat) {

	h := &v1IamHandler{
		IAMUseCase: iamUseCase,
		log:        logger,
	}

	//Implement JWT Middleware Here
	v1 := router.Group("/v1")
	v1.Use(h.IAMUseCase.AuthorizationHTTP())
	{
		v1.GET("/profile", h.GetProfile)
	}

}

func (v *v1IamHandler) GetProfile(g *gin.Context) {

	res, exist := g.Get("claim")
	if !exist {
		g.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Session Not Found",
			"data":    nil,
		})
	} else {
		g.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "OK",
			"data":    res,
		})
	}

}
