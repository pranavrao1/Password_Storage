package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"

	"golang.org/x/crypto/scrypt"
)

const (
	saltSize  = 16
	nonceSize = 12
	keySize   = 32
	scryptN   = 1 << 15
	scryptR   = 8
	scryptP   = 1
)

var vaultMagic = []byte("PWVAULT\x01")

func deriveKey(password string, salt []byte) ([]byte, error) {
	return scrypt.Key([]byte(password), salt, scryptN, scryptR, scryptP, keySize)
}

func encryptVault(plaintext []byte, password string) ([]byte, error) {
	salt := make([]byte, saltSize)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		return nil, err
	}
	key, err := deriveKey(password, salt)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, nonceSize)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	ct := gcm.Seal(nil, nonce, plaintext, nil)

	out := make([]byte, 0, len(vaultMagic)+saltSize+nonceSize+len(ct))
	out = append(out, vaultMagic...)
	out = append(out, salt...)
	out = append(out, nonce...)
	out = append(out, ct...)
	return out, nil
}

func decryptVault(data []byte, password string) ([]byte, error) {
	if len(data) < len(vaultMagic)+saltSize+nonceSize {
		return nil, errors.New("ciphertext too short")
	}
	if !bytes.Equal(data[:len(vaultMagic)], vaultMagic) {
		return nil, errors.New("not an encrypted vault")
	}
	rest := data[len(vaultMagic):]
	salt, rest := rest[:saltSize], rest[saltSize:]
	nonce, ct := rest[:nonceSize], rest[nonceSize:]

	key, err := deriveKey(password, salt)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	return gcm.Open(nil, nonce, ct, nil)
}

func isEncryptedVault(data []byte) bool {
	return len(data) >= len(vaultMagic) && bytes.Equal(data[:len(vaultMagic)], vaultMagic)
}
