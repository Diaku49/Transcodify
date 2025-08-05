package model

type UploadMetadata struct {
	Resolutions []string `json:"resolutions"`
	VideoName   string   `json:"video-name"`
}

type UploadedTempMetadata struct {
	Id          string   `json:"id"`
	VideoName   string   `json:"video-name"`
	Resolutions []string `json:"resolutions"`
	Path        string   `json:"path"`
}
