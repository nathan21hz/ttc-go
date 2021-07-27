package service

import (
	"strconv"
	"ttc-go/cache"
	"ttc-go/model"
	"ttc-go/serializer"
)

type QuitIslandService struct {
}

func (service QuitIslandService) Quit(seller *model.Seller) serializer.Response {
	if seller.Status == 0 {
		return serializer.Response{
			Status: 40005,
			Msg:    "Not in a island",
		}
	}
	cache.RedisClient.ZRem(strconv.Itoa(int(seller.IslandID)), strconv.Itoa(int(seller.ID)))
	err := model.DB.Model(&seller).Updates(
		map[string]interface{}{
			"IslandID": 0,
			"Status":   0,
		}).Error
	if err != nil {
		return serializer.Response{
			Status: 50001,
			Msg:    "Database Error",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: 0,
		Msg:    "Success",
	}
}
