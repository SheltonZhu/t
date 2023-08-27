package encryption

type noneEncryptor struct{}

// NewNoneEncryptor
func NewNoneEncryptor() *noneEncryptor {
	return &noneEncryptor{}
}

// Encode 实现了 Encryptor 接口的 Encode 方法
func (e *noneEncryptor) Encode(plainPwd string) (string, error) {
	return plainPwd, nil
}

// Verify 实现了 Encryptor 接口的 Verify 方法
func (e *noneEncryptor) Verify(hashedPwd string, plainPwd string) bool {
	hash, _ := e.Encode(plainPwd)
	return hash == hashedPwd
}

var _ Encryptor = (*noneEncryptor)(nil)
