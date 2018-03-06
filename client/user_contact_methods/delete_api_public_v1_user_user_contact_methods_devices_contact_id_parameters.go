// Code generated by go-swagger; DO NOT EDIT.

package user_contact_methods

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

// NewDeleteAPIPublicV1UserUserContactMethodsDevicesContactIDParams creates a new DeleteAPIPublicV1UserUserContactMethodsDevicesContactIDParams object
// with the default values initialized.
func NewDeleteAPIPublicV1UserUserContactMethodsDevicesContactIDParams() *DeleteAPIPublicV1UserUserContactMethodsDevicesContactIDParams {
	var ()
	return &DeleteAPIPublicV1UserUserContactMethodsDevicesContactIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAPIPublicV1UserUserContactMethodsDevicesContactIDParamsWithTimeout creates a new DeleteAPIPublicV1UserUserContactMethodsDevicesContactIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteAPIPublicV1UserUserContactMethodsDevicesContactIDParamsWithTimeout(timeout time.Duration) *DeleteAPIPublicV1UserUserContactMethodsDevicesContactIDParams {
	var ()
	return &DeleteAPIPublicV1UserUserContactMethodsDevicesContactIDParams{

		timeout: timeout,
	}
}

// NewDeleteAPIPublicV1UserUserContactMethodsDevicesContactIDParamsWithContext creates a new DeleteAPIPublicV1UserUserContactMethodsDevicesContactIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteAPIPublicV1UserUserContactMethodsDevicesContactIDParamsWithContext(ctx context.Context) *DeleteAPIPublicV1UserUserContactMethodsDevicesContactIDParams {
	var ()
	return &DeleteAPIPublicV1UserUserContactMethodsDevicesContactIDParams{

		Context: ctx,
	}
}

// NewDeleteAPIPublicV1UserUserContactMethodsDevicesContactIDParamsWithHTTPClient creates a new DeleteAPIPublicV1UserUserContactMethodsDevicesContactIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteAPIPublicV1UserUserContactMethodsDevicesContactIDParamsWithHTTPClient(client *http.Client) *DeleteAPIPublicV1UserUserContactMethodsDevicesContactIDParams {
	var ()
	return &DeleteAPIPublicV1UserUserContactMethodsDevicesContactIDParams{
		HTTPClient: client,
	}
}

/*DeleteAPIPublicV1UserUserContactMethodsDevicesContactIDParams contains all the parameters to send to the API endpoint
for the delete API public v1 user user contact methods devices contact ID operation typically these are written to a http.Request
*/
type DeleteAPIPublicV1UserUserContactMethodsDevicesContactIDParams struct {

	/*XVOAPIID
	  Your API ID

	*/
	XVOAPIID string
	/*XVOAPIKey
	  Your API Key

	*/
	XVOAPIKey string
	/*ContactID
	  The unique contact identifier

	*/
	ContactID string
	/*User
	  The VictorOps user ID

	*/
	User string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete API public v1 user user contact methods devices contact ID params
func (o *DeleteAPIPublicV1UserUserContactMethodsDevicesContactIDParams) WithTimeout(timeout time.Duration) *DeleteAPIPublicV1UserUserContactMethodsDevicesContactIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete API public v1 user user contact methods devices contact ID params
func (o *DeleteAPIPublicV1UserUserContactMethodsDevicesContactIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete API public v1 user user contact methods devices contact ID params
func (o *DeleteAPIPublicV1UserUserContactMethodsDevicesContactIDParams) WithContext(ctx context.Context) *DeleteAPIPublicV1UserUserContactMethodsDevicesContactIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete API public v1 user user contact methods devices contact ID params
func (o *DeleteAPIPublicV1UserUserContactMethodsDevicesContactIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete API public v1 user user contact methods devices contact ID params
func (o *DeleteAPIPublicV1UserUserContactMethodsDevicesContactIDParams) WithHTTPClient(client *http.Client) *DeleteAPIPublicV1UserUserContactMethodsDevicesContactIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete API public v1 user user contact methods devices contact ID params
func (o *DeleteAPIPublicV1UserUserContactMethodsDevicesContactIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXVOAPIID adds the xVOAPIID to the delete API public v1 user user contact methods devices contact ID params
func (o *DeleteAPIPublicV1UserUserContactMethodsDevicesContactIDParams) WithXVOAPIID(xVOAPIID string) *DeleteAPIPublicV1UserUserContactMethodsDevicesContactIDParams {
	o.SetXVOAPIID(xVOAPIID)
	return o
}

// SetXVOAPIID adds the xVOApiId to the delete API public v1 user user contact methods devices contact ID params
func (o *DeleteAPIPublicV1UserUserContactMethodsDevicesContactIDParams) SetXVOAPIID(xVOAPIID string) {
	o.XVOAPIID = xVOAPIID
}

// WithXVOAPIKey adds the xVOAPIKey to the delete API public v1 user user contact methods devices contact ID params
func (o *DeleteAPIPublicV1UserUserContactMethodsDevicesContactIDParams) WithXVOAPIKey(xVOAPIKey string) *DeleteAPIPublicV1UserUserContactMethodsDevicesContactIDParams {
	o.SetXVOAPIKey(xVOAPIKey)
	return o
}

// SetXVOAPIKey adds the xVOApiKey to the delete API public v1 user user contact methods devices contact ID params
func (o *DeleteAPIPublicV1UserUserContactMethodsDevicesContactIDParams) SetXVOAPIKey(xVOAPIKey string) {
	o.XVOAPIKey = xVOAPIKey
}

// WithContactID adds the contactID to the delete API public v1 user user contact methods devices contact ID params
func (o *DeleteAPIPublicV1UserUserContactMethodsDevicesContactIDParams) WithContactID(contactID string) *DeleteAPIPublicV1UserUserContactMethodsDevicesContactIDParams {
	o.SetContactID(contactID)
	return o
}

// SetContactID adds the contactId to the delete API public v1 user user contact methods devices contact ID params
func (o *DeleteAPIPublicV1UserUserContactMethodsDevicesContactIDParams) SetContactID(contactID string) {
	o.ContactID = contactID
}

// WithUser adds the user to the delete API public v1 user user contact methods devices contact ID params
func (o *DeleteAPIPublicV1UserUserContactMethodsDevicesContactIDParams) WithUser(user string) *DeleteAPIPublicV1UserUserContactMethodsDevicesContactIDParams {
	o.SetUser(user)
	return o
}

// SetUser adds the user to the delete API public v1 user user contact methods devices contact ID params
func (o *DeleteAPIPublicV1UserUserContactMethodsDevicesContactIDParams) SetUser(user string) {
	o.User = user
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAPIPublicV1UserUserContactMethodsDevicesContactIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param contactId
	if err := r.SetPathParam("contactId", o.ContactID); err != nil {
		return err
	}

	// path param user
	if err := r.SetPathParam("user", o.User); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}