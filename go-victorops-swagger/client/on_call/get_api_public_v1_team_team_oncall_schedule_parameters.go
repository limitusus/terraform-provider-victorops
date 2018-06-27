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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetAPIPublicV1TeamTeamOncallScheduleParams creates a new GetAPIPublicV1TeamTeamOncallScheduleParams object
// with the default values initialized.
func NewGetAPIPublicV1TeamTeamOncallScheduleParams() *GetAPIPublicV1TeamTeamOncallScheduleParams {
	var (
		daysForwardDefault = float64(14)
		daysSkipDefault    = float64(0)
		stepDefault        = float64(0)
	)
	return &GetAPIPublicV1TeamTeamOncallScheduleParams{
		DaysForward: &daysForwardDefault,
		DaysSkip:    &daysSkipDefault,
		Step:        &stepDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIPublicV1TeamTeamOncallScheduleParamsWithTimeout creates a new GetAPIPublicV1TeamTeamOncallScheduleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAPIPublicV1TeamTeamOncallScheduleParamsWithTimeout(timeout time.Duration) *GetAPIPublicV1TeamTeamOncallScheduleParams {
	var (
		daysForwardDefault = float64(14)
		daysSkipDefault    = float64(0)
		stepDefault        = float64(0)
	)
	return &GetAPIPublicV1TeamTeamOncallScheduleParams{
		DaysForward: &daysForwardDefault,
		DaysSkip:    &daysSkipDefault,
		Step:        &stepDefault,

		timeout: timeout,
	}
}

// NewGetAPIPublicV1TeamTeamOncallScheduleParamsWithContext creates a new GetAPIPublicV1TeamTeamOncallScheduleParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAPIPublicV1TeamTeamOncallScheduleParamsWithContext(ctx context.Context) *GetAPIPublicV1TeamTeamOncallScheduleParams {
	var (
		daysForwardDefault = float64(14)
		daysSkipDefault    = float64(0)
		stepDefault        = float64(0)
	)
	return &GetAPIPublicV1TeamTeamOncallScheduleParams{
		DaysForward: &daysForwardDefault,
		DaysSkip:    &daysSkipDefault,
		Step:        &stepDefault,

		Context: ctx,
	}
}

// NewGetAPIPublicV1TeamTeamOncallScheduleParamsWithHTTPClient creates a new GetAPIPublicV1TeamTeamOncallScheduleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAPIPublicV1TeamTeamOncallScheduleParamsWithHTTPClient(client *http.Client) *GetAPIPublicV1TeamTeamOncallScheduleParams {
	var (
		daysForwardDefault = float64(14)
		daysSkipDefault    = float64(0)
		stepDefault        = float64(0)
	)
	return &GetAPIPublicV1TeamTeamOncallScheduleParams{
		DaysForward: &daysForwardDefault,
		DaysSkip:    &daysSkipDefault,
		Step:        &stepDefault,
		HTTPClient:  client,
	}
}

/*GetAPIPublicV1TeamTeamOncallScheduleParams contains all the parameters to send to the API endpoint
for the get API public v1 team team oncall schedule operation typically these are written to a http.Request
*/
type GetAPIPublicV1TeamTeamOncallScheduleParams struct {

	/*XVOAPIID
	  Your API ID

	*/
	XVOAPIID string
	/*XVOAPIKey
	  Your API Key

	*/
	XVOAPIKey string
	/*DaysForward
	  Days to include in returned schedule (30 max)

	*/
	DaysForward *float64
	/*DaysSkip
	  Days to skip before computing schedule to return (90 max)

	*/
	DaysSkip *float64
	/*Step
	  Step of escalation policy (3 max)

	*/
	Step *float64
	/*Team
	  The VictorOps team 'slug'

	*/
	Team string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get API public v1 team team oncall schedule params
func (o *GetAPIPublicV1TeamTeamOncallScheduleParams) WithTimeout(timeout time.Duration) *GetAPIPublicV1TeamTeamOncallScheduleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API public v1 team team oncall schedule params
func (o *GetAPIPublicV1TeamTeamOncallScheduleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API public v1 team team oncall schedule params
func (o *GetAPIPublicV1TeamTeamOncallScheduleParams) WithContext(ctx context.Context) *GetAPIPublicV1TeamTeamOncallScheduleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API public v1 team team oncall schedule params
func (o *GetAPIPublicV1TeamTeamOncallScheduleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API public v1 team team oncall schedule params
func (o *GetAPIPublicV1TeamTeamOncallScheduleParams) WithHTTPClient(client *http.Client) *GetAPIPublicV1TeamTeamOncallScheduleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API public v1 team team oncall schedule params
func (o *GetAPIPublicV1TeamTeamOncallScheduleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXVOAPIID adds the xVOAPIID to the get API public v1 team team oncall schedule params
func (o *GetAPIPublicV1TeamTeamOncallScheduleParams) WithXVOAPIID(xVOAPIID string) *GetAPIPublicV1TeamTeamOncallScheduleParams {
	o.SetXVOAPIID(xVOAPIID)
	return o
}

// SetXVOAPIID adds the xVOApiId to the get API public v1 team team oncall schedule params
func (o *GetAPIPublicV1TeamTeamOncallScheduleParams) SetXVOAPIID(xVOAPIID string) {
	o.XVOAPIID = xVOAPIID
}

// WithXVOAPIKey adds the xVOAPIKey to the get API public v1 team team oncall schedule params
func (o *GetAPIPublicV1TeamTeamOncallScheduleParams) WithXVOAPIKey(xVOAPIKey string) *GetAPIPublicV1TeamTeamOncallScheduleParams {
	o.SetXVOAPIKey(xVOAPIKey)
	return o
}

// SetXVOAPIKey adds the xVOApiKey to the get API public v1 team team oncall schedule params
func (o *GetAPIPublicV1TeamTeamOncallScheduleParams) SetXVOAPIKey(xVOAPIKey string) {
	o.XVOAPIKey = xVOAPIKey
}

// WithDaysForward adds the daysForward to the get API public v1 team team oncall schedule params
func (o *GetAPIPublicV1TeamTeamOncallScheduleParams) WithDaysForward(daysForward *float64) *GetAPIPublicV1TeamTeamOncallScheduleParams {
	o.SetDaysForward(daysForward)
	return o
}

// SetDaysForward adds the daysForward to the get API public v1 team team oncall schedule params
func (o *GetAPIPublicV1TeamTeamOncallScheduleParams) SetDaysForward(daysForward *float64) {
	o.DaysForward = daysForward
}

// WithDaysSkip adds the daysSkip to the get API public v1 team team oncall schedule params
func (o *GetAPIPublicV1TeamTeamOncallScheduleParams) WithDaysSkip(daysSkip *float64) *GetAPIPublicV1TeamTeamOncallScheduleParams {
	o.SetDaysSkip(daysSkip)
	return o
}

// SetDaysSkip adds the daysSkip to the get API public v1 team team oncall schedule params
func (o *GetAPIPublicV1TeamTeamOncallScheduleParams) SetDaysSkip(daysSkip *float64) {
	o.DaysSkip = daysSkip
}

// WithStep adds the step to the get API public v1 team team oncall schedule params
func (o *GetAPIPublicV1TeamTeamOncallScheduleParams) WithStep(step *float64) *GetAPIPublicV1TeamTeamOncallScheduleParams {
	o.SetStep(step)
	return o
}

// SetStep adds the step to the get API public v1 team team oncall schedule params
func (o *GetAPIPublicV1TeamTeamOncallScheduleParams) SetStep(step *float64) {
	o.Step = step
}

// WithTeam adds the team to the get API public v1 team team oncall schedule params
func (o *GetAPIPublicV1TeamTeamOncallScheduleParams) WithTeam(team string) *GetAPIPublicV1TeamTeamOncallScheduleParams {
	o.SetTeam(team)
	return o
}

// SetTeam adds the team to the get API public v1 team team oncall schedule params
func (o *GetAPIPublicV1TeamTeamOncallScheduleParams) SetTeam(team string) {
	o.Team = team
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIPublicV1TeamTeamOncallScheduleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.DaysForward != nil {

		// query param daysForward
		var qrDaysForward float64
		if o.DaysForward != nil {
			qrDaysForward = *o.DaysForward
		}
		qDaysForward := swag.FormatFloat64(qrDaysForward)
		if qDaysForward != "" {
			if err := r.SetQueryParam("daysForward", qDaysForward); err != nil {
				return err
			}
		}

	}

	if o.DaysSkip != nil {

		// query param daysSkip
		var qrDaysSkip float64
		if o.DaysSkip != nil {
			qrDaysSkip = *o.DaysSkip
		}
		qDaysSkip := swag.FormatFloat64(qrDaysSkip)
		if qDaysSkip != "" {
			if err := r.SetQueryParam("daysSkip", qDaysSkip); err != nil {
				return err
			}
		}

	}

	if o.Step != nil {

		// query param step
		var qrStep float64
		if o.Step != nil {
			qrStep = *o.Step
		}
		qStep := swag.FormatFloat64(qrStep)
		if qStep != "" {
			if err := r.SetQueryParam("step", qStep); err != nil {
				return err
			}
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