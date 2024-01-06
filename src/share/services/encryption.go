package services

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"io"
	"log"

	"github.com/reddit-clone/src/share/config"
)

type EncryptionService struct {
	cfg *config.Config
}

func NewEncryptionService(cfg *config.Config) *EncryptionService {
	return &EncryptionService{
		cfg,
	}
}
func (s *EncryptionService) EncryptString(text string) (*[]byte, error) {
	key := []byte(s.cfg.JWT.Secret) // TODO: this can be changed in future
	plaintext := []byte(text)

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Panic(err)
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		log.Panic(err)
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		log.Panic(err)
	}

	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)

	return &ciphertext, nil
}

func (s *EncryptionService) DeEncryptString(encryptedString string) (string, error) {
	key := []byte(s.cfg.JWT.Secret) // TODO: this can be changed in future
	enc, err := hex.DecodeString(encryptedString)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := aesGCM.NonceSize()
	if len(enc) < nonceSize {
		return "", err
	}

	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		log.Panic(err)
	}
	return string(plaintext), nil
}
