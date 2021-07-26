package service

import (
	"ttc-go/model"
	"ttc-go/serializer"
)

type UpdateIslandService struct {
	Name      string `json:"name" binding:"required,max=20"`
	Price     uint   `json:"price" binding:"required,min=0,max=700"`
	Password  string `json:"password" binding:"required,max=10"`
	Remark    string `json:"remark" binding:"max=200"`
	MaxSeller uint   `json:"max_seller" binding:"required,min=1,max=6"`
}

func (service UpdateIslandService) Update(island *model.Island) serializer.Response {
	island.Name = service.Name
	island.Price = service.Price
	island.Password = service.Password
	island.Remark = service.Remark
	island.MaxSeller = service.MaxSeller
	model.DB.Save(&island)
	island.UpdateHeartbeat()
	return serializer.Response{
		Status: 0,
		Msg:    "Updated",
	}
}
