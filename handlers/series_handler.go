package handlers

import (
	"net/http"
	"strconv"

	"github.com/capybara120404/watch-tracker/models"
	"github.com/capybara120404/watch-tracker/repository"
	"github.com/gin-gonic/gin"
)

type SeriesHandler struct {
	repository *repository.SeriesInfoRepository
}

func NewSeriesHandler(repository *repository.SeriesInfoRepository) *SeriesHandler {
	return &SeriesHandler{
		repository: repository,
	}
}

func (handler *SeriesHandler) GetAllSeriesHandler(c *gin.Context) {
	series, err := handler.repository.GetAll()
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	if series == nil {
		series = []models.Series{}
	}

	c.JSON(http.StatusOK, map[string]interface{}{"series": series})
}

func (handler *SeriesHandler) GetSeriesByIdHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(http.StatusBadRequest, "invalid task Id format")
		return
	}

	series, err := handler.repository.GetById(id)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"series": series})
}
