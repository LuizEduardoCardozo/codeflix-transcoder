package repositories

import (
	"encoder/domain"
	"fmt"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type VideoRepository interface {
	Insert(video *domain.Video) (*domain.Video, error)
	FindById(id string) (*domain.Video, error)
}

type VideoRepositoryDb struct {
	DB *gorm.DB
}

func (repo *VideoRepositoryDb) Insert(video *domain.Video) (*domain.Video, error) {
	if video.ID == "" {
		video.ID = uuid.NewV4().String()
	}

	if err := repo.DB.Create(video); err.Error != nil {
		return nil, err.Error
	}

	return video, nil
}

func (repo *VideoRepositoryDb) FindById(id string) (*domain.Video, error) {
	var video domain.Video
	repo.DB.First(&video, "id = ?", id)

	if video.ID == id {
		return &video, nil
	}

	return nil, fmt.Errorf("video with id %s not found", id)
}

func NewVideoRepository(db *gorm.DB) *VideoRepositoryDb {
	return &VideoRepositoryDb{
		DB: db,
	}
}
