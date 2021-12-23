package repositories_test

import (
	"encoder/application/repositories"
	"encoder/domain"
	"encoder/framework/database"
	"testing"
	"time"

	"github.com/jaswdr/faker"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestIfVideoRepositoryCanInsertAVideo(t *testing.T) {
	fake := faker.New()

	db := database.NewTestDB()
	video := domain.NewVideo()

	video.ID = uuid.NewV4().String()
	video.FilePath = fake.File().FilenameWithExtension()
	video.CreatedAt = time.Now()

	repo := repositories.NewVideoRepository(db)

	_, err := repo.Insert(video)
	if err != nil {
		t.Errorf("error while inserting the video, %s", err.Error())
	}

	var foundVideo domain.Video
	repo.DB.First(&foundVideo, "id = ?", video.ID)

	require.Equal(t, video.ID, foundVideo.ID)
}

func TestIfVideoRepositoryDoesNotReturnErrorWhenInsert(t *testing.T) {
	fake := faker.New()

	db := database.NewTestDB()
	video := domain.NewVideo()

	video.ID = uuid.NewV4().String()
	video.FilePath = fake.File().FilenameWithExtension()
	video.CreatedAt = time.Now()

	repo := repositories.NewVideoRepository(db)

	_, err := repo.Insert(video)

	require.Nil(t, err)
}

func TestIfVideoRepositoryReturnVideoWhenItInserted(t *testing.T) {
	fake := faker.New()

	db := database.NewTestDB()
	video := domain.NewVideo()

	video.ID = uuid.NewV4().String()
	video.FilePath = fake.File().FilenameWithExtension()
	video.CreatedAt = time.Now()

	repo := repositories.NewVideoRepository(db)

	returnedVideo, err := repo.Insert(video)
	if err != nil {
		t.Errorf("error while inserting the video, %s", err.Error())
	}

	require.Equal(t, video.ID, returnedVideo.ID)
}

func TestIfVideoRepositoryReturnErrorWhenSearchForAnNonexistentVideo(t *testing.T) {
	db := database.NewTestDB()

	id := uuid.NewV4().String()

	repo := repositories.NewVideoRepository(db)

	_, err := repo.FindById(id)
	if err == nil {
		t.Error("error: find by id should return an error")
	}

	require.Error(t, err)
}

func TestIfVideoRepositoryReturnNilVideoWhenSearchForAnNonexistentVideo(t *testing.T) {
	db := database.NewTestDB()

	id := uuid.NewV4().String()

	repo := repositories.NewVideoRepository(db)

	foundVideo, err := repo.FindById(id)
	if err == nil {
		t.Error("error: find by id should return an error")
	}

	require.Nil(t, foundVideo)
}

func TestIfVideoRepositoryCanFindByIdAnExistentVideo(t *testing.T) {
	fake := faker.New()

	db := database.NewTestDB()
	video := domain.NewVideo()

	video.ID = uuid.NewV4().String()
	video.FilePath = fake.File().FilenameWithExtension()
	video.CreatedAt = time.Now()

	if err := db.Create(video); err.Error != nil {
		t.Errorf("error while creating a new video into database: %s", err.Error)
	}

	repo := repositories.NewVideoRepository(db)

	foundVideo, err := repo.FindById(video.ID)
	if err != nil {
		t.Errorf("error while searching for the video, %s", err.Error())
	}

	require.Equal(t, video.ID, foundVideo.ID)
}
