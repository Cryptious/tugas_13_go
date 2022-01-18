package main

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
)

func main() {
	// inisiasi variabel nama
	var name = "Martino Crypt Glorious Wagey"
	fmt.Println("nama : ", name)

	// Encoding base 64 variabel name
	var encodeString = base64.StdEncoding.EncodeToString([]byte(name))
	fmt.Println("Encoding String :", encodeString)

	// proses enkripsi Sha-1
	var sha = sha1.New()
	sha.Write([]byte(encodeString))
	var enkripsi = sha.Sum(nil)
	var stringenkripsi = fmt.Sprintf("%x", enkripsi)
	fmt.Println("Enkripsi SHA : ", stringenkripsi)

}
