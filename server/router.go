package server

import (
	"ttc-go/api"
	"ttc-go/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Cors())

	v1 := r.Group("/api/v1")
	{
		v1.POST("ping", api.Ping)

		v1.POST("island", api.CreateIsland)
		islandAuth := v1.Group("/")
		islandAuth.Use(middleware.IslandAuth())
		{
			islandAuth.GET("island/:id", api.ShowIslandInfo)
			islandAuth.POST("island/:id/open", api.OpenIsland)
			islandAuth.POST("island/:id/close", api.CloseIsland)
			// islandAuth.PUT("island/:id", api.UpdateIslandInfo)
			// islandAuth.POST("island/:id/report", api.ReportSeller)
			// islandAuth.POST("island/:id/kick", api.KickSeller)
		}
		// v1.GET("islands/:page", api.ShowIslands)

		v1.POST("/seller", api.CreateSeller)
		sellerAuth := v1.Group("/")
		sellerAuth.Use(middleware.SellerAuth())
		{
			sellerAuth.GET("/seller/:id", api.ShowSellerInfo)
			// 	v1.POST("/seller/:id/join", api.JoinIsland)
			// 	v1.POST("/seller/:id/quit", api.QuitIsland)
			//	v1.POST("/seller/:id/extend", api.SellerExtend)
			// 	v1.POST("/seller/:id/report", api.ReportIsland)
		}
	}

	return r
}
