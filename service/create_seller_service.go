package service

import (
	"time"
	"ttc-go/model"
	"ttc-go/serializer"
	"ttc-go/util"
)

type CreateSellerService struct {
	IP string
}

func (service *CreateSellerService) CreateSeller() serializer.Response {
	token := util.GenToken(service.IP)
	seller := model.Seller{
		Status:        0,
		Name:          "",
		LastHeartbeat: time.Now(),
		IslandID:      0,
		IP:            service.IP,
		Token:         token,
	}
	err := model.DB.Create(&seller).Error
	if err != nil {
		return serializer.Response{
			Status: 50001,
			Msg:    "创建失败",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: 0,
		Msg:    "成功",
		Data:   serializer.BuildSellerInit(seller),
	}
}
