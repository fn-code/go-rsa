package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./private.pem")
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
	key, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
	fmt.Println(key.D.Bytes())


}