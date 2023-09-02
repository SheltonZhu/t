package storage

import (
	"bytes"
	"io"
	"testing"

	mock_storage "t/mocks/storage"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestFileStorageBytes(t *testing.T) {
	// 创建内存存储实例
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockFileGetSaveCleaner := mock_storage.NewMockFileGetSaveCleaner(ctrl)
	storage := NewFileStorage(mockFileGetSaveCleaner)

	fileContent := []byte("This is a test file.")
	fileName := "test.txt"

	t.Run("Save file bytes", func(t *testing.T) {
		mockFileGetSaveCleaner.EXPECT().SaveFile(gomock.Any(), gomock.Any()).
			Return(fileName, nil).Times(1)

		_, err := storage.SaveFileBytes(fileContent, fileName)
		assert.NoError(t, err)
	})

	t.Run("Get file bytes", func(t *testing.T) {
		mockFileGetSaveCleaner.EXPECT().GetFile(gomock.Any()).
			Return(io.NopCloser(bytes.NewReader(fileContent)), nil).Times(1)

		storage.OnAfterResponse(func(*FileStorage, *[]byte) {})

		_, err := storage.GetFileBytes(fileName)
		assert.NoError(t, err)
	})

	t.Run("Failed to io.ReadAll", func(t *testing.T) {
		stubs := gomonkey.ApplyFunc(io.ReadAll, func(io.Reader) ([]byte, error) {
			return nil, assert.AnError
		})

		defer stubs.Reset()
		mockFileGetSaveCleaner.EXPECT().GetFile("read.txt").
			Return(io.NopCloser(bytes.NewReader([]byte("test"))), nil).AnyTimes()

		// 获取文件
		_, err := storage.GetFileBytes("read.txt")
		assert.Error(t, err)
	})

	t.Run("Operate file bytes error", func(t *testing.T) {
		mockFileGetSaveCleaner.EXPECT().SaveFile(gomock.Any(), gomock.Any()).
			Return("", assert.AnError).Times(1)
		mockFileGetSaveCleaner.EXPECT().GetFile(gomock.Any()).
			Return(nil, assert.AnError).Times(1)
		mockFileGetSaveCleaner.EXPECT().CleanFile(gomock.Any()).
			Return(assert.AnError).Times(1)

		// 保存文件
		_, err := storage.SaveFileBytes(fileContent, fileName)
		assert.Error(t, err)

		// 获取文件
		_, err = storage.GetFileBytes(fileName)
		assert.Error(t, err)

		// 清理文件
		err = storage.CleanFile(fileName)
		assert.Error(t, err)
	})
}

func TestBatchFileStorage(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFileGetSaveCleaner := mock_storage.NewMockFileGetSaveCleaner(ctrl)

	storage := NewFileStorage(mockFileGetSaveCleaner)
	batchFiles := map[string]io.Reader{
		"file1.txt": bytes.NewReader([]byte("File 1 content")),
		"file2.txt": bytes.NewReader([]byte("File 2 content")),
	}
	batchFilenames := []string{"file1.txt", "file2.txt"}

	t.Run("Batch save files", func(t *testing.T) {
		mockFileGetSaveCleaner.EXPECT().SaveFile(gomock.Any(), gomock.Any()).
			Return("test.txt", nil).Times(len(batchFiles))

		// 批量保存文件
		_, errs := storage.BatchSaveFiles(batchFiles)
		assert.Zero(t, len(errs), "Failed to batch save files")
	})

	t.Run("Batch get files", func(t *testing.T) {
		mockFileGetSaveCleaner.EXPECT().GetFile(gomock.Any()).
			Return(io.NopCloser(bytes.NewReader([]byte("This is a test file."))), nil).Times(1)
		mockFileGetSaveCleaner.EXPECT().GetFile(gomock.Any()).
			Return(io.NopCloser(bytes.NewReader([]byte("This is a test file."))), nil).Times(1)
		// 批量获取文件
		retrievedFiles, errs := storage.BatchGetFiles(batchFilenames)
		defer func() {
			for _, fileReadCloser := range retrievedFiles {
				err := fileReadCloser.Close()
				assert.NoError(t, err)
			}
		}()
		assert.Zero(t, len(errs))
		for idx, fileReader := range retrievedFiles {
			retrievedContent, err := io.ReadAll(fileReader)
			assert.NoError(t, err)
			expectedContent := []byte("This is a test file.")
			assert.Equal(t, expectedContent, retrievedContent, "Retrieved content does not match for file: %d", idx)
		}
	})

	t.Run("Batch clean files", func(t *testing.T) {
		mockFileGetSaveCleaner.EXPECT().CleanFile(gomock.Any()).
			Return(nil).Times(len(batchFilenames))
		// 批量清理
		errs := storage.BatchCleanFiles(batchFilenames)
		assert.Zero(t, len(errs), "Failed to batch clean files")
	})

	t.Run("Batch operate file reader error", func(t *testing.T) {
		mockFileGetSaveCleaner.EXPECT().SaveFile(gomock.Any(), gomock.Any()).
			Return("", assert.AnError).Times(len(batchFiles))
		mockFileGetSaveCleaner.EXPECT().GetFile(gomock.Any()).
			Return(nil, assert.AnError).Times(len(batchFilenames))
		mockFileGetSaveCleaner.EXPECT().CleanFile(gomock.Any()).
			Return(assert.AnError).Times(len(batchFilenames))
		// 批量保存文件
		_, errs := storage.BatchSaveFiles(batchFiles)
		assert.NotZero(t, len(errs))

		// 批量获取文件
		_, errs = storage.BatchGetFiles(batchFilenames)
		assert.NotZero(t, len(errs))

		// 批量清理
		errs = storage.BatchCleanFiles(batchFilenames)
		assert.NotZero(t, len(errs))
	})
}

func TestConcurrentBatchFileStorage(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFileGetSaveCleaner := mock_storage.NewMockFileGetSaveCleaner(ctrl)

	storage := NewFileStorage(mockFileGetSaveCleaner, func(fs *FileStorage) {})
	storage.SetConcurrencyLimit(10)

	batchFiles := map[string]io.Reader{
		"cfile1.txt": bytes.NewReader([]byte("File 1 content")),
		"cfile2.txt": bytes.NewReader([]byte("File 2 content")),
	}
	t.Run("Concurrent batch save file", func(t *testing.T) {
		mockFileGetSaveCleaner.EXPECT().
			SaveFile(gomock.Any(), gomock.Any()).
			Return("test.txt", nil).Times(len(batchFiles))

		// 批量并发保存文件
		_, err := storage.ConcurrentBatchSaveFiles(batchFiles)
		assert.NoError(t, err, "Failed to concurrent batch save files")
	})

	batchFilenames := []string{"cfile1.txt", "cfile2.txt"}

	t.Run("Concurrent batch get file", func(t *testing.T) {
		mockFileGetSaveCleaner.EXPECT().GetFile("cfile1.txt").
			Return(io.NopCloser(bytes.NewReader([]byte("This is a test file."))), nil).Times(1)

		mockFileGetSaveCleaner.EXPECT().GetFile("cfile2.txt").
			Return(io.NopCloser(bytes.NewReader([]byte("This is a test file."))), nil).Times(1)

		// 批量并发获取文件
		retrievedFiles, err := storage.ConcurrentBatchGetFiles(batchFilenames)
		defer func() {
			for _, fileReadCloser := range retrievedFiles {
				err := fileReadCloser.Close()
				assert.NoError(t, err)
			}
		}()
		assert.NoError(t, err, "Failed to concurrent batch get files")
		for idx, fileReader := range retrievedFiles {
			retrievedContent, err := io.ReadAll(fileReader)
			assert.NoError(t, err, "Failed to read file content")
			expectedContent := []byte("This is a test file.")
			assert.Equal(t, expectedContent, retrievedContent, "Retrieved content does not match for file: %d", idx)
		}
	})

	t.Run("Concurrent batch clean file", func(t *testing.T) {
		mockFileGetSaveCleaner.EXPECT().CleanFile(gomock.Any()).
			Return(nil).Times(len(batchFilenames))

		// 批量并发清理
		err := storage.ConcurrentBatchCleanFiles(batchFilenames)
		assert.NoError(t, err, "Failed to concurrent batch clean files")
	})

	t.Run("Concurrent batch operate file reader error", func(t *testing.T) {
		mockFileGetSaveCleaner.EXPECT().SaveFile(gomock.Any(), gomock.Any()).
			Return("", assert.AnError).Times(len(batchFiles))
		mockFileGetSaveCleaner.EXPECT().GetFile(gomock.Any()).
			Return(nil, assert.AnError).Times(len(batchFilenames))
		mockFileGetSaveCleaner.EXPECT().CleanFile(gomock.Any()).
			Return(assert.AnError).Times(len(batchFilenames))

		// 批量并发保存文件
		_, err := storage.ConcurrentBatchSaveFiles(batchFiles)
		assert.Error(t, err)

		// 批量并发保存文件
		_, err = storage.ConcurrentBatchGetFiles(batchFilenames)
		assert.Error(t, err)

		// 批量并发清理文件
		err = storage.ConcurrentBatchCleanFiles(batchFilenames)
		assert.Error(t, err)
	})
}
