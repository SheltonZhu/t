package storage

import (
	"io"
	"os"
	"path"
	"github.com/pkg/errors"
	"path/filepath"
)


type localDiskFileStorage struct {
	basePath string
}

// NewLocalDiskFileGetSaveCleaner
func NewLocalDiskFileGetSaveCleaner(basePath string) *localDiskFileStorage {
	return &localDiskFileStorage{basePath: basePath}
}

// IsDirExists 判断文件夹是否存在
func isDirExists(dirPath string) bool {
	info, err := os.Stat(dirPath)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return info.IsDir()
}

// SaveFile 实现了 FileSaver 接口的 SaveFile 方法
func (s *localDiskFileStorage) SaveFile(file io.Reader, filePath string) (string, error) {
	filePath = filepath.Join(s.basePath, filePath)
	fileDir := path.Dir(filePath)
	if !isDirExists(fileDir) {
		if err := os.MkdirAll(fileDir, 0755); err != nil {
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
