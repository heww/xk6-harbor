// Code generated by go-swagger; DO NOT EDIT.

package statistic

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

//go:generate mockery -name API -inpkg

// API is the interface of the statistic client
type API interface {
	/*
	   GetStatistic gets the statistic information about the projects and repositories

	   Get the statistic information about the projects and repositories*/
	GetStatistic(ctx context.Context, params *GetStatisticParams) (*GetStatisticOK, error)
}

// New creates a new statistic API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) *Client {
	return &Client{
		transport: transport,
		formats:   formats,
		authInfo:  authInfo,
	}
}

/*
Client for statistic API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
}

/*
GetStatistic gets the statistic information about the projects and repositories

Get the statistic information about the projects and repositories
*/
func (a *Client) GetStatistic(ctx context.Context, params *GetStatisticParams) (*GetStatisticOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getStatistic",
		Method:             "GET",
		PathPattern:        "/statistics",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetStatisticReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetStatisticOK), nil

}
