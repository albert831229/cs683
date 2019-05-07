package main

import (
	"bytes"
	"crypto/sha1"
	"file_encryption/utils"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"golang.org/x/crypto/pbkdf2"
)

var r *rand.Rand
var myDir = os.Getenv("HOME") + "/secret"
var dk []byte

func GetSalt(length int) []byte {
	const chars = "abcdef0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = chars[r.Intn(len(chars))]
	}
	return result
}

func WriteKey(dk []byte) {
	f, err := os.Create("key.txt")
	if err != nil {
		panic(err.Error())
	}

	_, err = f.Write(dk)
	if err != nil {
		f.Close()
		panic(err.Error())
	}

	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func ProcessFunc(pathX string, infoX os.FileInfo, errX error) error {

	// first thing to do, check error. and decide what to do about it
	if errX != nil {
		fmt.Printf("error 「%v」 at a path 「%q」\n", errX, pathX)
		return errX
	}

	// find out if it's a dir or file, if file, print info
	if infoX.IsDir() {
		fmt.Printf("is dir.\n")
	} else {
		path := filepath.Dir(pathX) + "/" + infoX.Name()
		fmt.Printf("Encrypting 「%v」\n", path)
		utils.EncryptFile(dk, path)
	}

	return nil
}

func main() {

	if len(os.Args) < 3 {
		panic("Invalid arguments.")
	}
	if os.Args[1] == "encrypt" {
		r = rand.New(rand.NewSource(time.Now().UnixNano()))
		pwd := []byte("I need 16 char")
		salt := GetSalt(8)
		dk := pbkdf2.Key(pwd, salt, 4096, 16, sha1.New)
		//fmt.Println(dk)
		WriteKey(dk)
		utils.EncryptFile(dk, os.Args[2])
	} else if os.Args[1] == "decrypt" {
		if len(os.Args) < 4 {
			panic("Invalid arguments.")
		}
		key, err := ioutil.ReadFile(os.Args[2])
		if err != nil {
			panic("Unable to read key.")
		}

		expect, _ := ioutil.ReadFile("right_text.txt")
		decryptedText := utils.DecryptFile(key, os.Args[3])

		fmt.Println(bytes.Equal(expect, decryptedText))
	}

	/**r = rand.New(rand.NewSource(time.Now().UnixNano()))
	pwd := []byte("I need 16 char")
	salt := GetSalt(8)
	dk = pbkdf2.Key(pwd, salt, 4096, 16, sha1.New)
	WriteKey(dk)

	err := filepath.Walk(myDir, ProcessFunc)
	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", myDir, err)
	}**/
}
