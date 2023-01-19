package keys

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
)

func CheckKeys() {
	// Loading the private and public keys from the respective .pem files
	privateKey := LoadPrivateKey()
	publicKey := LoadPublicKey()

	// Checking if the private key belongs to the public key
	if (privateKey.PublicKey.Equal(publicKey)) {
		fmt.Println("VALID: The private key BELONGS to the public key.")
	} else {
		fmt.Println("INVALID: The private key DOES NOT BELONG to the public key.")
	}
}

func LoadPrivateKey() *rsa.PrivateKey {
	// Load the private key inside the private.pem file
	bs, err := ioutil.ReadFile("./keys/private.pem");
	if err != nil {
		log.Fatalln("Error in loading the private.pem file")
	}

	// Parse the PEM-encoded data
	privateKeyPEMDecoded, _ := pem.Decode(bs)
	if privateKeyPEMDecoded == nil {
		log.Fatalln("Failed to parse PEM block containing the private key")
	}

	// Extract the private key from the PEM block
	privateKey, err := x509.ParsePKCS1PrivateKey(privateKeyPEMDecoded.Bytes)
	if err != nil {
		log.Fatalln("Error occurred in extracting the private key from the PEM block", err);
	}

	return privateKey;
}

func LoadPublicKey() *rsa.PublicKey {
	// Load the public key inside the private.pem file
	bs, err := ioutil.ReadFile("./keys/public.pem")
	if err != nil {
		log.Fatalln("Error in loading the public.pem file")
	}

	// Parse the PEM-encoded data
	publicKeyPEMDecoded, _ := pem.Decode(bs)
	if publicKeyPEMDecoded == nil {
		log.Fatalln("Failed to parse PEM block containing the private key")
	}

	// Extract the public key from the PEM block
	publicKey, err := x509.ParsePKIXPublicKey(publicKeyPEMDecoded.Bytes)
	if err != nil {
		log.Fatalln("Error occurred in extracting the private key from the PEM block", err);
	}

	return publicKey.(*rsa.PublicKey)
}