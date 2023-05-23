package security

import (
	"fmt"
	"testing"
)

func Test_StringEncryption(t *testing.T) {
	// output -> stdout in hex format
	var (
		secret    = "XXX"
		plaintext = "XXX"
	)

	e := encryptString(plaintext, hashSecretKey([]byte(secret)))
	fmt.Printf("%x\n", e)
}

func Test_StringDecryption(t *testing.T) {
	// output -> stdout
	var (
		secret    = "XXX"
		encrypted = "XXX"
	)

	e := decryptString(encrypted, hashSecretKey([]byte(secret)))
	fmt.Printf("%s\n", e)
}
