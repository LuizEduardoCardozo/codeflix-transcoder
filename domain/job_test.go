package domain_test

import (
	"encoder/domain"
	"testing"
	"time"

	"github.com/jaswdr/faker"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func stubVideo() (*domain.Video, error) {
	fake := faker.New()
	video := domain.NewVideo()

	video.ID = uuid.NewV4().String()
	video.ResourceID = uuid.NewV4().String()
	video.FilePath = fake.File().FilenameWithExtension()
	video.CreatedAt = time.Now()

	err := video.Validate()
	if err != nil {
		return nil, err
	}

	return video, nil
}

func TestIfJobIdIsNotAnUuid(t *testing.T) {

	fake := faker.New()

	video, err := stubVideo()
	if err != nil {
		t.Errorf("fail when creating the video: %s\n", err)
	}

	job, err := domain.NewJob(fake.File().FilenameWithExtension(), "status", video)
	if err != nil {
		t.Errorf("fail when creating the job: %s\n", err)
	}

	job.ID = "ola"

	err = job.Validate()

	require.Error(t, err)
}

func TestIfOutputBucketPathIsNotNull(t *testing.T) {

	fake := faker.New()

	video, err := stubVideo()
	if err != nil {
		t.Errorf("fail when creating the video: %s\n", err)
	}

	job, err := domain.NewJob(fake.File().FilenameWithExtension(), "status", video)
	if err != nil {
		t.Errorf("fail when creating the job: %s\n", err)
	}

	job.OutputBucketPath = ""

	err = job.Validate()

	require.Error(t, err)
}

func TestIfStatusIsNotNull(t *testing.T) {

	fake := faker.New()

	video, err := stubVideo()
	if err != nil {
		t.Errorf("fail when creating the video: %s\n", err)
	}

	job, err := domain.NewJob(fake.File().FilenameWithExtension(), "status", video)
	if err != nil {
		t.Errorf("fail when creating the job: %s\n", err)
	}

	job.Status = ""

	err = job.Validate()

	require.Error(t, err)
}
