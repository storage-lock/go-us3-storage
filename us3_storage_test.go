package us3_storage

import (
	"context"
	"testing"

	"github.com/storage-lock/go-storage"
	"github.com/stretchr/testify/assert"
)

func TestNewUS3Storage_NilOptions(t *testing.T) {
	assert.Panics(t, func() {
		NewUS3Storage(context.Background(), nil)
	})
}

func TestUS3Storage_GetName(t *testing.T) {
	s := &US3Storage{}
	name := s.GetName()
	assert.Equal(t, US3StorageName, name)
}

func TestUS3Storage_Capabilities(t *testing.T) {
	s := &US3Storage{}
	caps := s.Capabilities()
	assert.NotNil(t, caps)
	// US3 透传 S3 基类能力：CAS + ReliableTime + Watch。
	// 此前断言 Len==2 是旧发布版 S3 的残留（本地 S3 早已是 3 项），随 replace 补齐而修正。
	assert.Len(t, caps, 3)
	assert.Equal(t, storage.CapabilityCAS, caps[0])
	assert.Equal(t, storage.CapabilityReliableTime, caps[1])
	assert.Equal(t, storage.CapabilityWatch, caps[2])
}
