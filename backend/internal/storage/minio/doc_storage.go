package minio

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/kerilOvs/backend/internal/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

/*
accessKey: "ErKo2T8FdXxvR4phdFop"
  secretKey: "uV8044JzvasUXq6u62RD6DD3JVSsJq3x4w225AAl"
*/

type Client struct {
	Client       *minio.Client
	Bucket       string
	PublicHost   string // Добавляем поле для публичного хоста
	PublicPrefix string // Публичный префикс
}

func New(ctx context.Context, cfg config.MinioConfig) (*Client, error) {
	client, err := minio.New(cfg.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.AccessKey, cfg.SecretKey, ""),
		Secure: cfg.UseSSL,
	})
	if err != nil {
		return nil, fmt.Errorf("minio init error: %w", err)
	}

	exists, err := client.BucketExists(ctx, cfg.Bucket)
	if err != nil {
		return nil, fmt.Errorf("bucket check error: %w", err)
	}

	if !exists {
		if err = client.MakeBucket(ctx, cfg.Bucket, minio.MakeBucketOptions{}); err != nil {
			return nil, fmt.Errorf("bucket creation error: %w", err)
		}
	}

	return &Client{
		Client:       client,
		Bucket:       cfg.Bucket,
		PublicHost:   cfg.PublicHost,
		PublicPrefix: "pub",
	}, nil
}

func (c *Client) PutObject(ctx context.Context, objectName string, reader io.Reader, objectSize int64, opts string) (minio.UploadInfo, error) {
	return c.Client.PutObject(ctx, c.Bucket, objectName, reader, objectSize, minio.PutObjectOptions{ContentType: opts})
}

func (c *Client) GetObject(ctx context.Context, objectName string) (string, error) {
	if !strings.HasPrefix(objectName, c.PublicPrefix) {
		return "hui", fmt.Errorf("failed to get document URL")
	}

	return fmt.Sprintf("%s/%s/%s", c.PublicHost, c.Bucket, objectName), nil
}
