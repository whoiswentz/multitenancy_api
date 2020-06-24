package main

import (
	"fmt"
	"mongodb-test/validator"
)

func main() {
	v := validator.New()

	b := v.Encrypt([]byte("asdasdasd"))
	fmt.Printf("Ciphertext = %x\n", b)

	p := v.Decrypt([]byte(fmt.Sprintf("%x", b)))
	fmt.Printf("Decrypted Plaintext = %s\n", p)

}
