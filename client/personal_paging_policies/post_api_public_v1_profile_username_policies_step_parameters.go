// Code generated by go-swagger; DO NOT EDIT.

package personal_paging_policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/barkerd427/go-vo/models"
)

// NewPostAPIPublicV1ProfileUsernamePoliciesStepParams creates a new PostAPIPublicV1ProfileUsernamePoliciesStepParams object
// with the default values initialized.
func NewPostAPIPublicV1ProfileUsernamePoliciesStepParams() *PostAPIPublicV1ProfileUsernamePoliciesStepParams {
	var ()
	return &PostAPIPublicV1ProfileUsernamePoliciesStepParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostAPIPublicV1ProfileUsernamePoliciesStepParamsWithTimeout creates a new PostAPIPublicV1ProfileUsernamePoliciesStepParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostAPIPublicV1ProfileUsernamePoliciesStepParamsWithTimeout(timeout time.Duration) *PostAPIPublicV1ProfileUsernamePoliciesStepParams {
	var ()
	return &PostAPIPublicV1ProfileUsernamePoliciesStepParams{

		timeout: timeout,
	}
}

// NewPostAPIPublicV1ProfileUsernamePoliciesStepParamsWithContext creates a new PostAPIPublicV1ProfileUsernamePoliciesStepParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostAPIPublicV1ProfileUsernamePoliciesStepParamsWithContext(ctx context.Context) *PostAPIPublicV1ProfileUsernamePoliciesStepParams {
	var ()
	return &PostAPIPublicV1ProfileUsernamePoliciesStepParams{

		Context: ctx,
	}
}

// NewPostAPIPublicV1ProfileUsernamePoliciesStepParamsWithHTTPClient creates a new PostAPIPublicV1ProfileUsernamePoliciesStepParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostAPIPublicV1ProfileUsernamePoliciesStepParamsWithHTTPClient(client *http.Client) *PostAPIPublicV1ProfileUsernamePoliciesStepParams {
	var ()
	return &PostAPIPublicV1ProfileUsernamePoliciesStepParams{
		HTTPClient: client,
	}
}

/*PostAPIPublicV1ProfileUsernamePoliciesStepParams contains all the parameters to send to the API endpoint
for the post API public v1 profile username policies step operation typically these are written to a http.Request
*/
type PostAPIPublicV1ProfileUsernamePoliciesStepParams struct {

	/*XVOAPIID
	  Your API ID

	*/
	XVOAPIID string
	/*XVOAPIKey
	  Your API Key

	*/
	XVOAPIKey string
	/*Body*/
	Body *models.AddStepPayload
	/*Step
	  Paging policy step

	*/
	Step float64
	/*Username
	  Your username

	*/
	Username string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post API public v1 profile username policies step params
func (o *PostAPIPublicV1ProfileUsernamePoliciesStepParams) WithTimeout(timeout time.Duration) *PostAPIPublicV1ProfileUsernamePoliciesStepParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post API public v1 profile username policies step params
func (o *PostAPIPublicV1ProfileUsernamePoliciesStepParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post API public v1 profile username policies step params
func (o *PostAPIPublicV1ProfileUsernamePoliciesStepParams) WithContext(ctx context.Context) *PostAPIPublicV1ProfileUsernamePoliciesStepParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post API public v1 profile username policies step params
func (o *PostAPIPublicV1ProfileUsernamePoliciesStepParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post API public v1 profile username policies step params
func (o *PostAPIPublicV1ProfileUsernamePoliciesStepParams) WithHTTPClient(client *http.Client) *PostAPIPublicV1ProfileUsernamePoliciesStepParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post API public v1 profile username policies step params
func (o *PostAPIPublicV1ProfileUsernamePoliciesStepParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXVOAPIID adds the xVOAPIID to the post API public v1 profile username policies step params
func (o *PostAPIPublicV1ProfileUsernamePoliciesStepParams) WithXVOAPIID(xVOAPIID string) *PostAPIPublicV1ProfileUsernamePoliciesStepParams {
	o.SetXVOAPIID(xVOAPIID)
	return o
}

// SetXVOAPIID adds the xVOApiId to the post API public v1 profile username policies step params
func (o *PostAPIPublicV1ProfileUsernamePoliciesStepParams) SetXVOAPIID(xVOAPIID string) {
	o.XVOAPIID = xVOAPIID
}

// WithXVOAPIKey adds the xVOAPIKey to the post API public v1 profile username policies step params
func (o *PostAPIPublicV1ProfileUsernamePoliciesStepParams) WithXVOAPIKey(xVOAPIKey string) *PostAPIPublicV1ProfileUsernamePoliciesStepParams {
	o.SetXVOAPIKey(xVOAPIKey)
	return o
}

// SetXVOAPIKey adds the xVOApiKey to the post API public v1 profile username policies step params
func (o *PostAPIPublicV1ProfileUsernamePoliciesStepParams) SetXVOAPIKey(xVOAPIKey string) {
	o.XVOAPIKey = xVOAPIKey
}

// WithBody adds the body to the post API public v1 profile username policies step params
func (o *PostAPIPublicV1ProfileUsernamePoliciesStepParams) WithBody(body *models.AddStepPayload) *PostAPIPublicV1ProfileUsernamePoliciesStepParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post API public v1 profile username policies step params
func (o *PostAPIPublicV1ProfileUsernamePoliciesStepParams) SetBody(body *models.AddStepPayload) {
	o.Body = body
}

// WithStep adds the step to the post API public v1 profile username policies step params
func (o *PostAPIPublicV1ProfileUsernamePoliciesStepParams) WithStep(step float64) *PostAPIPublicV1ProfileUsernamePoliciesStepParams {
	o.SetStep(step)
	return o
}

// SetStep adds the step to the post API public v1 profile username policies step params
func (o *PostAPIPublicV1ProfileUsernamePoliciesStepParams) SetStep(step float64) {
	o.Step = step
}

// WithUsername adds the username to the post API public v1 profile username policies step params
func (o *PostAPIPublicV1ProfileUsernamePoliciesStepParams) WithUsername(username string) *PostAPIPublicV1ProfileUsernamePoliciesStepParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the post API public v1 profile username policies step params
func (o *PostAPIPublicV1ProfileUsernamePoliciesStepParams) SetUsername(username string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *PostAPIPublicV1ProfileUsernamePoliciesStepParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param step
	if err := r.SetPathParam("step", swag.FormatFloat64(o.Step)); err != nil {
		return err
	}

	// path param username
	if err := r.SetPathParam("username", o.Username); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
