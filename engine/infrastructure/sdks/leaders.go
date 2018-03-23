package sdks

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	cryptography "github.com/XMNBlockchain/openxmn/engine/domain/cryptography"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
	sdks "github.com/XMNBlockchain/openxmn/engine/domain/sdks"
	servers "github.com/XMNBlockchain/openxmn/engine/domain/data/types/servers"
	"github.com/go-resty/resty"
	uuid "github.com/satori/go.uuid"

	aggregated "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/transactions/signed/aggregated"
	concrete_aggregated "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/blockchains/transactions/signed/aggregated"
)

type leaders struct {
	sigBuilderFactory users.SignatureBuilderFactory
	routePrefix       string
	pk                cryptography.PrivateKey
	user              users.User
}

// CreateLeaders creates a new Leaders SDK instance
func CreateLeaders(sigBuilderFactory users.SignatureBuilderFactory, routePrefix string, pk cryptography.PrivateKey, user users.User) sdks.Leaders {
	out := leaders{
		sigBuilderFactory: sigBuilderFactory,
		routePrefix:       routePrefix,
		pk:                pk,
		user:              user,
	}
	return &out
}

// SaveTrs saves aggregated transactions to the leaders
func (sdklead *leaders) SaveTrs(serv servers.Server, trs aggregated.Transactions) (aggregated.SignedTransactions, error) {
	url := fmt.Sprintf("%s%s/aggregated-transactions", serv.String(), sdklead.routePrefix)
	js, jsErr := json.Marshal(trs)
	if jsErr != nil {
		str := fmt.Sprintf("the given transaction could not be converted to JSON: %s", jsErr.Error())
		return nil, errors.New(str)
	}

	formData := map[string]string{
		"transactions": string(js),
	}

	id := uuid.NewV4()
	crOn := time.Now().UTC()
	userSig, userSigErr := sdklead.sigBuilderFactory.Create().Create().WithID(&id).CreatedOn(crOn).WithUser(sdklead.user).WithPrivateKey(sdklead.pk).WithInterface(trs).Now()
	if userSigErr != nil {
		str := fmt.Sprintf("there was an error while building the user signature: %s", userSigErr.Error())
		return nil, errors.New(str)
	}

	resp, respErr := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("signature", userSig.String()).
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
