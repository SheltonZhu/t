// Code generated by MockGen. DO NOT EDIT.
// Source: storage.go

// Package mock_storage is a generated GoMock package.
package mock_storage

import (
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockFileSaver is a mock of FileSaver interface.
type MockFileSaver struct {
	ctrl     *gomock.Controller
	recorder *MockFileSaverMockRecorder
}

// MockFileSaverMockRecorder is the mock recorder for MockFileSaver.
type MockFileSaverMockRecorder struct {
	mock *MockFileSaver
}

// NewMockFileSaver creates a new mock instance.
func NewMockFileSaver(ctrl *gomock.Controller) *MockFileSaver {
	mock := &MockFileSaver{ctrl: ctrl}
	mock.recorder = &MockFileSaverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFileSaver) EXPECT() *MockFileSaverMockRecorder {
	return m.recorder
}

// SaveFile mocks base method.
func (m *MockFileSaver) SaveFile(file io.Reader, filePath string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveFile", file, filePath)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveFile indicates an expected call of SaveFile.
func (mr *MockFileSaverMockRecorder) SaveFile(file, filePath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveFile", reflect.TypeOf((*MockFileSaver)(nil).SaveFile), file, filePath)
}

// MockFileGetter is a mock of FileGetter interface.
type MockFileGetter struct {
	ctrl     *gomock.Controller
	recorder *MockFileGetterMockRecorder
}

// MockFileGetterMockRecorder is the mock recorder for MockFileGetter.
type MockFileGetterMockRecorder struct {
	mock *MockFileGetter
}

// NewMockFileGetter creates a new mock instance.
func NewMockFileGetter(ctrl *gomock.Controller) *MockFileGetter {
	mock := &MockFileGetter{ctrl: ctrl}
	mock.recorder = &MockFileGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFileGetter) EXPECT() *MockFileGetterMockRecorder {
	return m.recorder
}

// GetFile mocks base method.
func (m *MockFileGetter) GetFile(filePath string) (io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFile", filePath)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFile indicates an expected call of GetFile.
func (mr *MockFileGetterMockRecorder) GetFile(filePath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFile", reflect.TypeOf((*MockFileGetter)(nil).GetFile), filePath)
}

// MockFileCleaner is a mock of FileCleaner interface.
type MockFileCleaner struct {
	ctrl     *gomock.Controller
	recorder *MockFileCleanerMockRecorder
}

// MockFileCleanerMockRecorder is the mock recorder for MockFileCleaner.
type MockFileCleanerMockRecorder struct {
	mock *MockFileCleaner
}

// NewMockFileCleaner creates a new mock instance.
func NewMockFileCleaner(ctrl *gomock.Controller) *MockFileCleaner {
	mock := &MockFileCleaner{ctrl: ctrl}
	mock.recorder = &MockFileCleanerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFileCleaner) EXPECT() *MockFileCleanerMockRecorder {
	return m.recorder
}

// CleanFile mocks base method.
func (m *MockFileCleaner) CleanFile(filePath string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CleanFile", filePath)
	ret0, _ := ret[0].(error)
	return ret0
}

// CleanFile indicates an expected call of CleanFile.
func (mr *MockFileCleanerMockRecorder) CleanFile(filePath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CleanFile", reflect.TypeOf((*MockFileCleaner)(nil).CleanFile), filePath)
}

// MockFileStorageStreamHandler is a mock of FileStorageStreamHandler interface.
type MockFileStorageStreamHandler struct {
	ctrl     *gomock.Controller
	recorder *MockFileStorageStreamHandlerMockRecorder
}

// MockFileStorageStreamHandlerMockRecorder is the mock recorder for MockFileStorageStreamHandler.
type MockFileStorageStreamHandlerMockRecorder struct {
	mock *MockFileStorageStreamHandler
}

// NewMockFileStorageStreamHandler creates a new mock instance.
func NewMockFileStorageStreamHandler(ctrl *gomock.Controller) *MockFileStorageStreamHandler {
	mock := &MockFileStorageStreamHandler{ctrl: ctrl}
	mock.recorder = &MockFileStorageStreamHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFileStorageStreamHandler) EXPECT() *MockFileStorageStreamHandlerMockRecorder {
	return m.recorder
}

// GetFile mocks base method.
func (m *MockFileStorageStreamHandler) GetFile(filePath string) (io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFile", filePath)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFile indicates an expected call of GetFile.
func (mr *MockFileStorageStreamHandlerMockRecorder) GetFile(filePath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFile", reflect.TypeOf((*MockFileStorageStreamHandler)(nil).GetFile), filePath)
}

// SaveFile mocks base method.
func (m *MockFileStorageStreamHandler) SaveFile(file io.Reader, filePath string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveFile", file, filePath)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveFile indicates an expected call of SaveFile.
func (mr *MockFileStorageStreamHandlerMockRecorder) SaveFile(file, filePath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveFile", reflect.TypeOf((*MockFileStorageStreamHandler)(nil).SaveFile), file, filePath)
}

// MockFileGetSaveCleaner is a mock of FileGetSaveCleaner interface.
type MockFileGetSaveCleaner struct {
	ctrl     *gomock.Controller
	recorder *MockFileGetSaveCleanerMockRecorder
}

// MockFileGetSaveCleanerMockRecorder is the mock recorder for MockFileGetSaveCleaner.
type MockFileGetSaveCleanerMockRecorder struct {
	mock *MockFileGetSaveCleaner
}

// NewMockFileGetSaveCleaner creates a new mock instance.
func NewMockFileGetSaveCleaner(ctrl *gomock.Controller) *MockFileGetSaveCleaner {
	mock := &MockFileGetSaveCleaner{ctrl: ctrl}
	mock.recorder = &MockFileGetSaveCleanerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFileGetSaveCleaner) EXPECT() *MockFileGetSaveCleanerMockRecorder {
	return m.recorder
}

// CleanFile mocks base method.
func (m *MockFileGetSaveCleaner) CleanFile(filePath string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CleanFile", filePath)
	ret0, _ := ret[0].(error)
	return ret0
}

// CleanFile indicates an expected call of CleanFile.
func (mr *MockFileGetSaveCleanerMockRecorder) CleanFile(filePath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CleanFile", reflect.TypeOf((*MockFileGetSaveCleaner)(nil).CleanFile), filePath)
}

// GetFile mocks base method.
func (m *MockFileGetSaveCleaner) GetFile(filePath string) (io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFile", filePath)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFile indicates an expected call of GetFile.
func (mr *MockFileGetSaveCleanerMockRecorder) GetFile(filePath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFile", reflect.TypeOf((*MockFileGetSaveCleaner)(nil).GetFile), filePath)
}

// SaveFile mocks base method.
func (m *MockFileGetSaveCleaner) SaveFile(file io.Reader, filePath string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveFile", file, filePath)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveFile indicates an expected call of SaveFile.
func (mr *MockFileGetSaveCleanerMockRecorder) SaveFile(file, filePath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveFile", reflect.TypeOf((*MockFileGetSaveCleaner)(nil).SaveFile), file, filePath)
}

// MockfileStorageBytesHandler is a mock of fileStorageBytesHandler interface.
type MockfileStorageBytesHandler struct {
	ctrl     *gomock.Controller
	recorder *MockfileStorageBytesHandlerMockRecorder
}

// MockfileStorageBytesHandlerMockRecorder is the mock recorder for MockfileStorageBytesHandler.
type MockfileStorageBytesHandlerMockRecorder struct {
	mock *MockfileStorageBytesHandler
}

// NewMockfileStorageBytesHandler creates a new mock instance.
func NewMockfileStorageBytesHandler(ctrl *gomock.Controller) *MockfileStorageBytesHandler {
	mock := &MockfileStorageBytesHandler{ctrl: ctrl}
	mock.recorder = &MockfileStorageBytesHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockfileStorageBytesHandler) EXPECT() *MockfileStorageBytesHandlerMockRecorder {
	return m.recorder
}

// GetFileBytes mocks base method.
func (m *MockfileStorageBytesHandler) GetFileBytes(filePath string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFileBytes", filePath)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFileBytes indicates an expected call of GetFileBytes.
func (mr *MockfileStorageBytesHandlerMockRecorder) GetFileBytes(filePath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFileBytes", reflect.TypeOf((*MockfileStorageBytesHandler)(nil).GetFileBytes), filePath)
}

// SaveFileBytes mocks base method.
func (m *MockfileStorageBytesHandler) SaveFileBytes(fileBytes []byte, filePath string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveFileBytes", fileBytes, filePath)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveFileBytes indicates an expected call of SaveFileBytes.
func (mr *MockfileStorageBytesHandlerMockRecorder) SaveFileBytes(fileBytes, filePath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveFileBytes", reflect.TypeOf((*MockfileStorageBytesHandler)(nil).SaveFileBytes), fileBytes, filePath)
}

// MockbatchHandler is a mock of batchHandler interface.
type MockbatchHandler struct {
	ctrl     *gomock.Controller
	recorder *MockbatchHandlerMockRecorder
}

// MockbatchHandlerMockRecorder is the mock recorder for MockbatchHandler.
type MockbatchHandlerMockRecorder struct {
	mock *MockbatchHandler
}

// NewMockbatchHandler creates a new mock instance.
func NewMockbatchHandler(ctrl *gomock.Controller) *MockbatchHandler {
	mock := &MockbatchHandler{ctrl: ctrl}
	mock.recorder = &MockbatchHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockbatchHandler) EXPECT() *MockbatchHandlerMockRecorder {
	return m.recorder
}

// BatchCleanFiles mocks base method.
func (m *MockbatchHandler) BatchCleanFiles(filePaths []string) []error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchCleanFiles", filePaths)
	ret0, _ := ret[0].([]error)
	return ret0
}

// BatchCleanFiles indicates an expected call of BatchCleanFiles.
func (mr *MockbatchHandlerMockRecorder) BatchCleanFiles(filePaths interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchCleanFiles", reflect.TypeOf((*MockbatchHandler)(nil).BatchCleanFiles), filePaths)
}

// BatchGetFiles mocks base method.
func (m *MockbatchHandler) BatchGetFiles(filePaths []string) ([]io.ReadCloser, []error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchGetFiles", filePaths)
	ret0, _ := ret[0].([]io.ReadCloser)
	ret1, _ := ret[1].([]error)
	return ret0, ret1
}

// BatchGetFiles indicates an expected call of BatchGetFiles.
func (mr *MockbatchHandlerMockRecorder) BatchGetFiles(filePaths interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchGetFiles", reflect.TypeOf((*MockbatchHandler)(nil).BatchGetFiles), filePaths)
}

// BatchSaveFiles mocks base method.
func (m *MockbatchHandler) BatchSaveFiles(files map[string]io.Reader) (map[string]string, []error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchSaveFiles", files)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].([]error)
	return ret0, ret1
}

// BatchSaveFiles indicates an expected call of BatchSaveFiles.
func (mr *MockbatchHandlerMockRecorder) BatchSaveFiles(files interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchSaveFiles", reflect.TypeOf((*MockbatchHandler)(nil).BatchSaveFiles), files)
}

// ConcurrentBatchCleanFiles mocks base method.
func (m *MockbatchHandler) ConcurrentBatchCleanFiles(filePaths []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConcurrentBatchCleanFiles", filePaths)
	ret0, _ := ret[0].(error)
	return ret0
}

// ConcurrentBatchCleanFiles indicates an expected call of ConcurrentBatchCleanFiles.
func (mr *MockbatchHandlerMockRecorder) ConcurrentBatchCleanFiles(filePaths interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConcurrentBatchCleanFiles", reflect.TypeOf((*MockbatchHandler)(nil).ConcurrentBatchCleanFiles), filePaths)
}

// ConcurrentBatchGetFiles mocks base method.
func (m *MockbatchHandler) ConcurrentBatchGetFiles(filePaths []string) ([]io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConcurrentBatchGetFiles", filePaths)
	ret0, _ := ret[0].([]io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConcurrentBatchGetFiles indicates an expected call of ConcurrentBatchGetFiles.
func (mr *MockbatchHandlerMockRecorder) ConcurrentBatchGetFiles(filePaths interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConcurrentBatchGetFiles", reflect.TypeOf((*MockbatchHandler)(nil).ConcurrentBatchGetFiles), filePaths)
}

// ConcurrentBatchSaveFiles mocks base method.
func (m *MockbatchHandler) ConcurrentBatchSaveFiles(files map[string]io.Reader) (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConcurrentBatchSaveFiles", files)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConcurrentBatchSaveFiles indicates an expected call of ConcurrentBatchSaveFiles.
func (mr *MockbatchHandlerMockRecorder) ConcurrentBatchSaveFiles(files interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConcurrentBatchSaveFiles", reflect.TypeOf((*MockbatchHandler)(nil).ConcurrentBatchSaveFiles), files)
}

// MockIFileStorage is a mock of IFileStorage interface.
type MockIFileStorage struct {
	ctrl     *gomock.Controller
	recorder *MockIFileStorageMockRecorder
}

// MockIFileStorageMockRecorder is the mock recorder for MockIFileStorage.
type MockIFileStorageMockRecorder struct {
	mock *MockIFileStorage
}

// NewMockIFileStorage creates a new mock instance.
func NewMockIFileStorage(ctrl *gomock.Controller) *MockIFileStorage {
	mock := &MockIFileStorage{ctrl: ctrl}
	mock.recorder = &MockIFileStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIFileStorage) EXPECT() *MockIFileStorageMockRecorder {
	return m.recorder
}

// BatchCleanFiles mocks base method.
func (m *MockIFileStorage) BatchCleanFiles(filePaths []string) []error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchCleanFiles", filePaths)
	ret0, _ := ret[0].([]error)
	return ret0
}

// BatchCleanFiles indicates an expected call of BatchCleanFiles.
func (mr *MockIFileStorageMockRecorder) BatchCleanFiles(filePaths interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchCleanFiles", reflect.TypeOf((*MockIFileStorage)(nil).BatchCleanFiles), filePaths)
}

// BatchGetFiles mocks base method.
func (m *MockIFileStorage) BatchGetFiles(filePaths []string) ([]io.ReadCloser, []error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchGetFiles", filePaths)
	ret0, _ := ret[0].([]io.ReadCloser)
	ret1, _ := ret[1].([]error)
	return ret0, ret1
}

// BatchGetFiles indicates an expected call of BatchGetFiles.
func (mr *MockIFileStorageMockRecorder) BatchGetFiles(filePaths interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchGetFiles", reflect.TypeOf((*MockIFileStorage)(nil).BatchGetFiles), filePaths)
}

// BatchSaveFiles mocks base method.
func (m *MockIFileStorage) BatchSaveFiles(files map[string]io.Reader) (map[string]string, []error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchSaveFiles", files)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].([]error)
	return ret0, ret1
}

// BatchSaveFiles indicates an expected call of BatchSaveFiles.
func (mr *MockIFileStorageMockRecorder) BatchSaveFiles(files interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchSaveFiles", reflect.TypeOf((*MockIFileStorage)(nil).BatchSaveFiles), files)
}

// CleanFile mocks base method.
func (m *MockIFileStorage) CleanFile(filePath string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CleanFile", filePath)
	ret0, _ := ret[0].(error)
	return ret0
}

// CleanFile indicates an expected call of CleanFile.
func (mr *MockIFileStorageMockRecorder) CleanFile(filePath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CleanFile", reflect.TypeOf((*MockIFileStorage)(nil).CleanFile), filePath)
}

// ConcurrentBatchCleanFiles mocks base method.
func (m *MockIFileStorage) ConcurrentBatchCleanFiles(filePaths []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConcurrentBatchCleanFiles", filePaths)
	ret0, _ := ret[0].(error)
	return ret0
}

// ConcurrentBatchCleanFiles indicates an expected call of ConcurrentBatchCleanFiles.
func (mr *MockIFileStorageMockRecorder) ConcurrentBatchCleanFiles(filePaths interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConcurrentBatchCleanFiles", reflect.TypeOf((*MockIFileStorage)(nil).ConcurrentBatchCleanFiles), filePaths)
}

// ConcurrentBatchGetFiles mocks base method.
func (m *MockIFileStorage) ConcurrentBatchGetFiles(filePaths []string) ([]io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConcurrentBatchGetFiles", filePaths)
	ret0, _ := ret[0].([]io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConcurrentBatchGetFiles indicates an expected call of ConcurrentBatchGetFiles.
func (mr *MockIFileStorageMockRecorder) ConcurrentBatchGetFiles(filePaths interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConcurrentBatchGetFiles", reflect.TypeOf((*MockIFileStorage)(nil).ConcurrentBatchGetFiles), filePaths)
}

// ConcurrentBatchSaveFiles mocks base method.
func (m *MockIFileStorage) ConcurrentBatchSaveFiles(files map[string]io.Reader) (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConcurrentBatchSaveFiles", files)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConcurrentBatchSaveFiles indicates an expected call of ConcurrentBatchSaveFiles.
func (mr *MockIFileStorageMockRecorder) ConcurrentBatchSaveFiles(files interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConcurrentBatchSaveFiles", reflect.TypeOf((*MockIFileStorage)(nil).ConcurrentBatchSaveFiles), files)
}

// GetFile mocks base method.
func (m *MockIFileStorage) GetFile(filePath string) (io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFile", filePath)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFile indicates an expected call of GetFile.
func (mr *MockIFileStorageMockRecorder) GetFile(filePath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFile", reflect.TypeOf((*MockIFileStorage)(nil).GetFile), filePath)
}

// GetFileBytes mocks base method.
func (m *MockIFileStorage) GetFileBytes(filePath string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFileBytes", filePath)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFileBytes indicates an expected call of GetFileBytes.
func (mr *MockIFileStorageMockRecorder) GetFileBytes(filePath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFileBytes", reflect.TypeOf((*MockIFileStorage)(nil).GetFileBytes), filePath)
}

// SaveFile mocks base method.
func (m *MockIFileStorage) SaveFile(file io.Reader, filePath string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveFile", file, filePath)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveFile indicates an expected call of SaveFile.
func (mr *MockIFileStorageMockRecorder) SaveFile(file, filePath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveFile", reflect.TypeOf((*MockIFileStorage)(nil).SaveFile), file, filePath)
}

// SaveFileBytes mocks base method.
func (m *MockIFileStorage) SaveFileBytes(fileBytes []byte, filePath string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveFileBytes", fileBytes, filePath)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveFileBytes indicates an expected call of SaveFileBytes.
func (mr *MockIFileStorageMockRecorder) SaveFileBytes(fileBytes, filePath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveFileBytes", reflect.TypeOf((*MockIFileStorage)(nil).SaveFileBytes), fileBytes, filePath)
}
