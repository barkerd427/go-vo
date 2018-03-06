// Code generated by go-swagger; DO NOT EDIT.

package on_call

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

// NewPatchAPIPublicV1TeamTeamOncallUserParams creates a new PatchAPIPublicV1TeamTeamOncallUserParams object
// with the default values initialized.
func NewPatchAPIPublicV1TeamTeamOncallUserParams() *PatchAPIPublicV1TeamTeamOncallUserParams {
	var ()
	return &PatchAPIPublicV1TeamTeamOncallUserParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchAPIPublicV1TeamTeamOncallUserParamsWithTimeout creates a new PatchAPIPublicV1TeamTeamOncallUserParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchAPIPublicV1TeamTeamOncallUserParamsWithTimeout(timeout time.Duration) *PatchAPIPublicV1TeamTeamOncallUserParams {
	var ()
	return &PatchAPIPublicV1TeamTeamOncallUserParams{

		timeout: timeout,
	}
}

// NewPatchAPIPublicV1TeamTeamOncallUserParamsWithContext creates a new PatchAPIPublicV1TeamTeamOncallUserParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchAPIPublicV1TeamTeamOncallUserParamsWithContext(ctx context.Context) *PatchAPIPublicV1TeamTeamOncallUserParams {
	var ()
	return &PatchAPIPublicV1TeamTeamOncallUserParams{

		Context: ctx,
	}
}

// NewPatchAPIPublicV1TeamTeamOncallUserParamsWithHTTPClient creates a new PatchAPIPublicV1TeamTeamOncallUserParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchAPIPublicV1TeamTeamOncallUserParamsWithHTTPClient(client *http.Client) *PatchAPIPublicV1TeamTeamOncallUserParams {
	var ()
	return &PatchAPIPublicV1TeamTeamOncallUserParams{
		HTTPClient: client,
	}
}

/*PatchAPIPublicV1TeamTeamOncallUserParams contains all the parameters to send to the API endpoint
for the patch API public v1 team team oncall user operation typically these are written to a http.Request
*/
type PatchAPIPublicV1TeamTeamOncallUserParams struct {

	/*XVOAPIID
	  Your API ID

	*/
	XVOAPIID string
	/*XVOAPIKey
	  Your API Key

	*/
	XVOAPIKey string
	/*Body
	  The take on-call payload

	*/
	Body *models.TakeRequest
	/*Team
	  The VictorOps team 'slug'

	*/
	Team string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch API public v1 team team oncall user params
func (o *PatchAPIPublicV1TeamTeamOncallUserParams) WithTimeout(timeout time.Duration) *PatchAPIPublicV1TeamTeamOncallUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch API public v1 team team oncall user params
func (o *PatchAPIPublicV1TeamTeamOncallUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch API public v1 team team oncall user params
func (o *PatchAPIPublicV1TeamTeamOncallUserParams) WithContext(ctx context.Context) *PatchAPIPublicV1TeamTeamOncallUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch API public v1 team team oncall user params
func (o *PatchAPIPublicV1TeamTeamOncallUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch API public v1 team team oncall user params
func (o *PatchAPIPublicV1TeamTeamOncallUserParams) WithHTTPClient(client *http.Client) *PatchAPIPublicV1TeamTeamOncallUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch API public v1 team team oncall user params
func (o *PatchAPIPublicV1TeamTeamOncallUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXVOAPIID adds the xVOAPIID to the patch API public v1 team team oncall user params
func (o *PatchAPIPublicV1TeamTeamOncallUserParams) WithXVOAPIID(xVOAPIID string) *PatchAPIPublicV1TeamTeamOncallUserParams {
	o.SetXVOAPIID(xVOAPIID)
	return o
}

// SetXVOAPIID adds the xVOApiId to the patch API public v1 team team oncall user params
func (o *PatchAPIPublicV1TeamTeamOncallUserParams) SetXVOAPIID(xVOAPIID string) {
	o.XVOAPIID = xVOAPIID
}

// WithXVOAPIKey adds the xVOAPIKey to the patch API public v1 team team oncall user params
func (o *PatchAPIPublicV1TeamTeamOncallUserParams) WithXVOAPIKey(xVOAPIKey string) *PatchAPIPublicV1TeamTeamOncallUserParams {
	o.SetXVOAPIKey(xVOAPIKey)
	return o
}

// SetXVOAPIKey adds the xVOApiKey to the patch API public v1 team team oncall user params
func (o *PatchAPIPublicV1TeamTeamOncallUserParams) SetXVOAPIKey(xVOAPIKey string) {
	o.XVOAPIKey = xVOAPIKey
}

// WithBody adds the body to the patch API public v1 team team oncall user params
func (o *PatchAPIPublicV1TeamTeamOncallUserParams) WithBody(body *models.TakeRequest) *PatchAPIPublicV1TeamTeamOncallUserParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch API public v1 team team oncall user params
func (o *PatchAPIPublicV1TeamTeamOncallUserParams) SetBody(body *models.TakeRequest) {
	o.Body = body
}

// WithTeam adds the team to the patch API public v1 team team oncall user params
func (o *PatchAPIPublicV1TeamTeamOncallUserParams) WithTeam(team string) *PatchAPIPublicV1TeamTeamOncallUserParams {
	o.SetTeam(team)
	return o
}

// SetTeam adds the team to the patch API public v1 team team oncall user params
func (o *PatchAPIPublicV1TeamTeamOncallUserParams) SetTeam(team string) {
	o.Team = team
}

// WriteToRequest writes these params to a swagger request
func (o *PatchAPIPublicV1TeamTeamOncallUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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