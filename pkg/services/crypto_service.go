package services

import (
	"flexagent/models"
	"flexagent/pkg/config"
	"flexagent/pkg/crypto"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"

	"github.com/flexlet/utils"
)

var CryptoService *cryptoServiceImpl

type cryptoServiceImpl struct {
	mutex  *sync.Mutex
	Crypto crypto.Crypto
}

func NewCryptoService() *cryptoServiceImpl {
	cryptoInst, err := crypto.NewCrypto(config.AgentConfig.Keystores, config.AgentConfig.Algorithum)
	if err != nil {
		utils.LogPrintf(utils.LOG_ERROR, "crypto.NewCryptoService", "load crypto failed, reason: %s", err.Error())
		os.Exit(0)
	}
	if _, err = cryptoInst.Encrypt([]byte{' '}); err != nil {
		utils.LogPrintf(utils.LOG_ERROR, "crypto.NewCryptoService", "test crypto failed, reason: %s", err.Error())
		os.Exit(0)
	}
	return &cryptoServiceImpl{
		mutex:  &sync.Mutex{},
		Crypto: cryptoInst,
	}
}

func (s *cryptoServiceImpl) encrypt(plainData map[string]string, format string) (*map[string]string, error) {
	cipherData := make(map[string]string, len(plainData))

	for k, v := range plainData {
		if cipherText, err := s.Crypto.EncryptString(v, format); err != nil {
			return nil, err
		} else {
			cipherData[k] = cipherText
		}
	}

	return &cipherData, nil
}

func (s *cryptoServiceImpl) Encrypt(plainData *models.CryptoData) (*models.CryptoData, error) {
	cipherData, err := s.encrypt(plainData.Data, plainData.Format)
	if err != nil {
		return nil, err
	}

	return &models.CryptoData{
		Format: plainData.Format,
		Data:   *cipherData,
	}, nil
}

func (s *cryptoServiceImpl) SecretEncrypt(plainData *models.KubeSecret) (*models.KubeSecret, error) {
	cipherData, err := s.encrypt(plainData.Data, models.CryptoDataFormatBase64)
	if err != nil {
		return nil, err
	}

	return &models.KubeSecret{
		APIVersion: plainData.APIVersion,
		Kind:       plainData.Kind,
		Type:       plainData.Type,
		Metadata:   plainData.Metadata,
		Data:       *cipherData,
	}, nil
}

func (s *cryptoServiceImpl) decrypt(cipherData map[string]string, format string) (*map[string]string, error) {
	plainData := make(map[string]string, len(cipherData))

	for k, v := range cipherData {
		if plainText, err := s.Crypto.DecryptString(v, format); err != nil {
			return nil, err
		} else {
			plainData[k] = plainText
		}
	}
	return &plainData, nil
}

func (s *cryptoServiceImpl) Decrypt(cipherData *models.CryptoData) (*models.CryptoData, error) {
	plainData, err := s.decrypt(cipherData.Data, cipherData.Format)
	if err != nil {
		return nil, err
	}

	return &models.CryptoData{
		Format: cipherData.Format,
		Data:   *plainData,
	}, nil
}

func (s *cryptoServiceImpl) SecretDecrypt(cipherData *models.KubeSecret) (*models.KubeSecret, error) {
	plainData, err := s.decrypt(cipherData.Data, models.CryptoDataFormatBase64)
	if err != nil {
		return nil, err
	}

	return &models.KubeSecret{
		APIVersion: cipherData.APIVersion,
		Kind:       cipherData.Kind,
		Type:       cipherData.Type,
		Metadata:   cipherData.Metadata,
		Data:       *plainData,
	}, nil
}

func (s *cryptoServiceImpl) ListVaults(filter *string) (*[]string, error) {
	dirs, err := ioutil.ReadDir(config.AgentConfig.VaultsPath)
	if err != nil {
		return nil, err
	}

	var names []string
	for i := 0; i < len(dirs); i++ {
		name := dirs[i].Name()
		if filter != nil {
			if strings.Contains(strings.ToLower(name), strings.ToLower(*filter)) {
				names = append(names, name)
			}
		} else {
			names = append(names, name)
		}
	}

	return &names, nil
}

func (s *cryptoServiceImpl) CreateVault(name string, data map[string]string) error {
	vaultPath := config.AgentConfig.VaultsPath + "/" + name
	utils.MkdirIfNotExist(vaultPath)

	cipherData, err := s.encrypt(data, crypto.FORMAT_RAW)

	if err != nil {
		return nil
	}

	for k, v := range *cipherData {
		vaultFile := vaultPath + "/" + k
		if err := utils.WriteFile(vaultFile, os.O_CREATE|os.O_RDWR|os.O_TRUNC, utils.MODE_PERM_RW, v); err != nil {
			return err
		}
	}

	return nil
}

func (s *cryptoServiceImpl) UpdateVault(name string, data map[string]string) error {
	vaultPath := config.AgentConfig.VaultsPath + "/" + name
	if !utils.FileExist(vaultPath) {
		return fmt.Errorf("vault '%s' does not exist", name)
	}

	cipherData, err := s.encrypt(data, crypto.FORMAT_RAW)

	if err != nil {
		return nil
	}

	for k, v := range *cipherData {
		vaultFile := vaultPath + "/" + k
		if err := utils.WriteFile(vaultFile, os.O_CREATE|os.O_RDWR|os.O_TRUNC, utils.MODE_PERM_RW, v); err != nil {
			return err
		}
	}

	return nil
}

func (s *cryptoServiceImpl) QueryVault(name string, keys []string) (*map[string]string, error) {
	vaultPath := config.AgentConfig.VaultsPath + "/" + name
	if !utils.FileExist(vaultPath) {
		return nil, fmt.Errorf("vault '%s' does not exist", name)
	}

	cipherData := make(map[string]string)

	files, err := ioutil.ReadDir(vaultPath)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(files); i++ {
		key := files[i].Name()
		if len(keys) != 0 && !utils.ListContains(keys, key) {
			continue
		}
		filePath := vaultPath + "/" + key
		data, err := ioutil.ReadFile(filePath)
		if err != nil {
			return nil, err
		}
		cipherData[key] = string(data)
	}

	return s.decrypt(cipherData, crypto.FORMAT_RAW)
}

func (s *cryptoServiceImpl) DeleteVault(name string, keys []string) error {
	vaultPath := config.AgentConfig.VaultsPath + "/" + name
	if !utils.FileExist(vaultPath) {
		return fmt.Errorf("vault '%s' does not exist", name)
	}

	// delete vault
	if len(keys) == 0 {
		return os.RemoveAll(vaultPath)
	}

	// delete vault keys
	for i := 0; i < len(keys); i++ {
		file := vaultPath + "/" + keys[i]
		if err := os.Remove(file); err != nil {
			return fmt.Errorf("remove key '%s' faild", keys[i])
		}
	}

	return nil
}
