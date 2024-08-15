package main

import (
	"crypto/sha256"
	"encoding/hex"
	"golang.org/x/crypto/ssh"
	"strings"
)

func main() {

}

func FingerprintLegacySHA256(pubKey ssh.PublicKey) string {
	sha256sum := sha256.Sum256(pubKey.Marshal())
	hexarray := make([]string, len(sha256sum))
	for i, c := range sha256sum {
		hexarray[i] = hex.EncodeToString([]byte{c})
	}
	return strings.Join(hexarray, ":")
}

func PublicKeyGenerateFingerprint(publicKey string) (string, error) {
	pub, _, _, _, err := ssh.ParseAuthorizedKey([]byte(publicKey))
	if err != nil {
		return "", err
	}
	return FingerprintLegacySHA256(pub), nil
}
