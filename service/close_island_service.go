package service

import (
	"ttc-go/model"
	"ttc-go/serializer"
)

type CloseIslandService struct {
}

func (service CloseIslandService) Open(island *model.Island) serializer.Response {
	if island.Status == 0 {
		return serializer.Response{
			Status: 40006,
			Msg:    "Already Closed",
		}
	}
	model.DB.Model(&island).Update("status", 0)
	return serializer.Response{
		Status: 0,
		Msg:    "Closed",
	}
}
