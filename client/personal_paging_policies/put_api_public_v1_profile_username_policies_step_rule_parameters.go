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

// NewPutAPIPublicV1ProfileUsernamePoliciesStepRuleParams creates a new PutAPIPublicV1ProfileUsernamePoliciesStepRuleParams object
// with the default values initialized.
func NewPutAPIPublicV1ProfileUsernamePoliciesStepRuleParams() *PutAPIPublicV1ProfileUsernamePoliciesStepRuleParams {
	var ()
	return &PutAPIPublicV1ProfileUsernamePoliciesStepRuleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutAPIPublicV1ProfileUsernamePoliciesStepRuleParamsWithTimeout creates a new PutAPIPublicV1ProfileUsernamePoliciesStepRuleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutAPIPublicV1ProfileUsernamePoliciesStepRuleParamsWithTimeout(timeout time.Duration) *PutAPIPublicV1ProfileUsernamePoliciesStepRuleParams {
	var ()
	return &PutAPIPublicV1ProfileUsernamePoliciesStepRuleParams{

		timeout: timeout,
	}
}

// NewPutAPIPublicV1ProfileUsernamePoliciesStepRuleParamsWithContext creates a new PutAPIPublicV1ProfileUsernamePoliciesStepRuleParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutAPIPublicV1ProfileUsernamePoliciesStepRuleParamsWithContext(ctx context.Context) *PutAPIPublicV1ProfileUsernamePoliciesStepRuleParams {
	var ()
	return &PutAPIPublicV1ProfileUsernamePoliciesStepRuleParams{

		Context: ctx,
	}
}

// NewPutAPIPublicV1ProfileUsernamePoliciesStepRuleParamsWithHTTPClient creates a new PutAPIPublicV1ProfileUsernamePoliciesStepRuleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutAPIPublicV1ProfileUsernamePoliciesStepRuleParamsWithHTTPClient(client *http.Client) *PutAPIPublicV1ProfileUsernamePoliciesStepRuleParams {
	var ()
	return &PutAPIPublicV1ProfileUsernamePoliciesStepRuleParams{
		HTTPClient: client,
	}
}

/*PutAPIPublicV1ProfileUsernamePoliciesStepRuleParams contains all the parameters to send to the API endpoint
for the put API public v1 profile username policies step rule operation typically these are written to a http.Request
*/
type PutAPIPublicV1ProfileUsernamePoliciesStepRuleParams struct {

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
	/*Rule
	  Rule for a paging policy step

	*/
	Rule float64
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

// WithTimeout adds the timeout to the put API public v1 profile username policies step rule params
func (o *PutAPIPublicV1ProfileUsernamePoliciesStepRuleParams) WithTimeout(timeout time.Duration) *PutAPIPublicV1ProfileUsernamePoliciesStepRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put API public v1 profile username policies step rule params
func (o *PutAPIPublicV1ProfileUsernamePoliciesStepRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put API public v1 profile username policies step rule params
func (o *PutAPIPublicV1ProfileUsernamePoliciesStepRuleParams) WithContext(ctx context.Context) *PutAPIPublicV1ProfileUsernamePoliciesStepRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put API public v1 profile username policies step rule params
func (o *PutAPIPublicV1ProfileUsernamePoliciesStepRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put API public v1 profile username policies step rule params
func (o *PutAPIPublicV1ProfileUsernamePoliciesStepRuleParams) WithHTTPClient(client *http.Client) *PutAPIPublicV1ProfileUsernamePoliciesStepRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put API public v1 profile username policies step rule params
func (o *PutAPIPublicV1ProfileUsernamePoliciesStepRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXVOAPIID adds the xVOAPIID to the put API public v1 profile username policies step rule params
func (o *PutAPIPublicV1ProfileUsernamePoliciesStepRuleParams) WithXVOAPIID(xVOAPIID string) *PutAPIPublicV1ProfileUsernamePoliciesStepRuleParams {
	o.SetXVOAPIID(xVOAPIID)
	return o
}

// SetXVOAPIID adds the xVOApiId to the put API public v1 profile username policies step rule params
func (o *PutAPIPublicV1ProfileUsernamePoliciesStepRuleParams) SetXVOAPIID(xVOAPIID string) {
	o.XVOAPIID = xVOAPIID
}

// WithXVOAPIKey adds the xVOAPIKey to the put API public v1 profile username policies step rule params
func (o *PutAPIPublicV1ProfileUsernamePoliciesStepRuleParams) WithXVOAPIKey(xVOAPIKey string) *PutAPIPublicV1ProfileUsernamePoliciesStepRuleParams {
	o.SetXVOAPIKey(xVOAPIKey)
	return o
}

// SetXVOAPIKey adds the xVOApiKey to the put API public v1 profile username policies step rule params
func (o *PutAPIPublicV1ProfileUsernamePoliciesStepRuleParams) SetXVOAPIKey(xVOAPIKey string) {
	o.XVOAPIKey = xVOAPIKey
}

// WithBody adds the body to the put API public v1 profile username policies step rule params
func (o *PutAPIPublicV1ProfileUsernamePoliciesStepRuleParams) WithBody(body *models.AddStepPayload) *PutAPIPublicV1ProfileUsernamePoliciesStepRuleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put API public v1 profile username policies step rule params
func (o *PutAPIPublicV1ProfileUsernamePoliciesStepRuleParams) SetBody(body *models.AddStepPayload) {
	o.Body = body
}

// WithRule adds the rule to the put API public v1 profile username policies step rule params
func (o *PutAPIPublicV1ProfileUsernamePoliciesStepRuleParams) WithRule(rule float64) *PutAPIPublicV1ProfileUsernamePoliciesStepRuleParams {
	o.SetRule(rule)
	return o
}

// SetRule adds the rule to the put API public v1 profile username policies step rule params
func (o *PutAPIPublicV1ProfileUsernamePoliciesStepRuleParams) SetRule(rule float64) {
	o.Rule = rule
}

// WithStep adds the step to the put API public v1 profile username policies step rule params
func (o *PutAPIPublicV1ProfileUsernamePoliciesStepRuleParams) WithStep(step float64) *PutAPIPublicV1ProfileUsernamePoliciesStepRuleParams {
	o.SetStep(step)
	return o
}

// SetStep adds the step to the put API public v1 profile username policies step rule params
func (o *PutAPIPublicV1ProfileUsernamePoliciesStepRuleParams) SetStep(step float64) {
	o.Step = step
}

// WithUsername adds the username to the put API public v1 profile username policies step rule params
func (o *PutAPIPublicV1ProfileUsernamePoliciesStepRuleParams) WithUsername(username string) *PutAPIPublicV1ProfileUsernamePoliciesStepRuleParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the put API public v1 profile username policies step rule params
func (o *PutAPIPublicV1ProfileUsernamePoliciesStepRuleParams) SetUsername(username string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *PutAPIPublicV1ProfileUsernamePoliciesStepRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param rule
	if err := r.SetPathParam("rule", swag.FormatFloat64(o.Rule)); err != nil {
		return err
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
