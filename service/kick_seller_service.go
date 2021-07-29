package service

import (
	"strconv"
	"ttc-go/cache"
	"ttc-go/model"
	"ttc-go/serializer"

	"github.com/go-redis/redis"
)

type KickSellerService struct {
	SellerID uint `json:"seller_id" binding:"required"`
}

func (service KickSellerService) Kick(island *model.Island) serializer.Response {
	_, err := cache.RedisClient.ZScore(strconv.Itoa(int(island.ID)), strconv.Itoa(int(service.SellerID))).Result()
	switch {
	case err == redis.Nil:
		return serializer.Response{
			Status: 40004,
			Msg:    "Seller not in list",
			Error:  err.Error(),
		}
	case err != nil:
		return serializer.Response{
			Status: 50001,
			Msg:    "Redis error",
			Error:  err.Error(),
		}
	}
	var seller model.Seller
	err = model.DB.First(&seller, service.SellerID).Error
	if err != nil { // Island not exist
		return serializer.Response{
			Status: 40004,
			Msg:    "Seller not exist",
			Error:  err.Error(),
		}
	}
	cache.RedisClient.ZRem(strconv.Itoa(int(seller.IslandID)), strconv.Itoa(int(seller.ID)))
	err = model.DB.Model(&seller).Updates(
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
