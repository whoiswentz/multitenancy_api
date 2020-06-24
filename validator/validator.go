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
)

type Validator struct {
	nonce []byte
}

func New() *Validator {
	return &Validator{
		nonce: make([]byte, 12),
	}
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

	nonce := v.nonce
	fmt.Printf("Nonce E = %x\n", v.nonce)

	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		log.Fatalln(err.Error())
	}

	aesgcm, err := cipher.NewGCM(b)
	if err != nil {
		log.Fatalln(err.Error())
	}

	e := aesgcm.Seal(nil, nonce, data, nil)

	return e
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

	nonce, _ := hex.DecodeString(fmt.Sprintf("%x", v.nonce))
	fmt.Printf("Nonce D = %x\n", v.nonce)
	p, err := aesgcm.Open(nil, nonce, data, nil)
	if err != nil {
		panic(err.Error())
	}
	return p
}
