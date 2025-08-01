package main

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
)

// Load public key from file
func loadPublicKey(path string) (*rsa.PublicKey, error) {
	keyData, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read public key file: %v", err)
	}

	// Parse PEM block
	block, _ := pem.Decode(keyData)
	if block == nil {
		return nil, fmt.Errorf("failed to parse PEM block")
	}

	// Parse public key
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		// Try another format
		cert, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			return nil, fmt.Errorf("failed to parse public key: %v", err)
		}
		pub = cert.PublicKey
	}

	rsaPub, ok := pub.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("not an RSA public key")
	}

	return rsaPub, nil
}

// Verify signature
func verifySignature(message []byte, signature []byte, publicKey *rsa.PublicKey) error {
	// Calculate message hash
	hash := sha256.Sum256(message)

	// Verify signature
	err := rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hash[:], signature)
	if err != nil {
		return fmt.Errorf("signature verification failed: %v", err)
	}

	return nil
}

func main() {
	// Load public key
	publicKey, err := loadPublicKey("public_key.pem")
	if err != nil {
		log.Fatal(err)
	}

	message := []byte(`{"account":"0T02","agent":"X21006UAT","amount":1,"bet_id":"","feature_buy":"","game_id":"4009","owner_account":"X20001UAT","token":"d01e14944dcd4a0e9e207f4438bb37da","trans_id":"688b9ce601d64b0001f9ee15","valid_amount":0}`)
	// Convert sigBase64 to binary format
	sigBase64 := "Kg5aEcTdxcI772NZItxg6ZN27nGp4xLxxQqOsbONILxnrA/vFtqZKRnxrIp+/QkvYQR0Fc7uNiFdZLJyt4+qesVCov2y+vKVfclpxaKZ65nwrdKCP8yWJ8fJuko+t4UUtON8f/6yNshk0J3LF/9vZeZuu8hOcigNuPysWhOqDLE="
	signature, _ := base64.StdEncoding.DecodeString(sigBase64)

	// Read message
	//message2, err := ioutil.ReadFile("message.txt")
	//if err != nil {
	//	log.Fatal("Failed to read message:", err)
	//}
	//fmt.Println(message)
	//fmt.Println(message2)

	// Read signature (binary format)
	//signature, err := ioutil.ReadFile("signature.bin")
	//if err != nil {
	//	log.Fatal("Failed to read signature:", err)
	//}

	// Verify signature
	err = verifySignature(message, signature, publicKey)
	if err != nil {
		fmt.Println("Verification failed:", err)
	} else {
		fmt.Println("Signature verified successfully!")
	}
}
