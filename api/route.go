package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/swaggo/gin-swagger/example/basic/docs"
)

func SetRouting(route *gin.Engine) {
	v1 := route.Group("/v1")
	{
		match := v1.Group("/Match")
		match.POST("/NewPerson", AddSinglePersonAndMatch)
	}
}

func SetSwagger(route *gin.Engine) {
	if mode := gin.Mode(); mode == gin.DebugMode {
		url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
		route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	}
}
