package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func CheckDocument(document []byte, hash string) bool {
	documentHash := sha256.Sum256(document)
	documentHashString := hex.EncodeToString(documentHash[:])

	return documentHashString == hash
}

func main() {
	document := []byte("Hi dear, whats up?")
	hashString := "7df984fc8454b1789f1ad2561690872877bee8fd6a2c776d4e7082ee924b5b0e"
	if CheckDocument(document, hashString) {
		fmt.Println("сообщение прошло проверку")
	} else {
		fmt.Println("сообщение не прошло проверку")
	}
}
