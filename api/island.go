package api

import (
	"github.com/gin-gonic/gin"

	"ttc-go/model"
	"ttc-go/service"
)

func CreateIsland(c *gin.Context) {
	service := service.CreateIslandService{
		IP: c.ClientIP(),
	}
	if err := c.ShouldBind(&service); err == nil {
		res := service.CreateIsland()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func ShowIslandInfo(c *gin.Context) {
	service := service.ShowIslandInfoService{}
	islandGet, _ := c.Get("island")
	island, _ := islandGet.(*model.Island)
	res := service.Show(island)
	c.JSON(200, res)

}

func OpenIsland(c *gin.Context) {
	service := service.OpenIslandService{}
	islandGet, _ := c.Get("island")
	island, _ := islandGet.(*model.Island)
	res := service.Open(island)
	c.JSON(200, res)
}

func CloseIsland(c *gin.Context) {
	service := service.CloseIslandService{}
	islandGet, _ := c.Get("island")
	island, _ := islandGet.(*model.Island)
	res := service.Open(island)
	c.JSON(200, res)
}
