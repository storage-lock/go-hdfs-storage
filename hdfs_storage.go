package hdfs_storage

import (
	"context"

	s3_storage "github.com/storage-lock/go-s3-storage"
	"github.com/storage-lock/go-storage"
)

// HdfsStorage hdfs对象存储实现 (S3兼容)
type HdfsStorage struct {
	*s3_storage.S3Storage
}

func NewHdfsStorage(ctx context.Context, options *s3_storage.S3StorageOptions) (*HdfsStorage, error) {
	s3Storage, err := s3_storage.NewS3Storage(ctx, options)
	if err != nil {
		return nil, err
	}
	return &HdfsStorage{S3Storage: s3Storage}, nil
}

const StorageName = "hdfs-storage"

func (x *HdfsStorage) GetName() string {
	return StorageName
}

func (x *HdfsStorage) Capabilities() []storage.StorageCapability {
	return x.S3Storage.Capabilities()
}
