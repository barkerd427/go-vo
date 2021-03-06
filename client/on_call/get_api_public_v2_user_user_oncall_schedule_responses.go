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

// GetAPIPublicV2UserUserOncallScheduleReader is a Reader for the GetAPIPublicV2UserUserOncallSchedule structure.
type GetAPIPublicV2UserUserOncallScheduleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIPublicV2UserUserOncallScheduleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAPIPublicV2UserUserOncallScheduleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetAPIPublicV2UserUserOncallScheduleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetAPIPublicV2UserUserOncallScheduleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetAPIPublicV2UserUserOncallScheduleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetAPIPublicV2UserUserOncallScheduleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetAPIPublicV2UserUserOncallScheduleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAPIPublicV2UserUserOncallScheduleOK creates a GetAPIPublicV2UserUserOncallScheduleOK with default headers values
func NewGetAPIPublicV2UserUserOncallScheduleOK() *GetAPIPublicV2UserUserOncallScheduleOK {
	return &GetAPIPublicV2UserUserOncallScheduleOK{}
}

/*GetAPIPublicV2UserUserOncallScheduleOK handles this case with default header values.

The on-call schedule for the user
*/
type GetAPIPublicV2UserUserOncallScheduleOK struct {
	Payload *models.UserSchedule
}

func (o *GetAPIPublicV2UserUserOncallScheduleOK) Error() string {
	return fmt.Sprintf("[GET /api-public/v2/user/{user}/oncall/schedule][%d] getApiPublicV2UserUserOncallScheduleOK  %+v", 200, o.Payload)
}

func (o *GetAPIPublicV2UserUserOncallScheduleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserSchedule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIPublicV2UserUserOncallScheduleBadRequest creates a GetAPIPublicV2UserUserOncallScheduleBadRequest with default headers values
func NewGetAPIPublicV2UserUserOncallScheduleBadRequest() *GetAPIPublicV2UserUserOncallScheduleBadRequest {
	return &GetAPIPublicV2UserUserOncallScheduleBadRequest{}
}

/*GetAPIPublicV2UserUserOncallScheduleBadRequest handles this case with default header values.

Problem with the request arguments.  The response payload may include an error message.
*/
type GetAPIPublicV2UserUserOncallScheduleBadRequest struct {
}

func (o *GetAPIPublicV2UserUserOncallScheduleBadRequest) Error() string {
	return fmt.Sprintf("[GET /api-public/v2/user/{user}/oncall/schedule][%d] getApiPublicV2UserUserOncallScheduleBadRequest ", 400)
}

func (o *GetAPIPublicV2UserUserOncallScheduleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAPIPublicV2UserUserOncallScheduleUnauthorized creates a GetAPIPublicV2UserUserOncallScheduleUnauthorized with default headers values
func NewGetAPIPublicV2UserUserOncallScheduleUnauthorized() *GetAPIPublicV2UserUserOncallScheduleUnauthorized {
	return &GetAPIPublicV2UserUserOncallScheduleUnauthorized{}
}

/*GetAPIPublicV2UserUserOncallScheduleUnauthorized handles this case with default header values.

Authentication parameters missing
*/
type GetAPIPublicV2UserUserOncallScheduleUnauthorized struct {
}

func (o *GetAPIPublicV2UserUserOncallScheduleUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api-public/v2/user/{user}/oncall/schedule][%d] getApiPublicV2UserUserOncallScheduleUnauthorized ", 401)
}

func (o *GetAPIPublicV2UserUserOncallScheduleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAPIPublicV2UserUserOncallScheduleForbidden creates a GetAPIPublicV2UserUserOncallScheduleForbidden with default headers values
func NewGetAPIPublicV2UserUserOncallScheduleForbidden() *GetAPIPublicV2UserUserOncallScheduleForbidden {
	return &GetAPIPublicV2UserUserOncallScheduleForbidden{}
}

/*GetAPIPublicV2UserUserOncallScheduleForbidden handles this case with default header values.

Authentication failed or rate-limit reached
*/
type GetAPIPublicV2UserUserOncallScheduleForbidden struct {
}

func (o *GetAPIPublicV2UserUserOncallScheduleForbidden) Error() string {
	return fmt.Sprintf("[GET /api-public/v2/user/{user}/oncall/schedule][%d] getApiPublicV2UserUserOncallScheduleForbidden ", 403)
}

func (o *GetAPIPublicV2UserUserOncallScheduleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAPIPublicV2UserUserOncallScheduleNotFound creates a GetAPIPublicV2UserUserOncallScheduleNotFound with default headers values
func NewGetAPIPublicV2UserUserOncallScheduleNotFound() *GetAPIPublicV2UserUserOncallScheduleNotFound {
	return &GetAPIPublicV2UserUserOncallScheduleNotFound{}
}

/*GetAPIPublicV2UserUserOncallScheduleNotFound handles this case with default header values.

User not found
*/
type GetAPIPublicV2UserUserOncallScheduleNotFound struct {
}

func (o *GetAPIPublicV2UserUserOncallScheduleNotFound) Error() string {
	return fmt.Sprintf("[GET /api-public/v2/user/{user}/oncall/schedule][%d] getApiPublicV2UserUserOncallScheduleNotFound ", 404)
}

func (o *GetAPIPublicV2UserUserOncallScheduleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAPIPublicV2UserUserOncallScheduleInternalServerError creates a GetAPIPublicV2UserUserOncallScheduleInternalServerError with default headers values
func NewGetAPIPublicV2UserUserOncallScheduleInternalServerError() *GetAPIPublicV2UserUserOncallScheduleInternalServerError {
	return &GetAPIPublicV2UserUserOncallScheduleInternalServerError{}
}

/*GetAPIPublicV2UserUserOncallScheduleInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetAPIPublicV2UserUserOncallScheduleInternalServerError struct {
}

func (o *GetAPIPublicV2UserUserOncallScheduleInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api-public/v2/user/{user}/oncall/schedule][%d] getApiPublicV2UserUserOncallScheduleInternalServerError ", 500)
}

func (o *GetAPIPublicV2UserUserOncallScheduleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
