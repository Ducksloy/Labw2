package main

import (
	"bufio"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Enter a sentence: ")
	Reader := bufio.NewReader(os.Stdin)
	plaintext, _ := Reader.ReadString('\n')
	plaintext = strings.TrimSpace(plaintext)
	fmt.Println("Plaintext: " + plaintext)

	bin := []byte(plaintext)
	con := ""
	for i := range bin {
		con += fmt.Sprintf("%08b", bin[i])
	}
	fmt.Println("Binary: ", con)

	fmt.Println("Hex: " + hex.EncodeToString(bin))

	fmt.Println("Base64: " + base64.StdEncoding.EncodeToString(bin))
}