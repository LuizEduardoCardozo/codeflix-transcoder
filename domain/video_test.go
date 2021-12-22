package domain_test

import (
	"encoder/domain"
	"fmt"
	"time"

	"testing"

	"github.com/jaswdr/faker"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidateIfVideoIsEmpty(t *testing.T) {
	video := domain.Video{}
	err := video.Validate()

	require.Error(t, err)
}

func TestIfVideoIdIsNotAnUuid(t *testing.T) {
	fake := faker.New()

	video := domain.Video{
		ID:        fmt.Sprintf("%d", fake.Int()),
		FilePath:  fake.File().FilenameWithExtension(),
		CreatedAt: time.Now(),
	}

	err := video.Validate()

	assert.Error(t, err)
}

func TestValidateIfVideoResourceIdIsNotNull(t *testing.T) {
	fake := faker.New()

	video := domain.Video{
		ID:        uuid.NewV4().String(),
		FilePath:  fake.File().FilenameWithExtension(),
		CreatedAt: time.Now(),
	}

	err := video.Validate()

	assert.Error(t, err)
}

func TestValidateIfFilePathIsNotNull(t *testing.T) {
	video := domain.Video{
		ID:         uuid.NewV4().String(),
		ResourceID: uuid.NewV4().String(),
		CreatedAt:  time.Now(),
	}

	err := video.Validate()

	assert.Error(t, err)
}

func TestVideoValidation(t *testing.T) {
	fake := faker.New()

	video := domain.Video{
		ID:         uuid.NewV4().String(),
		ResourceID: uuid.NewV4().String(),
		FilePath:   fake.File().FilenameWithExtension(),
		CreatedAt:  time.Now(),
	}

	err := video.Validate()

	assert.Nil(t, err)
}
