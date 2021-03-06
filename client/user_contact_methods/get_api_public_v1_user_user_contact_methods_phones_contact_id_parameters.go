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

// NewGetAPIPublicV1UserUserContactMethodsPhonesContactIDParams creates a new GetAPIPublicV1UserUserContactMethodsPhonesContactIDParams object
// with the default values initialized.
func NewGetAPIPublicV1UserUserContactMethodsPhonesContactIDParams() *GetAPIPublicV1UserUserContactMethodsPhonesContactIDParams {
	var ()
	return &GetAPIPublicV1UserUserContactMethodsPhonesContactIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIPublicV1UserUserContactMethodsPhonesContactIDParamsWithTimeout creates a new GetAPIPublicV1UserUserContactMethodsPhonesContactIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAPIPublicV1UserUserContactMethodsPhonesContactIDParamsWithTimeout(timeout time.Duration) *GetAPIPublicV1UserUserContactMethodsPhonesContactIDParams {
	var ()
	return &GetAPIPublicV1UserUserContactMethodsPhonesContactIDParams{

		timeout: timeout,
	}
}

// NewGetAPIPublicV1UserUserContactMethodsPhonesContactIDParamsWithContext creates a new GetAPIPublicV1UserUserContactMethodsPhonesContactIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAPIPublicV1UserUserContactMethodsPhonesContactIDParamsWithContext(ctx context.Context) *GetAPIPublicV1UserUserContactMethodsPhonesContactIDParams {
	var ()
	return &GetAPIPublicV1UserUserContactMethodsPhonesContactIDParams{

		Context: ctx,
	}
}

// NewGetAPIPublicV1UserUserContactMethodsPhonesContactIDParamsWithHTTPClient creates a new GetAPIPublicV1UserUserContactMethodsPhonesContactIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAPIPublicV1UserUserContactMethodsPhonesContactIDParamsWithHTTPClient(client *http.Client) *GetAPIPublicV1UserUserContactMethodsPhonesContactIDParams {
	var ()
	return &GetAPIPublicV1UserUserContactMethodsPhonesContactIDParams{
		HTTPClient: client,
	}
}

/*GetAPIPublicV1UserUserContactMethodsPhonesContactIDParams contains all the parameters to send to the API endpoint
for the get API public v1 user user contact methods phones contact ID operation typically these are written to a http.Request
*/
type GetAPIPublicV1UserUserContactMethodsPhonesContactIDParams struct {

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

// WithTimeout adds the timeout to the get API public v1 user user contact methods phones contact ID params
func (o *GetAPIPublicV1UserUserContactMethodsPhonesContactIDParams) WithTimeout(timeout time.Duration) *GetAPIPublicV1UserUserContactMethodsPhonesContactIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API public v1 user user contact methods phones contact ID params
func (o *GetAPIPublicV1UserUserContactMethodsPhonesContactIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API public v1 user user contact methods phones contact ID params
func (o *GetAPIPublicV1UserUserContactMethodsPhonesContactIDParams) WithContext(ctx context.Context) *GetAPIPublicV1UserUserContactMethodsPhonesContactIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API public v1 user user contact methods phones contact ID params
func (o *GetAPIPublicV1UserUserContactMethodsPhonesContactIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API public v1 user user contact methods phones contact ID params
func (o *GetAPIPublicV1UserUserContactMethodsPhonesContactIDParams) WithHTTPClient(client *http.Client) *GetAPIPublicV1UserUserContactMethodsPhonesContactIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API public v1 user user contact methods phones contact ID params
func (o *GetAPIPublicV1UserUserContactMethodsPhonesContactIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXVOAPIID adds the xVOAPIID to the get API public v1 user user contact methods phones contact ID params
func (o *GetAPIPublicV1UserUserContactMethodsPhonesContactIDParams) WithXVOAPIID(xVOAPIID string) *GetAPIPublicV1UserUserContactMethodsPhonesContactIDParams {
	o.SetXVOAPIID(xVOAPIID)
	return o
}

// SetXVOAPIID adds the xVOApiId to the get API public v1 user user contact methods phones contact ID params
func (o *GetAPIPublicV1UserUserContactMethodsPhonesContactIDParams) SetXVOAPIID(xVOAPIID string) {
	o.XVOAPIID = xVOAPIID
}

// WithXVOAPIKey adds the xVOAPIKey to the get API public v1 user user contact methods phones contact ID params
func (o *GetAPIPublicV1UserUserContactMethodsPhonesContactIDParams) WithXVOAPIKey(xVOAPIKey string) *GetAPIPublicV1UserUserContactMethodsPhonesContactIDParams {
	o.SetXVOAPIKey(xVOAPIKey)
	return o
}

// SetXVOAPIKey adds the xVOApiKey to the get API public v1 user user contact methods phones contact ID params
func (o *GetAPIPublicV1UserUserContactMethodsPhonesContactIDParams) SetXVOAPIKey(xVOAPIKey string) {
	o.XVOAPIKey = xVOAPIKey
}

// WithContactID adds the contactID to the get API public v1 user user contact methods phones contact ID params
func (o *GetAPIPublicV1UserUserContactMethodsPhonesContactIDParams) WithContactID(contactID string) *GetAPIPublicV1UserUserContactMethodsPhonesContactIDParams {
	o.SetContactID(contactID)
	return o
}

// SetContactID adds the contactId to the get API public v1 user user contact methods phones contact ID params
func (o *GetAPIPublicV1UserUserContactMethodsPhonesContactIDParams) SetContactID(contactID string) {
	o.ContactID = contactID
}

// WithUser adds the user to the get API public v1 user user contact methods phones contact ID params
func (o *GetAPIPublicV1UserUserContactMethodsPhonesContactIDParams) WithUser(user string) *GetAPIPublicV1UserUserContactMethodsPhonesContactIDParams {
	o.SetUser(user)
	return o
}

// SetUser adds the user to the get API public v1 user user contact methods phones contact ID params
func (o *GetAPIPublicV1UserUserContactMethodsPhonesContactIDParams) SetUser(user string) {
	o.User = user
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIPublicV1UserUserContactMethodsPhonesContactIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
