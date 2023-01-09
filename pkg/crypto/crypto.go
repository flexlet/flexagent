package crypto

import (
	"encoding/base64"
	"fmt"
	"path/filepath"

	"gitee.com/yaohuiwang/utils"
)

const (
	FORMAT_BASE64 = "base64"
	FORMAT_RAW    = "raw"
	EMPTY_STRING  = ""
)

const (
	AES_128 = "AES_128"
	AES_192 = "AES_192"
	AES_256 = "AES_256"
)

type Crypto interface {
	Init() error
	Encrypt(plainText []byte) ([]byte, error)
	Decrypt(cipherText []byte) ([]byte, error)
	EncryptString(plainText string, format string) (string, error)
	DecryptString(cipherText string, format string) (string, error)
}

type CryptoEngine interface {
	Init() error
	Encrypt(plainText []byte) ([]byte, error)
	Decrypt(cipherText []byte) ([]byte, error)
}

type cryptoImpl struct {
	engine CryptoEngine
}

func NewCrypto(keystores map[string]string, algorithum string) (Crypto, error) {
	var engine CryptoEngine
	var err error

	engine, err = NewCryptoEngine_AES(keystores, algorithum)

	if err != nil {
		return nil, err
	}

	return &cryptoImpl{engine: engine}, nil
}

func (c *cryptoImpl) Init() error {
	return c.engine.Init()
}

func (c *cryptoImpl) Encrypt(plainText []byte) ([]byte, error) {
	return c.engine.Encrypt(plainText)
}

func (c *cryptoImpl) Decrypt(cipherText []byte) ([]byte, error) {
	return c.engine.Decrypt(cipherText)
}

func (c *cryptoImpl) EncryptString(plainText string, format string) (string, error) {
	if format == FORMAT_RAW {
		cipherText, err := c.Encrypt([]byte(plainText))
		if err != nil {
			return EMPTY_STRING, err
		}
		return string(cipherText), nil
	} else if format == FORMAT_BASE64 {
		decodedText, err := base64.StdEncoding.DecodeString(plainText)
		if err != nil {
			return EMPTY_STRING, err
		}
		cipherText, err := c.Encrypt(decodedText)
		if err != nil {
			return EMPTY_STRING, err
		}
		encodedText := base64.StdEncoding.EncodeToString(cipherText)
		return encodedText, nil
	} else {
		return EMPTY_STRING, fmt.Errorf("wrong encrypt string format: %s", format)
	}
}

func (c *cryptoImpl) DecryptString(cipherText string, format string) (string, error) {
	if format == FORMAT_RAW {
		plainText, err := c.Decrypt([]byte(cipherText))
		if err != nil {
			return EMPTY_STRING, err
		}
		return string(plainText), nil
	} else if format == FORMAT_BASE64 {
		decodedText, err := base64.StdEncoding.DecodeString(cipherText)
		if err != nil {
			return EMPTY_STRING, err
		}
		plainText, err := c.Decrypt(decodedText)
		if err != nil {
			return EMPTY_STRING, err
		}
		encodedText := base64.StdEncoding.EncodeToString(plainText)
		return encodedText, nil
	} else {
		return EMPTY_STRING, fmt.Errorf("wrong encrypt string format: %s", format)
	}
}

func checkKeystores(keystores map[string]string) bool {
	if ks, exist := keystores["primary"]; !exist {
		return false
	} else {
		if !utils.FileExist(ks) {
			utils.MkdirIfNotExist(filepath.Dir(ks))
		}
	}
	if ks, exist := keystores["standby"]; !exist {
		return false
	} else {
		if !utils.FileExist(ks) {
			utils.MkdirIfNotExist(filepath.Dir(ks))
		}
	}
	return true
}
