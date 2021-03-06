// Code generated by go-swagger; DO NOT EDIT.

package users

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

// NewPostAPIPublicV1UserParams creates a new PostAPIPublicV1UserParams object
// with the default values initialized.
func NewPostAPIPublicV1UserParams() *PostAPIPublicV1UserParams {
	var ()
	return &PostAPIPublicV1UserParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostAPIPublicV1UserParamsWithTimeout creates a new PostAPIPublicV1UserParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostAPIPublicV1UserParamsWithTimeout(timeout time.Duration) *PostAPIPublicV1UserParams {
	var ()
	return &PostAPIPublicV1UserParams{

		timeout: timeout,
	}
}

// NewPostAPIPublicV1UserParamsWithContext creates a new PostAPIPublicV1UserParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostAPIPublicV1UserParamsWithContext(ctx context.Context) *PostAPIPublicV1UserParams {
	var ()
	return &PostAPIPublicV1UserParams{

		Context: ctx,
	}
}

// NewPostAPIPublicV1UserParamsWithHTTPClient creates a new PostAPIPublicV1UserParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostAPIPublicV1UserParamsWithHTTPClient(client *http.Client) *PostAPIPublicV1UserParams {
	var ()
	return &PostAPIPublicV1UserParams{
		HTTPClient: client,
	}
}

/*PostAPIPublicV1UserParams contains all the parameters to send to the API endpoint
for the post API public v1 user operation typically these are written to a http.Request
*/
type PostAPIPublicV1UserParams struct {

	/*XVOAPIID
	  Your API ID

	*/
	XVOAPIID string
	/*XVOAPIKey
	  Your API Key

	*/
	XVOAPIKey string
	/*Body*/
	Body *models.AddUserPayload

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post API public v1 user params
func (o *PostAPIPublicV1UserParams) WithTimeout(timeout time.Duration) *PostAPIPublicV1UserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post API public v1 user params
func (o *PostAPIPublicV1UserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post API public v1 user params
func (o *PostAPIPublicV1UserParams) WithContext(ctx context.Context) *PostAPIPublicV1UserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post API public v1 user params
func (o *PostAPIPublicV1UserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post API public v1 user params
func (o *PostAPIPublicV1UserParams) WithHTTPClient(client *http.Client) *PostAPIPublicV1UserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post API public v1 user params
func (o *PostAPIPublicV1UserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXVOAPIID adds the xVOAPIID to the post API public v1 user params
func (o *PostAPIPublicV1UserParams) WithXVOAPIID(xVOAPIID string) *PostAPIPublicV1UserParams {
	o.SetXVOAPIID(xVOAPIID)
	return o
}

// SetXVOAPIID adds the xVOApiId to the post API public v1 user params
func (o *PostAPIPublicV1UserParams) SetXVOAPIID(xVOAPIID string) {
	o.XVOAPIID = xVOAPIID
}

// WithXVOAPIKey adds the xVOAPIKey to the post API public v1 user params
func (o *PostAPIPublicV1UserParams) WithXVOAPIKey(xVOAPIKey string) *PostAPIPublicV1UserParams {
	o.SetXVOAPIKey(xVOAPIKey)
	return o
}

// SetXVOAPIKey adds the xVOApiKey to the post API public v1 user params
func (o *PostAPIPublicV1UserParams) SetXVOAPIKey(xVOAPIKey string) {
	o.XVOAPIKey = xVOAPIKey
}

// WithBody adds the body to the post API public v1 user params
func (o *PostAPIPublicV1UserParams) WithBody(body *models.AddUserPayload) *PostAPIPublicV1UserParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post API public v1 user params
func (o *PostAPIPublicV1UserParams) SetBody(body *models.AddUserPayload) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostAPIPublicV1UserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
