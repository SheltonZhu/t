package storage

import (
	"errors"
	"io"
	"net/http"
	"path"
	"time"

	"github.com/go-resty/resty/v2"
)

type httpFileStorage struct {
	HttpClient *resty.Client
	APIConfig
	beforeSaveFile []func(*resty.Request)
	beforeGetFile  []func(*resty.Request)
}

type APIConfig struct {
	Schema      string
	Host        string
	UploadAPI   string
	DownloadAPI string
}

// HttpFileStorageOption http协议文件存储选项配置
type HttpFileStorageOption func(*httpFileStorage)

// WithHttpClient 基于http.Client创建resty.Client
func WithHttpClient(hc *http.Client) HttpFileStorageOption {
	return func(s *httpFileStorage) {
		s.HttpClient = resty.NewWithClient(hc)
	}
}

// WithHttpTrace 开启restyClient的trace功能
// 如果手动设置了restyClient需要放在其后
func WithHttpTrace() HttpFileStorageOption {
	return func(s *httpFileStorage) {
		s.HttpClient.EnableTrace()
	}
}

// SetDebug 开启restyClient的debug功能
// 如果手动设置了restyClient需要放在其后
func SetDebug(bodyLimit ...int64) HttpFileStorageOption {
	return func(s *httpFileStorage) {
		s.HttpClient.SetDebug(true)
		if len(bodyLimit) > 0 {
			s.HttpClient.SetDebugBodyLimit(bodyLimit[0])
		}
	}
}

// WithHttpHeader 设置公共请求头
// 如果手动设置了restyClient需要放在其后
func WithHttpHeader(key, val string) HttpFileStorageOption {
	return func(s *httpFileStorage) {
		s.HttpClient.SetHeader(key, val)
	}
}

// WithHttpHeaderHost 设置Host请求头
// 如果手动设置了restyClient需要放在其后
func WithHttpHeaderHost(host string) HttpFileStorageOption {
	return WithHttpHeader("host", host)
}

// SetBaseURL 设置base url
// 如果手动设置了restyClient需要放在其后
func SetBaseURL(domain string) HttpFileStorageOption {
	return func(s *httpFileStorage) {
		s.Host = domain
	}
}

// WithHttpBasicAuth 设置basic验证
// 如果手动设置了restyClient需要放在其后
func WithHttpBasicAuth(username, password string) HttpFileStorageOption {
	return func(s *httpFileStorage) {
		s.HttpClient.SetBasicAuth(username, password)
	}
}

// SetHttps 设置https
// 如果手动设置了restyClient需要放在其后
func SetHttps() HttpFileStorageOption {
	return func(s *httpFileStorage) {
		s.Schema = "https"
	}
}

// SetTimeout 设置超时时间
// 如果手动设置了restyClient需要放在其后
func SetTimeout(timeout time.Duration) HttpFileStorageOption {
	return func(s *httpFileStorage) {
		s.HttpClient.SetTimeout(timeout)
	}
}

// OnBeforeSaveFile 发送请求前可以做一些处理
func OnBeforeSaveFile(f func(*resty.Request)) HttpFileStorageOption {
	return func(s *httpFileStorage) {
		s.beforeSaveFile = append(s.beforeSaveFile, f)
	}
}

// OnBeforeGetFile 发送请求前可以做一些处理
func OnBeforeGetFile(f func(*resty.Request)) HttpFileStorageOption {
	return func(s *httpFileStorage) {
		s.beforeGetFile = append(s.beforeGetFile, f)
	}
}

// NewHttpFileGetSaveCleaner
func NewHttpFileGetSaveCleaner(apiConfig APIConfig, opts ...HttpFileStorageOption) *httpFileStorage {
	if apiConfig.Schema == "" {
		apiConfig.Schema = "http"
	}

	fs := &httpFileStorage{
		HttpClient: resty.New(),
		APIConfig:  apiConfig,
	}
	for _, o := range opts {
		o(fs)
	}
	fs.HttpClient.
		SetScheme(fs.Schema).
		SetBaseURL(fs.Host)

	return fs
}

// SaveFile 实现了 FileSaver 接口的 SaveFile 方法
func (s *httpFileStorage) SaveFile(file io.Reader, filePath string) error {
	uploadURL := path.Join(s.UploadAPI, filePath)
	req := s.HttpClient.R().SetBody(file)
	for _, h := range s.beforeSaveFile {
		h(req)
	}
	resp, err := req.Post(uploadURL)
	if err != nil {
		return err
	}
	if resp.StatusCode() != http.StatusOK {
		return errors.New("failed to save file")
	}
	return nil
}

// GetFile 实现了 FileGetter 接口的 GetFile 方法
func (s *httpFileStorage) GetFile(filePath string) (io.ReadCloser, error) {
	downloadURL := path.Join(s.DownloadAPI, filePath)
	req := s.HttpClient.R().SetDoNotParseResponse(true)
	for _, h := range s.beforeGetFile {
		h(req)
	}
	resp, err := req.Get(downloadURL)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New("failed to get file")
	}

	return resp.RawBody(), nil
}

// CleanFile 实现了 FileCleaner 接口的 CleanFile 方法
func (s *httpFileStorage) CleanFile(filePath string) error { return nil }

var _ FileGetSaveCleaner = (*httpFileStorage)(nil)
