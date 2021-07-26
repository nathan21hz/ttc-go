package service

import (
	"strconv"
	"time"
	"ttc-go/cache"
	"ttc-go/model"
	"ttc-go/serializer"

	"github.com/go-redis/redis"
)

type JoinIslandService struct {
	Name   string `json:"name" binding:"required,max=20"`
	Island uint   `json:"island" binding:"required"`
}

func (service *JoinIslandService) Join(seller *model.Seller) serializer.Response {
	var island model.Island
	err := model.DB.First(&island, service.Island).Error

	if err != nil { // Island not exist
		return serializer.Response{
			Status: 404,
			Msg:    "Island not exist",
			Error:  err.Error(),
		}
	}

	if seller.Status == 0 { // seller in idle
		if island.Status == 0 { // island not open
			return serializer.Response{
				Status: 40008,
				Msg:    "Island not open",
			}
		} else { // island open
			cache.RedisClient.ZAdd(
				strconv.Itoa(int(island.ID)),
				redis.Z{
					Score:  float64(time.Now().UnixNano()),
					Member: strconv.Itoa(int(seller.ID)),
				},
			)
			qLen := cache.RedisClient.ZCard(strconv.Itoa(int(island.ID))).Val()
			qStatus := 0
			if qLen > int64(island.MaxSeller) {
				qStatus = 2
			} else {
				qStatus = 1
			}
			err = model.DB.Model(&seller).Updates(
				map[string]interface{}{
					"IslandID": island.ID,
					"Status":   qStatus,
				}).Error
			if err != nil {
				return serializer.Response{
					Status: 50001,
					Msg:    "Database Error",
					Error:  err.Error(),
				}
			}
			return serializer.Response{
				Status: 200,
				Msg:    "Success",
				Data:   qStatus,
			}
		}

	} else {
		return serializer.Response{
			Status: 40005,
			Msg:    "In trading or queue now",
		}
	}
}