// Code generated by go-swagger; DO NOT EDIT.

package on_call

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/barkerd427/go-vo/models"
)

// GetAPIPublicV2TeamTeamOncallScheduleReader is a Reader for the GetAPIPublicV2TeamTeamOncallSchedule structure.
type GetAPIPublicV2TeamTeamOncallScheduleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIPublicV2TeamTeamOncallScheduleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAPIPublicV2TeamTeamOncallScheduleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetAPIPublicV2TeamTeamOncallScheduleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetAPIPublicV2TeamTeamOncallScheduleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetAPIPublicV2TeamTeamOncallScheduleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetAPIPublicV2TeamTeamOncallScheduleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetAPIPublicV2TeamTeamOncallScheduleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAPIPublicV2TeamTeamOncallScheduleOK creates a GetAPIPublicV2TeamTeamOncallScheduleOK with default headers values
func NewGetAPIPublicV2TeamTeamOncallScheduleOK() *GetAPIPublicV2TeamTeamOncallScheduleOK {
	return &GetAPIPublicV2TeamTeamOncallScheduleOK{}
}

/*GetAPIPublicV2TeamTeamOncallScheduleOK handles this case with default header values.

The on-call schedule for the team
*/
type GetAPIPublicV2TeamTeamOncallScheduleOK struct {
	Payload *models.TeamSchedule
}

func (o *GetAPIPublicV2TeamTeamOncallScheduleOK) Error() string {
	return fmt.Sprintf("[GET /api-public/v2/team/{team}/oncall/schedule][%d] getApiPublicV2TeamTeamOncallScheduleOK  %+v", 200, o.Payload)
}

func (o *GetAPIPublicV2TeamTeamOncallScheduleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TeamSchedule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIPublicV2TeamTeamOncallScheduleBadRequest creates a GetAPIPublicV2TeamTeamOncallScheduleBadRequest with default headers values
func NewGetAPIPublicV2TeamTeamOncallScheduleBadRequest() *GetAPIPublicV2TeamTeamOncallScheduleBadRequest {
	return &GetAPIPublicV2TeamTeamOncallScheduleBadRequest{}
}

/*GetAPIPublicV2TeamTeamOncallScheduleBadRequest handles this case with default header values.

Problem with the request arguments.  The response payload may include an error message.
*/
type GetAPIPublicV2TeamTeamOncallScheduleBadRequest struct {
}

func (o *GetAPIPublicV2TeamTeamOncallScheduleBadRequest) Error() string {
	return fmt.Sprintf("[GET /api-public/v2/team/{team}/oncall/schedule][%d] getApiPublicV2TeamTeamOncallScheduleBadRequest ", 400)
}

func (o *GetAPIPublicV2TeamTeamOncallScheduleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAPIPublicV2TeamTeamOncallScheduleUnauthorized creates a GetAPIPublicV2TeamTeamOncallScheduleUnauthorized with default headers values
func NewGetAPIPublicV2TeamTeamOncallScheduleUnauthorized() *GetAPIPublicV2TeamTeamOncallScheduleUnauthorized {
	return &GetAPIPublicV2TeamTeamOncallScheduleUnauthorized{}
}

/*GetAPIPublicV2TeamTeamOncallScheduleUnauthorized handles this case with default header values.

Authentication parameters missing
*/
type GetAPIPublicV2TeamTeamOncallScheduleUnauthorized struct {
}

func (o *GetAPIPublicV2TeamTeamOncallScheduleUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api-public/v2/team/{team}/oncall/schedule][%d] getApiPublicV2TeamTeamOncallScheduleUnauthorized ", 401)
}

func (o *GetAPIPublicV2TeamTeamOncallScheduleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAPIPublicV2TeamTeamOncallScheduleForbidden creates a GetAPIPublicV2TeamTeamOncallScheduleForbidden with default headers values
func NewGetAPIPublicV2TeamTeamOncallScheduleForbidden() *GetAPIPublicV2TeamTeamOncallScheduleForbidden {
	return &GetAPIPublicV2TeamTeamOncallScheduleForbidden{}
}

/*GetAPIPublicV2TeamTeamOncallScheduleForbidden handles this case with default header values.

Authentication failed or rate-limit reached
*/
type GetAPIPublicV2TeamTeamOncallScheduleForbidden struct {
}

func (o *GetAPIPublicV2TeamTeamOncallScheduleForbidden) Error() string {
	return fmt.Sprintf("[GET /api-public/v2/team/{team}/oncall/schedule][%d] getApiPublicV2TeamTeamOncallScheduleForbidden ", 403)
}

func (o *GetAPIPublicV2TeamTeamOncallScheduleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAPIPublicV2TeamTeamOncallScheduleNotFound creates a GetAPIPublicV2TeamTeamOncallScheduleNotFound with default headers values
func NewGetAPIPublicV2TeamTeamOncallScheduleNotFound() *GetAPIPublicV2TeamTeamOncallScheduleNotFound {
	return &GetAPIPublicV2TeamTeamOncallScheduleNotFound{}
}

/*GetAPIPublicV2TeamTeamOncallScheduleNotFound handles this case with default header values.

Team not found
*/
type GetAPIPublicV2TeamTeamOncallScheduleNotFound struct {
}

func (o *GetAPIPublicV2TeamTeamOncallScheduleNotFound) Error() string {
	return fmt.Sprintf("[GET /api-public/v2/team/{team}/oncall/schedule][%d] getApiPublicV2TeamTeamOncallScheduleNotFound ", 404)
}

func (o *GetAPIPublicV2TeamTeamOncallScheduleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAPIPublicV2TeamTeamOncallScheduleInternalServerError creates a GetAPIPublicV2TeamTeamOncallScheduleInternalServerError with default headers values
func NewGetAPIPublicV2TeamTeamOncallScheduleInternalServerError() *GetAPIPublicV2TeamTeamOncallScheduleInternalServerError {
	return &GetAPIPublicV2TeamTeamOncallScheduleInternalServerError{}
}

/*GetAPIPublicV2TeamTeamOncallScheduleInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetAPIPublicV2TeamTeamOncallScheduleInternalServerError struct {
}

func (o *GetAPIPublicV2TeamTeamOncallScheduleInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api-public/v2/team/{team}/oncall/schedule][%d] getApiPublicV2TeamTeamOncallScheduleInternalServerError ", 500)
}

func (o *GetAPIPublicV2TeamTeamOncallScheduleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
