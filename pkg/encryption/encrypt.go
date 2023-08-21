package encryption

type Encryptor interface {
	Encode(plainPwd string) (string, error)
	Verify(hashedPwd, plainPwd string) bool
}
