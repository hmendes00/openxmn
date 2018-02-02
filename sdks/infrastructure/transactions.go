package infrastructure

import (
	"encoding/json"
	"errors"
	"fmt"

	cryptography "github.com/XMNBlockchain/core/packages/cryptography/domain"
	servers "github.com/XMNBlockchain/core/packages/servers/domain"
	signed_transactions "github.com/XMNBlockchain/core/packages/transactions/signed/domain"
	concrete_signed_transactions "github.com/XMNBlockchain/core/packages/transactions/signed/infrastructure"
	trs "github.com/XMNBlockchain/core/packages/transactions/transactions/domain"
	users "github.com/XMNBlockchain/core/packages/users/domain"
	sdk "github.com/XMNBlockchain/core/sdks/domain"
	"github.com/go-resty/resty"
	uuid "github.com/satori/go.uuid"
)

type transactions struct {
	sigBuilderFactory users.SignatureBuilderFactory
	pk                cryptography.PrivateKey
	userID            *uuid.UUID
}

// CreateTransactions creates a new Transactions SDK instance
func CreateTransactions(sigBuilderFactory users.SignatureBuilderFactory, pk cryptography.PrivateKey, userID *uuid.UUID) sdk.Transactions {
	out := transactions{
		sigBuilderFactory: sigBuilderFactory,
		pk:                pk,
		userID:            userID,
	}
	return &out
}

// SaveTrs save a Transaction on the Transactions API
func (sdktrs *transactions) SaveTrs(serv servers.Server, trs trs.Transaction) (signed_transactions.Transaction, error) {
	url := fmt.Sprintf("%s/transaction", serv.String())

	js, jsErr := json.Marshal(trs)
	if jsErr != nil {
		str := fmt.Sprintf("the given transaction could not be converted to JSON: %s", jsErr.Error())
		return nil, errors.New(str)
	}

	formData := map[string]string{
		"transaction": string(js),
	}

	userSig, userSigErr := sdktrs.sigBuilderFactory.Create().Create().WithUserID(*sdktrs.userID).WithPrivateKey(sdktrs.pk).WithInterface(trs).Now()
	if userSigErr != nil {
		str := fmt.Sprintf("there was an error while building the user signature: %s", userSigErr.Error())
		return nil, errors.New(str)
	}

	resp, respErr := resty.R().
		SetHeader("signature", userSig.GetSig().String()).
		SetHeader("user_id", userSig.GetUser().GetID().String()).
		SetHeader("Content-Type", "application/json").
		SetFormData(formData).
		Post(url)

	if respErr != nil {
		str := fmt.Sprintf("there was a problem while executing the http query: %s", respErr.Error())
		return nil, errors.New(str)
	}

	statusCode := resp.StatusCode()
	if statusCode < 200 || statusCode >= 300 {
		return nil, errors.New(string(resp.Body()))
	}

	out := new(concrete_signed_transactions.Transaction)
	outErr := json.Unmarshal(resp.Body(), out)
	if outErr != nil {
		str := fmt.Sprintf("there was a problem while converting output to a signed transaction: %s", outErr.Error())
		return nil, errors.New(str)
	}

	return out, nil
}

// SaveAtomicTrs save an AtomicTransaction on the Transactions API
func (sdktrs *transactions) SaveAtomicTrs(serv servers.Server, trs []trs.Transaction) (signed_transactions.AtomicTransaction, error) {
	url := fmt.Sprintf("%s/atomic-transaction", serv.String())

	js, jsErr := json.Marshal(trs)
	if jsErr != nil {
		str := fmt.Sprintf("the given transactions could not be converted to JSON: %s", jsErr.Error())
		return nil, errors.New(str)
	}

	formData := map[string]string{
		"transactions": string(js),
	}

	userSig, userSigErr := sdktrs.sigBuilderFactory.Create().Create().WithUserID(*sdktrs.userID).WithPrivateKey(sdktrs.pk).WithInterface(trs).Now()
	if userSigErr != nil {
		str := fmt.Sprintf("there was an error while building the user signature: %s", userSigErr.Error())
		return nil, errors.New(str)
	}

	resp, respErr := resty.R().
		SetHeader("signature", userSig.GetSig().String()).
		SetHeader("user_id", userSig.GetUser().GetID().String()).
		SetHeader("Content-Type", "application/json").
		SetFormData(formData).
		Post(url)

	if respErr != nil {
		str := fmt.Sprintf("there was a problem while executing the http query: %s", respErr.Error())
		return nil, errors.New(str)
	}

	statusCode := resp.StatusCode()
	if statusCode < 200 || statusCode >= 300 {
		return nil, errors.New(string(resp.Body()))
	}

	out := new(concrete_signed_transactions.AtomicTransaction)
	outErr := json.Unmarshal(resp.Body(), out)
	if outErr != nil {
		str := fmt.Sprintf("there was a problem while converting output to an atomic signed transaction: %s", outErr.Error())
		return nil, errors.New(str)
	}

	return out, nil
}
