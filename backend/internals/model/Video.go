package model

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	Name         string         `json:"name"`
	VideoVariant []VideoVariant `json:"videoVariants"`
}

type VideoVariant struct {
	gorm.Model
	VideoID    uint   `json:"videoId"`
	Resolution string `json:"resolution"`
	URL        string `json:"url"`
}

// Response

type VideoVariantResp struct {
	ID         uint   `json:"id"`
	Resolution string `json:"resolution"`
	URL        string `json:"url"`
}

type VideoResp struct {
	ID           uint               `json:"id"`
	Name         string             `json:"name"`
	VideoVariant []VideoVariantResp `json:"videoVariants"`
}

type GetAllVideosResp struct {
	Videos  []VideoResp `json:"videos"`
	Message string      `json:"message"`
}

type UploadVideoResp struct {
	ID      string `json:"id"`
	Message string `json:"message"`
}
