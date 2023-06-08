package storages

import (
	"context"
	"io"
	"log"
	"mime/multipart"
	"strings"
	"time"

	"cloud.google.com/go/storage"
	"github.com/GroupProject2-Kelompok4/BE/app/config"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/option"
)

type StorageGCPConfig struct {
	GCPClient  *storage.Client
	ProjectID  string
	BucketName string
	Path       string
}

func InitGCPClient() *storage.Client {
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(config.GCP_CREDENTIAL))
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func UploadImage(c echo.Context, file *multipart.FileHeader) (string, error) {
	if file == nil {
		return "", nil
	}
	image, err := file.Open()
	if err != nil {
		return "", err
	}
	defer image.Close()
	sgcp := StorageGCPConfig{
		GCPClient:  InitGCPClient(),
		ProjectID:  config.GCP_PROJECTID,
		BucketName: config.GCP_BUCKETNAME,
		Path:       config.GCP_PATH,
	}

	imageURL, err := sgcp.UploadFile(image, file.Filename)
	if err != nil {
		return "", err
	}
	return imageURL, nil
}

func (s *StorageGCPConfig) UploadFile(file io.Reader, fileName string) (string, error) {
	rand := uuid.New().String()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	wc := s.GCPClient.Bucket(s.BucketName).Object(s.Path + fileName).NewWriter(ctx)
	if _, err := io.Copy(wc, file); err != nil {
		return "", err
	}
	if err := wc.Close(); err != nil {
		return "", err
	}

	escapedFileName := strings.ReplaceAll(fileName, " ", "%20")
	fileURL := "https://storage.googleapis.com/" + s.BucketName + "/" + s.Path + escapedFileName + rand
	return fileURL, nil
}
