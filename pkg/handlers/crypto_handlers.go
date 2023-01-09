package handlers

import (
	"flexagent/pkg/services"
	"flexagent/restapi/operations/crypto"

	"github.com/go-openapi/runtime/middleware"
)

type CryptoEncryptHandlerImpl struct {
}

func (h *CryptoEncryptHandlerImpl) Handle(params crypto.EncryptParams) middleware.Responder {
	cipherData, err := services.CryptoService.Encrypt(params.Data)
	if err != nil {
		errMsg := err.Error()
		body := &crypto.EncryptBadRequestBody{Message: &errMsg}
		return crypto.NewEncryptBadRequest().WithPayload(body)
	}
	return crypto.NewEncryptOK().WithPayload(cipherData)
}

type CryptoDecryptHandlerImpl struct {
}

func (h *CryptoDecryptHandlerImpl) Handle(params crypto.DecryptParams) middleware.Responder {
	plainData, err := services.CryptoService.Decrypt(params.Data)
	if err != nil {
		errMsg := err.Error()
		body := &crypto.DecryptBadRequestBody{Message: &errMsg}
		return crypto.NewDecryptBadRequest().WithPayload(body)
	}
	return crypto.NewDecryptOK().WithPayload(plainData)
}

type CryptoSecretEncryptHandlerImpl struct {
}

func (h *CryptoSecretEncryptHandlerImpl) Handle(params crypto.SecretEncryptParams) middleware.Responder {
	secret, err := services.CryptoService.SecretEncrypt(params.Secret)
	if err != nil {
		errMsg := err.Error()
		body := &crypto.SecretEncryptBadRequestBody{Message: &errMsg}
		return crypto.NewSecretEncryptBadRequest().WithPayload(body)
	}
	return crypto.NewSecretEncryptOK().WithPayload(secret)
}

type CryptoSecretDecryptHandlerImpl struct {
}

func (h *CryptoSecretDecryptHandlerImpl) Handle(params crypto.SecretDecryptParams) middleware.Responder {
	secret, err := services.CryptoService.SecretDecrypt(params.Secret)
	if err != nil {
		errMsg := err.Error()
		body := &crypto.SecretDecryptBadRequestBody{Message: &errMsg}
		return crypto.NewSecretDecryptBadRequest().WithPayload(body)
	}
	return crypto.NewSecretDecryptOK().WithPayload(secret)
}

type CryptoListVaultsHandlerImpl struct {
}

func (h *CryptoListVaultsHandlerImpl) Handle(params crypto.ListVaultsParams) middleware.Responder {
	names, err := services.CryptoService.ListVaults(params.Name)
	if err != nil {
		errMsg := err.Error()
		body := &crypto.ListVaultsBadRequestBody{Message: &errMsg}
		return crypto.NewListVaultsBadRequest().WithPayload(body)
	}
	return crypto.NewListVaultsOK().WithPayload(*names)
}

type CryptoCreateVaultHandlerImpl struct {
}

func (h *CryptoCreateVaultHandlerImpl) Handle(params crypto.CreateVaultParams) middleware.Responder {
	err := services.CryptoService.CreateVault(params.Name, params.Data)
	if err != nil {
		errMsg := err.Error()
		body := &crypto.CreateVaultBadRequestBody{Message: &errMsg}
		return crypto.NewCreateVaultBadRequest().WithPayload(body)
	}
	return crypto.NewCreateVaultOK().WithPayload("ok")
}

type CryptoUpdateVaultHandlerImpl struct {
}

func (h *CryptoUpdateVaultHandlerImpl) Handle(params crypto.UpdateVaultParams) middleware.Responder {
	err := services.CryptoService.UpdateVault(params.Name, params.Data)
	if err != nil {
		errMsg := err.Error()
		body := &crypto.UpdateVaultBadRequestBody{Message: &errMsg}
		return crypto.NewUpdateVaultBadRequest().WithPayload(body)
	}
	return crypto.NewUpdateVaultOK().WithPayload("ok")
}

type CryptoQueryVaultHandlerImpl struct {
}

func (h *CryptoQueryVaultHandlerImpl) Handle(params crypto.QueryVaultParams) middleware.Responder {
	data, err := services.CryptoService.QueryVault(params.Name, params.Keys)
	if err != nil {
		errMsg := err.Error()
		body := &crypto.QueryVaultBadRequestBody{Message: &errMsg}
		return crypto.NewQueryVaultBadRequest().WithPayload(body)
	}
	return crypto.NewQueryVaultOK().WithPayload(*data)
}

type CryptoDeleteVaultHandlerImpl struct {
}

func (h *CryptoDeleteVaultHandlerImpl) Handle(params crypto.DeleteVaultParams) middleware.Responder {
	err := services.CryptoService.DeleteVault(params.Name, params.Keys)
	if err != nil {
		errMsg := err.Error()
		body := &crypto.DeleteVaultBadRequestBody{Message: &errMsg}
		return crypto.NewDeleteVaultBadRequest().WithPayload(body)
	}
	return crypto.NewDeleteVaultOK().WithPayload("ok")
}
