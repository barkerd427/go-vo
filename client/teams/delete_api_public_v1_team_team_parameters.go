// Code generated by go-swagger; DO NOT EDIT.

package teams

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

// NewDeleteAPIPublicV1TeamTeamParams creates a new DeleteAPIPublicV1TeamTeamParams object
// with the default values initialized.
func NewDeleteAPIPublicV1TeamTeamParams() *DeleteAPIPublicV1TeamTeamParams {
	var ()
	return &DeleteAPIPublicV1TeamTeamParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAPIPublicV1TeamTeamParamsWithTimeout creates a new DeleteAPIPublicV1TeamTeamParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteAPIPublicV1TeamTeamParamsWithTimeout(timeout time.Duration) *DeleteAPIPublicV1TeamTeamParams {
	var ()
	return &DeleteAPIPublicV1TeamTeamParams{

		timeout: timeout,
	}
}

// NewDeleteAPIPublicV1TeamTeamParamsWithContext creates a new DeleteAPIPublicV1TeamTeamParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteAPIPublicV1TeamTeamParamsWithContext(ctx context.Context) *DeleteAPIPublicV1TeamTeamParams {
	var ()
	return &DeleteAPIPublicV1TeamTeamParams{

		Context: ctx,
	}
}

// NewDeleteAPIPublicV1TeamTeamParamsWithHTTPClient creates a new DeleteAPIPublicV1TeamTeamParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteAPIPublicV1TeamTeamParamsWithHTTPClient(client *http.Client) *DeleteAPIPublicV1TeamTeamParams {
	var ()
	return &DeleteAPIPublicV1TeamTeamParams{
		HTTPClient: client,
	}
}

/*DeleteAPIPublicV1TeamTeamParams contains all the parameters to send to the API endpoint
for the delete API public v1 team team operation typically these are written to a http.Request
*/
type DeleteAPIPublicV1TeamTeamParams struct {

	/*XVOAPIID
	  Your API ID

	*/
	XVOAPIID string
	/*XVOAPIKey
	  Your API Key

	*/
	XVOAPIKey string
	/*Team
	  The VictorOps team to be deleted

	*/
	Team string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete API public v1 team team params
func (o *DeleteAPIPublicV1TeamTeamParams) WithTimeout(timeout time.Duration) *DeleteAPIPublicV1TeamTeamParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete API public v1 team team params
func (o *DeleteAPIPublicV1TeamTeamParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete API public v1 team team params
func (o *DeleteAPIPublicV1TeamTeamParams) WithContext(ctx context.Context) *DeleteAPIPublicV1TeamTeamParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete API public v1 team team params
func (o *DeleteAPIPublicV1TeamTeamParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete API public v1 team team params
func (o *DeleteAPIPublicV1TeamTeamParams) WithHTTPClient(client *http.Client) *DeleteAPIPublicV1TeamTeamParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete API public v1 team team params
func (o *DeleteAPIPublicV1TeamTeamParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXVOAPIID adds the xVOAPIID to the delete API public v1 team team params
func (o *DeleteAPIPublicV1TeamTeamParams) WithXVOAPIID(xVOAPIID string) *DeleteAPIPublicV1TeamTeamParams {
	o.SetXVOAPIID(xVOAPIID)
	return o
}

// SetXVOAPIID adds the xVOApiId to the delete API public v1 team team params
func (o *DeleteAPIPublicV1TeamTeamParams) SetXVOAPIID(xVOAPIID string) {
	o.XVOAPIID = xVOAPIID
}

// WithXVOAPIKey adds the xVOAPIKey to the delete API public v1 team team params
func (o *DeleteAPIPublicV1TeamTeamParams) WithXVOAPIKey(xVOAPIKey string) *DeleteAPIPublicV1TeamTeamParams {
	o.SetXVOAPIKey(xVOAPIKey)
	return o
}

// SetXVOAPIKey adds the xVOApiKey to the delete API public v1 team team params
func (o *DeleteAPIPublicV1TeamTeamParams) SetXVOAPIKey(xVOAPIKey string) {
	o.XVOAPIKey = xVOAPIKey
}

// WithTeam adds the team to the delete API public v1 team team params
func (o *DeleteAPIPublicV1TeamTeamParams) WithTeam(team string) *DeleteAPIPublicV1TeamTeamParams {
	o.SetTeam(team)
	return o
}

// SetTeam adds the team to the delete API public v1 team team params
func (o *DeleteAPIPublicV1TeamTeamParams) SetTeam(team string) {
	o.Team = team
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAPIPublicV1TeamTeamParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param team
	if err := r.SetPathParam("team", o.Team); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}