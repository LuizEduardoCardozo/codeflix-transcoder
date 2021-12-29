package services

import (
	"context"
	"encoder/application/repositories"
	"encoder/domain"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"cloud.google.com/go/storage"
)

type VideoService struct {
	Video           *domain.Video
	VideoRepository repositories.VideoRepository
	StorageClient   *storage.Client
}

func (v *VideoService) Download(bucketName string) (string, error) {
	reader, err := v.StorageClient.Bucket(bucketName).Object(v.Video.FilePath).NewReader(context.TODO())
	if err != nil {
		return "", err
	}
	defer reader.Close()

	fileContent, err := ioutil.ReadAll(reader)
	if err != nil {
		return "", err
	}

	localVideoPath := os.Getenv("LOCAL_STORAGE_PATH")
	filePath := fmt.Sprintf("%s/%s.mp4", localVideoPath, v.Video.ID)

	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}

	_, err = file.Write(fileContent)
	if err != nil {
		return "", err
	}
	defer file.Close()

	return filePath, err
}

func (v *VideoService) Fragment() error {
	localVideoPath := os.Getenv("LOCAL_STORAGE_PATH")
	targetDirPath := fmt.Sprintf("%s/%s", localVideoPath, v.Video.ID)

	err := os.Mkdir(targetDirPath, os.ModePerm)
	if err != nil {
		return err
	}

	sourcePath := fmt.Sprintf("%s/%s.mp4", localVideoPath, v.Video.ID)
	targetPath := fmt.Sprintf("%s/%s.frag", targetDirPath, v.Video.ID)

	cmd := exec.Command("mp4fragment", sourcePath, targetPath)
	_, err = cmd.CombinedOutput()
	if err != nil {
		return err
	}

	return nil
}

func (v *VideoService) Encode() error {
	localVideoPath := os.Getenv("LOCAL_STORAGE_PATH")

	targetDirPath := fmt.Sprintf("%s/%s", localVideoPath, v.Video.ID)
	fragmentedVideoPath := fmt.Sprintf("%s/%s.frag", targetDirPath, v.Video.ID)

	cmdArgs := []string{}
	cmdArgs = append(cmdArgs, fragmentedVideoPath)
	cmdArgs = append(cmdArgs, "-f")
	cmdArgs = append(cmdArgs, "--use-segment-timeline")
	cmdArgs = append(cmdArgs, "-o")
	cmdArgs = append(cmdArgs, targetDirPath)

	cmd := exec.CommandContext(context.TODO(), "mp4dash", cmdArgs...)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("output:", string(out))
		return err
	}

	return nil
}

func NewVideoService(video *domain.Video, videoRepository repositories.VideoRepository, storageClient *storage.Client) *VideoService {
	return &VideoService{
		Video:           video,
		VideoRepository: videoRepository,
		StorageClient:   storageClient,
	}
}
