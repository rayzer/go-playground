package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

func main() {
	waitCertificateSeconds := 10
	for {
		mountedKeyFile, _ := ioutil.ReadFile("public")
		if len(mountedKeyFile) > 0 {
			log.Printf("valid ssh key file mounted: %v", len(mountedKeyFile))
			return
		}

		log.Printf("no ssh key file ready, wait...")
		time.Sleep(5 * time.Second)
		waitCertificateSeconds = waitCertificateSeconds - 2
		if waitCertificateSeconds <= 0 {
			fmt.Printf("wait ssh key pair timeout")
			return
		}
	}
}
