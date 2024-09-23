package controller

import (
	"github.com/gin-gonic/gin"

	"auto-monitoring/internal/adapter/gin-http/controller/response"
	"auto-monitoring/internal/application/usecase"
	"auto-monitoring/internal/domain"
)

type StationController struct {
	response response.IResponse
	station  usecase.StationUsecase
}

func NewStationController(response response.IResponse, usecase *usecase.StationUsecase) *StationController {
	return &StationController{
		response: response,
		station:  *usecase,
	}
}

func (sc *StationController) GetList(c *gin.Context) {
	stations, err := sc.station.List(domain.Station{IsEnable: true})
	if err != nil {
		sc.response.FailWithError(c, queryFail, err)
		return
	}

	sc.response.SuccessWithData(c, querySuccess, stations)
}

type StationRequest struct {
	UUID        string  `json:"uuid" binding:"required,uuid"`
	ID          string  `json:"id" binding:"required"`
	Name        string  `json:"name" binding:"required"`
	Priority    int     `json:"priority" `
	Address     string  `json:"address"`
	IsEnable    bool    `json:"is_enable" `
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Description string  `json:"description"`
}

func (sc *StationController) PostCreate(c *gin.Context) {
	var req StationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		sc.response.ValidatorFail(c, paramError)
		return
	}

	station := domain.Station{
		UUID:        req.UUID,
		ID:          req.ID,
		Name:        req.Name,
		Priority:    req.Priority,
		Address:     req.Address,
		IsEnable:    req.IsEnable,
		Lat:         req.Lat,
		Lon:         req.Lon,
		Description: req.Description,
	}

	if err := sc.station.Create(station); err != nil {
		sc.response.FailWithError(c, createFail, err)
		return
	}

	sc.response.Success(c, createSuccess)
}

func (sc *StationController) PutUpdate(c *gin.Context) {
	var req StationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		sc.response.ValidatorFail(c, paramError)
		return
	}

	station := domain.Station{
		UUID:        req.UUID,
		ID:          req.ID,
		Name:        req.Name,
		Priority:    req.Priority,
		Address:     req.Address,
		IsEnable:    req.IsEnable,
		Lat:         req.Lat,
		Lon:         req.Lon,
		Description: req.Description,
	}

	if err := sc.station.Update(station); err != nil {
		sc.response.FailWithError(c, updateFail, err)
		return
	}

	sc.response.Success(c, updateSuccess)
}

func (sc *StationController) Delete(c *gin.Context) {
	uuid := c.Param("uuid")

	if err := sc.station.Delete(uuid); err != nil {
		sc.response.FailWithError(c, deleteFail, err)
		return
	}

	sc.response.Success(c, deleteSuccess)
}
