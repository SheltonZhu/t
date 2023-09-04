package storage

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

type MyLogger struct{}

func (l *MyLogger) Errorf(format string, v ...interface{}) {}
func (l *MyLogger) Warnf(format string, v ...interface{})  {}
func (l *MyLogger) Debugf(format string, v ...interface{}) {}

func TestNewHttpFileGetSaveCleaner(t *testing.T) {
	t.Parallel()
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

	myLogger := &MyLogger{}
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
		SetLogger(myLogger),
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
	fileContent := []byte("This is a test file.")
	fileName := "test.txt"

	t.Run("Save file", func(t *testing.T) {
		// 保存文件
		_, err := storage.SaveFile(bytes.NewReader(fileContent), fileName)
		assert.NoError(t, err, "Failed to save file")
	})

	t.Run("Get file", func(t *testing.T) {
		// 获取文件
		retrievedContentReader, err := storage.GetFile(fileName)
		assert.NoError(t, err, "Failed to read file content")
		retrievedContent, err := io.ReadAll(retrievedContentReader)
		assert.NoError(t, err)
		assert.Equal(t, fileContent, retrievedContent, "Retrieved file content does not match")
	})

	t.Run("Clean file", func(t *testing.T) {
		// 清理文件
		err := storage.CleanFile(fileName)
		assert.NoError(t, err)
	})
}

func TestHttpFileStorageErr(t *testing.T) {
	t.Parallel()
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

		_, err := storage.SaveFile(bytes.NewReader([]byte("This is a test file.")), "test.txt")
		assert.Error(t, err)

		_, err = storage.GetFile("test.txt")
		assert.Error(t, err)
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

		_, err := storage.SaveFile(bytes.NewReader([]byte("This is a test file.")), "test.txt")
		assert.Error(t, err)

		_, err = storage.GetFile("test.txt")
		assert.Error(t, err)
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
		assert.Error(t, err)

		_, err = storage.GetFileBytes("test.txt")
		assert.Error(t, err)
	})
}
