package storage

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	mock_storage "t/mocks/storage"

	"github.com/go-resty/resty/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestLocalDiskFileStorage(t *testing.T) {
	// 创建临时目录作为存储路径
	tempDir, err := os.MkdirTemp(".", "local_disk_storage_test")
	if err != nil {
		t.Fatalf("Failed to create temporary directory: %v", err)
	}
	defer func() {
		err := os.RemoveAll(tempDir)
		assert.NoError(t, err)
	}()

	// 创建本地磁盘存储实例
	storage := NewLocalDiskFileStorage(tempDir)

	// 保存文件
	fileContent := []byte("This is a test file.")
	fileName := "test.txt"
	_, err = storage.SaveFileBytes(fileContent, fileName)
	assert.NoError(t, err, "Failed to save file")

	// 获取文件
	retrievedContent, err := storage.GetFileBytes(fileName)
	assert.NoError(t, err, "Failed to read file content")
	assert.Equal(t, fileContent, retrievedContent, "Retrieved file content does not match")

	// 修改响应值
	newBytes := []byte("12345")
	storage.OnAfterResponse(func(fs *FileStorage, bs *[]byte) { *bs = newBytes })
	retrievedContent, err = storage.GetFileBytes(fileName)
	assert.NoError(t, err, "Failed to read file content")
	assert.Equal(t, newBytes, retrievedContent, "Retrieved file content does not match")

	// 清理文件
	err = storage.CleanFile(fileName)
	assert.NoError(t, err)
}

func TestLocalDiskFileStorageErr(t *testing.T) {
	// 创建本地磁盘存储实例
	storage := NewLocalDiskFileStorage("not_exists")

	f := isDirExists("not_exists")
	assert.False(t, f)
	
	_, err := storage.GetFile("not_exists")
	assert.Error(t, err, "Failed to read file content")

	err = storage.CleanFile("not_exists")
	assert.Error(t, err)
}

func TestHttpFileStorage(t *testing.T) {
	// mock 实现
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 这里构造 mock 的具体处理细节
		w.WriteHeader(http.StatusOK)
		if r.Method != "GET" {
			assert.Equal(t, "/upload/test.txt", r.URL.Path)
			_, _ = w.Write([]byte("This is a test file."))
			return
		}
		assert.Equal(t, "/download/test.txt", r.URL.Path)
		_, _ = w.Write([]byte("This is a test file."))
	}))
	defer ts.Close()

	// 创建http存储实例
	storage := NewHttpFileStorage(
		APIConfig{
			Host:        "example.com",
			UploadAPI:   "/upload",
			DownloadAPI: "/download",
		},
		WithHttpClient(http.DefaultClient),
		SetBaseURL(ts.URL),
		WithHttpHeaderHost("test.com"),
		SetTimeout(time.Second),
		WithHttpTrace(),
		WithHttpBasicAuth("user", "user"),
		SetDebug(1000),
		SetHttps(),
		OnBeforeGetFile(func(s *httpFileStorage, r *resty.Request) {
			r.SetQueryParam("test", "1")
		}),
		OnAfterSaveFile(func(hfs *httpFileStorage, r *resty.Response, filePath *string) {
			assert.Equal(t, r.Body(), []byte("This is a test file."))
		}),
		OnBeforeSaveFile(func(s *httpFileStorage, r *resty.Request, file io.Reader) {
			r.SetQueryParam("test", "1")
			r.SetFileReader("file", "test.txt", file)
		}),
	)
	// 保存文件
	fileContent := []byte("This is a test file.")
	fileName := "test.txt"
	_, err := storage.SaveFileBytes(fileContent, fileName)
	assert.NoError(t, err, "Failed to save file")

	// 获取文件
	retrievedContent, err := storage.GetFileBytes(fileName)
	assert.NoError(t, err, "Failed to read file content", retrievedContent)
	assert.Equal(t, fileContent, retrievedContent, "Retrieved file content does not match")

	// 清理文件
	err = storage.CleanFile(fileName)
	assert.NoError(t, err)
}

func TestHttpFileStorageErr(t *testing.T) {
	// mock 实现
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 这里构造 mock 的具体处理细节
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer ts.Close()

	t.Run("Response is not 200", func(t *testing.T) {
		// 响应不为200
		// 创建http存储实例
		storage := NewHttpFileStorage(
			APIConfig{
				Host:        "example.com",
				UploadAPI:   "/upload",
				DownloadAPI: "/download",
			},
			SetBaseURL(ts.URL),
			SetTimeout(time.Second),
			SetHttps(),
		)

		_, err := storage.SaveFileBytes([]byte("This is a test file."), "test.txt")
		assert.Error(t, err, "Failed to save file")

		_, err = storage.GetFileBytes("test.txt")
		assert.Error(t, err, "Failed to save file")
	})

	t.Run("Response error", func(t *testing.T) {
		// 服务器返回错误
		// 创建http存储实例
		storage := NewHttpFileStorage(
			APIConfig{
				Host:        "!@#!@$!@!@%@!%",
				UploadAPI:   "/upload",
				DownloadAPI: "/download",
			},
			SetTimeout(time.Second),
			SetHttps(),
		)

		_, err := storage.SaveFileBytes([]byte("This is a test file."), "test.txt")
		assert.Error(t, err, "Failed to save file")

		_, err = storage.GetFileBytes("test.txt")
		assert.Error(t, err, "Failed to save file")
	})

	t.Run("Request api parse error", func(t *testing.T) {
		// API 解析错误
		// 创建http存储实例
		storage := NewHttpFileStorage(
			APIConfig{
				Host:        "example.com",
				UploadAPI:   "!@$!@%!@!5",
				DownloadAPI: "!$!@$!%@%!@%!@%!",
			},
			SetBaseURL(ts.URL),
			SetTimeout(time.Second),
			SetHttps(),
		)

		_, err := storage.SaveFileBytes([]byte("This is a test file."), "test.txt")
		assert.Error(t, err, "Failed to save file")

		_, err = storage.GetFileBytes("test.txt")
		assert.Error(t, err, "Failed to save file")
	})
}

func TestFileStorageErr(t *testing.T) {
	// 创建内存存储实例
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFileGetSaveCleaner := mock_storage.NewMockFileGetSaveCleaner(ctrl)
	mockFileGetSaveCleaner.EXPECT().
		SaveFile(gomock.Any(), gomock.Any()).
		Return("", assert.AnError).AnyTimes()
	mockFileGetSaveCleaner.EXPECT().
		GetFile(gomock.Any()).
		Return(nil, assert.AnError).AnyTimes()
	mockFileGetSaveCleaner.EXPECT().
		CleanFile(gomock.Any()).
		Return(assert.AnError).AnyTimes()

	storage := NewFileStorage(mockFileGetSaveCleaner)

	t.Run("Operate file bytes error", func(t *testing.T) {
		// 保存文件
		fileContent := []byte("This is a test file.")
		fileName := "test.txt"
		_, err := storage.SaveFileBytes(fileContent, fileName)
		assert.Error(t, err, "Failed to save file")

		// 获取文件
		_, err = storage.GetFileBytes(fileName)
		assert.Error(t, err, "Failed to read file content")

		// 清理文件
		err = storage.CleanFile(fileName)
		assert.Error(t, err)
	})

	batchFiles := map[string]io.Reader{
		"file1.txt": bytes.NewReader([]byte("File 1 content")),
		"file2.txt": bytes.NewReader([]byte("File 2 content")),
	}
	batchFilenames := []string{"file1.txt", "file2.txt"}

	t.Run("batch operate file reader error", func(t *testing.T) {
		// 批量保存文件
		_, errs := storage.BatchSaveFiles(batchFiles)
		assert.NotZero(t, len(errs), "Failed to batch save files")

		// 批量获取文件
		_, errs = storage.BatchGetFiles(batchFilenames)
		assert.NotZero(t, len(errs), "Failed to batch get files")

		// 批量清理
		errs = storage.BatchCleanFiles(batchFilenames)
		assert.NotZero(t, len(errs), "Failed to batch clean files")
	})

	t.Run("Concurrent batch operate file reader error", func(t *testing.T) {
		// 批量并发保存文件
		_, err := storage.ConcurrentBatchSaveFiles(batchFiles)
		assert.Error(t, err, "Failed to concurrent batch save files")

		// 批量并发保存文件
		_, err = storage.ConcurrentBatchGetFiles(batchFilenames)
		assert.Error(t, err, "Failed to concurrent batch save files")

		// 批量并发清理文件
		err = storage.ConcurrentBatchCleanFiles(batchFilenames)
		assert.Error(t, err, "Failed to concurrent batch clean files")
	})
}

func TestBatchFileStorage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFileGetSaveCleaner := mock_storage.NewMockFileGetSaveCleaner(ctrl)
	gomock.InOrder(
		mockFileGetSaveCleaner.EXPECT().
			SaveFile(gomock.Any(), gomock.Any()).
			Return("test.txt", nil).AnyTimes(),
		mockFileGetSaveCleaner.EXPECT().
			GetFile("file1.txt").
			Return(io.NopCloser(bytes.NewReader([]byte("This is a test file."))), nil),
		mockFileGetSaveCleaner.EXPECT().
			GetFile("file2.txt").
			Return(io.NopCloser(bytes.NewReader([]byte("This is a test file."))), nil),
		mockFileGetSaveCleaner.EXPECT().
			CleanFile(gomock.Any()).
			Return(nil).AnyTimes(),
	)

	storage := NewFileStorage(mockFileGetSaveCleaner)

	// 批量保存文件
	batchFiles := map[string]io.Reader{
		"file1.txt": bytes.NewReader([]byte("File 1 content")),
		"file2.txt": bytes.NewReader([]byte("File 2 content")),
	}
	_, errs := storage.BatchSaveFiles(batchFiles)
	assert.Zero(t, len(errs), "Failed to batch save files")

	// 批量获取文件
	batchFilenames := []string{"file1.txt", "file2.txt"}
	retrievedFiles, errs := storage.BatchGetFiles(batchFilenames)
	defer func() {
		for _, fileReadCloser := range retrievedFiles {
			err := fileReadCloser.Close()
			assert.NoError(t, err)
		}
		// 批量清理
		errs := storage.BatchCleanFiles(batchFilenames)
		assert.Zero(t, len(errs), "Failed to batch clean files")
	}()
	assert.Zero(t, len(errs), "Failed to batch get files")
	for idx, fileReader := range retrievedFiles {
		retrievedContent, err := io.ReadAll(fileReader)
		assert.NoError(t, err, "Failed to read file content")
		expectedContent := []byte("This is a test file.")
		assert.Equal(t, expectedContent, retrievedContent, "Retrieved content does not match for file: %d", idx)
	}
}

func TestConcurrentBatchFileStorage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFileGetSaveCleaner := mock_storage.NewMockFileGetSaveCleaner(ctrl)
	mockFileGetSaveCleaner.EXPECT().
		SaveFile(gomock.Any(), gomock.Any()).
		Return("test.txt", nil).AnyTimes()
	mockFileGetSaveCleaner.EXPECT().
		GetFile("cfile1.txt").
		Return(io.NopCloser(bytes.NewReader([]byte("This is a test file."))), nil)
	mockFileGetSaveCleaner.EXPECT().
		GetFile("cfile2.txt").
		Return(io.NopCloser(bytes.NewReader([]byte("This is a test file."))), nil)
	mockFileGetSaveCleaner.EXPECT().
		CleanFile(gomock.Any()).
		Return(nil).AnyTimes()

	storage := NewFileStorage(mockFileGetSaveCleaner, func(fs *FileStorage) {
		fs.SetConcurrencyLimit(10)
	})
	storage.SetConcurrencyLimit(10)
	// 批量并发保存文件
	batchFiles := map[string]io.Reader{
		"cfile1.txt": bytes.NewReader([]byte("File 1 content")),
		"cfile2.txt": bytes.NewReader([]byte("File 2 content")),
	}
	_, err := storage.ConcurrentBatchSaveFiles(batchFiles)
	assert.NoError(t, err, "Failed to concurrent batch save files")

	// 批量并发获取文件
	batchFilenames := []string{"cfile1.txt", "cfile2.txt"}
	retrievedFiles, err := storage.ConcurrentBatchGetFiles(batchFilenames)
	defer func() {
		for _, fileReadCloser := range retrievedFiles {
			err := fileReadCloser.Close()
			assert.NoError(t, err)
		}
		// 批量并发清理
		err := storage.ConcurrentBatchCleanFiles(batchFilenames)
		assert.NoError(t, err, "Failed to concurrent batch clean files")
	}()
	assert.NoError(t, err, "Failed to concurrent batch get files")
	for idx, fileReader := range retrievedFiles {
		retrievedContent, err := io.ReadAll(fileReader)
		assert.NoError(t, err, "Failed to read file content")
		expectedContent := []byte("This is a test file.")
		assert.Equal(t, expectedContent, retrievedContent, "Retrieved content does not match for file: %d", idx)
	}
}
