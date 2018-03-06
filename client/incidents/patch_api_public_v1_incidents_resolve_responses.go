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

// PatchAPIPublicV1IncidentsResolveReader is a Reader for the PatchAPIPublicV1IncidentsResolve structure.
type PatchAPIPublicV1IncidentsResolveReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchAPIPublicV1IncidentsResolveReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPatchAPIPublicV1IncidentsResolveOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPatchAPIPublicV1IncidentsResolveBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPatchAPIPublicV1IncidentsResolveUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPatchAPIPublicV1IncidentsResolveForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPatchAPIPublicV1IncidentsResolveNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPatchAPIPublicV1IncidentsResolveInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchAPIPublicV1IncidentsResolveOK creates a PatchAPIPublicV1IncidentsResolveOK with default headers values
func NewPatchAPIPublicV1IncidentsResolveOK() *PatchAPIPublicV1IncidentsResolveOK {
	return &PatchAPIPublicV1IncidentsResolveOK{}
}

/*PatchAPIPublicV1IncidentsResolveOK handles this case with default header values.

The result of each resolve incident action.
*/
type PatchAPIPublicV1IncidentsResolveOK struct {
	Payload *models.AckOrResolveResponse
}

func (o *PatchAPIPublicV1IncidentsResolveOK) Error() string {
	return fmt.Sprintf("[PATCH /api-public/v1/incidents/resolve][%d] patchApiPublicV1IncidentsResolveOK  %+v", 200, o.Payload)
}

func (o *PatchAPIPublicV1IncidentsResolveOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AckOrResolveResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAPIPublicV1IncidentsResolveBadRequest creates a PatchAPIPublicV1IncidentsResolveBadRequest with default headers values
func NewPatchAPIPublicV1IncidentsResolveBadRequest() *PatchAPIPublicV1IncidentsResolveBadRequest {
	return &PatchAPIPublicV1IncidentsResolveBadRequest{}
}

/*PatchAPIPublicV1IncidentsResolveBadRequest handles this case with default header values.

Problem with the request arguments.  The response payload may include an error message.
*/
type PatchAPIPublicV1IncidentsResolveBadRequest struct {
}

func (o *PatchAPIPublicV1IncidentsResolveBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api-public/v1/incidents/resolve][%d] patchApiPublicV1IncidentsResolveBadRequest ", 400)
}

func (o *PatchAPIPublicV1IncidentsResolveBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchAPIPublicV1IncidentsResolveUnauthorized creates a PatchAPIPublicV1IncidentsResolveUnauthorized with default headers values
func NewPatchAPIPublicV1IncidentsResolveUnauthorized() *PatchAPIPublicV1IncidentsResolveUnauthorized {
	return &PatchAPIPublicV1IncidentsResolveUnauthorized{}
}

/*PatchAPIPublicV1IncidentsResolveUnauthorized handles this case with default header values.

Authentication parameters missing
*/
type PatchAPIPublicV1IncidentsResolveUnauthorized struct {
}

func (o *PatchAPIPublicV1IncidentsResolveUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api-public/v1/incidents/resolve][%d] patchApiPublicV1IncidentsResolveUnauthorized ", 401)
}

func (o *PatchAPIPublicV1IncidentsResolveUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchAPIPublicV1IncidentsResolveForbidden creates a PatchAPIPublicV1IncidentsResolveForbidden with default headers values
func NewPatchAPIPublicV1IncidentsResolveForbidden() *PatchAPIPublicV1IncidentsResolveForbidden {
	return &PatchAPIPublicV1IncidentsResolveForbidden{}
}

/*PatchAPIPublicV1IncidentsResolveForbidden handles this case with default header values.

Authentication failed or rate-limit reached
*/
type PatchAPIPublicV1IncidentsResolveForbidden struct {
}

func (o *PatchAPIPublicV1IncidentsResolveForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api-public/v1/incidents/resolve][%d] patchApiPublicV1IncidentsResolveForbidden ", 403)
}

func (o *PatchAPIPublicV1IncidentsResolveForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchAPIPublicV1IncidentsResolveNotFound creates a PatchAPIPublicV1IncidentsResolveNotFound with default headers values
func NewPatchAPIPublicV1IncidentsResolveNotFound() *PatchAPIPublicV1IncidentsResolveNotFound {
	return &PatchAPIPublicV1IncidentsResolveNotFound{}
}

/*PatchAPIPublicV1IncidentsResolveNotFound handles this case with default header values.

User not found
*/
type PatchAPIPublicV1IncidentsResolveNotFound struct {
}

func (o *PatchAPIPublicV1IncidentsResolveNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api-public/v1/incidents/resolve][%d] patchApiPublicV1IncidentsResolveNotFound ", 404)
}

func (o *PatchAPIPublicV1IncidentsResolveNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchAPIPublicV1IncidentsResolveInternalServerError creates a PatchAPIPublicV1IncidentsResolveInternalServerError with default headers values
func NewPatchAPIPublicV1IncidentsResolveInternalServerError() *PatchAPIPublicV1IncidentsResolveInternalServerError {
	return &PatchAPIPublicV1IncidentsResolveInternalServerError{}
}

/*PatchAPIPublicV1IncidentsResolveInternalServerError handles this case with default header values.

Internal Server Error
*/
type PatchAPIPublicV1IncidentsResolveInternalServerError struct {
}

func (o *PatchAPIPublicV1IncidentsResolveInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api-public/v1/incidents/resolve][%d] patchApiPublicV1IncidentsResolveInternalServerError ", 500)
}

func (o *PatchAPIPublicV1IncidentsResolveInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}