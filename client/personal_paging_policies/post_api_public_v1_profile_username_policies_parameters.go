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

	strfmt "github.com/go-openapi/strfmt"

	"github.com/barkerd427/go-vo/models"
)

// NewPostAPIPublicV1ProfileUsernamePoliciesParams creates a new PostAPIPublicV1ProfileUsernamePoliciesParams object
// with the default values initialized.
func NewPostAPIPublicV1ProfileUsernamePoliciesParams() *PostAPIPublicV1ProfileUsernamePoliciesParams {
	var ()
	return &PostAPIPublicV1ProfileUsernamePoliciesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostAPIPublicV1ProfileUsernamePoliciesParamsWithTimeout creates a new PostAPIPublicV1ProfileUsernamePoliciesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostAPIPublicV1ProfileUsernamePoliciesParamsWithTimeout(timeout time.Duration) *PostAPIPublicV1ProfileUsernamePoliciesParams {
	var ()
	return &PostAPIPublicV1ProfileUsernamePoliciesParams{

		timeout: timeout,
	}
}

// NewPostAPIPublicV1ProfileUsernamePoliciesParamsWithContext creates a new PostAPIPublicV1ProfileUsernamePoliciesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostAPIPublicV1ProfileUsernamePoliciesParamsWithContext(ctx context.Context) *PostAPIPublicV1ProfileUsernamePoliciesParams {
	var ()
	return &PostAPIPublicV1ProfileUsernamePoliciesParams{

		Context: ctx,
	}
}

// NewPostAPIPublicV1ProfileUsernamePoliciesParamsWithHTTPClient creates a new PostAPIPublicV1ProfileUsernamePoliciesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostAPIPublicV1ProfileUsernamePoliciesParamsWithHTTPClient(client *http.Client) *PostAPIPublicV1ProfileUsernamePoliciesParams {
	var ()
	return &PostAPIPublicV1ProfileUsernamePoliciesParams{
		HTTPClient: client,
	}
}

/*PostAPIPublicV1ProfileUsernamePoliciesParams contains all the parameters to send to the API endpoint
for the post API public v1 profile username policies operation typically these are written to a http.Request
*/
type PostAPIPublicV1ProfileUsernamePoliciesParams struct {

	/*XVOAPIID
	  Your API ID

	*/
	XVOAPIID string
	/*XVOAPIKey
	  Your API Key

	*/
	XVOAPIKey string
	/*Body*/
	Body *models.AddGroupPayload
	/*Username
	  Your username

	*/
	Username string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post API public v1 profile username policies params
func (o *PostAPIPublicV1ProfileUsernamePoliciesParams) WithTimeout(timeout time.Duration) *PostAPIPublicV1ProfileUsernamePoliciesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post API public v1 profile username policies params
func (o *PostAPIPublicV1ProfileUsernamePoliciesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post API public v1 profile username policies params
func (o *PostAPIPublicV1ProfileUsernamePoliciesParams) WithContext(ctx context.Context) *PostAPIPublicV1ProfileUsernamePoliciesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post API public v1 profile username policies params
func (o *PostAPIPublicV1ProfileUsernamePoliciesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post API public v1 profile username policies params
func (o *PostAPIPublicV1ProfileUsernamePoliciesParams) WithHTTPClient(client *http.Client) *PostAPIPublicV1ProfileUsernamePoliciesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post API public v1 profile username policies params
func (o *PostAPIPublicV1ProfileUsernamePoliciesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXVOAPIID adds the xVOAPIID to the post API public v1 profile username policies params
func (o *PostAPIPublicV1ProfileUsernamePoliciesParams) WithXVOAPIID(xVOAPIID string) *PostAPIPublicV1ProfileUsernamePoliciesParams {
	o.SetXVOAPIID(xVOAPIID)
	return o
}

// SetXVOAPIID adds the xVOApiId to the post API public v1 profile username policies params
func (o *PostAPIPublicV1ProfileUsernamePoliciesParams) SetXVOAPIID(xVOAPIID string) {
	o.XVOAPIID = xVOAPIID
}

// WithXVOAPIKey adds the xVOAPIKey to the post API public v1 profile username policies params
func (o *PostAPIPublicV1ProfileUsernamePoliciesParams) WithXVOAPIKey(xVOAPIKey string) *PostAPIPublicV1ProfileUsernamePoliciesParams {
	o.SetXVOAPIKey(xVOAPIKey)
	return o
}

// SetXVOAPIKey adds the xVOApiKey to the post API public v1 profile username policies params
func (o *PostAPIPublicV1ProfileUsernamePoliciesParams) SetXVOAPIKey(xVOAPIKey string) {
	o.XVOAPIKey = xVOAPIKey
}

// WithBody adds the body to the post API public v1 profile username policies params
func (o *PostAPIPublicV1ProfileUsernamePoliciesParams) WithBody(body *models.AddGroupPayload) *PostAPIPublicV1ProfileUsernamePoliciesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post API public v1 profile username policies params
func (o *PostAPIPublicV1ProfileUsernamePoliciesParams) SetBody(body *models.AddGroupPayload) {
	o.Body = body
}

// WithUsername adds the username to the post API public v1 profile username policies params
func (o *PostAPIPublicV1ProfileUsernamePoliciesParams) WithUsername(username string) *PostAPIPublicV1ProfileUsernamePoliciesParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the post API public v1 profile username policies params
func (o *PostAPIPublicV1ProfileUsernamePoliciesParams) SetUsername(username string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *PostAPIPublicV1ProfileUsernamePoliciesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param username
	if err := r.SetPathParam("username", o.Username); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}