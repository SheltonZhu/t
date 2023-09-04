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
