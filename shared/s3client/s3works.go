package s3client

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/aws/aws-sdk-go/service/s3"
)

func (sc *StorageClient) Upload(id, resolution, localPath string) (string, error) {
	file, err := os.Open(localPath)
	if err != nil {
		return "", fmt.Errorf("open file error: %w", err)
	}
	defer file.Close()

	buf := bytes.NewBuffer(nil)
	io.Copy(buf, file)

	key := id + ": " + resolution
	_, err = sc.Client.PutObject(&s3.PutObjectInput{
		Bucket: &sc.Bucket,
		Key:    &key,
		Body:   bytes.NewReader(buf.Bytes()),
	})
	if err != nil {
		return "", fmt.Errorf("uploading to liara failed: %v", err)
	}

	return key, nil
}

func (sc *StorageClient) Delete(key string) error {
	_, err := sc.Client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: &sc.Bucket,
		Key:    &key,
	})
	if err != nil {
		return fmt.Errorf("failed to download video: %v", err)
	}

	return nil
}
