package main

import (
	"crypto/tls"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetCert() tls.Certificate {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	SSL_CERT := os.Getenv("SSL_CERT")
	SSL_KEY := os.Getenv("SSL_KEY")

	cert, err := tls.LoadX509KeyPair(SSL_CERT, SSL_KEY)
	if err != nil {
		log.Fatal(err)
	}
	return cert
}

func GetTLSConfig(cert tls.Certificate) tls.Config {
	config := tls.Config{
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: false,
	}
	return config
}

func GetTLSConn() *tls.Conn {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	IRC_SERVER := os.Getenv("IRC_SERVER")
	cert := GetCert()                                 // Load certs
	config := GetTLSConfig(cert)                      // Create TLS configuration
	conn, err := tls.Dial("tcp", IRC_SERVER, &config) // Connect TLS connection
	if err != nil {
		log.Fatalf("Client: dial: %s", err)
	}
	return conn
}
