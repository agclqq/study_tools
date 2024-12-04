package router

import (
	"github.com/agclqq/prow-framework/http/restful/router"
	"github.com/gin-gonic/gin"

	"github.com/agclqq/study_tools/app/http/controller"
)

func Register(eng *gin.Engine) {
	apiGroup := eng.Group("/api")
	{
		router.ApiResource(apiGroup, "/demo", &controller.Demo{})
	}
}
