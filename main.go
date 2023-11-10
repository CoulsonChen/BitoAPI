package main

import (
	"github.com/CoulsonChen/BitoAPI/api"
	"github.com/CoulsonChen/BitoAPI/internal/service"
	"github.com/CoulsonChen/BitoAPI/third-party/redis"
	"github.com/gin-gonic/gin"
)

// @title Swagger API
// @version 1.0
// @description API documentation.
// @contact.name API Support
// @contact.url https://CoulsonChen.github.io/
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// schemes http
// @BasePath /v1
func main() {
	// storage initialize
	redis.Initial()
	// service initialize
	service.InitMatchService()

	// routing set
	r := gin.Default()
	api.SetRouting(r)
	api.SetSwagger(r)
	r.Run()

}
