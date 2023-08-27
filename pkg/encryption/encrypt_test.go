package encryption

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBcryptEncryptor(t *testing.T) {
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

func TestNoneEncryptor(t *testing.T) {
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

func TestMD5Encryptor(t *testing.T) {
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

func TestArgon2Encryptor(t *testing.T) {
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
