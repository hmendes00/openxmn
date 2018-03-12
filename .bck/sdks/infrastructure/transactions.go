package infrastructure

import (
	"encoding/json"
	"errors"
	"fmt"

	signed_transactions "github.com/XMNBlockchain/core/packages/blockchains/transactions/signed/domain"
	concrete_signed_transactions "github.com/XMNBlockchain/core/packages/blockchains/transactions/signed/infrastructure"
	trs "github.com/XMNBlockchain/core/packages/blockchains/transactions/transactions/domain"
	users "github.com/XMNBlockchain/core/packages/blockchains/users/domain"
	cryptography "github.com/XMNBlockchain/core/packages/cryptography/domain"
	servers "github.com/XMNBlockchain/core/packages/servers/domain"
	sdk "github.com/XMNBlockchain/core/sdks/domain"
	"github.com/go-resty/resty"
)

type transactions struct {
	sigBuilderFactory users.SignatureBuilderFactory
	pk                cryptography.PrivateKey
	user              users.User
}

// CreateTransactions creates a new Transactions SDK instance
func CreateTransactions(sigBuilderFactory users.SignatureBuilderFactory, pk cryptography.PrivateKey, user users.User) sdk.Transactions {
	out := transactions{
		sigBuilderFactory: sigBuilderFactory,
		pk:                pk,
		user:              user,
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

	userSig, userSigErr := sdktrs.sigBuilderFactory.Create().Create().WithUser(sdktrs.user).WithPrivateKey(sdktrs.pk).WithInterface(trs).Now()
	if userSigErr != nil {
		str := fmt.Sprintf("there was an error while building the user signature: %s", userSigErr.Error())
		return nil, errors.New(str)
	}

	resp, respErr := resty.R().
		SetHeader("signature", userSig.GetSignature().String()).
		SetHeader("user_id", userSig.GetUser().GetMetaData().GetID().String()).
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
func (sdktrs *transactions) SaveAtomicTrs(serv servers.Server, trs trs.Transactions) (signed_transactions.AtomicTransaction, error) {
	url := fmt.Sprintf("%s/atomic-transaction", serv.String())

	js, jsErr := json.Marshal(trs)
	if jsErr != nil {
		str := fmt.Sprintf("the given transactions could not be converted to JSON: %s", jsErr.Error())
		return nil, errors.New(str)
	}

	formData := map[string]string{
		"transactions": string(js),
	}

	userSig, userSigErr := sdktrs.sigBuilderFactory.Create().Create().WithUser(sdktrs.user).WithPrivateKey(sdktrs.pk).WithInterface(trs).Now()
	if userSigErr != nil {
		str := fmt.Sprintf("there was an error while building the user signature: %s", userSigErr.Error())
		return nil, errors.New(str)
	}

	resp, respErr := resty.R().
		SetHeader("signature", userSig.GetSignature().String()).
		SetHeader("user_id", userSig.GetUser().GetMetaData().GetID().String()).
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
