package service

import (
	"strconv"
	"ttc-go/cache"
	"ttc-go/model"
	"ttc-go/serializer"
)

type CloseIslandService struct {
}

func (service CloseIslandService) Open(island *model.Island) serializer.Response {
	if island.Status == 0 { //Already closed
		return serializer.Response{
			Status: 40006,
			Msg:    "Already Closed",
		}
	}
	model.DB.Model(&island).Update("Status", 0)
	cache.RedisClient.Del(strconv.Itoa(int(island.ID)))
	model.DB.Model(model.Seller{}).Where("island_id = ?", island.ID).Updates(map[string]interface{}{"Status": 0, "IslandID": 0})

	return serializer.Response{
		Status: 0,
		Msg:    "Closed",
	}
}
