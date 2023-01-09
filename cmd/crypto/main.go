package main

import (
	"bufio"
	"encoding/base64"
	"flag"
	"flexagent/pkg/crypto"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/flexlet/utils"
)

func main() {
	var (
		init            bool
		algo            string
		encrypt         bool
		decrypt         bool
		format          string
		sourceFile      string
		targetFile      string
		primaryKeystore string
		standbyKeystore string
	)
	flag.BoolVar(&init, "init", false, "Initialize crypto")
	flag.StringVar(&algo, "algo", "AES-256", "Crypto algorithum: AES-128, AES-192, AES-256")
	flag.BoolVar(&encrypt, "encrypt", false, "Encrypt data")
	flag.BoolVar(&decrypt, "decrypt", false, "Decrypt data")
	flag.StringVar(&format, "format", "raw", "Format: raw, base64")
	flag.StringVar(&sourceFile, "sourcefile", "", "Source file")
	flag.StringVar(&targetFile, "targetfile", "", "Target file")
	flag.StringVar(&primaryKeystore, "primaryks", "/opt/flexagent/keystore/primary.ks", "Primary keystore file")
	flag.StringVar(&standbyKeystore, "standbyks", "/opt/flexagent/keystore/standby.ks", "Standby keystore file")
	flag.Parse()

	keystores := map[string]string{
		"primary": primaryKeystore,
		"standby": standbyKeystore,
	}

	// load crypto engine
	engine, err := crypto.NewCrypto(keystores, algo)

	if err != nil {
		panic(err)
	}

	if init {
		if err := engine.Init(); err != nil {
			panic(err)
		} else {
			println("crypto initialized")
		}
	} else if encrypt {
		var plainText []byte
		if sourceFile != "" {
			plainText = readFile(sourceFile)
		} else {
			if format == crypto.FORMAT_BASE64 {
				plainText = inputBase64()
			} else {
				plainText = inputRaw()
			}
		}

		cipherText, err := engine.Encrypt(plainText)
		if err != nil {
			panic(err)
		}

		if targetFile != "" {
			writeFile(targetFile, cipherText)
		} else {
			if format == crypto.FORMAT_BASE64 {
				printBase64(cipherText)
			} else {
				printRaw(cipherText)
			}
		}
	} else if decrypt {
		var cipherText []byte
		if sourceFile != "" {
			cipherText = readFile(sourceFile)
		} else {
			if format == crypto.FORMAT_BASE64 {
				cipherText = inputBase64()
			} else {
				cipherText = inputRaw()
			}
		}

		plainText, err := engine.Decrypt(cipherText)
		if err != nil {
			panic(err)
		}

		if targetFile != "" {
			writeFile(targetFile, plainText)
		} else {
			if format == crypto.FORMAT_BASE64 {
				printBase64(plainText)
			} else {
				printRaw(plainText)
			}
		}
	} else {
		panic(flag.ErrHelp)
	}
}

func readFile(filepath string) []byte {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	return data
}

func writeFile(filepath string, data []byte) {
	err := ioutil.WriteFile(filepath, data, utils.MODE_PERM_RW)
	if err != nil {
		panic(err)
	}
}

func inputRaw() []byte {
	reader := bufio.NewReader(os.Stdin)
	rawText, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	return []byte(rawText)
}

func inputBase64() []byte {
	reader := bufio.NewReader(os.Stdin)
	base64Text, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	data, err := base64.StdEncoding.DecodeString(base64Text)
	if err != nil {
		panic(err)
	}
	return data
}

func printRaw(data []byte) {
	fmt.Println(string(data))
}

func printBase64(data []byte) {
	base64Text := base64.StdEncoding.EncodeToString(data)
	fmt.Println(base64Text)
}
