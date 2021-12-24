package services

import (
	"context"
	"encoder/application/repositories"
	"encoder/domain"
	"fmt"
	"io/ioutil"
	"os"

	"cloud.google.com/go/storage"
)

type VideoService struct {
	Video           *domain.Video
	VideoRepository repositories.VideoRepository
	StorageClient   *storage.Client
}

func (v *VideoService) Download(bucketName string) error {
	reader, err := v.StorageClient.Bucket(bucketName).Object(v.Video.FilePath).NewReader(context.TODO())
	if err != nil {
		return err
	}
	defer reader.Close()

	fileContent, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}

	localPath := os.Getenv("LOCAL_STORAGE_PATH")
	fileName := fmt.Sprintf("%s/%s.mp4", localPath, v.Video.ID)

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}

	_, err = file.Write(fileContent)
	if err != nil {
		return err
	}
	defer file.Close()

	return nil
}

func NewVideoService(video *domain.Video, videoRepository repositories.VideoRepository, storageClient *storage.Client) *VideoService {
	return &VideoService{
		Video:           video,
		VideoRepository: videoRepository,
		StorageClient:   storageClient,
	}
}
