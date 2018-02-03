package infrastructure

import (
	"encoding/json"
	"errors"
	"fmt"

	blocks "github.com/XMNBlockchain/core/packages/blocks/blocks/domain"
	concrete_blocks "github.com/XMNBlockchain/core/packages/blocks/blocks/infrastructure"
	cryptography "github.com/XMNBlockchain/core/packages/cryptography/domain"
	servers "github.com/XMNBlockchain/core/packages/servers/domain"
	users "github.com/XMNBlockchain/core/packages/users/domain"
	sdk "github.com/XMNBlockchain/core/sdks/domain"
	"github.com/go-resty/resty"
	uuid "github.com/satori/go.uuid"
)

type databases struct {
	sigBuilderFactory users.SignatureBuilderFactory
	pk                cryptography.PrivateKey
	userID            *uuid.UUID
}

// CreateDatabases creates a new Databases SDK instance
func CreateDatabases(sigBuilderFactory users.SignatureBuilderFactory, pk cryptography.PrivateKey, userID *uuid.UUID) sdk.Databases {
	out := databases{
		sigBuilderFactory: sigBuilderFactory,
		pk:                pk,
		userID:            userID,
	}
	return &out
}

// SaveBlock saves a block to the database
func (sdkdb *databases) SaveBlock(serv servers.Server, blk blocks.Block) (blocks.SignedBlock, error) {
	url := fmt.Sprintf("%s/block", serv.String())
	js, jsErr := json.Marshal(blk)
	if jsErr != nil {
		str := fmt.Sprintf("the given block could not be converted to JSON: %s", jsErr.Error())
		return nil, errors.New(str)
	}

	formData := map[string]string{
		"block": string(js),
	}

	userSig, userSigErr := sdkdb.sigBuilderFactory.Create().Create().WithUserID(*sdkdb.userID).WithPrivateKey(sdkdb.pk).WithInterface(blk).Now()
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

	out := new(concrete_blocks.SignedBlock)
	outErr := json.Unmarshal(resp.Body(), out)
	if outErr != nil {
		str := fmt.Sprintf("there was a problem while converting output to a signed block: %s", outErr.Error())
		return nil, errors.New(str)
	}

	return out, nil
}
