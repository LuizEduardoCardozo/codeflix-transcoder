package services_test

import (
	"encoder/application/repositories"
	"encoder/application/services"
	"encoder/domain"
	"encoder/framework/database"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/fsouza/fake-gcs-server/fakestorage"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestVideoServiceDownload(t *testing.T) {
	os.Setenv("LOCAL_STORAGE_PATH", "/tmp")
	os.Setenv("VIDEO_MP4_TEST_LOCATION", "../../test/assets/video.mp4")

	stubVideoPath := os.Getenv("VIDEO_MP4_TEST_LOCATION")
	stubVideo, err := os.Open(stubVideoPath)
	if err != nil {
		t.Errorf("error while opening stub video: %s\n", err.Error())
		assert.Nil(t, err)
	}
	stubVideoContent, err := ioutil.ReadAll(stubVideo)
	if err != nil {
		t.Errorf("error while reading stub video bytes: %s\n", err.Error())
		assert.Nil(t, err)
	}

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.ResourceID = uuid.NewV4().String()
	video.FilePath = "video.mp4"
	video.CreatedAt = time.Now()

	storageServer := fakestorage.NewServer([]fakestorage.Object{
		{
			ObjectAttrs: fakestorage.ObjectAttrs{
				BucketName: "bucket_teste",
				Name:       video.FilePath,
			},
			Content: stubVideoContent,
		},
	})
	defer storageServer.Stop()

	storageClient := storageServer.Client()
	db := database.NewTestDB()
	videoRepository := repositories.NewVideoRepository(db)
	videoService := services.NewVideoService(video, videoRepository, storageClient)

	savedVideoFilePath, err := videoService.Download("bucket_teste")
	assert.Nil(t, err)

	_, err = os.Open(savedVideoFilePath)
	assert.Nil(t, err)

}

func TestVideoFragment(t *testing.T) {
	os.Setenv("LOCAL_STORAGE_PATH", "/tmp")
	os.Setenv("VIDEO_MP4_TEST_LOCATION", "../../test/assets/video.mp4")

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.ResourceID = uuid.NewV4().String()
	video.FilePath = "video.mp4"
	video.CreatedAt = time.Now()

	stubVideoPath := os.Getenv("VIDEO_MP4_TEST_LOCATION")

	localVideoPath := os.Getenv("LOCAL_STORAGE_PATH")
	sourcePath := fmt.Sprintf("%s/%s.mp4", localVideoPath, video.ID)

	cmd := exec.Command("cp", stubVideoPath, sourcePath)
	_, err := cmd.CombinedOutput()
	if err != nil {
		assert.Nil(t, err)
		t.Error(err)
	}

	db := database.NewTestDB()
	videoRepository := repositories.NewVideoRepository(db)
	videoService := services.NewVideoService(video, videoRepository, nil)

	err = videoService.Fragment()
	if err != nil {
		assert.Nil(t, err)
		t.Error(err.Error())
	}
}

func TestDownloadAndFragmentVideo(t *testing.T) {
	os.Setenv("LOCAL_STORAGE_PATH", "/tmp")
	os.Setenv("VIDEO_MP4_TEST_LOCATION", "../../test/assets/video.mp4")

	stubVideoPath := os.Getenv("VIDEO_MP4_TEST_LOCATION")
	stubVideo, err := os.Open(stubVideoPath)
	if err != nil {
		t.Errorf("error while opening stub video: %s\n", err.Error())
		assert.Nil(t, err)
	}
	stubVideoContent, err := ioutil.ReadAll(stubVideo)
	if err != nil {
		t.Errorf("error while reading stub video bytes: %s\n", err.Error())
		assert.Nil(t, err)
	}

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.ResourceID = uuid.NewV4().String()
	video.FilePath = "video.mp4"
	video.CreatedAt = time.Now()

	storageServer := fakestorage.NewServer([]fakestorage.Object{
		{
			ObjectAttrs: fakestorage.ObjectAttrs{
				BucketName: "bucket_teste",
				Name:       video.FilePath,
			},
			Content: stubVideoContent,
		},
	})
	defer storageServer.Stop()

	storageClient := storageServer.Client()
	db := database.NewTestDB()
	videoRepository := repositories.NewVideoRepository(db)
	videoService := services.NewVideoService(video, videoRepository, storageClient)

	savedVideoFilePath, err := videoService.Download("bucket_teste")
	assert.Nil(t, err)

	_, err = os.Open(savedVideoFilePath)
	assert.Nil(t, err)

	localVideoPath := os.Getenv("LOCAL_STORAGE_PATH")
	sourcePath := fmt.Sprintf("%s/%s.mp4", localVideoPath, video.ID)

	cmd := exec.Command("cp", stubVideoPath, sourcePath)
	_, err = cmd.CombinedOutput()
	assert.Nil(t, err)

	err = videoService.Fragment()
	assert.Nil(t, err)
}
