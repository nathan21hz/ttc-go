package service

import (
	"ttc-go/model"
	"ttc-go/serializer"
)

type ShowSellerInfoService struct {
}

func (service *ShowSellerInfoService) Show(seller *model.Seller) serializer.Response {
	var island model.Island
	seller.UpdateHeartbeat()
	if seller.Status != 0 {
		island.ID = seller.IslandID
		err := model.DB.First(&island).Error
		if err != nil {
			return serializer.Response{
				Status: 404,
				Msg:    "Island not exist",
				Error:  err.Error(),
			}
		}
	}

	return serializer.Response{
		Data: serializer.BuildSellerInfo(*seller, island),
	}
}
