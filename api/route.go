package api

import (
	_ "github.com/CoulsonChen/BitoAPI/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetRouting(route *gin.Engine) {
	v1 := route.Group("/v1")
	{
		match := v1.Group("/Match")
		match.POST("/NewPerson", AddSinglePersonAndMatch)
		match.DELETE("/Remove/:id", RemovePerson)
		match.GET("/QueryMatches/:id", QueryMatches)
		match.PUT("/:id/:idm", Match)
	}
}

func SetSwagger(route *gin.Engine) {
	if mode := gin.Mode(); mode == gin.DebugMode {
		url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
		route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	}
}
