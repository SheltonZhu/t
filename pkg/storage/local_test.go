package storage

import (
	"bytes"
	"io"
	"os"
	"testing"
	"time"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
)

func Test_localDiskFileStorage_SaveFile(t *testing.T) {
	// 创建本地磁盘存储实例
	storage := NewLocalDiskFileStorage("test")
	fileContent := []byte("This is a test file.")
	fileName := "test.txt"
	t.Run("Save file successfully", func(t *testing.T) {
		isDirExistsStubs := gomonkey.ApplyFuncReturn(isDirExists, true, nil)
		osCreateStubs := gomonkey.ApplyFuncReturn(os.Create, nil, nil)
		ioCopyStubs := gomonkey.ApplyFuncReturn(io.Copy, int64(0), nil)
		defer func() {
			isDirExistsStubs.Reset()
			osCreateStubs.Reset()
			ioCopyStubs.Reset()
		}()
		_, err := storage.SaveFile(bytes.NewReader(fileContent), fileName)
		assert.NoError(t, err, "Failed to save file")
	})

	t.Run("IsDirExists error", func(t *testing.T) {
		isDirExistsStubs := gomonkey.ApplyFuncReturn(isDirExists, false, assert.AnError)
		defer func() {
			isDirExistsStubs.Reset()
		}()
		_, err := storage.SaveFile(bytes.NewReader(fileContent), fileName)
		assert.Error(t, err)
	})

	t.Run("Failed to create file", func(t *testing.T) {
		isDirExistsStubs := gomonkey.ApplyFuncReturn(isDirExists, true, nil)
		osCreateStubs := gomonkey.ApplyFuncReturn(os.Create, nil, assert.AnError)
		defer func() {
			isDirExistsStubs.Reset()
			osCreateStubs.Reset()
		}()

		// 保存文件
		_, err := storage.SaveFile(bytes.NewReader(fileContent), fileName)
		assert.Error(t, err)
	})

	t.Run("Failed to copy file", func(t *testing.T) {
		isDirExistsStubs := gomonkey.ApplyFuncReturn(isDirExists, true, nil)
		osCreateStubs := gomonkey.ApplyFuncReturn(os.Create, nil, nil)
		ioCopyStubs := gomonkey.ApplyFuncReturn(io.Copy, int64(0), assert.AnError)
		defer func() {
			isDirExistsStubs.Reset()
			osCreateStubs.Reset()
			ioCopyStubs.Reset()
		}()
		_, err := storage.SaveFile(bytes.NewReader(fileContent), fileName)
		assert.Error(t, err)
	})

	t.Run("Failed to makedirall", func(t *testing.T) {
		isDirExistsStubs := gomonkey.ApplyFuncReturn(isDirExists, false, nil)
		osMkdirAll := gomonkey.ApplyFuncReturn(os.MkdirAll, assert.AnError)
		defer func() {
			isDirExistsStubs.Reset()
			osMkdirAll.Reset()
		}()
		_, err := storage.SaveFile(bytes.NewReader(fileContent), fileName)
		assert.Error(t, err)
	})
}

func Test_localDiskFileStorage_GetFile(t *testing.T) {
	storage := NewLocalDiskFileStorage("test")
	fileName := "test.txt"
	t.Run("Get file successfully", func(t *testing.T) {
		osOpen := gomonkey.ApplyFuncReturn(os.Open, nil, nil)
		defer osOpen.Reset()
		// 获取文件
		_, err := storage.GetFile(fileName)
		assert.NoError(t, err, "Failed to read file content")
	})

	t.Run("Failed to open file", func(t *testing.T) {
		osOpen := gomonkey.ApplyFuncReturn(os.Open, nil, assert.AnError)
		defer osOpen.Reset()
		// 获取文件
		_, err := storage.GetFile(fileName)
		assert.Error(t, err)
	})
}

func Test_localDiskFileStorage_CleanFile(t *testing.T) {
	storage := NewLocalDiskFileStorage("test")
	fileName := "test.txt"
	t.Run("Clean file successfully", func(t *testing.T) {
		osRemove := gomonkey.ApplyFuncReturn(os.Remove, nil)
		defer osRemove.Reset()
		// 清理文件
		err := storage.CleanFile(fileName)
		assert.NoError(t, err, "Failed to read file")
	})

	t.Run("Failed to remove file", func(t *testing.T) {
		osRemove := gomonkey.ApplyFuncReturn(os.Remove, assert.AnError)
		defer osRemove.Reset()
		// 清理文件
		err := storage.CleanFile(fileName)
		assert.Error(t, err)
	})
}

type FakeFileInfo struct{}

func (f *FakeFileInfo) IsDir() bool        { return true }
func (f *FakeFileInfo) Mode() os.FileMode  { return 0 }
func (f *FakeFileInfo) ModTime() time.Time { return time.Now() }
func (f *FakeFileInfo) Name() string       { return "" }
func (f *FakeFileInfo) Size() int64        { return 0 }
func (f *FakeFileInfo) Sys() interface{}   { return nil }

func Test_isDirExists(t *testing.T) {
	t.Run("Dir existed", func(t *testing.T) {
		ff := &FakeFileInfo{}
		osStatStubs := gomonkey.ApplyFuncReturn(os.Stat, ff, nil)
		isDirStubs := gomonkey.ApplyMethodReturn(ff, "IsDir", true)
		defer func() {
			osStatStubs.Reset()
			isDirStubs.Reset()
		}()
		ok, err := isDirExists("test")
		assert.True(t, ok)
		assert.NoError(t, err)
	})

	t.Run("Dir is not dir", func(t *testing.T) {
		ff := &FakeFileInfo{}
		osStatStubs := gomonkey.ApplyFuncReturn(os.Stat, ff, nil)
		isDirStubs := gomonkey.ApplyMethodReturn(ff, "IsDir", false)
		defer func() {
			osStatStubs.Reset()
			isDirStubs.Reset()
		}()
		ok, err := isDirExists("test")
		assert.False(t, ok)
		assert.NoError(t, err)
	})

	t.Run("Os.stat os.ErrNotExist error", func(t *testing.T) {
		osStatStubs := gomonkey.ApplyFuncReturn(os.Stat, nil, os.ErrNotExist)
		defer osStatStubs.Reset()
		ok, err := isDirExists("test")
		assert.False(t, ok)
		assert.NoError(t, err)
	})

	t.Run("Os.stat error", func(t *testing.T) {
		osStatStubs := gomonkey.ApplyFuncReturn(os.Stat, nil, assert.AnError)
		defer osStatStubs.Reset()
		ok, err := isDirExists("test")
		assert.False(t, ok)
		assert.Error(t, err)
	})
}
