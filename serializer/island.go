package serializer

import (
	"strconv"
	"ttc-go/model"
)

type IslandInitResponse struct {
	IslandID string
	Token    string
}

type SellerInfoItem struct {
	SellerID uint
	Name     string
	JoinTime interface{}
}

type IslandInfoResponse struct {
	IslandID string
	Status   uint
	Sellers  []SellerInfoItem
	Queue    []SellerInfoItem
}

func BuildIslandInit(island model.Island) IslandInitResponse {
	return IslandInitResponse{
		IslandID: strconv.FormatUint(uint64(island.ID), 10),
		Token:    island.Token,
	}
}

func BuildSellerInfoItem(item model.Seller) SellerInfoItem {
	return SellerInfoItem{
		SellerID: item.ID,
		Name:     item.Name,
	}
}

func BuildIslandInfo(island model.Island, sellers []model.Seller) IslandInfoResponse {
	var activeList []SellerInfoItem
	var queueList []SellerInfoItem
	for i, item := range sellers {
		if i < int(island.MaxSeller) {
			seller := BuildSellerInfoItem(item)
			activeList = append(activeList, seller)
		} else {
			seller := BuildSellerInfoItem(item)
			queueList = append(queueList, seller)
		}
	}
	return IslandInfoResponse{
		IslandID: strconv.FormatUint(uint64(island.ID), 10),
		Status:   island.Status,
		Sellers:  activeList,
		Queue:    queueList,
	}
}
