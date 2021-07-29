package serializer

import (
	"ttc-go/model"
)

type IslandInitResponse struct {
	IslandID uint   `json:"island_id"`
	Token    string `json:"token"`
}

type SellerInfoItem struct {
	SellerID uint        `json:"seller_id"`
	Name     string      `json:"name"`
	JoinTime interface{} `json:"join_time"`
}

type IslandInfoResponse struct {
	IslandID uint             `json:"island_id"`
	Status   uint             `json:"status"`
	Sellers  []SellerInfoItem `json:"sellers"`
	Queue    []SellerInfoItem `json:"queue"`
}

type IslandInfoItem struct {
	IslandID    uint   `json:"island_id"`
	Name        string `json:"name"`
	Price       uint   `json:"price"`
	Remark      string `json:"remark"`
	SellerCount uint   `json:"seller_count"`
	QueueLength uint   `json:"queue_length"`
	MaxSeller   uint   `json:"max_seller"`
}

func BuildIslandInit(island model.Island) IslandInitResponse {
	return IslandInitResponse{
		IslandID: island.ID,
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
		IslandID: island.ID,
		Status:   island.Status,
		Sellers:  activeList,
		Queue:    queueList,
	}
}

func BuildIslandInfoItem(island model.Island) IslandInfoItem {
	sellerCount, queueLength := island.GetQueueInfo()
	return IslandInfoItem{
		IslandID:    island.ID,
		Name:        island.Name,
		Price:       island.Price,
		Remark:      island.Remark,
		SellerCount: sellerCount,
		QueueLength: queueLength,
		MaxSeller:   island.MaxSeller,
	}
}

func BuildIslandInfoList(items []model.Island) []IslandInfoItem {
	var islands []IslandInfoItem
	for _, item := range items {
		island := BuildIslandInfoItem(item)
		islands = append(islands, island)
	}
	return islands
}
