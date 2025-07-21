package repository

import (
	"github.com/Diaku49/FoodOrderSystem/backend/internals/model"
	"gorm.io/gorm"
)

type VideoRepository struct {
	DB *gorm.DB
}

func (vr *VideoRepository) GetAllVideos(limit, offset int) ([]model.Video, error) {
	var videos []model.Video

	err := vr.DB.Limit(limit).Offset(offset).
		Select("id", "name").
		Preload("VideoVariant", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "video_id", "resolution", "url")
		}).
		Find(&videos).Error

	if err != nil {
		return nil, err
	}

	return videos, nil
}
