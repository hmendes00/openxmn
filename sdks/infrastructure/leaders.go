package infrastructure

import (
	"encoding/json"
	"errors"
	"fmt"

	cryptography "github.com/XMNBlockchain/core/packages/cryptography/domain"
	servers "github.com/XMNBlockchain/core/packages/servers/domain"
	users "github.com/XMNBlockchain/core/packages/lives/users/domain"
	sdk "github.com/XMNBlockchain/core/sdks/domain"
	"github.com/go-resty/resty"
	uuid "github.com/satori/go.uuid"

	aggregated "github.com/XMNBlockchain/core/packages/lives/transactions/aggregated/domain"
	concrete_aggregated "github.com/XMNBlockchain/core/packages/lives/transactions/aggregated/infrastructure"
)

type leaders struct {
	sigBuilderFactory users.SignatureBuilderFactory
	pk                cryptography.PrivateKey
	userID            *uuid.UUID
}

// CreateLeaders creates a new Leaders SDK instance
func CreateLeaders(sigBuilderFactory users.SignatureBuilderFactory, pk cryptography.PrivateKey, userID *uuid.UUID) sdk.Leaders {
	out := leaders{
		sigBuilderFactory: sigBuilderFactory,
		pk:                pk,
		userID:            userID,
	}
	return &out
}

// SaveTrs saves aggregated transactions to the leaders
func (sdklead *leaders) SaveTrs(serv servers.Server, trs aggregated.Transactions) (aggregated.SignedTransactions, error) {
	url := fmt.Sprintf("%s/aggregated-transactions", serv.String())
	js, jsErr := json.Marshal(trs)
	if jsErr != nil {
		str := fmt.Sprintf("the given transaction could not be converted to JSON: %s", jsErr.Error())
		return nil, errors.New(str)
	}

	formData := map[string]string{
		"transactions": string(js),
	}

	userSig, userSigErr := sdklead.sigBuilderFactory.Create().Create().WithUserID(sdklead.userID).WithPrivateKey(sdklead.pk).WithInterface(trs).Now()
	if userSigErr != nil {
		str := fmt.Sprintf("there was an error while building the user signature: %s", userSigErr.Error())
		return nil, errors.New(str)
	}

	resp, respErr := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("signature", userSig.GetSig().String()).
		SetHeader("user_id", userSig.GetUser().GetID().String()).
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

	out := new(concrete_aggregated.SignedTransactions)
	outErr := json.Unmarshal(resp.Body(), out)
	if outErr != nil {
		str := fmt.Sprintf("there was a problem while converting output to a signed aggregated transaction: %s", outErr.Error())
		return nil, errors.New(str)
	}

	return out, nil
}
