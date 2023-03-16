package minioc

import (
	"context"
	"dousheng/common/config"
	"fmt"
	"mime/multipart"
	"path/filepath"

	"github.com/hashicorp/go-uuid"
	"github.com/minio/minio-go/v7"
)

var (
	videoOpt = minio.PutObjectOptions{ContentType: "video/mp4"}
)

func SaveVideoFile(ctx context.Context, authorID uint64, title string, videoData *multipart.FileHeader) (videoUrl string, err error) {
	filename := filepath.Base(title)
	ext := filepath.Ext(videoData.Filename)
	uuidStr, err := uuid.GenerateUUID()
	if err != nil {
		return "", err
	}
	finalName := fmt.Sprintf("%d_%s_%s%s", authorID, filename, uuidStr, ext)

	// endpointURL := MinIOClient.EndpointURL().String()
	endpointURL := config.Configs.MinioConfig.MinioAccessUrl
	videoFileURL := endpointURL + "/" + VideoBucketName + "/" + finalName
	data, err := videoData.Open()
	if err != nil {
		return "", err
	}
	_, err = MinIOClient.PutObject(ctx, VideoBucketName, finalName, data, videoData.Size, videoOpt)
	if err != nil {
		return "", err
	}

	return videoFileURL, nil
}
