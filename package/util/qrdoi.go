package util

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/asn1"
	"encoding/hex"
	"fmt"
	"gateway/package/log"
	"math/big"
	"strings"
)

//TODO (jwm): generate and store the private key in the config file
const key = "307702010104208655afb19646eb46c8b4b5f153e9809ba1a4eba169a6d4b989e6afe536aad2f5a00a06082a8648ce3d030107a14403420004039dbc5881387a4e44d4ece0aa7b687d0a04e074d57f787d13ba19144bd1d11dba2055d1d8011c97a752ad7b19f7584071486fe080238e4c7842254876698fcc"

type ecdsaSignature struct {
	R, S *big.Int
}

func getPrivateKeyfromString(private string) (*ecdsa.PrivateKey, error) {
	decoded, err := hex.DecodeString(string(private))
	if err != nil {
		return nil, err
	}
	priv, err := x509.ParseECPrivateKey(decoded)
	if err != nil {
		return nil, err
	}
	return priv, nil
}

//getPrivateString returns private  string
func getPrivateString(priv *ecdsa.PrivateKey) string {
	x509EncodedPriv, _ := x509.MarshalECPrivateKey(priv)
	return hex.EncodeToString(x509EncodedPriv)
}

func getPublicString(pub *ecdsa.PublicKey) string {
	x509EncodedPub, _ := x509.MarshalPKIXPublicKey(pub)
	return hex.EncodeToString(x509EncodedPub)
}

// GetQRString generates the qr string with the doi for the given data
func GetQRString(data []byte) (string, string, error) {
	privKey, _ := getPrivateKeyfromString(key)
	hash := sha256.Sum256(data)
	sig, err := ecdsa.SignASN1(rand.Reader, privKey, hash[:])
	if err != nil {
		log.Errorf("error signing the data: %v", err)
		return "", "", err
	}
	pub := getPublicString(&privKey.PublicKey)

	hexHash := hex.EncodeToString(hash[:])
	l := len(hexHash)
	doi := strings.ToUpper(hexHash[0:4] + " " + hexHash[l-4:l])
	//fmt.Printf("doi= %s\n", doi)
	//st := fmt.Sprintf("%s|%s|%s|%s", doi, pub, hex.EncodeToString(sig), hex.EncodeToString(hash[:]))
	st := fmt.Sprintf("https://soma.dit.ac.tz/verify/%s/%s/%s", hex.EncodeToString(sig), pub, hex.EncodeToString(hash[:]))
	return st, doi, nil
}

func Verify(signature, publicKey, hash string) (bool, string) {

	privKey, _ := getPrivateKeyfromString(key)
	hashByte, err := hex.DecodeString(hash)
	if err != nil {
		return false, ""
	}
	decoded, err := hex.DecodeString(signature)
	if err != nil {
		return false, ""
	}

	var signEcdsa ecdsaSignature

	_, err = asn1.Unmarshal(decoded, &signEcdsa)
	if err != nil {
		return false, ""
	}
	isValid := ecdsa.Verify(&privKey.PublicKey, hashByte, signEcdsa.R, signEcdsa.S)

	di := doi(hash)
	return isValid, di
}

func doi(hash string) string {
	l := len(hash)
	return strings.ToUpper(hash[0:4] + " " + hash[l-4:l])
}
