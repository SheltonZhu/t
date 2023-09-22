package storage

import (
	"io"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

type localDiskFileStorage struct {
	basePath string
}

// NewLocalDiskFileGetSaveCleaner creates a new local disk file storage for the given path.
func NewLocalDiskFileGetSaveCleaner(basePath string) *localDiskFileStorage {
	return &localDiskFileStorage{basePath: basePath}
}

// IsDirExists 判断文件夹是否存在
func isDirExists(dirPath string) (bool, error) {
	info, err := os.Stat(dirPath)
	if err == nil {
		return info.IsDir(), nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// SaveFile 实现了 FileSaver 接口的 SaveFile 方法
func (s *localDiskFileStorage) SaveFile(file io.Reader, filePath string) (string, error) {
	if filePath == "" {
		return "", errors.New("file path is empty")
	}
	filePath = filepath.Join(s.basePath, filePath)
	fileDir := filepath.Dir(filePath)
	dirExists, err := isDirExists(fileDir)
	if err != nil {
		return "", errors.Wrap(err, "failed to save file")
	}

	if !dirExists {
		if err := os.MkdirAll(fileDir, 0o755); err != nil {
			return "", errors.Wrap(err, "failed to save file")
		}
	}
	outputFile, err := os.Create(filePath)
	if err != nil {
		return "", errors.Wrap(err, "failed to save file")
	}
	defer outputFile.Close()

	_, err = io.Copy(outputFile, file)
	if err != nil {
		return "", errors.Wrap(err, "failed to save file")
	}

	return filePath, nil
}

// GetFile 实现了 FileGetter 接口的 GetFile 方法
func (s *localDiskFileStorage) GetFile(filePath string) (io.ReadCloser, error) {
	if filePath == "" {
		return nil, errors.New("file path is empty")
	}
	filePath = filepath.Join(s.basePath, filePath)
	file, err := os.Open(filePath)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get file")
	}

	return file, nil
}

// CleanFile 实现了 FileCleaner 接口的 CleanFile 方法
func (s *localDiskFileStorage) CleanFile(filePath string) error {
	filePath = filepath.Join(s.basePath, filePath)
	err := os.Remove(filePath)
	if err != nil {
		return errors.Wrap(err, "failed to clean file")
	}
	return nil
}

var _ FileGetSaveCleaner = (*localDiskFileStorage)(nil)
