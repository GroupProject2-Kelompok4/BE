package storages

import (
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strings"
	"time"

	"cloud.google.com/go/storage"
	"github.com/GroupProject2-Kelompok4/BE/app/config"
)

type StorageGCP struct {
	ClG        *storage.Client
	ProjectID  string
	BucketName string
	Path       string
}

func NewStorage() (*StorageGCP, error) {
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", config.GCP_CREDENTIAL)
	client, err := storage.NewClient(context.Background())
	if err != nil {
		return nil, err
	}
	return &StorageGCP{
		ClG:        client,
		ProjectID:  config.GCP_PROJECTID,
		BucketName: config.GCP_BUCKETNAME,
		Path:       config.GCP_PATH,
	}, nil
}

func (s *StorageGCP) UploadFile(file multipart.File, fileName string) error {
	if !strings.Contains(strings.ToLower(fileName), ".jpg") && !strings.Contains(strings.ToLower(fileName), ".png") && !strings.Contains(strings.ToLower(fileName), ".jpeg") {
		fmt.Println(strings.Contains(strings.ToLower(fileName), ".jpg"))
		return errors.New("file type not allowed")
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	wc := s.ClG.Bucket(s.BucketName).Object(s.Path + fileName).NewWriter(ctx)
	if _, err := io.Copy(wc, file); err != nil {
		return err
	}

	if err := wc.Close(); err != nil {
		return err
	}
	return nil
}
