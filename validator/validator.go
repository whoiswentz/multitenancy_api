package validator

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Validator struct {
	nonce []byte
}

func New() *Validator {
	return &Validator{}
}

func (v *Validator) Encrypt(data []byte) []byte {
	base64Key := os.Getenv("key")

	k, err := base64.StdEncoding.DecodeString(base64Key)
	if err != nil {
		log.Fatalln(err.Error())
	}

	if _, err := io.ReadFull(rand.Reader, k); err != nil {
		log.Fatalln(err.Error())
	}

	b, err := aes.NewCipher(k)
	if err != nil {
		log.Fatalln(err.Error())
	}

	nonce := make([]byte, 12)

	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		log.Fatalln(err.Error())
	}

	aesgcm, err := cipher.NewGCM(b)
	if err != nil {
		log.Fatalln(err.Error())
	}

	e := aesgcm.Seal(nil, nonce, data, nil)
	token := fmt.Sprintf("%x:%x", e, nonce)

	return []byte(token)
}

func (v *Validator) Decrypt(data []byte) []byte {
	base64Key := os.Getenv("key")

	k, err := base64.StdEncoding.DecodeString(base64Key)
	if err != nil {
		log.Fatalln(err.Error())
	}

	if _, err := io.ReadFull(rand.Reader, k); err != nil {
		log.Fatalln(err.Error())
	}

	b, err := aes.NewCipher(k)
	if err != nil {
		log.Fatalln(err.Error())
	}

	aesgcm, err := cipher.NewGCM(b)
	if err != nil {
		log.Fatalln(err.Error())
	}

	s := strings.Split(fmt.Sprintf("%s", data), ":")

	token, _ := hex.DecodeString(s[0])
	nonce, _ := hex.DecodeString(s[1])

	p, err := aesgcm.Open(nil, nonce, token, nil)
	if err != nil {
		panic(err.Error())
	}
	return p
}
