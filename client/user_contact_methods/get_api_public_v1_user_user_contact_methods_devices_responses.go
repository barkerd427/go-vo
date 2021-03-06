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

// GetAPIPublicV1UserUserContactMethodsDevicesReader is a Reader for the GetAPIPublicV1UserUserContactMethodsDevices structure.
type GetAPIPublicV1UserUserContactMethodsDevicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIPublicV1UserUserContactMethodsDevicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAPIPublicV1UserUserContactMethodsDevicesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetAPIPublicV1UserUserContactMethodsDevicesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetAPIPublicV1UserUserContactMethodsDevicesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetAPIPublicV1UserUserContactMethodsDevicesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetAPIPublicV1UserUserContactMethodsDevicesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAPIPublicV1UserUserContactMethodsDevicesOK creates a GetAPIPublicV1UserUserContactMethodsDevicesOK with default headers values
func NewGetAPIPublicV1UserUserContactMethodsDevicesOK() *GetAPIPublicV1UserUserContactMethodsDevicesOK {
	return &GetAPIPublicV1UserUserContactMethodsDevicesOK{}
}

/*GetAPIPublicV1UserUserContactMethodsDevicesOK handles this case with default header values.

The list of contact devices for the user
*/
type GetAPIPublicV1UserUserContactMethodsDevicesOK struct {
	Payload models.GetAPIPublicV1UserUserContactMethodsDevicesOKBody
}

func (o *GetAPIPublicV1UserUserContactMethodsDevicesOK) Error() string {
	return fmt.Sprintf("[GET /api-public/v1/user/{user}/contact-methods/devices][%d] getApiPublicV1UserUserContactMethodsDevicesOK  %+v", 200, o.Payload)
}

func (o *GetAPIPublicV1UserUserContactMethodsDevicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIPublicV1UserUserContactMethodsDevicesUnauthorized creates a GetAPIPublicV1UserUserContactMethodsDevicesUnauthorized with default headers values
func NewGetAPIPublicV1UserUserContactMethodsDevicesUnauthorized() *GetAPIPublicV1UserUserContactMethodsDevicesUnauthorized {
	return &GetAPIPublicV1UserUserContactMethodsDevicesUnauthorized{}
}

/*GetAPIPublicV1UserUserContactMethodsDevicesUnauthorized handles this case with default header values.

Authentication parameters missing
*/
type GetAPIPublicV1UserUserContactMethodsDevicesUnauthorized struct {
}

func (o *GetAPIPublicV1UserUserContactMethodsDevicesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api-public/v1/user/{user}/contact-methods/devices][%d] getApiPublicV1UserUserContactMethodsDevicesUnauthorized ", 401)
}

func (o *GetAPIPublicV1UserUserContactMethodsDevicesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAPIPublicV1UserUserContactMethodsDevicesForbidden creates a GetAPIPublicV1UserUserContactMethodsDevicesForbidden with default headers values
func NewGetAPIPublicV1UserUserContactMethodsDevicesForbidden() *GetAPIPublicV1UserUserContactMethodsDevicesForbidden {
	return &GetAPIPublicV1UserUserContactMethodsDevicesForbidden{}
}

/*GetAPIPublicV1UserUserContactMethodsDevicesForbidden handles this case with default header values.

Authentication failed or rate-limit reached
*/
type GetAPIPublicV1UserUserContactMethodsDevicesForbidden struct {
}

func (o *GetAPIPublicV1UserUserContactMethodsDevicesForbidden) Error() string {
	return fmt.Sprintf("[GET /api-public/v1/user/{user}/contact-methods/devices][%d] getApiPublicV1UserUserContactMethodsDevicesForbidden ", 403)
}

func (o *GetAPIPublicV1UserUserContactMethodsDevicesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAPIPublicV1UserUserContactMethodsDevicesNotFound creates a GetAPIPublicV1UserUserContactMethodsDevicesNotFound with default headers values
func NewGetAPIPublicV1UserUserContactMethodsDevicesNotFound() *GetAPIPublicV1UserUserContactMethodsDevicesNotFound {
	return &GetAPIPublicV1UserUserContactMethodsDevicesNotFound{}
}

/*GetAPIPublicV1UserUserContactMethodsDevicesNotFound handles this case with default header values.

User not found
*/
type GetAPIPublicV1UserUserContactMethodsDevicesNotFound struct {
}

func (o *GetAPIPublicV1UserUserContactMethodsDevicesNotFound) Error() string {
	return fmt.Sprintf("[GET /api-public/v1/user/{user}/contact-methods/devices][%d] getApiPublicV1UserUserContactMethodsDevicesNotFound ", 404)
}

func (o *GetAPIPublicV1UserUserContactMethodsDevicesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAPIPublicV1UserUserContactMethodsDevicesInternalServerError creates a GetAPIPublicV1UserUserContactMethodsDevicesInternalServerError with default headers values
func NewGetAPIPublicV1UserUserContactMethodsDevicesInternalServerError() *GetAPIPublicV1UserUserContactMethodsDevicesInternalServerError {
	return &GetAPIPublicV1UserUserContactMethodsDevicesInternalServerError{}
}

/*GetAPIPublicV1UserUserContactMethodsDevicesInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetAPIPublicV1UserUserContactMethodsDevicesInternalServerError struct {
}

func (o *GetAPIPublicV1UserUserContactMethodsDevicesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api-public/v1/user/{user}/contact-methods/devices][%d] getApiPublicV1UserUserContactMethodsDevicesInternalServerError ", 500)
}

func (o *GetAPIPublicV1UserUserContactMethodsDevicesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
