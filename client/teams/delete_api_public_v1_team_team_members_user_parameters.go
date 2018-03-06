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

// NewDeleteAPIPublicV1TeamTeamMembersUserParams creates a new DeleteAPIPublicV1TeamTeamMembersUserParams object
// with the default values initialized.
func NewDeleteAPIPublicV1TeamTeamMembersUserParams() *DeleteAPIPublicV1TeamTeamMembersUserParams {
	var ()
	return &DeleteAPIPublicV1TeamTeamMembersUserParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAPIPublicV1TeamTeamMembersUserParamsWithTimeout creates a new DeleteAPIPublicV1TeamTeamMembersUserParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteAPIPublicV1TeamTeamMembersUserParamsWithTimeout(timeout time.Duration) *DeleteAPIPublicV1TeamTeamMembersUserParams {
	var ()
	return &DeleteAPIPublicV1TeamTeamMembersUserParams{

		timeout: timeout,
	}
}

// NewDeleteAPIPublicV1TeamTeamMembersUserParamsWithContext creates a new DeleteAPIPublicV1TeamTeamMembersUserParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteAPIPublicV1TeamTeamMembersUserParamsWithContext(ctx context.Context) *DeleteAPIPublicV1TeamTeamMembersUserParams {
	var ()
	return &DeleteAPIPublicV1TeamTeamMembersUserParams{

		Context: ctx,
	}
}

// NewDeleteAPIPublicV1TeamTeamMembersUserParamsWithHTTPClient creates a new DeleteAPIPublicV1TeamTeamMembersUserParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteAPIPublicV1TeamTeamMembersUserParamsWithHTTPClient(client *http.Client) *DeleteAPIPublicV1TeamTeamMembersUserParams {
	var ()
	return &DeleteAPIPublicV1TeamTeamMembersUserParams{
		HTTPClient: client,
	}
}

/*DeleteAPIPublicV1TeamTeamMembersUserParams contains all the parameters to send to the API endpoint
for the delete API public v1 team team members user operation typically these are written to a http.Request
*/
type DeleteAPIPublicV1TeamTeamMembersUserParams struct {

	/*XVOAPIID
	  Your API ID

	*/
	XVOAPIID string
	/*XVOAPIKey
	  Your API Key

	*/
	XVOAPIKey string
	/*Body
	  The user information

	*/
	Body *models.RemoveTeamMemberPayload
	/*Team
	  The VictorOps team to be deleted

	*/
	Team string
	/*User
	  The team member username

	*/
	User string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete API public v1 team team members user params
func (o *DeleteAPIPublicV1TeamTeamMembersUserParams) WithTimeout(timeout time.Duration) *DeleteAPIPublicV1TeamTeamMembersUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete API public v1 team team members user params
func (o *DeleteAPIPublicV1TeamTeamMembersUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete API public v1 team team members user params
func (o *DeleteAPIPublicV1TeamTeamMembersUserParams) WithContext(ctx context.Context) *DeleteAPIPublicV1TeamTeamMembersUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete API public v1 team team members user params
func (o *DeleteAPIPublicV1TeamTeamMembersUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete API public v1 team team members user params
func (o *DeleteAPIPublicV1TeamTeamMembersUserParams) WithHTTPClient(client *http.Client) *DeleteAPIPublicV1TeamTeamMembersUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete API public v1 team team members user params
func (o *DeleteAPIPublicV1TeamTeamMembersUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXVOAPIID adds the xVOAPIID to the delete API public v1 team team members user params
func (o *DeleteAPIPublicV1TeamTeamMembersUserParams) WithXVOAPIID(xVOAPIID string) *DeleteAPIPublicV1TeamTeamMembersUserParams {
	o.SetXVOAPIID(xVOAPIID)
	return o
}

// SetXVOAPIID adds the xVOApiId to the delete API public v1 team team members user params
func (o *DeleteAPIPublicV1TeamTeamMembersUserParams) SetXVOAPIID(xVOAPIID string) {
	o.XVOAPIID = xVOAPIID
}

// WithXVOAPIKey adds the xVOAPIKey to the delete API public v1 team team members user params
func (o *DeleteAPIPublicV1TeamTeamMembersUserParams) WithXVOAPIKey(xVOAPIKey string) *DeleteAPIPublicV1TeamTeamMembersUserParams {
	o.SetXVOAPIKey(xVOAPIKey)
	return o
}

// SetXVOAPIKey adds the xVOApiKey to the delete API public v1 team team members user params
func (o *DeleteAPIPublicV1TeamTeamMembersUserParams) SetXVOAPIKey(xVOAPIKey string) {
	o.XVOAPIKey = xVOAPIKey
}

// WithBody adds the body to the delete API public v1 team team members user params
func (o *DeleteAPIPublicV1TeamTeamMembersUserParams) WithBody(body *models.RemoveTeamMemberPayload) *DeleteAPIPublicV1TeamTeamMembersUserParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the delete API public v1 team team members user params
func (o *DeleteAPIPublicV1TeamTeamMembersUserParams) SetBody(body *models.RemoveTeamMemberPayload) {
	o.Body = body
}

// WithTeam adds the team to the delete API public v1 team team members user params
func (o *DeleteAPIPublicV1TeamTeamMembersUserParams) WithTeam(team string) *DeleteAPIPublicV1TeamTeamMembersUserParams {
	o.SetTeam(team)
	return o
}

// SetTeam adds the team to the delete API public v1 team team members user params
func (o *DeleteAPIPublicV1TeamTeamMembersUserParams) SetTeam(team string) {
	o.Team = team
}

// WithUser adds the user to the delete API public v1 team team members user params
func (o *DeleteAPIPublicV1TeamTeamMembersUserParams) WithUser(user string) *DeleteAPIPublicV1TeamTeamMembersUserParams {
	o.SetUser(user)
	return o
}

// SetUser adds the user to the delete API public v1 team team members user params
func (o *DeleteAPIPublicV1TeamTeamMembersUserParams) SetUser(user string) {
	o.User = user
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAPIPublicV1TeamTeamMembersUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param user
	if err := r.SetPathParam("user", o.User); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}