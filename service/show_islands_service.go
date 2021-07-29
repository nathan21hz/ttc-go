package service

import (
	"ttc-go/model"
	"ttc-go/serializer"
)

type ShowIslandsService struct {
}

func (service ShowIslandsService) Show() serializer.Response {
	var islands []model.Island
	err := model.DB.Where("status = ?", 1).Find(&islands).Error
	if err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "Database error",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Data: serializer.BuildIslandInfoList(islands),
	}
}
