// Code generated by go-swagger; DO NOT EDIT.

package user_contact_methods

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/barkerd427/go-vo/models"
)

// PutAPIPublicV1UserUserContactMethodsDevicesContactIDReader is a Reader for the PutAPIPublicV1UserUserContactMethodsDevicesContactID structure.
type PutAPIPublicV1UserUserContactMethodsDevicesContactIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutAPIPublicV1UserUserContactMethodsDevicesContactIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutAPIPublicV1UserUserContactMethodsDevicesContactIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPutAPIPublicV1UserUserContactMethodsDevicesContactIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPutAPIPublicV1UserUserContactMethodsDevicesContactIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPutAPIPublicV1UserUserContactMethodsDevicesContactIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPutAPIPublicV1UserUserContactMethodsDevicesContactIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPutAPIPublicV1UserUserContactMethodsDevicesContactIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutAPIPublicV1UserUserContactMethodsDevicesContactIDOK creates a PutAPIPublicV1UserUserContactMethodsDevicesContactIDOK with default headers values
func NewPutAPIPublicV1UserUserContactMethodsDevicesContactIDOK() *PutAPIPublicV1UserUserContactMethodsDevicesContactIDOK {
	return &PutAPIPublicV1UserUserContactMethodsDevicesContactIDOK{}
}

/*PutAPIPublicV1UserUserContactMethodsDevicesContactIDOK handles this case with default header values.

The list of contact devices for the user
*/
type PutAPIPublicV1UserUserContactMethodsDevicesContactIDOK struct {
	Payload *models.ContactDevice
}

func (o *PutAPIPublicV1UserUserContactMethodsDevicesContactIDOK) Error() string {
	return fmt.Sprintf("[PUT /api-public/v1/user/{user}/contact-methods/devices/{contactId}][%d] putApiPublicV1UserUserContactMethodsDevicesContactIdOK  %+v", 200, o.Payload)
}

func (o *PutAPIPublicV1UserUserContactMethodsDevicesContactIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContactDevice)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAPIPublicV1UserUserContactMethodsDevicesContactIDBadRequest creates a PutAPIPublicV1UserUserContactMethodsDevicesContactIDBadRequest with default headers values
func NewPutAPIPublicV1UserUserContactMethodsDevicesContactIDBadRequest() *PutAPIPublicV1UserUserContactMethodsDevicesContactIDBadRequest {
	return &PutAPIPublicV1UserUserContactMethodsDevicesContactIDBadRequest{}
}

/*PutAPIPublicV1UserUserContactMethodsDevicesContactIDBadRequest handles this case with default header values.

Problem with the request arguments.  The response payload may include an error message.
*/
type PutAPIPublicV1UserUserContactMethodsDevicesContactIDBadRequest struct {
}

func (o *PutAPIPublicV1UserUserContactMethodsDevicesContactIDBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api-public/v1/user/{user}/contact-methods/devices/{contactId}][%d] putApiPublicV1UserUserContactMethodsDevicesContactIdBadRequest ", 400)
}

func (o *PutAPIPublicV1UserUserContactMethodsDevicesContactIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutAPIPublicV1UserUserContactMethodsDevicesContactIDUnauthorized creates a PutAPIPublicV1UserUserContactMethodsDevicesContactIDUnauthorized with default headers values
func NewPutAPIPublicV1UserUserContactMethodsDevicesContactIDUnauthorized() *PutAPIPublicV1UserUserContactMethodsDevicesContactIDUnauthorized {
	return &PutAPIPublicV1UserUserContactMethodsDevicesContactIDUnauthorized{}
}

/*PutAPIPublicV1UserUserContactMethodsDevicesContactIDUnauthorized handles this case with default header values.

Authentication parameters missing
*/
type PutAPIPublicV1UserUserContactMethodsDevicesContactIDUnauthorized struct {
}

func (o *PutAPIPublicV1UserUserContactMethodsDevicesContactIDUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api-public/v1/user/{user}/contact-methods/devices/{contactId}][%d] putApiPublicV1UserUserContactMethodsDevicesContactIdUnauthorized ", 401)
}

func (o *PutAPIPublicV1UserUserContactMethodsDevicesContactIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutAPIPublicV1UserUserContactMethodsDevicesContactIDForbidden creates a PutAPIPublicV1UserUserContactMethodsDevicesContactIDForbidden with default headers values
func NewPutAPIPublicV1UserUserContactMethodsDevicesContactIDForbidden() *PutAPIPublicV1UserUserContactMethodsDevicesContactIDForbidden {
	return &PutAPIPublicV1UserUserContactMethodsDevicesContactIDForbidden{}
}

/*PutAPIPublicV1UserUserContactMethodsDevicesContactIDForbidden handles this case with default header values.

Authentication failed or rate-limit reached
*/
type PutAPIPublicV1UserUserContactMethodsDevicesContactIDForbidden struct {
}

func (o *PutAPIPublicV1UserUserContactMethodsDevicesContactIDForbidden) Error() string {
	return fmt.Sprintf("[PUT /api-public/v1/user/{user}/contact-methods/devices/{contactId}][%d] putApiPublicV1UserUserContactMethodsDevicesContactIdForbidden ", 403)
}

func (o *PutAPIPublicV1UserUserContactMethodsDevicesContactIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutAPIPublicV1UserUserContactMethodsDevicesContactIDNotFound creates a PutAPIPublicV1UserUserContactMethodsDevicesContactIDNotFound with default headers values
func NewPutAPIPublicV1UserUserContactMethodsDevicesContactIDNotFound() *PutAPIPublicV1UserUserContactMethodsDevicesContactIDNotFound {
	return &PutAPIPublicV1UserUserContactMethodsDevicesContactIDNotFound{}
}

/*PutAPIPublicV1UserUserContactMethodsDevicesContactIDNotFound handles this case with default header values.

User not found
*/
type PutAPIPublicV1UserUserContactMethodsDevicesContactIDNotFound struct {
}

func (o *PutAPIPublicV1UserUserContactMethodsDevicesContactIDNotFound) Error() string {
	return fmt.Sprintf("[PUT /api-public/v1/user/{user}/contact-methods/devices/{contactId}][%d] putApiPublicV1UserUserContactMethodsDevicesContactIdNotFound ", 404)
}

func (o *PutAPIPublicV1UserUserContactMethodsDevicesContactIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutAPIPublicV1UserUserContactMethodsDevicesContactIDInternalServerError creates a PutAPIPublicV1UserUserContactMethodsDevicesContactIDInternalServerError with default headers values
func NewPutAPIPublicV1UserUserContactMethodsDevicesContactIDInternalServerError() *PutAPIPublicV1UserUserContactMethodsDevicesContactIDInternalServerError {
	return &PutAPIPublicV1UserUserContactMethodsDevicesContactIDInternalServerError{}
}

/*PutAPIPublicV1UserUserContactMethodsDevicesContactIDInternalServerError handles this case with default header values.

Internal Server Error
*/
type PutAPIPublicV1UserUserContactMethodsDevicesContactIDInternalServerError struct {
}

func (o *PutAPIPublicV1UserUserContactMethodsDevicesContactIDInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api-public/v1/user/{user}/contact-methods/devices/{contactId}][%d] putApiPublicV1UserUserContactMethodsDevicesContactIdInternalServerError ", 500)
}

func (o *PutAPIPublicV1UserUserContactMethodsDevicesContactIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
