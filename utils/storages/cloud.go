package storages

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"cloud.google.com/go/storage"
	"github.com/GroupProject2-Kelompok4/BE/app/config"
	"github.com/google/uuid"
)

type ClientUploader struct {
	storageClient *storage.Client
	projectID     string
	bucketName    string
	path          string
}

func InitGCPClient() *storage.Client {
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", config.GCP_CREDENTIAL)
	client, err := storage.NewClient(context.Background())
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	image, err := file.Open()
	if err != nil {
		return "", err
	}
	defer image.Close()

	sgcp := ClientUploader{
		storageClient: InitGCPClient(),
		projectID:     config.GCP_PROJECTID,
		bucketName:    config.GCP_BUCKETNAME,
		path:          config.GCP_PATH,
	}

	imageURL, err := sgcp.UploadFile(image, file.Filename)
	if err != nil {
		return "", err
	}

	return imageURL, nil
}

func (s *ClientUploader) UploadFile(file io.Reader, fileName string) (string, error) {
	rand := uuid.New().String()
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	wc := s.storageClient.Bucket(s.bucketName).Object(s.path + fileName + rand).NewWriter(ctx)

	if _, err := io.Copy(wc, file); err != nil {
		return "", fmt.Errorf("io.Copy: %v", err)
	}
	if err := wc.Close(); err != nil {
		return "", fmt.Errorf("Writer.Close: %v", err)
	}

	escapedFileName := strings.ReplaceAll(fileName, " ", "%20")
	fileURL := "https://storage.googleapis.com/" + s.bucketName + "/" + s.path + escapedFileName + rand
	return fileURL, nil
}
