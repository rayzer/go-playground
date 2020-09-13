package main

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"

	"golang.org/x/crypto/ssh"
)

func generateKeyPair() (pair map[string][]byte, err error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return
	}

	privateKeyPEM := &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)}
	var private bytes.Buffer
	if err = pem.Encode(&private, privateKeyPEM); err != nil {
		return
	}

	pair = make(map[string][]byte, 2)
	pub, err := ssh.NewPublicKey(&privateKey.PublicKey)
	if err != nil {
		return pair, err
	}

	public := ssh.MarshalAuthorizedKey(pub)
	pair["public"] = public
	pair["private"] = private.Bytes()
	return pair, nil
}

func main() {
	data, err := generateKeyPair()
	if err != nil {
		log.Println("cannot generate key pair")
		return
	}
	fmt.Println(data["private"])
	fmt.Println(data["public"])
}
