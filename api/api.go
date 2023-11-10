package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/CoulsonChen/BitoAPI/internal/model"
	"github.com/CoulsonChen/BitoAPI/internal/service"
	"github.com/gin-gonic/gin"
)

// @Summary Add a new pet to the store
// @Description get string by ID
// @Accept  json
// @Produce  json
// @Param Name body model.Person true "Person"
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

	data, err := match.AddPerson(*person)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, data)
}

// @Summary Remove person from sorted set
// @Description Remove person from sorted set
// @Accept  json
// @Produce  json
// @Param person_id path integer true "Person ID"
// @Success 200 {string} string	"ok"
// @Failure 400 {object} web.APIError "We need ID!!"
// @Router /Match/Remove/{person_id} [delete]
func RemovePerson(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	match := service.GetMatchServiceInstance()
	match.RemovePerson(id)
}

// @Summary Query matches for person
// @Description Query matches for person
// @Accept  json
// @Produce  json
// @Param person_id path integer true "Person ID"
// @Param top_n query integer false "Top N Matches"
// @Success 200 {string} string	"ok"
// @Failure 400 {object} web.APIError "We need ID!!"
// @Router /Match/QueryMatches/{person_id} [get]
func QueryMatches(c *gin.Context) {
	id_str := c.Param("id")
	id, err := strconv.Atoi(id_str)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	topn_str := c.Request.URL.Query().Get("top_n")
	topn, err := strconv.Atoi(topn_str)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	match := service.GetMatchServiceInstance()
	result, err := match.QueryMatches(id, topn)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, result)
}

// @Summary Match people
// @Description Match people
// @Accept  json
// @Produce  json
// @Param person_id path integer true "Person ID"
// @Param person_to_match_id path integer true "Person-to-match ID"
// @Success 200 {string} string	"ok"
// @Failure 400 {object} web.APIError "We need ID!!"
// @Router /Match/{person_id}/{person_to_match_id} [put]
func Match(c *gin.Context) {
	id_str := c.Param("id")
	id, err := strconv.Atoi(id_str)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	idm_str := c.Param("idm")
	idm, err := strconv.Atoi(idm_str)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	match := service.GetMatchServiceInstance()
	match.PersonMatch(id, idm)
}
