package services_test

import (
	"encoder/application/repositories"
	"encoder/application/services"
	"encoder/domain"
	"encoder/framework/database"
	"os"
	"testing"
	"time"

	"github.com/fsouza/fake-gcs-server/fakestorage"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestVideoServiceDownload(t *testing.T) {
	os.Setenv("LOCAL_STORAGE_PATH", "/tmp")

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
			Content: []byte("inside the file"),
		},
	})
	defer storageServer.Stop()

	storageClient := storageServer.Client()
	db := database.NewTestDB()
	videoRepository := repositories.NewVideoRepository(db)
	videoService := services.NewVideoService(video, videoRepository, storageClient)

	err := videoService.Download("bucket_teste")

	assert.Nil(t, err)

}
