package service

import (
	"context"
	"fmt"
	"io"

	"github.com/google/uuid"
	"github.com/kerilOvs/backend/internal/storage/minio"
	//"github.com/minio/minio-go/v7"
)

type DocService struct {
	minioClient *minio.Client
	bucket      string
}

func NewDocService(client *minio.Client, bucket string) *DocService {
	return &DocService{
		minioClient: client,
		bucket:      bucket,
	}
}

func (s *DocService) UploadDoc(ctx context.Context, file io.Reader, size int64, ext string) (string, error) {
	// Генерируем уникальное имя файла с правильным расширением
	objectName := s.minioClient.PublicPrefix + "/" + uuid.New().String() + ext

	opts := "application/vnd.openxmlformats-officedocument.wordprocessingml.document"

	_, err := s.minioClient.PutObject(
		ctx,
		objectName,
		file,
		size,
		opts,
	)

	if err != nil {
		return "", fmt.Errorf("failed to upload document to MinIO: %w", err)
	}

	return objectName, nil
}

func (s *DocService) GetDocURL(ctx context.Context, objectName string) (string, error) {
	return s.minioClient.GetObject(ctx, objectName)
}
