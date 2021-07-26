package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Seller 岛民模型
type Seller struct {
	gorm.Model
	Status        uint
	Name          string
	LastHeartbeat time.Time
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
