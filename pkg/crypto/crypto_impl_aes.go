package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io/ioutil"

	"github.com/flexlet/utils"
)

type CryptoEngine_AES struct {
	PrimaryKeyStoreFile string
	StandbyKeyStoreFile string
	algorithum          string
	key                 []byte
	block               cipher.Block
}

func NewCryptoEngine_AES(keystores map[string]string, algorithum string) (CryptoEngine, error) {
	if !checkKeystores(keystores) {
		return nil, fmt.Errorf("keystores not set: primary and standby")
	}
	return &CryptoEngine_AES{
		PrimaryKeyStoreFile: keystores["primary"],
		StandbyKeyStoreFile: keystores["standby"],
		algorithum:          algorithum,
	}, nil
}

func (c *CryptoEngine_AES) Init() error {
	// generate key
	var key []byte
	if c.algorithum == AES_128 {
		key = make([]byte, 16)
		rand.Read(key)
	} else if c.algorithum == AES_192 {
		key = make([]byte, 24)
		rand.Read(key)
	} else {
		key = make([]byte, 32)
		rand.Read(key)
	}

	// try load block
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	// write to primary keystore file
	if err := ioutil.WriteFile(c.PrimaryKeyStoreFile, key, utils.MODE_PERM_RW); err != nil {
		return err
	}

	// write to standby keystore file
	if err := ioutil.WriteFile(c.StandbyKeyStoreFile, key, utils.MODE_PERM_RW); err != nil {
		return err
	}

	// set key & block
	c.block = block
	c.key = key

	return nil
}

func (c *CryptoEngine_AES) loadKey(keyFile string) error {
	if c.block == nil {
		key, err := ioutil.ReadFile(keyFile)
		if err != nil {
			return err
		}
		// try load block
		block, err := aes.NewCipher(key)
		if err != nil {
			return err
		}
		c.block = block
		c.key = key
	}
	return nil
}

func (c *CryptoEngine_AES) Encrypt(plainText []byte) ([]byte, error) {
	// load primary key file
	if err := c.loadKey(c.PrimaryKeyStoreFile); err != nil {
		// load standby key file
		if err := c.loadKey(c.StandbyKeyStoreFile); err != nil {
			return nil, err
		}
	}

	blockSize := c.block.BlockSize()
	plainText = PKCS7Padding(plainText, blockSize)
	blockMode := cipher.NewCBCEncrypter(c.block, c.key[:blockSize])
	cipherText := make([]byte, len(plainText))
	blockMode.CryptBlocks(cipherText, plainText)

	return cipherText, nil
}

func (c *CryptoEngine_AES) Decrypt(cipherText []byte) ([]byte, error) {
	// load primary key file
	if err := c.loadKey(c.PrimaryKeyStoreFile); err != nil {
		// load standby key file
		if err := c.loadKey(c.StandbyKeyStoreFile); err != nil {
			return nil, err
		}
	}

	blockSize := c.block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(c.block, c.key[:blockSize])
	plainText := make([]byte, len(cipherText))
	blockMode.CryptBlocks(plainText, cipherText)
	plainText = PKCS7UnPadding(plainText)

	return plainText, nil
}

//补码
func PKCS7Padding(ciphertext []byte, blocksize int) []byte {
	padding := blocksize - len(ciphertext)%blocksize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

//去码
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
