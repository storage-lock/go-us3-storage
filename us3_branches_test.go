package us3_storage

import (
	"context"
	"testing"

	s3_storage "github.com/storage-lock/go-s3-storage"
	"github.com/stretchr/testify/assert"
)

func TestNewUS3Storage_InvalidOptions(t *testing.T) {
	_, err := NewUS3Storage(context.Background(), &s3_storage.S3StorageOptions{})
	assert.Error(t, err)
}

func TestNewUS3Storage_Success(t *testing.T) {
	old := newS3
	newS3 = func(context.Context, *s3_storage.S3StorageOptions) (*s3_storage.S3Storage, error) {
		return &s3_storage.S3Storage{}, nil
	}
	t.Cleanup(func() { newS3 = old })
	s, err := NewUS3Storage(context.Background(), &s3_storage.S3StorageOptions{})
	assert.NoError(t, err)
	assert.NotNil(t, s.S3Storage)
}

func TestUS3Storage_Watch_EmptyLockID(t *testing.T) {
	_, err := (&US3Storage{}).Watch(context.Background(), "")
	assert.Error(t, err)
}
