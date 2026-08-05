package us3_storage

// US3 Storage UCloud 对象存储
// UCloud US3 兼容 S3 协议，复用 S3 实现

import (
	"context"

	s3_storage "github.com/storage-lock/go-s3-storage"
	"github.com/storage-lock/go-storage"
)

// US3Storage UCloud US3 对象存储实现
type US3Storage struct {
	*s3_storage.S3Storage
}

var _ storage.Storage = &US3Storage{}

func NewUS3Storage(ctx context.Context, options *s3_storage.S3StorageOptions) (*US3Storage, error) {
	s3Storage, err := s3_storage.NewS3Storage(ctx, options)
	if err != nil {
		return nil, err
	}
	return &US3Storage{S3Storage: s3Storage}, nil
}

const US3StorageName = "us3-storage"

func (x *US3Storage) GetName() string {
	return US3StorageName
}

func (x *US3Storage) Capabilities() []storage.StorageCapability {
	return x.S3Storage.Capabilities()
}