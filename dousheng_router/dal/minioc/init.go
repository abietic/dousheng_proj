package minioc

import (
	"context"
	"dousheng/common/config"
	"log"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

const (
	// Endpoint        = "127.0.0.1:9000"
	// AccessKeyID     = "minioadmin"
	// SecretAccessKey = "minioadmin"
	// UseSSL          = false
	VideoBucketName = "videos"
	ImageBucketName = "images"
)

var MinIOClient *minio.Client

func Init() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// endpoint := Endpoint
	endpoint := config.Configs.MinioConfig.Endpoint
	// accessKeyID := AccessKeyID
	accessKeyID := config.Configs.MinioConfig.AccessKeyID
	// secretAccessKey := SecretAccessKey
	secretAccessKey := config.Configs.MinioConfig.SecretAccessKey
	// useSSL := UseSSL
	useSSL := config.Configs.MinioConfig.UseSSL
	var err error
	// Initialize minio client object.
	MinIOClient, err = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	// Make two new buckets
	bucketNames := []string{VideoBucketName, ImageBucketName}
	// location := "local"
	location := config.Configs.MinioConfig.Region

	for _, bucketName := range bucketNames {
		err = MinIOClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
		if err != nil {
			// Check to see if we already own this bucket (which happens if you run this twice)
			exists, errBucketExists := MinIOClient.BucketExists(ctx, bucketName)
			if errBucketExists == nil && exists {
				log.Printf("We already own %s\n", bucketName)
			} else {
				log.Fatalln(err)
			}
		} else {
			log.Printf("Bucket Successfully created %s\n", bucketName)
		}
	}
}
