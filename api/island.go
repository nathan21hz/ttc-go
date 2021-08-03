package api

import (
	"github.com/gin-gonic/gin"

	"ttc-go/model"
	"ttc-go/service"
)

// CreateIsland 
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

// ShowIslandInfo return island info for island owner, Token Required
func ShowIslandInfo(c *gin.Context) {
	service := service.ShowIslandInfoService{}
	islandGet, _ := c.Get("island")
	island, _ := islandGet.(*model.Island)
	res := service.Show(island)
	c.JSON(200, res)

}

//OpenIsland, Token Required
func OpenIsland(c *gin.Context) {
	service := service.OpenIslandService{}
	islandGet, _ := c.Get("island")
	island, _ := islandGet.(*model.Island)
	if err := c.ShouldBind(&service); err == nil {
		res := service.Open(island)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

//CloseIsland, Token Required
func CloseIsland(c *gin.Context) {
	service := service.CloseIslandService{}
	islandGet, _ := c.Get("island")
	island, _ := islandGet.(*model.Island)
	res := service.Open(island)
	c.JSON(200, res)
}

//UpdateIslandInfo, Token Required
func UpdateIslandInfo(c *gin.Context) {
	service := service.UpdateIslandService{}
	islandGet, _ := c.Get("island")
	island, _ := islandGet.(*model.Island)
	if err := c.ShouldBind(&service); err == nil {
		res := service.Update(island)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

//ShowIslands, Island list for seller
func ShowIslands(c *gin.Context) {
	service := service.ShowIslandsService{}
	res := service.Show()
	c.JSON(200, res)
}

//KickSeller, Token Required
func KickSeller(c *gin.Context) {
	service := service.KickSellerService{}
	islandGet, _ := c.Get("island")
	island, _ := islandGet.(*model.Island)
	if err := c.ShouldBind(&service); err == nil {
		res := service.Kick(island)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
