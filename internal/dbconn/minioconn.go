package dbconn

import (
	"context"
	"errors"
	"fmt"
	"gameplatform/internal/config"
	"mime/multipart"

	"log/slog"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var (
	S3ErrorIncorrectFormat = errors.New("incorrect format")
)

type MinioConnection struct {
	MinioClient *minio.Client
	appBucket   string
}

func NewMinioConnection(config *config.Config) MinioConnection {
	endpoint := config.MinioHost
	accessKeyID := config.MinioAccessKey
	secretAccessKey := config.MinioSecretKey

	var err error
	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: config.MinioSecure,
	})
	if err != nil {
		panic(err.Error())
	}

	slog.Info("✅ MinioClient client connected successfully...")

	return MinioConnection{
		MinioClient: minioClient,
		appBucket:   config.AppBucket,
	}
}

func (m *MinioConnection) PutObject(dstPath string, fileReader multipart.File) (path string, err error) {
	ctx := context.Background()

	info, err := m.MinioClient.PutObject(
		ctx, m.appBucket,
		dstPath,
		fileReader,
		-1, minio.PutObjectOptions{})

	return info.Location, err
}

func (m *MinioConnection) RemoveObject(path string) error {
	ctx := context.Background()
	err := m.MinioClient.RemoveObject(ctx, m.appBucket, path, minio.RemoveObjectOptions{})
	return err
}

func (m *MinioConnection) RemoveFolder(path string) {
	ctx := context.Background()

	objectsCh := make(chan minio.ObjectInfo)

	go func() {
		defer close(objectsCh)
		opts := minio.ListObjectsOptions{
			Prefix:    path,
			Recursive: true,
		}

		for object := range m.MinioClient.ListObjects(ctx, m.appBucket, opts) {
			if object.Err != nil {
				slog.Error(object.Err.Error())
				continue
			}
			objectsCh <- object
		}
	}()

	errorCh := m.MinioClient.RemoveObjects(ctx, m.appBucket, objectsCh, minio.RemoveObjectsOptions{})

	for e := range errorCh {
		slog.Error(fmt.Sprintf("Failed to remove %s, error: %v", e.ObjectName, e.Err))
	}
}
