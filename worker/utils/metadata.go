package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Diaku49/FoodOrderSystem/worker/model"
)

const metaDir = "tmp/meta"

func LoadMetadata(id string) (model.UploadedTempMetadata, error) {
	filePath := filepath.Join(metaDir, id+".json")

	data, err := os.ReadFile(filePath)
	if err != nil {
		return model.UploadedTempMetadata{}, fmt.Errorf("failed to read metadata file: %w", err)
	}

	var meta model.UploadedTempMetadata
	err = json.Unmarshal(data, &meta)
	if err != nil {
		return model.UploadedTempMetadata{}, fmt.Errorf("failed to parse metadata: %w", err)
	}

	return meta, nil
}

func DeleteMetadata(id string) error {
	filePath := filepath.Join(metaDir, id+".json")
	return os.Remove(filePath)
}
