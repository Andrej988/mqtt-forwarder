package main

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"os"
	"path/filepath"
)

func NewTlsConfig(caCert string, clientCert string, clientKey string) *tls.Config {
	certDir := filepath.Join(".", "certs")
	caCertPath := filepath.Join(certDir, caCert)
	clientCertPath := filepath.Join(certDir, clientCert)
	clientKeyPath := filepath.Join(certDir, clientKey)

	certpool := x509.NewCertPool()
	ca, err := os.ReadFile(caCertPath)
	if err != nil {
		log.Fatalln(err.Error())
	}
	certpool.AppendCertsFromPEM(ca)
	clientKeyPair, err := tls.LoadX509KeyPair(clientCertPath, clientKeyPath)
	if err != nil {
		panic(err)
	}
	return &tls.Config{
		RootCAs:            certpool,
		ClientAuth:         tls.RequireAndVerifyClientCert,
		ClientCAs:          nil,
		InsecureSkipVerify: true,
		MinVersion:         tls.VersionTLS13,
		Certificates:       []tls.Certificate{clientKeyPair},
	}
}