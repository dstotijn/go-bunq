package bunq

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"sort"
)

type header struct {
	Key   string
	Value string
}

// byKey implements sort.Interface for []header based on the key field.
type byKey []header

func (k byKey) Len() int           { return len(k) }
func (k byKey) Swap(i, j int)      { k[i], k[j] = k[j], k[i] }
func (k byKey) Less(i, j int) bool { return k[i].Key < k[j].Key }

func (c *Client) addSignature(req *http.Request, endpoint, body string) error {
	if c.PrivateKey == nil {
		return errors.New("bunq: private key cannot be nil")
	}
	var headers []header
	for key, val := range req.Header {
		for j := range val {
			headers = append(headers, header{
				Key:   key,
				Value: val[j],
			})
		}
	}
	sort.Sort(byKey(headers))

	message := endpoint + "\n"
	for i := range headers {
		if string([]byte(headers[i].Key)[:7]) == "X-Bunq-" || headers[i].Key == "User-Agent" || headers[i].Key == "Cache-Control" {
			message += headers[i].Key + ": " + headers[i].Value + "\n"
		}
	}
	message += "\n" + body

	h := sha256.New()
	_, err := h.Write([]byte(message))
	if err != nil {
		return err
	}

	signature, err := rsa.SignPKCS1v15(rand.Reader, c.PrivateKey, crypto.SHA256, h.Sum(nil))
	if err != nil {
		return err
	}

	req.Header.Set("X-Bunq-Client-Signature", base64.StdEncoding.EncodeToString(signature))

	return nil
}

// SetPrivateKey reads and parses private key data into a private key.
func (c *Client) SetPrivateKey(r io.Reader) error {
	pemData, err := ioutil.ReadAll(r)
	if err != nil {
		return fmt.Errorf("bunq: error reading PEM data: %v", err)
	}

	pemBlock, _ := pem.Decode(pemData)
	if pemBlock == nil {
		return errors.New("bunq: no PEM block found in data")
	}
	if pemBlock.Type != "RSA PRIVATE KEY" {
		return errors.New("bunq: invalid key type found, expected `RSA PRIVATE KEY`")
	}
	privKey, err := x509.ParsePKCS1PrivateKey(pemBlock.Bytes)
	if err != nil {
		return fmt.Errorf("bunq: error parsing PEM block into private key")
	}
	c.PrivateKey = privKey

	return nil
}

func (c *Client) publicKey() ([]byte, error) {
	if c.PrivateKey == nil {
		return nil, errors.New("private key cannot be nil")
	}
	pubKeyDer, err := x509.MarshalPKIXPublicKey(c.PrivateKey.Public())
	if err != nil {
		return nil, fmt.Errorf("error serializing public key to DER-encoded PKIX format: %v", err)
	}
	pubKeyPemBlock := pem.Block{
		Type:    "PUBLIC KEY",
		Headers: nil,
		Bytes:   pubKeyDer,
	}

	return pem.EncodeToMemory(&pubKeyPemBlock), nil
}
