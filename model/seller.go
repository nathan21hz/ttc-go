package model

import (
	"strconv"
	"time"
	"ttc-go/cache"

	"github.com/jinzhu/gorm"
)

// Seller 岛民模型
type Seller struct {
	gorm.Model
	Status        uint
	Name          string
	LastHeartbeat time.Time
	LastTrade     time.Time
	IslandID      uint
	Password      string
	IP            string
	Token         string
}

func GetSeller(id uint, token string) (Seller, error) {
	var seller Seller
	err := DB.Where("id = ? AND token = ?", id, token).First(&seller).Error
	return seller, err
}

// UpdateHeartbeat Update Last Heartbeat Time
func (seller *Seller) UpdateHeartbeat() {
	DB.Model(&seller).Update("LastHeartbeat", time.Now())
}

func (seller *Seller) QueueStatus() (int, int, int) {
	if seller.Status == 0 {
		return 0, 0, 0
	}
	var island Island
	err := DB.First(&island, seller.IslandID).Error
	if err != nil {
		return 0, 0, 0
	}
	status, pos, queueLength := 0, 0, 0
	qPos := int(cache.RedisClient.ZRank(strconv.Itoa(int(seller.IslandID)), strconv.Itoa(int(seller.ID))).Val())
	if qPos < int(island.MaxSeller) {
		status = 1
		pos = qPos
		queueLength = int(island.MaxSeller)
	} else {
		status = 2
		pos = qPos - int(island.MaxSeller) + 1
		queueLength = int(cache.RedisClient.ZCard(strconv.Itoa(int(seller.IslandID))).Val()) - int(island.MaxSeller)
	}
	return status, pos, queueLength
}
