// Code generated by go-swagger; DO NOT EDIT.

package webhookjob

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewListWebhookJobsParams creates a new ListWebhookJobsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListWebhookJobsParams() *ListWebhookJobsParams {
	return &ListWebhookJobsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListWebhookJobsParamsWithTimeout creates a new ListWebhookJobsParams object
// with the ability to set a timeout on a request.
func NewListWebhookJobsParamsWithTimeout(timeout time.Duration) *ListWebhookJobsParams {
	return &ListWebhookJobsParams{
		timeout: timeout,
	}
}

// NewListWebhookJobsParamsWithContext creates a new ListWebhookJobsParams object
// with the ability to set a context for a request.
func NewListWebhookJobsParamsWithContext(ctx context.Context) *ListWebhookJobsParams {
	return &ListWebhookJobsParams{
		Context: ctx,
	}
}

// NewListWebhookJobsParamsWithHTTPClient creates a new ListWebhookJobsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListWebhookJobsParamsWithHTTPClient(client *http.Client) *ListWebhookJobsParams {
	return &ListWebhookJobsParams{
		HTTPClient: client,
	}
}

/*
ListWebhookJobsParams contains all the parameters to send to the API endpoint

	for the list webhook jobs operation.

	Typically these are written to a http.Request.
*/
type ListWebhookJobsParams struct {

	/* XIsResourceName.

	   The flag to indicate whether the parameter which supports both name and id in the path is the name of the resource. When the X-Is-Resource-Name is false and the parameter can be converted to an integer, the parameter will be as an id, otherwise, it will be as a name.
	*/
	XIsResourceName *bool `js:"xIsResourceName"`

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string `js:"xRequestID"`

	/* Page.

	   The page number

	   Format: int64
	   Default: 1
	*/
	Page *int64 `js:"page"`

	/* PageSize.

	   The size of per page

	   Format: int64
	   Default: 10
	*/
	PageSize *int64 `js:"pageSize"`

	/* PolicyID.

	   The policy ID.

	   Format: int64
	*/
	PolicyID int64 `js:"policyID"`

	/* ProjectNameOrID.

	   The name or id of the project
	*/
	ProjectNameOrID string `js:"projectNameOrID"`

	/* Q.

	   Query string to query resources. Supported query patterns are "exact match(k=v)", "fuzzy match(k=~v)", "range(k=[min~max])", "list with union releationship(k={v1 v2 v3})" and "list with intersetion relationship(k=(v1 v2 v3))". The value of range and list can be string(enclosed by " or '), integer or time(in format "2020-04-09 02:36:00"). All of these query patterns should be put in the query string "q=xxx" and splitted by ",". e.g. q=k1=v1,k2=~v2,k3=[min~max]
	*/
	Q *string `js:"q"`

	/* Sort.

	   Sort the resource list in ascending or descending order. e.g. sort by field1 in ascending orderr and field2 in descending order with "sort=field1,-field2"
	*/
	Sort *string `js:"sort"`

	/* Status.

	   The status of webhook job.
	*/
	Status []string `js:"status"`

	timeout    time.Duration
	Context    context.Context `js:"context"`
	HTTPClient *http.Client    `js:"httpClient"`
}

// WithDefaults hydrates default values in the list webhook jobs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListWebhookJobsParams) WithDefaults() *ListWebhookJobsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list webhook jobs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListWebhookJobsParams) SetDefaults() {
	var (
		xIsResourceNameDefault = bool(false)

		pageDefault = int64(1)

		pageSizeDefault = int64(10)
	)

	val := ListWebhookJobsParams{
		XIsResourceName: &xIsResourceNameDefault,
		Page:            &pageDefault,
		PageSize:        &pageSizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the list webhook jobs params
func (o *ListWebhookJobsParams) WithTimeout(timeout time.Duration) *ListWebhookJobsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list webhook jobs params
func (o *ListWebhookJobsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list webhook jobs params
func (o *ListWebhookJobsParams) WithContext(ctx context.Context) *ListWebhookJobsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list webhook jobs params
func (o *ListWebhookJobsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list webhook jobs params
func (o *ListWebhookJobsParams) WithHTTPClient(client *http.Client) *ListWebhookJobsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list webhook jobs params
func (o *ListWebhookJobsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXIsResourceName adds the xIsResourceName to the list webhook jobs params
func (o *ListWebhookJobsParams) WithXIsResourceName(xIsResourceName *bool) *ListWebhookJobsParams {
	o.SetXIsResourceName(xIsResourceName)
	return o
}

// SetXIsResourceName adds the xIsResourceName to the list webhook jobs params
func (o *ListWebhookJobsParams) SetXIsResourceName(xIsResourceName *bool) {
	o.XIsResourceName = xIsResourceName
}

// WithXRequestID adds the xRequestID to the list webhook jobs params
func (o *ListWebhookJobsParams) WithXRequestID(xRequestID *string) *ListWebhookJobsParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the list webhook jobs params
func (o *ListWebhookJobsParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithPage adds the page to the list webhook jobs params
func (o *ListWebhookJobsParams) WithPage(page *int64) *ListWebhookJobsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the list webhook jobs params
func (o *ListWebhookJobsParams) SetPage(page *int64) {
	o.Page = page
}

// WithPageSize adds the pageSize to the list webhook jobs params
func (o *ListWebhookJobsParams) WithPageSize(pageSize *int64) *ListWebhookJobsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the list webhook jobs params
func (o *ListWebhookJobsParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithPolicyID adds the policyID to the list webhook jobs params
func (o *ListWebhookJobsParams) WithPolicyID(policyID int64) *ListWebhookJobsParams {
	o.SetPolicyID(policyID)
	return o
}

// SetPolicyID adds the policyId to the list webhook jobs params
func (o *ListWebhookJobsParams) SetPolicyID(policyID int64) {
	o.PolicyID = policyID
}

// WithProjectNameOrID adds the projectNameOrID to the list webhook jobs params
func (o *ListWebhookJobsParams) WithProjectNameOrID(projectNameOrID string) *ListWebhookJobsParams {
	o.SetProjectNameOrID(projectNameOrID)
	return o
}

// SetProjectNameOrID adds the projectNameOrId to the list webhook jobs params
func (o *ListWebhookJobsParams) SetProjectNameOrID(projectNameOrID string) {
	o.ProjectNameOrID = projectNameOrID
}

// WithQ adds the q to the list webhook jobs params
func (o *ListWebhookJobsParams) WithQ(q *string) *ListWebhookJobsParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the list webhook jobs params
func (o *ListWebhookJobsParams) SetQ(q *string) {
	o.Q = q
}

// WithSort adds the sort to the list webhook jobs params
func (o *ListWebhookJobsParams) WithSort(sort *string) *ListWebhookJobsParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the list webhook jobs params
func (o *ListWebhookJobsParams) SetSort(sort *string) {
	o.Sort = sort
}

// WithStatus adds the status to the list webhook jobs params
func (o *ListWebhookJobsParams) WithStatus(status []string) *ListWebhookJobsParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the list webhook jobs params
func (o *ListWebhookJobsParams) SetStatus(status []string) {
	o.Status = status
}

// WriteToRequest writes these params to a swagger request
func (o *ListWebhookJobsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XIsResourceName != nil {

		// header param X-Is-Resource-Name
		if err := r.SetHeaderParam("X-Is-Resource-Name", swag.FormatBool(*o.XIsResourceName)); err != nil {
			return err
		}
	}

	if o.XRequestID != nil {

		// header param X-Request-Id
		if err := r.SetHeaderParam("X-Request-Id", *o.XRequestID); err != nil {
			return err
		}
	}

	if o.Page != nil {

		// query param page
		var qrPage int64

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	if o.PageSize != nil {

		// query param page_size
		var qrPageSize int64

		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt64(qrPageSize)
		if qPageSize != "" {

			if err := r.SetQueryParam("page_size", qPageSize); err != nil {
				return err
			}
		}
	}

	// query param policy_id
	qrPolicyID := o.PolicyID
	qPolicyID := swag.FormatInt64(qrPolicyID)
	if qPolicyID != "" {

		if err := r.SetQueryParam("policy_id", qPolicyID); err != nil {
			return err
		}
	}

	// path param project_name_or_id
	if err := r.SetPathParam("project_name_or_id", o.ProjectNameOrID); err != nil {
		return err
	}

	if o.Q != nil {

		// query param q
		var qrQ string

		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {

			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}
	}

	if o.Sort != nil {

		// query param sort
		var qrSort string

		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {

			if err := r.SetQueryParam("sort", qSort); err != nil {
				return err
			}
		}
	}

	if o.Status != nil {

		// binding items for status
		joinedStatus := o.bindParamStatus(reg)

		// query array param status
		if err := r.SetQueryParam("status", joinedStatus...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamListWebhookJobs binds the parameter status
func (o *ListWebhookJobsParams) bindParamStatus(formats strfmt.Registry) []string {
	statusIR := o.Status

	var statusIC []string
	for _, statusIIR := range statusIR { // explode []string

		statusIIV := statusIIR // string as string
		statusIC = append(statusIC, statusIIV)
	}

	// items.CollectionFormat: ""
	statusIS := swag.JoinByFormat(statusIC, "")

	return statusIS
}
