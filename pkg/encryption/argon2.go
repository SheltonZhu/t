package encryption

import (
	"fmt"
	"strings"

	"github.com/tvdburgt/go-argon2"
)

// Argon2Option 选项配置
type Argon2Option func(*argon2Encryptor)

type argon2Encryptor struct {
	RandomSaltGenerator
	Ctx *argon2.Context
}

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
	return e.encodeWithSalt(e.GetRandSalt(), plainPwd)
}

// Verify 实现了 Encryptor 接口的 Verify 方法
func (e *argon2Encryptor) Verify(hashedPwd string, plainPwd string) bool {
	// 获取盐
	pairs := strings.Split(hashedPwd, "$")
	if len(pairs) != 7 {
		return false
	}
	randSalt := pairs[len(pairs)-1]
	hash, err := e.encodeWithSalt(randSalt, plainPwd)
	if err != nil {
		return false
	}
	return hash == hashedPwd
}

func (e *argon2Encryptor) encodeWithSalt(salt string, plainPwd string) (string, error) {
	hash, err := argon2.HashEncoded(e.Ctx, []byte(plainPwd), []byte(salt))
	return fmt.Sprintf("%s$%s", hash, salt), err
}

var _ Encryptor = (*argon2Encryptor)(nil)
