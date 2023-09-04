package encryption

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBcryptEncryptor(t *testing.T) {
	t.Parallel()
	bcryptEncryptor := NewBcryptEncryptor(WithCost(10))

	t.Run("Encode and Verify", func(t *testing.T) {
		plainPwd := "password123"
		hashedPwd, err := bcryptEncryptor.Encode(plainPwd)
		assert.NoError(t, err)
		assert.NotEmpty(t, hashedPwd)

		ok := bcryptEncryptor.Verify(hashedPwd, plainPwd)
		assert.True(t, ok)

		plainPwd = "d8c816180dfa507cb880a867243fe7fa"
		hashedPwd = "$2b$09$6jnvSyEiquxeiB/aooX4PedfQeCEhLmRPEHiXFvTET81X9J4hu4XC"
		ok = bcryptEncryptor.Verify(hashedPwd, plainPwd)
		assert.True(t, ok)
	})

	t.Run("Verify with wrong password", func(t *testing.T) {
		hashedPwd := "$2a$10$T/4xM5V5mHfNjNczbp8p6uXQZz2zJH8m9ZL6x2fXNQWqHj/ji1ZGm"
		ok := bcryptEncryptor.Verify(hashedPwd, "wrongpassword")
		assert.False(t, ok)
	})

	t.Run("Encode with wrong cost", func(t *testing.T) {
		bcryptEncryptor.Cost = 10000000000000
		plainPwd := "password123"
		_, err := bcryptEncryptor.Encode(plainPwd)
		assert.Error(t, err)
	})
}
