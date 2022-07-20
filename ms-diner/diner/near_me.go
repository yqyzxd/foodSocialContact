package diner

import (
	"context"
	"github.com/go-redis/redis/v8"
	"strconv"
)

type NearMeService struct {
	Redis *redis.Client
}

func (s *NearMeService) UpdateDinerLocation(dinerID int, lon, lat float64) {

	key := "diner:location"
	s.Redis.GeoAdd(context.Background(), key, &redis.GeoLocation{
		Name:      strconv.Itoa(dinerID),
		Latitude:  lat,
		Longitude: lon,
	})
}
