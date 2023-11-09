package api

import (
	"net/http"

	"github.com/CoulsonChen/BitoAPI/internal/model"
	"github.com/CoulsonChen/BitoAPI/internal/service"
	"github.com/gin-gonic/gin"
)

// @Summary Add a new pet to the store
// @Description get string by ID
// @Accept  json
// @Produce  json
// @Param   some_id     path    int     true        "Some ID"
// @Success 200 {string} string	"ok"
// @Failure 400 {object} web.APIError "We need ID!!"
// @Failure 404 {object} web.APIError "Can not find ID"
// @Router /Match/NewPerson [post]
func AddSinglePersonAndMatch(c *gin.Context) {

	person := &model.Person{}
	err := c.BindJSON(person)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	match := service.GetMatchServiceInstance()

	match.AddPerson(*person)
}

// @Description get struct array by ID
// @Accept  json
// @Produce  json
// @Param   some_id     path    string     true        "Some ID"
// @Param   offset     query    int     true        "Offset"
// @Param   limit      query    int     true        "Limit"
// @Success 200 {string} string	"ok"
// @Failure 400 {object} web.APIError "We need ID!!"
// @Failure 404 {object} web.APIError "Can not find ID"
// @Router /testapi/get-struct-array-by-string/{some_id} [get]
func GetStructArrayByString(c *gin.Context) {

}
