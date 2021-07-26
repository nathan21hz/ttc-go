package service

import (
	"ttc-go/model"
	"ttc-go/serializer"
)

type ShowIslandInfoService struct {
}

func (service *ShowIslandInfoService) Show(island *model.Island) serializer.Response {
	island.UpdateHeartbeat()
	return serializer.Response{
		Data: serializer.BuildIslandInfo(*island),
	}
}
