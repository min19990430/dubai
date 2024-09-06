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
	stations, err := sc.station.List(domain.Station{})
	if err != nil {
		sc.response.FailWithError(c, "fail to get stations", err)
		return
	}

	sc.response.SuccessWithData(c, "success", stations)
}
