package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/agclqq/study_tools/app/http/controller/response"
	"github.com/agclqq/study_tools/domain/demo"
)

type Demo struct {
}

func (d *Demo) Index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "http controller index demo"})
}
func (d *Demo) Show(ctx *gin.Context) {
	agg := demo.NewAgg()
	resp, err := agg.GetTest(ctx, 1)
	if err != nil {
		response.Failure(ctx, response.SERVER_ERROR, err.Error())
		return
	}
	response.Success(ctx, resp)
}
func (d *Demo) Update(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "http controller update demo"})
}
func (d *Demo) Store(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "http controller store demo"})
}
func (d *Demo) Destroy(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "http controller destroy demo"})
}
