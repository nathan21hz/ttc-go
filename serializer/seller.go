package serializer

import (
	"ttc-go/model"
)

type SellerInitResponse struct {
	SellerID uint   `json:"seller_id"`
	Token    string `json:"token"`
}

type SellerInfoResponse struct {
	SellerID   uint        `json:"seller_id"`
	Status     uint        `json:"status"`
	IslandInfo interface{} `json:"island_info"`
}

type IslandInfoForSeller struct {
	IslandID uint   `json:"island_id"`
	Name     string `json:"name"`
	Price    uint   `json:"price"`
	Remark   string `json:"remark"`
	Password string `json:"password"`
}

type IslandInfoForQueue struct {
	IslandID    uint   `json:"island_id"`
	Name        string `json:"name"`
	Price       uint   `json:"price"`
	Remark      string `json:"remark"`
	QueuePos    uint   `json:"queue_pos"`
	QueueLength uint   `json:"queue_length"`
}

func BuildSellerInit(seller model.Seller) SellerInitResponse {
	return SellerInitResponse{
		SellerID: seller.ID,
		Token:    seller.Token,
	}
}

func BuildSellerInfo(seller model.Seller, island model.Island) SellerInfoResponse {
	status, pos, qLength := seller.QueueStatus()
	var islandInfo interface{}
	if status == 0 {
	} else if status == 1 {
		islandInfo = IslandInfoForSeller{
			IslandID: island.ID,
			Name:     island.Name,
			Price:    island.Price,
			Remark:   island.Remark,
			Password: island.Password,
		}
	} else {
		islandInfo = IslandInfoForQueue{
			IslandID:    island.ID,
			Name:        island.Name,
			Price:       island.Price,
			Remark:      island.Remark,
			QueuePos:    uint(pos),
			QueueLength: uint(qLength),
		}
	}
	return SellerInfoResponse{
		SellerID:   seller.ID,
		Status:     uint(status),
		IslandInfo: islandInfo,
	}
}
