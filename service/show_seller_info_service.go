package service

import (
	"ttc-go/model"
	"ttc-go/serializer"
)

type ShowSellerInfoService struct {
}

func (service *ShowSellerInfoService) Show(seller *model.Seller) serializer.Response {
	// seller.UpdateHeartbeat()
	return serializer.Response{
		Data: serializer.BuildSellerInfo(*seller),
	}
}
