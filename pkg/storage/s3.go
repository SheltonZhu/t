package storage

import (
	"io"

	"github.com/minio/minio-go"
	"github.com/pkg/errors"
)

type S3Config struct {
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
	BucketName      string
	UseSSL          bool
}

type s3FileStorage struct {
	S3Config
	client               *minio.Client
	GetObjectOptionsFunc func(storage *s3FileStorage, filePath string) minio.GetObjectOptions
	PutObjectOptionsFunc func(storage *s3FileStorage, file io.Reader, filePath string) minio.PutObjectOptions
}

type S3FileStorageOption func(*s3FileStorage)

// S3Debug 开启s3Client的debug功能
func S3Debug(w io.Writer) S3FileStorageOption {
	return func(s *s3FileStorage) {
		s.client.TraceOn(w)
	}
}

// NewS3FileGetSaveCleaner create a new s3 file storage for the given path.
func NewS3FileGetSaveCleaner(config S3Config, opts ...S3FileStorageOption) (*s3FileStorage, error) {
	minioClient, err := minio.New(config.Endpoint, config.AccessKeyID, config.SecretAccessKey, config.UseSSL)
	if err != nil {
		return nil, err
	}

	s3 := &s3FileStorage{
		S3Config: config,
		client: minioClient,
		GetObjectOptionsFunc: func(storage *s3FileStorage, filePath string) minio.GetObjectOptions {
			return minio.GetObjectOptions{}
		},
		PutObjectOptionsFunc: func(storage *s3FileStorage, file io.Reader, filePath string) minio.PutObjectOptions {
			return minio.PutObjectOptions{ContentType: "application/octet-stream"}
		},
	}

	for _, opt := range opts {
		opt(s3)
	}

	return s3, nil
}

// SaveFile 实现了 FileSaver 接口的 SaveFile 方法
func (s *s3FileStorage) SaveFile(file io.Reader, filePath string) (string, error) {
	opts := s.PutObjectOptionsFunc(s, file, filePath)
	if _, err := s.client.PutObject(s.BucketName, filePath, file, -1, opts); err != nil {
		return "", errors.Wrap(err, "failed to save file")
	}

	return filePath, nil
}

// GetFile 实现了 FileGetter 接口的 GetFile 方法
func (s *s3FileStorage) GetFile(filePath string) (io.ReadCloser, error) {
	opts := s.GetObjectOptionsFunc(s, filePath)
	object, err := s.client.GetObject(s.BucketName, filePath, opts)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get file")
	}

	return object, nil
}

// CleanFile 实现了 FileCleaner 接口的 CleanFile 方法
func (s *s3FileStorage) CleanFile(filePath string) error {
	if err := s.client.RemoveObject(s.BucketName, filePath); err != nil {
		return errors.Wrap(err, "failed to clean file")
	}
	return nil
}

var _ FileGetSaveCleaner = (*s3FileStorage)(nil)
