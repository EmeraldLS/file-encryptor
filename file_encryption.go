package main

import (
	"crypto/aes"
	"crypto/cipher"
	"os"
)

func ReadFile(fileName string) (fileByte []byte, err error) {
	fileByte, err = os.ReadFile("src/" + fileName)
	if err != nil {
		return []byte{}, err
	}
	return fileByte, nil
}

func CreateBlockCipher(key []byte) (cipher.Block, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	return block, nil
}

func CreateGCMCipher(block cipher.Block) (cipher.AEAD, error) {
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	return gcm, nil
}
