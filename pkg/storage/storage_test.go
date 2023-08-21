package storage

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/go-resty/resty/v2"
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
	err = storage.SaveFileBytes(fileContent, fileName)
	assert.NoError(t, err, "Failed to save file")

	// 获取文件
	retrievedContent, err := storage.GetFileBytes(fileName)
	assert.NoError(t, err, "Failed to read file content")
	assert.Equal(t, fileContent, retrievedContent, "Retrieved file content does not match")

	// 修改响应值
	newBytes := []byte("12345")
	storage.OnAfterResponse(func(bs *[]byte) { *bs = newBytes })
	retrievedContent, err = storage.GetFileBytes(fileName)
	assert.NoError(t, err, "Failed to read file content")
	assert.Equal(t, newBytes, retrievedContent, "Retrieved file content does not match")

	// 清理文件
	err = storage.CleanFile(fileName)
	assert.NoError(t, err)

	// 批量保存文件
	batchFiles := map[string]io.Reader{
		"file1.txt": bytes.NewReader([]byte("File 1 content")),
		"file2.txt": bytes.NewReader([]byte("File 2 content")),
	}
	errs := storage.BatchSaveFiles(batchFiles)
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
		expectedContent := []byte(fmt.Sprintf("File %d content", idx+1))
		assert.Equal(t, expectedContent, retrievedContent, "Retrieved content does not match for file: %d", idx)
	}

	storage.SetConcurrencyLimit(1)
	// 批量并发保存文件
	batchFiles1 := map[string]io.Reader{
		"cfile1.txt": bytes.NewReader([]byte("File 1 content")),
		"cfile2.txt": bytes.NewReader([]byte("File 2 content")),
	}
	err = storage.ConcurrentBatchSaveFiles(batchFiles1)
	assert.NoError(t, err, "Failed to concurrent batch save files")

	// 批量并发获取文件
	batchFilenames1 := []string{"cfile1.txt", "cfile2.txt"}
	retrievedFiles1, err := storage.ConcurrentBatchGetFiles(batchFilenames1)
	defer func() {
		for _, fileReadCloser := range retrievedFiles1 {
			err := fileReadCloser.Close()
			assert.NoError(t, err)
		}
		// 批量并发清理
		err := storage.ConcurrentBatchCleanFiles(batchFilenames1)
		assert.NoError(t, err, "Failed to concurrent batch clean files")
	}()
	assert.NoError(t, err, "Failed to concurrent batch get files")
	for idx, fileReader := range retrievedFiles1 {
		retrievedContent, err := io.ReadAll(fileReader)
		assert.NoError(t, err, "Failed to read file content")
		expectedContent := []byte(fmt.Sprintf("File %d content", idx+1))
		assert.Equal(t, expectedContent, retrievedContent, "Retrieved content does not match for file: %d", idx)
	}
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
		[]HttpFileStorageOption{
			WithHttpClient(http.DefaultClient),
			SetBaseURL(ts.URL),
			WithHttpHeaderHost("test.com"),
			SetTimeout(time.Second),
			WithHttpTrace(),
			WithHttpBasicAuth("user", "user"),
			SetDebug(),
			SetHttps(),
			OnBeforeGetFile(func(r *resty.Request) {
				r.SetQueryParam("test", "1")
			}),
			OnBeforeSaveFile(func(r *resty.Request) {
				r.SetQueryParam("test", "1")
			}),
		}...,
	)
	// 保存文件
	fileContent := []byte("This is a test file.")
	fileName := "test.txt"
	err := storage.SaveFileBytes(fileContent, fileName)
	assert.NoError(t, err, "Failed to save file")

	// 获取文件
	retrievedContent, err := storage.GetFileBytes(fileName)
	assert.NoError(t, err, "Failed to read file content", retrievedContent)
	assert.Equal(t, fileContent, retrievedContent, "Retrieved file content does not match")
}
