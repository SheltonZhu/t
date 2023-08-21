package encryption

import (
	"golang.org/x/crypto/bcrypt"
)

// BcryptOption 选项配置
type BcryptOption func(*bcryptEncryptor)

// WithCost 设置 cost
func WithCost(cost int) BcryptOption {
	return func(be *bcryptEncryptor) {
		be.Cost = cost
	}
}

type bcryptEncryptor struct {
	Cost int
}

// NewBcryptEncryptor
func NewBcryptEncryptor(opts ...BcryptOption) *bcryptEncryptor {
	defaultConfig := &bcryptEncryptor{
		Cost: bcrypt.DefaultCost,
	}
	for _, o := range opts {
		o(defaultConfig)
	}
	return defaultConfig
}

// Encode 实现了 Encryptor 接口的 Encode 方法
func (b *bcryptEncryptor) Encode(plainPwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(plainPwd), b.Cost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// Verify 实现了 Encryptor 接口的 Verify 方法
func (b *bcryptEncryptor) Verify(hashedPwd string, plainPwd string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd)); err != nil {
		return false
	}
	return true
}

var _ Encryptor = (*bcryptEncryptor)(nil)
