package api

import (
	"github.com/gin-gonic/gin"

	"ttc-go/model"
	"ttc-go/service"
)

//CreateSeller Create/Init a seller
func CreateSeller(c *gin.Context) {
	service := service.CreateSellerService{
		IP: c.ClientIP(),
	}
	if err := c.ShouldBind(&service); err == nil {
		res := service.CreateSeller()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

//ShowSellerInfo seller info for seller, Login required
func ShowSellerInfo(c *gin.Context) {
	service := service.ShowSellerInfoService{}
	sellerGet, _ := c.Get("seller")
	seller, _ := sellerGet.(*model.Seller)
	res := service.Show(seller)
	c.JSON(200, res)

}

//JoinIsland, Login required
func JoinIsland(c *gin.Context) {
	service := service.JoinIslandService{}
	sellerGet, _ := c.Get("seller")
	seller, _ := sellerGet.(*model.Seller)
	if err := c.ShouldBind(&service); err == nil {
		res := service.Join(seller)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}

}

//QuitIsland, Login required
func QuitIsland(c *gin.Context) {
	service := service.QuitIslandService{}
	sellerGet, _ := c.Get("seller")
	seller, _ := sellerGet.(*model.Seller)
	res := service.Quit(seller)
	c.JSON(200, res)
}
