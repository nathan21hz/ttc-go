package serializer

import (
	"strconv"
	"ttc-go/model"
)

type IslandInitResponse struct {
	IslandID string
	Token    string
}

type IslandInfoResponse struct {
	IslandID string
	Status   uint
}

func BuildIslandInit(island model.Island) IslandInitResponse {
	return IslandInitResponse{
		IslandID: strconv.FormatUint(uint64(island.ID), 10),
		Token:    island.Token,
	}
}

func BuildIslandInfo(island model.Island) IslandInfoResponse {
	return IslandInfoResponse{
		IslandID: strconv.FormatUint(uint64(island.ID), 10),
		Status:   island.Status,
	}
}
