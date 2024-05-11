// Code generated by go-swagger; DO NOT EDIT.

package jobservice

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

//go:generate mockery --name API --keeptree --with-expecter --case underscore

// API is the interface of the jobservice client
type API interface {
	/*
	   ActionGetJobLog gets job log by job id

	   Get job log by job id, it is only used by administrator*/
	ActionGetJobLog(ctx context.Context, params *ActionGetJobLogParams) (*ActionGetJobLogOK, error)
	/*
	   ActionPendingJobs stops and clean pause resume pending jobs in the queue

	   stop and clean, pause, resume pending jobs in the queue*/
	ActionPendingJobs(ctx context.Context, params *ActionPendingJobsParams) (*ActionPendingJobsOK, error)
	/*
	   GetWorkerPools gets worker pools

	   Get worker pools*/
	GetWorkerPools(ctx context.Context, params *GetWorkerPoolsParams) (*GetWorkerPoolsOK, error)
	/*
	   GetWorkers gets workers

	   Get workers in current pool*/
	GetWorkers(ctx context.Context, params *GetWorkersParams) (*GetWorkersOK, error)
	/*
	   ListJobQueues lists job queues

	   list job queue*/
	ListJobQueues(ctx context.Context, params *ListJobQueuesParams) (*ListJobQueuesOK, error)
	/*
	   StopRunningJob stops running job

	   Stop running job*/
	StopRunningJob(ctx context.Context, params *StopRunningJobParams) (*StopRunningJobOK, error)
}

// New creates a new jobservice API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) *Client {
	return &Client{
		transport: transport,
		formats:   formats,
		authInfo:  authInfo,
	}
}

/*
Client for jobservice API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
}

/*
ActionGetJobLog gets job log by job id

Get job log by job id, it is only used by administrator
*/
func (a *Client) ActionGetJobLog(ctx context.Context, params *ActionGetJobLogParams) (*ActionGetJobLogOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "actionGetJobLog",
		Method:             "GET",
		PathPattern:        "/jobservice/jobs/{job_id}/log",
		ProducesMediaTypes: []string{"text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ActionGetJobLogReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	switch value := result.(type) {
	case *ActionGetJobLogOK:
		return value, nil
	case *ActionGetJobLogUnauthorized:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *ActionGetJobLogForbidden:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *ActionGetJobLogNotFound:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *ActionGetJobLogInternalServerError:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for actionGetJobLog: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ActionPendingJobs stops and clean pause resume pending jobs in the queue

stop and clean, pause, resume pending jobs in the queue
*/
func (a *Client) ActionPendingJobs(ctx context.Context, params *ActionPendingJobsParams) (*ActionPendingJobsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "actionPendingJobs",
		Method:             "PUT",
		PathPattern:        "/jobservice/queues/{job_type}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ActionPendingJobsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	switch value := result.(type) {
	case *ActionPendingJobsOK:
		return value, nil
	case *ActionPendingJobsUnauthorized:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *ActionPendingJobsForbidden:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *ActionPendingJobsNotFound:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *ActionPendingJobsUnprocessableEntity:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *ActionPendingJobsInternalServerError:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for actionPendingJobs: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetWorkerPools gets worker pools

Get worker pools
*/
func (a *Client) GetWorkerPools(ctx context.Context, params *GetWorkerPoolsParams) (*GetWorkerPoolsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getWorkerPools",
		Method:             "GET",
		PathPattern:        "/jobservice/pools",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetWorkerPoolsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	switch value := result.(type) {
	case *GetWorkerPoolsOK:
		return value, nil
	case *GetWorkerPoolsUnauthorized:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *GetWorkerPoolsForbidden:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *GetWorkerPoolsInternalServerError:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getWorkerPools: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetWorkers gets workers

Get workers in current pool
*/
func (a *Client) GetWorkers(ctx context.Context, params *GetWorkersParams) (*GetWorkersOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getWorkers",
		Method:             "GET",
		PathPattern:        "/jobservice/pools/{pool_id}/workers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetWorkersReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	switch value := result.(type) {
	case *GetWorkersOK:
		return value, nil
	case *GetWorkersUnauthorized:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *GetWorkersForbidden:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *GetWorkersNotFound:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *GetWorkersInternalServerError:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getWorkers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListJobQueues lists job queues

list job queue
*/
func (a *Client) ListJobQueues(ctx context.Context, params *ListJobQueuesParams) (*ListJobQueuesOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listJobQueues",
		Method:             "GET",
		PathPattern:        "/jobservice/queues",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListJobQueuesReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	switch value := result.(type) {
	case *ListJobQueuesOK:
		return value, nil
	case *ListJobQueuesUnauthorized:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *ListJobQueuesForbidden:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *ListJobQueuesNotFound:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *ListJobQueuesInternalServerError:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listJobQueues: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
StopRunningJob stops running job

Stop running job
*/
func (a *Client) StopRunningJob(ctx context.Context, params *StopRunningJobParams) (*StopRunningJobOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "stopRunningJob",
		Method:             "PUT",
		PathPattern:        "/jobservice/jobs/{job_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StopRunningJobReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	switch value := result.(type) {
	case *StopRunningJobOK:
		return value, nil
	case *StopRunningJobUnauthorized:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *StopRunningJobForbidden:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *StopRunningJobNotFound:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *StopRunningJobInternalServerError:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for stopRunningJob: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}