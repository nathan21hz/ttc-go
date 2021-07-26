package api

import (
	"github.com/gin-gonic/gin"

	"ttc-go/model"
	"ttc-go/service"
)

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

func ShowSellerInfo(c *gin.Context) {
	service := service.ShowSellerInfoService{}
	sellerGet, _ := c.Get("seller")
	seller, _ := sellerGet.(*model.Seller)
	res := service.Show(seller)
	c.JSON(200, res)

}
