package station

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lacsapadnan/mrt-schedules/common/response"
)

func Initiate(router *gin.RouterGroup) {
	stationService := NewService()

	station := router.Group("/stations")
	station.GET("", func(c *gin.Context) {
		// service code
		GetAllStations(c, stationService)
	})
}

func GetAllStations(c *gin.Context, service Service) {
	datas, err := service.GetAllStations()
	if err != nil {
		c.JSON(http.StatusBadRequest,
			response.APIResponse{
				Success: false,
				Message: err.Error(),
				Data:    nil,
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		response.APIResponse{
			Success: true,
			Message: "Successfully get all stations",
			Data:    datas,
		},
	)
}
