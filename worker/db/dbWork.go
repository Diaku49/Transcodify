package db

import (
	"github.com/Diaku49/FoodOrderSystem/worker/model"
	"gorm.io/gorm"
)

func CreateVideo(db *gorm.DB, videoName string, variantMap []model.VideoVariant) error {
	// create the record
	video := model.Video{
		Name:         videoName,
		VideoVariant: variantMap,
	}

	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&video).Error; err != nil {
			return err
		}
		return nil
	})
}
