package encryption

import (
	"fmt"
	"github.com/tvdburgt/go-argon2"
	"strings"
)

type argon2Encryptor struct {
	RandomSaltGenerator
	Ctx      *argon2.Context
	randSalt string
}

// Argon2Option 选项配置
type Argon2Option func(*argon2Encryptor)

// NewArgon2Encryptor
func NewArgon2Encryptor(opts ...Argon2Option) *argon2Encryptor {
	ctx := &argon2.Context{
		Iterations:  5,
		Memory:      1 << 16,
		Parallelism: 2,
		HashLen:     32,
		Mode:        argon2.ModeArgon2i,
		Version:     argon2.Version13,
	}
	defaultConfig := &argon2Encryptor{
		RandomSaltGenerator: &RandSaltGen{},
		Ctx:                 ctx,
	}
	for _, o := range opts {
		o(defaultConfig)
	}
	return defaultConfig
}

// Encode 实现了 Encryptor 接口的 Encode 方法
func (e *argon2Encryptor) Encode(plainPwd string) (string, error) {
	if e.randSalt == "" {
		e.randSalt = e.GetRandSalt()
	}
	hash, err := argon2.HashEncoded(e.Ctx, []byte(plainPwd), []byte(e.randSalt))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s$%s", hash, e.randSalt), nil
}

// Verify 实现了 Encryptor 接口的 Verify 方法
func (e *argon2Encryptor) Verify(hashedPwd string, plainPwd string) bool {
	// 获取盐
	pairs := strings.Split(hashedPwd, "$")
	if len(pairs) != 7 {
		return false
	}
	e.randSalt = pairs[len(pairs)-1]
	hash, err := e.Encode(plainPwd)
	if err != nil {
		return false
	}
	e.randSalt = ""
	return hash == hashedPwd
}

var _ Encryptor = (*argon2Encryptor)(nil)
