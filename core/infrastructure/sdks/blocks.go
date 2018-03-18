package sdks

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	cryptography "github.com/XMNBlockchain/exmachina-network/core/domain/cryptography"
	dblocks "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/blocks"
	users "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/users"
	sdks "github.com/XMNBlockchain/exmachina-network/core/domain/sdks"
	servers "github.com/XMNBlockchain/exmachina-network/core/domain/servers"
	concrete_blocks "github.com/XMNBlockchain/exmachina-network/core/infrastructure/projects/blockchains/blocks"
	"github.com/go-resty/resty"
	uuid "github.com/satori/go.uuid"
)

type blocks struct {
	sigBuilderFactory users.SignatureBuilderFactory
	routePrefix       string
	pk                cryptography.PrivateKey
	user              users.User
}

// CreateBlocks creates a new Blocks SDK instance
func CreateBlocks(sigBuilderFactory users.SignatureBuilderFactory, routePrefix string, pk cryptography.PrivateKey, user users.User) sdks.Blocks {
	out := blocks{
		sigBuilderFactory: sigBuilderFactory,
		routePrefix:       routePrefix,
		pk:                pk,
		user:              user,
	}
	return &out
}

// SaveBlock saves a block to the blocks
func (sdkblks *blocks) SaveBlock(serv servers.Server, blk dblocks.Block) (dblocks.SignedBlock, error) {
	url := fmt.Sprintf("%s%s/block", serv.String(), sdkblks.routePrefix)
	js, jsErr := json.Marshal(blk)
	if jsErr != nil {
		str := fmt.Sprintf("the given block could not be converted to JSON: %s", jsErr.Error())
		return nil, errors.New(str)
	}

	formData := map[string]string{
		"block": string(js),
	}

	id := uuid.NewV4()
	crOn := time.Now().UTC()
	userSig, userSigErr := sdkblks.sigBuilderFactory.Create().Create().WithID(&id).CreatedOn(crOn).WithUser(sdkblks.user).WithPrivateKey(sdkblks.pk).WithInterface(blk).Now()
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

	out := new(concrete_blocks.SignedBlock)
	outErr := json.Unmarshal(resp.Body(), out)
	if outErr != nil {
		str := fmt.Sprintf("there was a problem while converting output to a signed block: %s", outErr.Error())
		return nil, errors.New(str)
	}

	return out, nil
}
