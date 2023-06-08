package storages

import (
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"strings"
	"time"

	"cloud.google.com/go/storage"
	"github.com/google/uuid"
)

const (
	projectID  = "be-alta-17" // GANTI DENGAN PROJECT ID ANDA
	bucketName = "immersive"  // GANTI DENGAN NAMA BUCKET ANDA
)

type ClientUploader struct {
	cl         *storage.Client
	projectID  string
	bucketName string
	uploadPath string
}

var Uploader *ClientUploader

func init() {
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "./credential.json") // GANTI DENGAN JALUR FILE KREDENSIAL JSON ANDA
	client, err := storage.NewClient(context.Background())
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	Uploader = &ClientUploader{
		cl:         client,
		bucketName: bucketName,
		projectID:  projectID,
		uploadPath: "latihan/",
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
