package storage

import (
	"bytes"
	"context"
	"io"

	"github.com/go-kratos/kratos/pkg/sync/errgroup"
	"github.com/pkg/errors"
)

type FileSaver interface {
	SaveFile(file io.Reader, filePath string) (string, error)
}

type FileGetter interface {
	GetFile(filePath string) (io.ReadCloser, error)
}

type FileCleaner interface {
	CleanFile(filePath string) error
}

type FileStorageStreamHandler interface {
	FileSaver
	FileGetter
}

type FileGetSaveCleaner interface {
	FileStorageStreamHandler
	FileCleaner
}

type fileStorageBytesHandler interface {
	SaveFileBytes(fileBytes []byte, filePath string) (string, error)
	GetFileBytes(filePath string) ([]byte, error)
}

type batchHandler interface {
	BatchSaveFiles(files map[string]io.Reader) (map[string]string, []error)
	BatchGetFiles(filePaths []string) ([]io.ReadCloser, []error)
	BatchCleanFiles(filePaths []string) []error
	ConcurrentBatchSaveFiles(files map[string]io.Reader) (map[string]string, error)
	ConcurrentBatchGetFiles(filePaths []string) ([]io.ReadCloser, error)
	ConcurrentBatchCleanFiles(filePaths []string) error
}

// FileStorage 定义文件仓储接口
// 实现一个新的存储只需要实现FileGetSaveCleaner接口
//
//go:generate mockgen -source=storage.go -destination=../../mocks/storage/mock_storage.go -package=mock_storage
type IFileStorage interface {
	FileGetSaveCleaner
	fileStorageBytesHandler
	batchHandler
}

type FileStorage struct {
	FileGetSaveCleaner
	ConcurrencyLimit uint
	afterResponse    []func(*FileStorage, *[]byte)
}

// SaveFileBytes 通过传入bytes保存文件
func (fs *FileStorage) SaveFileBytes(fileBytes []byte, filePath string) (string, error) {
	reader := bytes.NewReader(fileBytes)
	filePath, err := fs.SaveFile(reader, filePath)
	if err != nil {
		return "", errors.Wrap(err, "failed to save file bytes")
	}

	return filePath, nil
}

// GetFileBytes 获取保存文件bytes
func (fs *FileStorage) GetFileBytes(filePath string) ([]byte, error) {
	readCloser, err := fs.GetFile(filePath)
	defer func() {
		_ = readCloser.Close()
	}()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get file bytes")
	}
	fileBytes, err := io.ReadAll(readCloser)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get file bytes")
	}
	for _, h := range fs.afterResponse {
		h(fs, &fileBytes)
	}
	return fileBytes, nil
}

// BatchSaveFiles 批量保存文件
func (fs *FileStorage) BatchSaveFiles(files map[string]io.Reader) (map[string]string, []error) {
	errs := make([]error, 0)
	filePaths := make(map[string]string, len(files))
	for filePath, file := range files {
		var (
			fp  string
			err error
		)
		if fp, err = fs.SaveFile(file, filePath); err != nil {
			errs = append(errs, errors.Wrapf(err, "failed to save file, path: %s", filePath))
		}
		filePaths[filePath] = fp
	}
	return filePaths, errs
}

// BatchGetFiles 批量获取文件
func (fs *FileStorage) BatchGetFiles(filePaths []string) ([]io.ReadCloser, []error) {
	fileReadClosers := make([]io.ReadCloser, len(filePaths))
	errs := make([]error, 0)
	for idx, filePath := range filePaths {
		fileReadCloser, err := fs.GetFile(filePath)
		if err != nil {
			errs = append(errs, errors.Wrapf(err, "failed to get file, path: %s", filePath))
		}
		fileReadClosers[idx] = fileReadCloser
	}
	return fileReadClosers, errs
}

// BatchCleanFiles 批量清理文件
func (fs *FileStorage) BatchCleanFiles(filePaths []string) []error {
	errs := make([]error, 0)
	for _, filePath := range filePaths {
		if err := fs.CleanFile(filePath); err != nil {
			errs = append(errs, errors.Wrapf(err, "failed to clean file, path: %s", filePath))
		}
	}
	return errs
}

// SetResponseHandle 设置响应体hook函数
func (fs *FileStorage) OnAfterResponse(f func(*FileStorage, *[]byte)) *FileStorage {
	fs.afterResponse = append(fs.afterResponse, f)
	return fs
}

// SetConcurrencyLimit 设置并发数
func (fs *FileStorage) SetConcurrencyLimit(limit uint) *FileStorage {
	fs.ConcurrencyLimit = limit
	return fs
}

// ConcurrentBatchSaveFiles 批量并发保存文件
func (fs *FileStorage) ConcurrentBatchSaveFiles(files map[string]io.Reader) (map[string]string, error) {
	g := &errgroup.Group{}
	g.GOMAXPROCS(int(fs.ConcurrencyLimit))

	filePaths := make(map[string]string, len(files))
	for filePath, file := range files {
		filePath, file := filePath, file
		g.Go(func(ctx context.Context) error {
			fp, err := fs.SaveFile(file, filePath)
			filePaths[filePath] = fp
			if err != nil {
				return errors.Wrapf(err, "failed to concurrent save file, path: %s", filePath)
			}
			return nil
		})
	}
	return filePaths, g.Wait()
}

// ConcurrentBatchGetFiles 批量并发获取文件
func (fs *FileStorage) ConcurrentBatchGetFiles(filePaths []string) ([]io.ReadCloser, error) {
	fileReadClosers := make([]io.ReadCloser, len(filePaths))
	g := &errgroup.Group{}
	g.GOMAXPROCS(int(fs.ConcurrencyLimit))
	for idx, filePath := range filePaths {
		idx, filePath := idx, filePath // https://golang.org/doc/faq#closures_and_goroutines
		g.Go(func(ctx context.Context) error {
			fileReadCloser, err := fs.GetFile(filePath)
			if err != nil {
				return errors.Wrapf(err, "failed to con current get file, path: %s", filePath)
			}
			fileReadClosers[idx] = fileReadCloser
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}
	return fileReadClosers, nil
}

// ConcurrentBatchCleanFiles 批量并发清理文件
func (fs *FileStorage) ConcurrentBatchCleanFiles(filePaths []string) error {
	g := &errgroup.Group{}
	g.GOMAXPROCS(int(fs.ConcurrencyLimit))
	for _, filePath := range filePaths {
		filePath := filePath // https://golang.org/doc/faq#closures_and_goroutines
		g.Go(func(ctx context.Context) error {
			if err := fs.CleanFile(filePath); err != nil {
				return errors.Wrapf(err, "failed to concurrent clean file, path: %s", filePath)
			}
			return nil
		})
	}
	return g.Wait()
}

func defaultFileStorageConfig() *FileStorage {
	return &FileStorage{
		ConcurrencyLimit: 10,
	}
}

// FileStorageOption 文件存储选项配置
type FileStorageOption func(*FileStorage)

// NewFileStorage 创建一个文件存储
func NewFileStorage(fgsc FileGetSaveCleaner, opts ...FileStorageOption) *FileStorage {
	fs := defaultFileStorageConfig()
	fs.FileGetSaveCleaner = fgsc
	for _, o := range opts {
		o(fs)
	}
	return fs
}

// NewHttpFileStorage 创建一个http协议传输文件仓储
func NewHttpFileStorage(apiConfig APIConfig, opts ...HttpFileStorageOption) *FileStorage {
	return NewFileStorage(NewHttpFileGetSaveCleaner(apiConfig, opts...))
}

// NewlocalDiskFileStorage 创建一个本地文件仓储
func NewLocalDiskFileStorage(basePath string) *FileStorage {
	return NewFileStorage(NewLocalDiskFileGetSaveCleaner(basePath))
}

// 检查是否实现了接口
var _ IFileStorage = (*FileStorage)(nil)
