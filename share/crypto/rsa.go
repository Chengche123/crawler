package crypto

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"

	"fmt"
)

func DecryptComicDetail(cipher string) ([]byte, error) {
	pck, err := base64.StdEncoding.DecodeString(PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("invalid Private key: %v", err)
	}

	key, err := x509.ParsePKCS8PrivateKey(pck)
	if err != nil {
		return nil, fmt.Errorf("invalid Private key: %v", err)
	}

	bcipher, err := base64.StdEncoding.DecodeString(cipher)
	if err != nil {
		return nil, fmt.Errorf("invalid cipher: %v", err)
	}

	rsaKey, ok := key.(*rsa.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("assert error: key.(*rsa.PrivateKey)")
	}

	numBlocks := len(bcipher) / rsaKey.PublicKey.Size()
	res := make([]byte, 0)

	for i := 0; i < numBlocks; i++ {
		dec, err := rsa.DecryptPKCS1v15(nil, rsaKey, bcipher[128*i:128*(i+1)])
		if err != nil {
			return nil, fmt.Errorf("failed to decrypt cipher: %v", err)
		}

		res = append(res, dec...)
	}

	return res, nil
}
