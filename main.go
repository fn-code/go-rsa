/*
 * Genarate rsa keys.
 */

package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/gob"
	"encoding/pem"
	"fmt"
	"log"
	"os"
)

func main() {
	reader := rand.Reader
	//bitSize := 2048
	bitSize := 4096

	key, err := rsa.GenerateKey(reader, bitSize)
	checkError(err)

	publicKey := &key.PublicKey

	saveGobKey("private.key", key)
	savePEMKey("private.pem", key)

	saveGobKey("public.key", publicKey)
	savePublicPEMKey("public.pem", publicKey)
}

func saveGobKey(fileName string, key interface{}) {
	outFile, err := os.Create(fileName)
	checkError(err)
	defer outFile.Close()

	encoder := gob.NewEncoder(outFile)
	err = encoder.Encode(key)
	checkError(err)
}

func savePEMKey(fileName string, key *rsa.PrivateKey) {
	outFile, err := os.Create(fileName)
	checkError(err)
	defer outFile.Close()

	var privateKey = &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}
	err = pem.Encode(outFile, privateKey)
	checkError(err)
}

func savePublicPEMKey(fileName string, pubkey *rsa.PublicKey) {

	dix, err := x509.MarshalPKIXPublicKey(pubkey)
	if err != nil {
		log.Println("1 ", err)
		os.Exit(1)
	}

	var pemkey = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: dix,
	}

	pemfile, err := os.Create(fileName)
	checkError(err)
	defer pemfile.Close()

	err = pem.Encode(pemfile, pemkey)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}