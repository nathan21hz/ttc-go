package service

import (
	"ttc-go/model"
	"ttc-go/serializer"
)

type OpenIslandService struct {
}

func (service OpenIslandService) Open(island *model.Island) serializer.Response {
	if island.Status == 1 {
		return serializer.Response{
			Status: 40005,
			Msg:    "Already Opened",
		}
	}
	model.DB.Model(&island).Update("status", 1)
	return serializer.Response{
		Status: 0,
		Msg:    "Opened",
	}
}
