package sdks

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	projs "github.com/XMNBlockchain/exmachina-network/core/domain/projects"
	sdks "github.com/XMNBlockchain/exmachina-network/core/domain/sdks"
	servers "github.com/XMNBlockchain/exmachina-network/core/domain/servers"
	concrete_projects "github.com/XMNBlockchain/exmachina-network/core/infrastructure/projects"
	"github.com/go-resty/resty"
	uuid "github.com/satori/go.uuid"
)

type projects struct {
	routePrefix string
}

// CreateProjects creates a new Projects instance
func CreateProjects(routePrefix string) sdks.Projects {
	out := projects{
		routePrefix: routePrefix,
	}
	return &out
}

// Retrieve retrieve projects
func (proj *projects) Retrieve(serv servers.Server, index int, amount int) (projs.Projects, error) {
	//create the url:
	urlAsString := fmt.Sprintf("%s%s/", serv.String(), proj.routePrefix)

	//create query:
	resp, respErr := resty.R().
		SetHeader("Content-Type", "application/json").
		SetQueryParam("index", strconv.Itoa(index)).
		SetQueryParam("amount", strconv.Itoa(amount)).
		Get(urlAsString)

	if respErr != nil {
		str := fmt.Sprintf("there was a problem while executing the http query: %s", respErr.Error())
		return nil, errors.New(str)
	}

	//make sure the status code is good:
	statusCode := resp.StatusCode()
	if statusCode < 200 || statusCode >= 300 {
		return nil, errors.New(string(resp.Body()))
	}

	//create the instance:
	out := new(concrete_projects.Projects)
	outErr := json.Unmarshal(resp.Body(), out)
	if outErr != nil {
		str := fmt.Sprintf("there was a problem while converting output to an instance: %s", outErr.Error())
		return nil, errors.New(str)
	}

	//returns:
	return out, nil
}

// RetrieveByID retrieve project by ID
func (proj *projects) RetrieveByID(serv servers.Server, id *uuid.UUID) (projs.Project, error) {
	//create the url:
	urlAsString := fmt.Sprintf("%s%s/%s", serv.String(), proj.routePrefix, id.String())

	//create query:
	resp, respErr := resty.R().
		SetHeader("Content-Type", "application/json").
		Get(urlAsString)

	if respErr != nil {
		str := fmt.Sprintf("there was a problem while executing the http query: %s", respErr.Error())
		return nil, errors.New(str)
	}

	//make sure the status code is good:
	statusCode := resp.StatusCode()
	if statusCode < 200 || statusCode >= 300 {
		return nil, errors.New(string(resp.Body()))
	}

	//create the instance:
	out := new(concrete_projects.Project)
	outErr := json.Unmarshal(resp.Body(), out)
	if outErr != nil {
		str := fmt.Sprintf("there was a problem while converting output to an instance: %s", outErr.Error())
		return nil, errors.New(str)
	}

	//returns:
	return out, nil
}
