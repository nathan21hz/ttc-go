package service

import (
	"ttc-go/model"
	"ttc-go/serializer"
)

type OpenIslandService struct {
	Name      string `json:"name" binding:"required,max=20"`
	Price     uint   `json:"price" binding:"required,min=0,max=700"`
	Password  string `json:"password" binding:"required,max=10"`
	Remark    string `json:"remark" binding:"max=200"`
	MaxSeller uint   `json:"max_seller" binding:"required,min=1,max=6"`
}

func (service OpenIslandService) Open(island *model.Island) serializer.Response {
	if island.Status == 1 {
		return serializer.Response{
			Status: 40005,
			Msg:    "Already Opened",
		}
	}
	island.Name = service.Name
	island.Price = service.Price
	island.Password = service.Password
	island.Remark = service.Remark
	island.MaxSeller = service.MaxSeller
	island.Status = 1
	model.DB.Save(&island)
	return serializer.Response{
		Status: 0,
		Msg:    "Opened",
	}
}
