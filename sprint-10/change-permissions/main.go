package main

import (
	"log"
	"os"
)

func main() {
	fileName := "/home/margo/scripts/readme.txt"
	err := os.Chmod(fileName, 0777)
	if err != nil {
		log.Fatal(err)
	}
	fi, err := os.Stat(fileName)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Права доступа: 0%o\n", fi.Mode().Perm())
}
