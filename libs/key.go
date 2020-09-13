package main

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/gob"
	"encoding/pem"
	"fmt"
	"log"
	"os"
	"reflect"

	"golang.org/x/crypto/ssh"
)

func main() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return
	}

	privateKeyPEM := &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)}
	var private bytes.Buffer
	if err = pem.Encode(&private, privateKeyPEM); err != nil {
		return
	}

	pair := make(map[string][]byte, 2)
	pub, err := ssh.NewPublicKey(&privateKey.PublicKey)
	if err != nil {
		panic(err)
	}

	public := ssh.MarshalAuthorizedKey(pub)
	pair["public"] = public
	pair["private"] = private.Bytes()

	b := new(bytes.Buffer)
	e := gob.NewEncoder(b)
	// Encoding the map
	err = e.Encode(pair)
	if err != nil {
		panic(err)
	}

	outFile, err := os.Create("encoded")
	defer outFile.Close()
	wrote, _ := outFile.Write(b.Bytes())
	fmt.Printf("%v bytes wrote\n", wrote)

	encodedFile, _ := os.OpenFile("encoded", os.O_RDONLY, 0755)
	info, _ := encodedFile.Stat()
	log.Printf("received %d bytes from %s", info.Size(), "encoded")

	var decodedMap map[string][]byte
	d := gob.NewDecoder(encodedFile)
	err = d.Decode(&decodedMap)
	if err != nil {
		panic(err)
	}

	// fmt.Printf("%v\n", decodedMap)
	// fmt.Println(decodedMap["public"])
	fmt.Println(reflect.DeepEqual(decodedMap["public"], pair["public"]))
}
