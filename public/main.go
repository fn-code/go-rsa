package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./public.pem")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	buf, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	block, _ := pem.Decode(buf)
	key, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	pubKey := key.(*rsa.PublicKey)
	log.Println(pubKey.N)


}