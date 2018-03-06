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

	"github.com/barkerd427/go-vo/models"
)

// NewPostAPIPublicV1TeamTeamMembersParams creates a new PostAPIPublicV1TeamTeamMembersParams object
// with the default values initialized.
func NewPostAPIPublicV1TeamTeamMembersParams() *PostAPIPublicV1TeamTeamMembersParams {
	var ()
	return &PostAPIPublicV1TeamTeamMembersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostAPIPublicV1TeamTeamMembersParamsWithTimeout creates a new PostAPIPublicV1TeamTeamMembersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostAPIPublicV1TeamTeamMembersParamsWithTimeout(timeout time.Duration) *PostAPIPublicV1TeamTeamMembersParams {
	var ()
	return &PostAPIPublicV1TeamTeamMembersParams{

		timeout: timeout,
	}
}

// NewPostAPIPublicV1TeamTeamMembersParamsWithContext creates a new PostAPIPublicV1TeamTeamMembersParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostAPIPublicV1TeamTeamMembersParamsWithContext(ctx context.Context) *PostAPIPublicV1TeamTeamMembersParams {
	var ()
	return &PostAPIPublicV1TeamTeamMembersParams{

		Context: ctx,
	}
}

// NewPostAPIPublicV1TeamTeamMembersParamsWithHTTPClient creates a new PostAPIPublicV1TeamTeamMembersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostAPIPublicV1TeamTeamMembersParamsWithHTTPClient(client *http.Client) *PostAPIPublicV1TeamTeamMembersParams {
	var ()
	return &PostAPIPublicV1TeamTeamMembersParams{
		HTTPClient: client,
	}
}

/*PostAPIPublicV1TeamTeamMembersParams contains all the parameters to send to the API endpoint
for the post API public v1 team team members operation typically these are written to a http.Request
*/
type PostAPIPublicV1TeamTeamMembersParams struct {

	/*XVOAPIID
	  Your API ID

	*/
	XVOAPIID string
	/*XVOAPIKey
	  Your API Key

	*/
	XVOAPIKey string
	/*Body*/
	Body *models.AddTeamMemberPayload
	/*Team
	  The VictorOps team to fetch

	*/
	Team string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post API public v1 team team members params
func (o *PostAPIPublicV1TeamTeamMembersParams) WithTimeout(timeout time.Duration) *PostAPIPublicV1TeamTeamMembersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post API public v1 team team members params
func (o *PostAPIPublicV1TeamTeamMembersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post API public v1 team team members params
func (o *PostAPIPublicV1TeamTeamMembersParams) WithContext(ctx context.Context) *PostAPIPublicV1TeamTeamMembersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post API public v1 team team members params
func (o *PostAPIPublicV1TeamTeamMembersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post API public v1 team team members params
func (o *PostAPIPublicV1TeamTeamMembersParams) WithHTTPClient(client *http.Client) *PostAPIPublicV1TeamTeamMembersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post API public v1 team team members params
func (o *PostAPIPublicV1TeamTeamMembersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXVOAPIID adds the xVOAPIID to the post API public v1 team team members params
func (o *PostAPIPublicV1TeamTeamMembersParams) WithXVOAPIID(xVOAPIID string) *PostAPIPublicV1TeamTeamMembersParams {
	o.SetXVOAPIID(xVOAPIID)
	return o
}

// SetXVOAPIID adds the xVOApiId to the post API public v1 team team members params
func (o *PostAPIPublicV1TeamTeamMembersParams) SetXVOAPIID(xVOAPIID string) {
	o.XVOAPIID = xVOAPIID
}

// WithXVOAPIKey adds the xVOAPIKey to the post API public v1 team team members params
func (o *PostAPIPublicV1TeamTeamMembersParams) WithXVOAPIKey(xVOAPIKey string) *PostAPIPublicV1TeamTeamMembersParams {
	o.SetXVOAPIKey(xVOAPIKey)
	return o
}

// SetXVOAPIKey adds the xVOApiKey to the post API public v1 team team members params
func (o *PostAPIPublicV1TeamTeamMembersParams) SetXVOAPIKey(xVOAPIKey string) {
	o.XVOAPIKey = xVOAPIKey
}

// WithBody adds the body to the post API public v1 team team members params
func (o *PostAPIPublicV1TeamTeamMembersParams) WithBody(body *models.AddTeamMemberPayload) *PostAPIPublicV1TeamTeamMembersParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post API public v1 team team members params
func (o *PostAPIPublicV1TeamTeamMembersParams) SetBody(body *models.AddTeamMemberPayload) {
	o.Body = body
}

// WithTeam adds the team to the post API public v1 team team members params
func (o *PostAPIPublicV1TeamTeamMembersParams) WithTeam(team string) *PostAPIPublicV1TeamTeamMembersParams {
	o.SetTeam(team)
	return o
}

// SetTeam adds the team to the post API public v1 team team members params
func (o *PostAPIPublicV1TeamTeamMembersParams) SetTeam(team string) {
	o.Team = team
}

// WriteToRequest writes these params to a swagger request
func (o *PostAPIPublicV1TeamTeamMembersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param team
	if err := r.SetPathParam("team", o.Team); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
