package storages

import (
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"strings"
	"time"

	"cloud.google.com/go/storage"
	"github.com/GroupProject2-Kelompok4/BE/app/config"
	"github.com/google/uuid"
	"google.golang.org/api/option"
)

type ClientUploader struct {
	cl         *storage.Client
	projectID  string
	bucketName string
	uploadPath string
}

var Uploader *ClientUploader

func init() {
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(config.GCP_CREDENTIAL))
	if err != nil {
		log.Fatal(err)
	}

	Uploader = &ClientUploader{
		cl:         client,
		bucketName: config.GCP_BUCKETNAME,
		projectID:  config.GCP_PROJECTID,
		uploadPath: config.GCP_PATH,
	}

}

// UploadFile uploads an object
func (c *ClientUploader) UploadFile(file multipart.File, object string) (string, error) {
	rand := uuid.New().String()
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	// Upload an object with storage.Writer.
	wc := c.cl.Bucket(c.bucketName).Object(c.uploadPath + object + rand).NewWriter(ctx)
	if _, err := io.Copy(wc, file); err != nil {
		return "", fmt.Errorf("io.Copy: %v", err)
	}
	if err := wc.Close(); err != nil {
		return "", fmt.Errorf("Writer.Close: %v", err)
	}

	escapedObject := strings.ReplaceAll(object, " ", "%20")
	fileURL := "https://storage.googleapis.com/" + c.bucketName + "/" + c.uploadPath + escapedObject + rand
	return fileURL, nil
}
