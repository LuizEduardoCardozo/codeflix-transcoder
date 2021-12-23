package repositories

import (
	"encoder/domain"
	"fmt"

	"gorm.io/gorm"
)

type JobRepository interface {
	Insert(job *domain.Job) (*domain.Job, error)
	FindById(id string) (*domain.Job, error)
	Update(new_job *domain.Job) (*domain.Job, error)
}

type JobRepositoryDB struct {
	DB *gorm.DB
}

func (repo *JobRepositoryDB) Insert(job *domain.Job) (*domain.Job, error) {
	if err := repo.DB.Create(job); err.Error != nil {
		return nil, err.Error
	}
	return job, nil
}

func (repo *JobRepositoryDB) FindById(id string) (*domain.Job, error) {
	var job domain.Job
	repo.DB.Preload("Video").First(&job, "id = ?", id)

	if job.ID == id {
		return &job, nil
	}

	return nil, fmt.Errorf("job with id %s not found", id)
}

func (repo *JobRepositoryDB) Update(new_job *domain.Job) (*domain.Job, error) {

	_, err := repo.FindById(new_job.ID)
	if err != nil {
		return nil, err
	}

	if err := repo.DB.Save(new_job); err.Error != nil {
		return nil, err.Error
	}

	return new_job, nil
}

func NewJobRepository(db *gorm.DB) *JobRepositoryDB {
	return &JobRepositoryDB{
		DB: db,
	}
}
