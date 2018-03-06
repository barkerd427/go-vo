// Code generated by go-swagger; DO NOT EDIT.

package personal_paging_policy_values

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetAPIPublicV1PoliciesTypesTimeoutsParams creates a new GetAPIPublicV1PoliciesTypesTimeoutsParams object
// with the default values initialized.
func NewGetAPIPublicV1PoliciesTypesTimeoutsParams() *GetAPIPublicV1PoliciesTypesTimeoutsParams {
	var ()
	return &GetAPIPublicV1PoliciesTypesTimeoutsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIPublicV1PoliciesTypesTimeoutsParamsWithTimeout creates a new GetAPIPublicV1PoliciesTypesTimeoutsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAPIPublicV1PoliciesTypesTimeoutsParamsWithTimeout(timeout time.Duration) *GetAPIPublicV1PoliciesTypesTimeoutsParams {
	var ()
	return &GetAPIPublicV1PoliciesTypesTimeoutsParams{

		timeout: timeout,
	}
}

// NewGetAPIPublicV1PoliciesTypesTimeoutsParamsWithContext creates a new GetAPIPublicV1PoliciesTypesTimeoutsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAPIPublicV1PoliciesTypesTimeoutsParamsWithContext(ctx context.Context) *GetAPIPublicV1PoliciesTypesTimeoutsParams {
	var ()
	return &GetAPIPublicV1PoliciesTypesTimeoutsParams{

		Context: ctx,
	}
}

// NewGetAPIPublicV1PoliciesTypesTimeoutsParamsWithHTTPClient creates a new GetAPIPublicV1PoliciesTypesTimeoutsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAPIPublicV1PoliciesTypesTimeoutsParamsWithHTTPClient(client *http.Client) *GetAPIPublicV1PoliciesTypesTimeoutsParams {
	var ()
	return &GetAPIPublicV1PoliciesTypesTimeoutsParams{
		HTTPClient: client,
	}
}

/*GetAPIPublicV1PoliciesTypesTimeoutsParams contains all the parameters to send to the API endpoint
for the get API public v1 policies types timeouts operation typically these are written to a http.Request
*/
type GetAPIPublicV1PoliciesTypesTimeoutsParams struct {

	/*XVOAPIID
	  Your API ID

	*/
	XVOAPIID string
	/*XVOAPIKey
	  Your API Key

	*/
	XVOAPIKey string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get API public v1 policies types timeouts params
func (o *GetAPIPublicV1PoliciesTypesTimeoutsParams) WithTimeout(timeout time.Duration) *GetAPIPublicV1PoliciesTypesTimeoutsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API public v1 policies types timeouts params
func (o *GetAPIPublicV1PoliciesTypesTimeoutsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API public v1 policies types timeouts params
func (o *GetAPIPublicV1PoliciesTypesTimeoutsParams) WithContext(ctx context.Context) *GetAPIPublicV1PoliciesTypesTimeoutsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API public v1 policies types timeouts params
func (o *GetAPIPublicV1PoliciesTypesTimeoutsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API public v1 policies types timeouts params
func (o *GetAPIPublicV1PoliciesTypesTimeoutsParams) WithHTTPClient(client *http.Client) *GetAPIPublicV1PoliciesTypesTimeoutsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API public v1 policies types timeouts params
func (o *GetAPIPublicV1PoliciesTypesTimeoutsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXVOAPIID adds the xVOAPIID to the get API public v1 policies types timeouts params
func (o *GetAPIPublicV1PoliciesTypesTimeoutsParams) WithXVOAPIID(xVOAPIID string) *GetAPIPublicV1PoliciesTypesTimeoutsParams {
	o.SetXVOAPIID(xVOAPIID)
	return o
}

// SetXVOAPIID adds the xVOApiId to the get API public v1 policies types timeouts params
func (o *GetAPIPublicV1PoliciesTypesTimeoutsParams) SetXVOAPIID(xVOAPIID string) {
	o.XVOAPIID = xVOAPIID
}

// WithXVOAPIKey adds the xVOAPIKey to the get API public v1 policies types timeouts params
func (o *GetAPIPublicV1PoliciesTypesTimeoutsParams) WithXVOAPIKey(xVOAPIKey string) *GetAPIPublicV1PoliciesTypesTimeoutsParams {
	o.SetXVOAPIKey(xVOAPIKey)
	return o
}

// SetXVOAPIKey adds the xVOApiKey to the get API public v1 policies types timeouts params
func (o *GetAPIPublicV1PoliciesTypesTimeoutsParams) SetXVOAPIKey(xVOAPIKey string) {
	o.XVOAPIKey = xVOAPIKey
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIPublicV1PoliciesTypesTimeoutsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-VO-Api-Id
	if err := r.SetHeaderParam("X-VO-Api-Id", o.XVOAPIID); err != nil {
		return err
	}

	// header param X-VO-Api-Key
	if err := r.SetHeaderParam("X-VO-Api-Key", o.XVOAPIKey); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}