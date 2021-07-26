package service

import (
	"fmt"
	"strconv"
	"strings"
	"ttc-go/cache"
	"ttc-go/model"
	"ttc-go/serializer"
)

type ShowIslandInfoService struct {
}

func (service *ShowIslandInfoService) Show(island *model.Island) serializer.Response {
	var sellers []model.Seller

	queue := cache.RedisClient.ZRange(strconv.Itoa(int(island.ID)), 0, -1).Val()
	if len(queue) > 0 {
		order := fmt.Sprintf("FIELD(id, %s)", strings.Join(queue, ","))
		err := model.DB.Where("id in (?)", queue).Order(order).Find(&sellers).Error
		if err != nil {
			return serializer.Response{
				Status: 50000,
				Msg:    "Database error",
				Error:  err.Error(),
			}
		}
	}

	island.UpdateHeartbeat()
	return serializer.Response{
		Data: serializer.BuildIslandInfo(*island, sellers),
	}
}
