package serializer

import (
	"strconv"
	"ttc-go/model"
)

type SellerInitResponse struct {
	SellerID string
	Token    string
}

type SellerInfoResponse struct {
	SellerID string
	Status   uint
	Island   interface{}
}

func BuildSellerInit(seller model.Seller) SellerInitResponse {
	return SellerInitResponse{
		SellerID: strconv.FormatUint(uint64(seller.ID), 10),
		Token:    seller.Token,
	}
}

func BuildSellerInfo(seller model.Seller) SellerInfoResponse {
	return SellerInfoResponse{
		SellerID: strconv.FormatUint(uint64(seller.ID), 10),
		Status:   seller.Status,
	}
}
