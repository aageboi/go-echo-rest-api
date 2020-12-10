package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"encoding/json"

	"github.com/aageboi/go-echo-rest-api/models"
	"github.com/labstack/echo"

	"gopkg.in/mgo.v2/bson"
)

// FindAllArticle get list latest article
func (h *Handler) FindAllArticle(c echo.Context) (err error) {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	// Defaults
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 10
	}

	// Retrive articles from database
	articles := []*models.ArticleList{}
	db := h.DB.Clone()
	rds := h.REDIS

	const objectPrefix string = "X:article:all"

	// get data from redis
	s, err := rds.Get(objectPrefix).Result()
	if err != nil {
		// fmt.Println(rds, err)
	}

	if s != "" {
		json.Unmarshal([]byte(s), &articles)
		return c.JSON(http.StatusOK, articles)
	}

	// get data from mongodb
	if err = db.DB("X").C("article").
		Find(bson.M{flag_active: 1}).
		Sort("-date_news").
		Skip((page - 1) * limit).
		Limit(limit).
		All(&articles); err != nil {
		var e []string

		// return empty data
		return c.JSON(http.StatusNoContent, e)
	}
	defer db.Close()

	// save data from mongo to redis
	j, err := json.Marshal(articles)
	fmt.Println("Setting redis data into key " + objectPrefix)
	rds.Set(objectPrefix, j, 0)
	if err != nil {
	}

	// return data
	return c.JSON(http.StatusOK, articles)
}

// FindArticleByID get detail article
func (h *Handler) FindArticleByID(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	article := &models.ArticleDetail{}
	db := h.DB.Clone()
	if err = db.DB("X").C("article").
		Find(bson.M{"id": id}).
		One(article); err != nil {
		var e []string
		return c.JSON(http.StatusNoContent, e)
	}
	defer db.Close()

	// save data from mongo to redis
	j, err := json.Marshal(articles)
	fmt.Println("Setting redis data into key " + objectPrefix)
	rds.Set(objectPrefix, j, 0)
	if err != nil {
	}
	return c.JSON(http.StatusOK, article)
}
