package storages

import (
	"context"
	"io"
	"log"
	"mime/multipart"
	"os"
	"time"

	"cloud.google.com/go/storage"
	"github.com/GroupProject2-Kelompok4/BE/app/config"
	"github.com/labstack/echo/v4"
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
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	wc := s.storageClient.Bucket(s.bucketName).Object(s.path + fileName).NewWriter(ctx)

	if _, err := io.Copy(wc, file); err != nil {
		return "", err
	}

	if err := wc.Close(); err != nil {
		return "", err
	}

	fileURL := "https://storage.googleapis.com/" + s.bucketName + "/" + s.path + fileName
	return fileURL, nil
}
