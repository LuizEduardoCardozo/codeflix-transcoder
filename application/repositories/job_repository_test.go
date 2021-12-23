package repositories_test

import (
	"encoder/application/repositories"
	"encoder/domain"
	"encoder/framework/database"
	"testing"
	"time"

	"github.com/jaswdr/faker"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIfJobRepositoryCanInsertAJob(t *testing.T) {
	fake := faker.New()

	db := database.NewTestDB()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.ResourceID = uuid.NewV4().String()
	video.FilePath = fake.File().FilenameWithExtension()
	video.CreatedAt = time.Now()

	if err := video.Validate(); err != nil {
		t.Errorf("error while validating the video: %s", err.Error())
	}

	job, err := domain.NewJob(fake.File().FilenameWithExtension(), "OnGoing", video)
	if err != nil {
		t.Errorf("error while creating the new job: %s", err.Error())
	}

	repo := repositories.NewJobRepository(db)

	_, err = repo.Insert(job)
	if err != nil {
		t.Errorf("error while inserting the new job: %s", err.Error())
	}

	var foundJob domain.Job
	db.Raw("SELECT * FROM jobs WHERE id = ?", job.ID).Scan(&foundJob)

	require.Equal(t, job.ID, foundJob.ID)
}

func TestIfJobRepositoryDoesNotReturnErrorWhenInsert(t *testing.T) {
	fake := faker.New()

	db := database.NewTestDB()
	video := domain.NewVideo()

	video.ID = uuid.NewV4().String()
	video.FilePath = fake.File().FilenameWithExtension()
	video.CreatedAt = time.Now()

	job, err := domain.NewJob(fake.File().FilenameWithExtension(), "status", video)
	if err != nil {
		t.Errorf("error while creating the new job: %s", err.Error())
	}

	repo := repositories.NewJobRepository(db)

	_, err = repo.Insert(job)

	require.Nil(t, err)
}

func TestIJobRepositoryReturnVideoWhenItInserted(t *testing.T) {
	fake := faker.New()

	db := database.NewTestDB()
	video := domain.NewVideo()

	video.ID = uuid.NewV4().String()
	video.FilePath = fake.File().FilenameWithExtension()
	video.CreatedAt = time.Now()

	job, err := domain.NewJob(fake.File().FilenameWithExtension(), "status", video)
	if err != nil {
		t.Errorf("error while creating the new job: %s", err.Error())
	}

	repo := repositories.NewJobRepository(db)

	returnedJob, err := repo.Insert(job)
	if err != nil {
		t.Errorf("error while inserting the video, %s", err.Error())
	}

	require.Equal(t, job.ID, returnedJob.ID)
}

func TestIfJobRepositoryReturnErrorWhenSearchForAnNonexistentVideo(t *testing.T) {
	db := database.NewTestDB()

	id := uuid.NewV4().String()

	repo := repositories.NewJobRepository(db)

	_, err := repo.FindById(id)
	if err == nil {
		t.Error("error: find by id should return an error")
	}

	require.Error(t, err)
}

func TestIfJobRepositoryReturnNilVideoWhenSearchForAnNonexistentVideo(t *testing.T) {
	db := database.NewTestDB()

	id := uuid.NewV4().String()

	repo := repositories.NewJobRepository(db)

	foundJob, err := repo.FindById(id)
	if err == nil {
		t.Error("error: find by id should return an error")
	}

	require.Nil(t, foundJob)
}

func TestIfJobRepositoryCanFindByIdAnExistentVideo(t *testing.T) {
	fake := faker.New()

	db := database.NewTestDB()
	video := domain.NewVideo()

	video.ID = uuid.NewV4().String()
	video.FilePath = fake.File().FilenameWithExtension()
	video.CreatedAt = time.Now()

	job, err := domain.NewJob(fake.File().FilenameWithExtension(), "status", video)
	if err != nil {
		t.Errorf("error while creating the new job: %s", err.Error())
	}

	if err := db.Create(job); err.Error != nil {
		t.Errorf("error while inserting the video, %s", err.Error)
	}

	repo := repositories.NewJobRepository(db)

	foundJob, err := repo.FindById(job.ID)
	if err != nil {
		t.Errorf("error while searching for the video, %s", err.Error())
	}

	require.Equal(t, job.ID, foundJob.ID)
}

func TestIfJobRepositoryCanPopulateVideoWhenFind(t *testing.T) {
	fake := faker.New()

	db := database.NewTestDB()
	video := domain.NewVideo()

	video.ID = uuid.NewV4().String()
	video.FilePath = fake.File().FilenameWithExtension()
	video.CreatedAt = time.Now()

	job, err := domain.NewJob(fake.File().FilenameWithExtension(), "status", video)
	if err != nil {
		t.Errorf("error while creating the new job: %s", err.Error())
	}

	if err := db.Create(job); err.Error != nil {
		t.Errorf("error while inserting the video, %s", err.Error)
	}

	repo := repositories.NewJobRepository(db)

	foundJob, err := repo.FindById(job.ID)
	if err != nil {
		t.Errorf("error while searching for the video, %s", err.Error())
	}

	require.Equal(t, job.Video.ResourceID, foundJob.Video.ResourceID)
}
func TestUpdateShouldReturnAnErrorIfJobDoesNotExists(t *testing.T) {
	fake := faker.New()

	db := database.NewTestDB()
	video := domain.NewVideo()

	video.ID = uuid.NewV4().String()
	video.FilePath = fake.File().FilenameWithExtension()
	video.CreatedAt = time.Now()

	job, err := domain.NewJob(fake.File().FilenameWithExtension(), "status", video)
	if err != nil {
		t.Errorf("error while creating the new job: %s", err.Error())
	}

	repo := repositories.NewJobRepository(db)

	_, err = repo.Update(job)

	require.Error(t, err)
}

func TestUpdateShouldReturnAnNullJobIfJobDoesNotExists(t *testing.T) {
	fake := faker.New()

	db := database.NewTestDB()
	video := domain.NewVideo()

	video.ID = uuid.NewV4().String()
	video.FilePath = fake.File().FilenameWithExtension()
	video.CreatedAt = time.Now()

	job, err := domain.NewJob(fake.File().FilenameWithExtension(), "status", video)
	if err != nil {
		t.Errorf("error while creating the new job: %s", err.Error())
	}

	repo := repositories.NewJobRepository(db)

	updatedJob, _ := repo.Update(job)

	assert.Nil(t, updatedJob)
}

func TestShouldNotReturnErrorIfJobExistsWhenUpdating(t *testing.T) {
	fake := faker.New()

	db := database.NewTestDB()
	video := domain.NewVideo()

	video.ID = uuid.NewV4().String()
	video.FilePath = fake.File().FilenameWithExtension()
	video.CreatedAt = time.Now()

	job, err := domain.NewJob(fake.File().FilenameWithExtension(), "status", video)
	if err != nil {
		t.Errorf("error while creating the new job: %s", err.Error())
	}

	repo := repositories.NewJobRepository(db)

	_, err = repo.Insert(job)
	if err != nil {
		t.Errorf("error while inserting the video, %s", err.Error())
	}

	job.Status = "altered status"

	_, err = repo.Update(job)

	assert.Nil(t, err)
}

func TestShouldNotReturnAnNilJobIfJobExistsWhenUpdating(t *testing.T) {
	fake := faker.New()

	db := database.NewTestDB()
	video := domain.NewVideo()

	video.ID = uuid.NewV4().String()
	video.FilePath = fake.File().FilenameWithExtension()
	video.CreatedAt = time.Now()

	job, err := domain.NewJob(fake.File().FilenameWithExtension(), "status", video)
	if err != nil {
		t.Errorf("error while creating the new job: %s", err.Error())
	}

	repo := repositories.NewJobRepository(db)

	_, err = repo.Insert(job)
	if err != nil {
		t.Errorf("error while inserting the video, %s", err.Error())
	}

	job.Status = "altered status"

	updatedJob, err := repo.Update(job)
	if err != nil {
		t.Errorf("error while updating the job: %s\n", err.Error())
	}

	assert.NotNil(t, updatedJob)
}

func TestShouldUpdateStatusWhenJobExists(t *testing.T) {
	fake := faker.New()

	db := database.NewTestDB()
	video := domain.NewVideo()

	video.ID = uuid.NewV4().String()
	video.FilePath = fake.File().FilenameWithExtension()
	video.CreatedAt = time.Now()

	job, err := domain.NewJob(fake.File().FilenameWithExtension(), "status", video)
	if err != nil {
		t.Errorf("error while creating the new job: %s", err.Error())
	}

	repo := repositories.NewJobRepository(db)

	_, err = repo.Insert(job)
	if err != nil {
		t.Errorf("error while inserting the video, %s", err.Error())
	}

	job.Status = "altered status"

	if _, err = repo.Update(job); err != nil {
		t.Errorf("error while updating the job: %s\n", err.Error())
	}

	var foundJob domain.Job
	db.First(&foundJob, "id = ?", job.ID)

	assert.Equal(t, "altered status", foundJob.Status)
}
