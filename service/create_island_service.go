package service

import (
	"time"
	"ttc-go/model"
	"ttc-go/serializer"
	"ttc-go/util"
)

type CreateIslandService struct {
	IP string
}

func (service *CreateIslandService) CreateIsland() serializer.Response {
	token := util.GenToken(service.IP)
	island := model.Island{
		Status:        0,
		Name:          "",
		LastHeartbeat: time.Now(),
		Price:         0,
		Password:      "",
		Remark:        "",
		IP:            service.IP,
		Token:         token,
		MaxSeller:     0,
	}
	err := model.DB.Create(&island).Error
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
		Data:   serializer.BuildIslandInit(island),
	}
}
