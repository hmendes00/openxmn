package sdks

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	cryptography "github.com/XMNBlockchain/exmachina-network/core/domain/cryptography"
	trs "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/blockchains/transactions"
	signed_transactions "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/blockchains/transactions/signed"
	users "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/blockchains/users"
	sdks "github.com/XMNBlockchain/exmachina-network/core/domain/sdks"
	servers "github.com/XMNBlockchain/exmachina-network/core/domain/servers"
	concrete_signed_transactions "github.com/XMNBlockchain/exmachina-network/core/infrastructure/data/types/blockchains/transactions/signed"
	"github.com/go-resty/resty"
	uuid "github.com/satori/go.uuid"
)

type transactions struct {
	sigBuilderFactory users.SignatureBuilderFactory
	routePrefix       string
	pk                cryptography.PrivateKey
	user              users.User
}

// CreateTransactions creates a new Transactions SDK instance
func CreateTransactions(sigBuilderFactory users.SignatureBuilderFactory, routePrefix string, pk cryptography.PrivateKey, user users.User) sdks.Transactions {
	out := transactions{
		sigBuilderFactory: sigBuilderFactory,
		routePrefix:       routePrefix,
		pk:                pk,
		user:              user,
	}
	return &out
}

// SaveTrs save a Transaction on the Transactions API
func (sdktrs *transactions) SaveTrs(serv servers.Server, trs trs.Transaction) (signed_transactions.Transaction, error) {
	url := fmt.Sprintf("%s%s/transaction", serv.String(), sdktrs.routePrefix)

	js, jsErr := json.Marshal(trs)
	if jsErr != nil {
		str := fmt.Sprintf("the given transaction could not be converted to JSON: %s", jsErr.Error())
		return nil, errors.New(str)
	}

	formData := map[string]string{
		"transaction": string(js),
	}

	id := uuid.NewV4()
	crOn := time.Now().UTC()
	userSig, userSigErr := sdktrs.sigBuilderFactory.Create().Create().WithID(&id).CreatedOn(crOn).WithUser(sdktrs.user).WithPrivateKey(sdktrs.pk).WithInterface(trs).Now()
	if userSigErr != nil {
		str := fmt.Sprintf("there was an error while building the user signature: %s", userSigErr.Error())
		return nil, errors.New(str)
	}

	resp, respErr := resty.R().
		SetHeader("signature", userSig.String()).
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
	url := fmt.Sprintf("%s%s/atomic-transaction", serv.String(), sdktrs.routePrefix)

	js, jsErr := json.Marshal(trs)
	if jsErr != nil {
		str := fmt.Sprintf("the given transactions could not be converted to JSON: %s", jsErr.Error())
		return nil, errors.New(str)
	}

	formData := map[string]string{
		"transactions": string(js),
	}

	id := uuid.NewV4()
	crOn := time.Now().UTC()
	userSig, userSigErr := sdktrs.sigBuilderFactory.Create().Create().WithID(&id).CreatedOn(crOn).WithUser(sdktrs.user).WithPrivateKey(sdktrs.pk).WithInterface(trs).Now()
	if userSigErr != nil {
		str := fmt.Sprintf("there was an error while building the user signature: %s", userSigErr.Error())
		return nil, errors.New(str)
	}

	resp, respErr := resty.R().
		SetHeader("signature", userSig.String()).
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
