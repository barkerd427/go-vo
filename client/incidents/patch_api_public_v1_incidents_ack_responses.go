// Code generated by go-swagger; DO NOT EDIT.

package incidents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/barkerd427/go-vo/models"
)

// PatchAPIPublicV1IncidentsAckReader is a Reader for the PatchAPIPublicV1IncidentsAck structure.
type PatchAPIPublicV1IncidentsAckReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchAPIPublicV1IncidentsAckReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPatchAPIPublicV1IncidentsAckOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPatchAPIPublicV1IncidentsAckBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPatchAPIPublicV1IncidentsAckUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPatchAPIPublicV1IncidentsAckForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPatchAPIPublicV1IncidentsAckNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPatchAPIPublicV1IncidentsAckInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchAPIPublicV1IncidentsAckOK creates a PatchAPIPublicV1IncidentsAckOK with default headers values
func NewPatchAPIPublicV1IncidentsAckOK() *PatchAPIPublicV1IncidentsAckOK {
	return &PatchAPIPublicV1IncidentsAckOK{}
}

/*PatchAPIPublicV1IncidentsAckOK handles this case with default header values.

The result of each acknowledge incident action.
*/
type PatchAPIPublicV1IncidentsAckOK struct {
	Payload *models.AckOrResolveResponse
}

func (o *PatchAPIPublicV1IncidentsAckOK) Error() string {
	return fmt.Sprintf("[PATCH /api-public/v1/incidents/ack][%d] patchApiPublicV1IncidentsAckOK  %+v", 200, o.Payload)
}

func (o *PatchAPIPublicV1IncidentsAckOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AckOrResolveResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAPIPublicV1IncidentsAckBadRequest creates a PatchAPIPublicV1IncidentsAckBadRequest with default headers values
func NewPatchAPIPublicV1IncidentsAckBadRequest() *PatchAPIPublicV1IncidentsAckBadRequest {
	return &PatchAPIPublicV1IncidentsAckBadRequest{}
}

/*PatchAPIPublicV1IncidentsAckBadRequest handles this case with default header values.

Problem with the request arguments.  The response payload may include an error message.
*/
type PatchAPIPublicV1IncidentsAckBadRequest struct {
}

func (o *PatchAPIPublicV1IncidentsAckBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api-public/v1/incidents/ack][%d] patchApiPublicV1IncidentsAckBadRequest ", 400)
}

func (o *PatchAPIPublicV1IncidentsAckBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchAPIPublicV1IncidentsAckUnauthorized creates a PatchAPIPublicV1IncidentsAckUnauthorized with default headers values
func NewPatchAPIPublicV1IncidentsAckUnauthorized() *PatchAPIPublicV1IncidentsAckUnauthorized {
	return &PatchAPIPublicV1IncidentsAckUnauthorized{}
}

/*PatchAPIPublicV1IncidentsAckUnauthorized handles this case with default header values.

Authentication parameters missing
*/
type PatchAPIPublicV1IncidentsAckUnauthorized struct {
}

func (o *PatchAPIPublicV1IncidentsAckUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api-public/v1/incidents/ack][%d] patchApiPublicV1IncidentsAckUnauthorized ", 401)
}

func (o *PatchAPIPublicV1IncidentsAckUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchAPIPublicV1IncidentsAckForbidden creates a PatchAPIPublicV1IncidentsAckForbidden with default headers values
func NewPatchAPIPublicV1IncidentsAckForbidden() *PatchAPIPublicV1IncidentsAckForbidden {
	return &PatchAPIPublicV1IncidentsAckForbidden{}
}

/*PatchAPIPublicV1IncidentsAckForbidden handles this case with default header values.

Authentication failed or rate-limit reached
*/
type PatchAPIPublicV1IncidentsAckForbidden struct {
}

func (o *PatchAPIPublicV1IncidentsAckForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api-public/v1/incidents/ack][%d] patchApiPublicV1IncidentsAckForbidden ", 403)
}

func (o *PatchAPIPublicV1IncidentsAckForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchAPIPublicV1IncidentsAckNotFound creates a PatchAPIPublicV1IncidentsAckNotFound with default headers values
func NewPatchAPIPublicV1IncidentsAckNotFound() *PatchAPIPublicV1IncidentsAckNotFound {
	return &PatchAPIPublicV1IncidentsAckNotFound{}
}

/*PatchAPIPublicV1IncidentsAckNotFound handles this case with default header values.

User not found
*/
type PatchAPIPublicV1IncidentsAckNotFound struct {
}

func (o *PatchAPIPublicV1IncidentsAckNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api-public/v1/incidents/ack][%d] patchApiPublicV1IncidentsAckNotFound ", 404)
}

func (o *PatchAPIPublicV1IncidentsAckNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchAPIPublicV1IncidentsAckInternalServerError creates a PatchAPIPublicV1IncidentsAckInternalServerError with default headers values
func NewPatchAPIPublicV1IncidentsAckInternalServerError() *PatchAPIPublicV1IncidentsAckInternalServerError {
	return &PatchAPIPublicV1IncidentsAckInternalServerError{}
}

/*PatchAPIPublicV1IncidentsAckInternalServerError handles this case with default header values.

Internal Server Error
*/
type PatchAPIPublicV1IncidentsAckInternalServerError struct {
}

func (o *PatchAPIPublicV1IncidentsAckInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api-public/v1/incidents/ack][%d] patchApiPublicV1IncidentsAckInternalServerError ", 500)
}

func (o *PatchAPIPublicV1IncidentsAckInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
