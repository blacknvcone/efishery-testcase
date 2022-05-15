package http

import (
	"net/http"

	"github.com/blacknvcone/efishery-testcase/svc-fetch/common/logger"
	"github.com/blacknvcone/efishery-testcase/svc-fetch/domain"
	"github.com/gin-gonic/gin"
)

type v1FetchHandler struct {
	IAMUseCase   domain.IIamUseCase
	FetchUseCase domain.IFetchUseCase
	log          logger.LogInfoFormat
	jwtSecret    string
}

func NewV1FetchHandler(router *gin.Engine, iamUseCase domain.IIamUseCase, fetchUseCase domain.IFetchUseCase, logger logger.LogInfoFormat) {

	h := &v1FetchHandler{
		IAMUseCase:   iamUseCase,
		FetchUseCase: fetchUseCase,
		log:          logger,
	}

	//Implement JWT Middleware Here
	v1 := router.Group("/v1")
	v1.Use(h.IAMUseCase.AuthorizationHTTP())
	{
		v1.GET("/ping", h.Ping)
		v1.GET("/fetch", h.Fetch)
		v1.GET("/aggregate", h.Aggregate)
	}

}

func (v *v1FetchHandler) Ping(g *gin.Context) {
	g.JSON(http.StatusOK, "OK")
}

func (v *v1FetchHandler) Fetch(g *gin.Context) {

	ctx := g.Request.Context()
	res, err := v.FetchUseCase.FetchAndCustom(ctx)

	if err != nil {
		v.log.Info(err.Error())
		g.JSON(http.StatusBadGateway, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "OK",
		"data":    res,
	})

}

func (v *v1FetchHandler) Aggregate(g *gin.Context) {
	ctx := g.Request.Context()
	res, err := v.FetchUseCase.SumAggregate(ctx)

	if err != nil {
		v.log.Info(err.Error())
		g.JSON(http.StatusBadGateway, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "OK",
		"data":    res,
	})

}
