package main

import (
	"bufio"
	"os"
	"strings"
	"fmt"
)

func xorEncryptDecrypt(text string, key []byte) []byte {
	textBytes := []byte(text)
	cipherBytes := []byte{}

	for i := range textBytes {
		cipherBytes = append(cipherBytes, textBytes[i]^key[i%len(key)])
	}

	return cipherBytes
}

func main() {
	Reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a sentence: ")
	plaintext, _ := Reader.ReadString('\n')
	plaintext = strings.TrimSpace(plaintext)

	fmt.Print("Enter your key: ")
	key, _ := Reader.ReadString('\n')
	key = strings.TrimSpace(key)

	keyByte := []byte(key)
	ciphertext := xorEncryptDecrypt(plaintext, keyByte)
	cipher := string(ciphertext)
	fmt.Println("Encrypted/Decrypted text: ", cipher)

}
