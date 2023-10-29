package main

import (
	"crypto/rand"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

func RunCli() error {
	rootCmd := &cobra.Command{
		Use:   "encrypt",
		Short: "A CLI application to encrypt files",
		Long: `A command-line interface (CLI) program designed for the purpose of file encryption.
This is how to use the program. 

./cli-app "fileToEncrypt.ext" "key.txt" "encryptedFileName.bin"

1. The first argument is the file you wanna encrypt. Note that the file should be in the "src" folder initially

2. The second argument is the key you wanna use for encryption, it can be aes-128 bit(16 bytes), aes-192 bit(24 bytes) or aes-256 bit(32bytes). All this mean is that the lenght of the content in the key file can be 16, 24 or 32. And the key file must be in the "src" folder also. 

3. The third argument is the whatever name you wanna give your encrypted file with it ".bin" extension preferably, since it's a binary content. 

Note: All file should be with their extension, e.g "myself.txt" not "myself"

		`,
		Args: cobra.ExactArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			fileName := args[0]
			keyForEncryption := args[1]
			cipherFileName := args[2]
			plainTxtByte, err := ReadFile(fileName)

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			key, err := ReadFile(keyForEncryption)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			block, err := CreateBlockCipher(key)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			gcm, err := CreateGCMCipher(block)
			if err != nil {
				fmt.Println(err)
			}
			nonce := make([]byte, gcm.NonceSize())
			if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
				fmt.Println(err)
			}

			cipherTxt := gcm.Seal(nonce, nonce, plainTxtByte, nil)
			cipherPath := fmt.Sprintf("encrypted/%s", cipherFileName)
			if err := os.WriteFile(cipherPath, cipherTxt, 0644); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}

	if err := rootCmd.Execute(); err != nil {
		return err
	}
	return nil
}
