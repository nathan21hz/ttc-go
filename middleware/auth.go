package middleware

import (
	"ttc-go/model"
	"ttc-go/serializer"

	"github.com/gin-gonic/gin"
)

type IslandID struct {
	ID uint `uri:"id" binding:"required"`
}

type SellerID struct {
	ID uint `uri:"id" binding:"required"`
}

// IslandAuth Island Owner Auth
func IslandAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var islandID IslandID
		if err := c.ShouldBindUri(&islandID); err != nil {
			c.JSON(200, serializer.Response{
				Msg:    "Island ID Error",
				Status: 40001,
			})
			c.Abort()
			return
		}
		token := c.Query("token")
		if token == "" {
			c.JSON(200, serializer.Response{
				Msg:    "No Token",
				Status: 40002,
			})
			c.Abort()
			return
		}
		island, err := model.GetIsland(islandID.ID, token)

		if err != nil {
			c.JSON(200, serializer.Response{
				Msg:    "Token Error",
				Status: 40003,
			})
			c.Abort()
			return
		} else {
			c.Set("island", &island)
			c.Next()
			return
		}

	}
}

// SellerAuth Island Owner Auth
func SellerAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var sellerID SellerID
		if err := c.ShouldBindUri(&sellerID); err != nil {
			c.JSON(200, serializer.Response{
				Msg:    "Seller ID Error",
				Status: 40001,
			})
			c.Abort()
			return
		}
		token := c.Query("token")
		if token == "" {
			c.JSON(200, serializer.Response{
				Msg:    "No Token",
				Status: 40002,
			})
			c.Abort()
			return
		}
		seller, err := model.GetSeller(sellerID.ID, token)

		if err != nil {
			c.JSON(200, serializer.Response{
				Msg:    "Token Error",
				Status: 40003,
			})
			c.Abort()
			return
		} else {
			c.Set("seller", &seller)
			c.Next()
			return
		}

	}
}
