package model

import (
	"strconv"
	"time"
	"ttc-go/cache"

	"github.com/jinzhu/gorm"
)

// Island 岛主模型
type Island struct {
	gorm.Model
	Status        uint
	Name          string
	LastHeartbeat time.Time
	Price         uint
	Password      string
	Remark        string
	IP            string
	Token         string
	MaxSeller     uint
}

// GetIsland get island by id and token
func GetIsland(id uint, token string) (Island, error) {
	var island Island
	err := DB.Where("id = ? AND token = ?", id, token).First(&island).Error
	return island, err
}

// UpdateHeartbeat Update Last Heartbeat Time
func (island *Island) UpdateHeartbeat() {
	DB.Model(&island).Update("LastHeartbeat", time.Now())
}

func (island *Island) GetQueueInfo() (uint, uint) {
	listLength := uint(cache.RedisClient.ZCard(strconv.Itoa(int(island.ID))).Val())
	if listLength <= island.MaxSeller {
		return listLength, 0
	} else {
		return island.MaxSeller, listLength - island.MaxSeller
	}

}
