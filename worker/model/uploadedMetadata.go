package model

import "gorm.io/gorm"

type UploadedTempMetadata struct {
	Id          string   `json:"id"`
	VideoName   string   `json:"video-name"`
	Resolutions []string `json:"resolutions"`
	Path        string   `json:"path"`
}

type Video struct {
	gorm.Model
	Name         string         `json:"name"`
	VideoVariant []VideoVariant `json:"videoVariants"`
}

type VideoVariant struct {
	gorm.Model
	VideoID    uint   `json:"videoId"`
	Resolution string `json:"resolution"`
	Key        string `json:"key"`
}
