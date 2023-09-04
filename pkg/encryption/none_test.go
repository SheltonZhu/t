package encryption

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNoneEncryptor(t *testing.T) {
	t.Parallel()
	noneEncryptor := NewNoneEncryptor()

	t.Run("Encode and Verify", func(t *testing.T) {
		plainPwd := "password123"
		hashedPwd, err := noneEncryptor.Encode(plainPwd)
		assert.NoError(t, err)
		assert.Equal(t, plainPwd, hashedPwd)

		ok := noneEncryptor.Verify(hashedPwd, plainPwd)
		assert.True(t, ok)
	})

	t.Run("Verify with wrong password", func(t *testing.T) {
		hashedPwd := "password123"
		ok := noneEncryptor.Verify(hashedPwd, "wrongpassword")
		assert.False(t, ok)
	})
}
