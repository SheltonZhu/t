package encryption

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMD5Encryptor(t *testing.T) {
	t.Parallel()
	md5Encryptor := NewMD5Encryptor(WithConstSalt("&^@*("), WithRandSaltGen(&RandSaltGen{}))

	t.Run("Encode and Verify", func(t *testing.T) {
		plainPwd := "password123"
		hashedPwd, err := md5Encryptor.Encode(plainPwd)
		assert.NoError(t, err)
		assert.NotEmpty(t, hashedPwd)

		ok := md5Encryptor.Verify(hashedPwd, plainPwd)
		assert.True(t, ok)

		plainPwd = "d8c816180dfa507cb880a867243fe7fa"
		hashedPwd = "$md5$mqyl8HvsUsB7gBeb$3961cbf15b03e0946e7a364145611576"
		ok = md5Encryptor.Verify(hashedPwd, plainPwd)
		assert.True(t, ok)
	})

	t.Run("Verify with wrong password", func(t *testing.T) {
		hashedPwd := "password123"
		ok := md5Encryptor.Verify(hashedPwd, "wrongpassword")
		assert.False(t, ok)
	})
}
