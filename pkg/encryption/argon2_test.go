package encryption

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArgon2Encryptor(t *testing.T) {
	t.Parallel()
	argon2Encryptor := NewArgon2Encryptor(func(ae *argon2Encryptor) {})

	t.Run("Encode and Verify", func(t *testing.T) {
		plainPwd := "password123"
		hashedPwd, err := argon2Encryptor.Encode(plainPwd)
		assert.NoError(t, err)
		assert.NotEmpty(t, hashedPwd)

		ok := argon2Encryptor.Verify(hashedPwd, plainPwd)
		assert.True(t, ok)
	})

	t.Run("Verify with wrong password", func(t *testing.T) {
		hashedPwd := "password123"
		ok := argon2Encryptor.Verify(hashedPwd, "wrongpassword")
		assert.False(t, ok)
	})
}

func TestXxx(t *testing.T) {
	e := NewArgon2Encryptor()
	hash, _ := e.Encode("123123123123")
	_ = hash
	ok := e.Verify(
		"$argon2i$v=19$m=12,t=2,p=1$/bJp+TDWOO0KdyriRkILvdkm6GoGdl8t5RqDreWAX3s$DznndClCzojpMe8A1Egv6GejMpFJDiO0GVF4gL2L1wA",
		"12312",
	)
	_ = ok
}
