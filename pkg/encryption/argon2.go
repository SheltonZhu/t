package encryption

import (
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
		Iterations:  2,
		Memory:      12,
		Parallelism: 1,
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
	ok, _ := argon2.VerifyEncoded(hashedPwd, []byte(plainPwd))
	return ok
}

func (e *argon2Encryptor) encodeWithSalt(salt string, plainPwd string) (string, error) {
	return argon2.HashEncoded(e.Ctx, []byte(plainPwd), []byte(salt))
}

var _ Encryptor = (*argon2Encryptor)(nil)
