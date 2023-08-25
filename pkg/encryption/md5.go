package encryption

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"strings"
)

type RandomSaltGenerator interface {
	GetRandSalt() string
}

type RandSaltGen struct{}

func (r *RandSaltGen) GetRandSalt() string {
	unencodedSalt := make([]byte, 16)
	_, _ = io.ReadFull(rand.Reader, unencodedSalt)
	sEnc := base64.StdEncoding.EncodeToString(unencodedSalt)
	return sEnc[:len(sEnc)-2]
}

// MD5Option 选项配置
type MD5Option func(*md5Encryptor)

// WithRandSaltGen 设置随机盐生成器
func WithRandSaltGen(gen RandomSaltGenerator) MD5Option {
	return func(e *md5Encryptor) {
		e.RandomSaltGenerator = gen
	}
}

func WithConstSalt(salt string) MD5Option {
	return func(e *md5Encryptor) {
		e.constSalt = salt
	}
}

type md5Encryptor struct {
	RandomSaltGenerator
	constSalt string
	randSalt  string
}

// NewMD5Encryptor
func NewMD5Encryptor(opts ...MD5Option) *md5Encryptor {
	defaultConfig := &md5Encryptor{
		RandomSaltGenerator: &RandSaltGen{},
	}
	for _, o := range opts {
		o(defaultConfig)
	}
	return defaultConfig
}

// Encode 实现了 Encryptor 接口的 Encode 方法
func (e *md5Encryptor) Encode(plainPwd string) (string, error) {
	if e.randSalt == "" {
		e.randSalt = e.GetRandSalt()
		defer func() {
			e.randSalt = ""
		}()
	}
	data := []byte(plainPwd + e.constSalt + e.randSalt)
	h := md5.New()
	if _, err := h.Write(data); err != nil {
		return "", err
	}
	hash := hex.EncodeToString(h.Sum(nil))
	return fmt.Sprintf("$md5$%s$%s", e.randSalt, hash), nil

}

// Verify 实现了 Encryptor 接口的 Verify 方法
func (e *md5Encryptor) Verify(hashedPwd string, plainPwd string) bool {
	// 获取盐
	pairs := strings.Split(hashedPwd, "$")
	if len(pairs) < 3 {
		return false
	}
	e.randSalt = pairs[2]
	defer func() {
		e.randSalt = ""
	}()
	hash, err := e.Encode(plainPwd)
	if err != nil {
		return false
	}
	return hash == hashedPwd
}

var _ Encryptor = (*md5Encryptor)(nil)
