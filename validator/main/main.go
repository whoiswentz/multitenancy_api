package main

import (
	"fmt"
	"mongodb-test/validator"
)

func main() {
	v := validator.New()

	e := v.Encrypt([]byte("asdasdasdassd"))
	fmt.Printf("e = %s\n", e)

	d := v.Decrypt(e)
	fmt.Printf("d = %", d)
}
