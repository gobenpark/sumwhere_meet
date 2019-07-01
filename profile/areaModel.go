package profile

import (
	"context"
	"sumwhere_meet/factory"
	"time"
)

type Area struct {
	ID        int       `json:"id" xorm:"id pk autoincr"`
	City      string    `json:"city" xorm:"city"`
	District  string    `json:"district" xml:"district"`
	CreatedAt time.Time `json:"createdAt" xorm:"created"`
	UpdatedAt time.Time `json:"updatedAt" xorm:"updated"`
}

func (Area) GetCity(ctx context.Context) ([]Area, error) {
	var v []Area
	err := factory.DB(ctx).GroupBy("city").Find(&v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (Area) GetDistrict(ctx context.Context, city string) ([]Area, error) {
	var v []Area
	err := factory.DB(ctx).Where("city = ?", city).Asc("id").Find(&v)
	if err != nil {
		return nil, err
	}

	return v, nil
}
