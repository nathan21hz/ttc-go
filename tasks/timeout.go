package tasks

import (
	"fmt"
	"strconv"
	"time"
	"ttc-go/cache"
	"ttc-go/model"

	"github.com/go-redis/redis"
)

func s2i(s []string) []int {
	var t2 = []int{}
	for _, i := range s {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		t2 = append(t2, j)
	}
	return t2
}

func DelTimeOutUsers() error {
	var err error
	now := time.Now().Unix()
	// delete time out sellers
	opt := redis.ZRangeBy{
		Min: "0",
		Max: strconv.FormatInt(now-90, 10),
	}
	sellerIDs := cache.RedisClient.ZRangeByScore("hb:seller", opt).Val()
	fmt.Println("%v", sellerIDs)
	if len(sellerIDs) > 0 {
		var sellers []model.Seller
		model.DB.Find(&sellers, sellerIDs)
		for _, seller := range sellers {
			cache.RedisClient.ZRem(strconv.Itoa(int(seller.IslandID)), strconv.Itoa(int(seller.ID)))
		}
		err = model.DB.Delete(&model.Seller{}, sellerIDs).Error
		if err != nil {
			return err
		}
	}
	cache.RedisClient.ZRemRangeByScore("hb:seller", "0", strconv.FormatInt(now-90, 10))
	// delete time out islands
	islandIDs := cache.RedisClient.ZRangeByScore("hb:island", opt).Val()
	fmt.Println("%v", islandIDs)
	if len(islandIDs) > 0 {
		err = model.DB.Table("sellers").Where("island_id IN (?)", s2i(islandIDs)).Updates(map[string]interface{}{"island_id": 0, "status": 0}).Error
		if err != nil {
			return err
		}
		err = model.DB.Delete(&model.Island{}, islandIDs).Error
		if err != nil {
			return err
		}
	}
	cache.RedisClient.ZRemRangeByScore("hb:island", "0", strconv.FormatInt(now-90, 10))
	for _, islandID := range islandIDs {
		cache.RedisClient.Del(islandID)
	}
	return err
}
