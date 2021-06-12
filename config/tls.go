package config

import (
	"astropetal/embeded"
	"crypto/tls"
	"io/fs"
	"log"
)

func mustReadTlsCert() *tls.Certificate {
	certPemBlock, err := fs.ReadFile(embeded.ClientCertFS, "identity.crt")
	if err != nil {
		log.Fatal(err)
	}
	keyPemBlock, err := fs.ReadFile(embeded.ClientCertFS, "identity.key")
	if err != nil {
		log.Fatal(err)
	}
	tlsCert, err := tls.X509KeyPair(certPemBlock, keyPemBlock)
	if err != nil {
		log.Fatal(err)
	}
	return &tlsCert
}
