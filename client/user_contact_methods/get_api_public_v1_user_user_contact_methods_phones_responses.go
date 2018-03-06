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

// GetAPIPublicV1UserUserContactMethodsPhonesReader is a Reader for the GetAPIPublicV1UserUserContactMethodsPhones structure.
type GetAPIPublicV1UserUserContactMethodsPhonesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIPublicV1UserUserContactMethodsPhonesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAPIPublicV1UserUserContactMethodsPhonesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetAPIPublicV1UserUserContactMethodsPhonesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetAPIPublicV1UserUserContactMethodsPhonesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetAPIPublicV1UserUserContactMethodsPhonesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetAPIPublicV1UserUserContactMethodsPhonesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAPIPublicV1UserUserContactMethodsPhonesOK creates a GetAPIPublicV1UserUserContactMethodsPhonesOK with default headers values
func NewGetAPIPublicV1UserUserContactMethodsPhonesOK() *GetAPIPublicV1UserUserContactMethodsPhonesOK {
	return &GetAPIPublicV1UserUserContactMethodsPhonesOK{}
}

/*GetAPIPublicV1UserUserContactMethodsPhonesOK handles this case with default header values.

The list of contact phones for the user
*/
type GetAPIPublicV1UserUserContactMethodsPhonesOK struct {
	Payload models.GetAPIPublicV1UserUserContactMethodsPhonesOKBody
}

func (o *GetAPIPublicV1UserUserContactMethodsPhonesOK) Error() string {
	return fmt.Sprintf("[GET /api-public/v1/user/{user}/contact-methods/phones][%d] getApiPublicV1UserUserContactMethodsPhonesOK  %+v", 200, o.Payload)
}

func (o *GetAPIPublicV1UserUserContactMethodsPhonesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIPublicV1UserUserContactMethodsPhonesUnauthorized creates a GetAPIPublicV1UserUserContactMethodsPhonesUnauthorized with default headers values
func NewGetAPIPublicV1UserUserContactMethodsPhonesUnauthorized() *GetAPIPublicV1UserUserContactMethodsPhonesUnauthorized {
	return &GetAPIPublicV1UserUserContactMethodsPhonesUnauthorized{}
}

/*GetAPIPublicV1UserUserContactMethodsPhonesUnauthorized handles this case with default header values.

Authentication parameters missing
*/
type GetAPIPublicV1UserUserContactMethodsPhonesUnauthorized struct {
}

func (o *GetAPIPublicV1UserUserContactMethodsPhonesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api-public/v1/user/{user}/contact-methods/phones][%d] getApiPublicV1UserUserContactMethodsPhonesUnauthorized ", 401)
}

func (o *GetAPIPublicV1UserUserContactMethodsPhonesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAPIPublicV1UserUserContactMethodsPhonesForbidden creates a GetAPIPublicV1UserUserContactMethodsPhonesForbidden with default headers values
func NewGetAPIPublicV1UserUserContactMethodsPhonesForbidden() *GetAPIPublicV1UserUserContactMethodsPhonesForbidden {
	return &GetAPIPublicV1UserUserContactMethodsPhonesForbidden{}
}

/*GetAPIPublicV1UserUserContactMethodsPhonesForbidden handles this case with default header values.

Authentication failed or rate-limit reached
*/
type GetAPIPublicV1UserUserContactMethodsPhonesForbidden struct {
}

func (o *GetAPIPublicV1UserUserContactMethodsPhonesForbidden) Error() string {
	return fmt.Sprintf("[GET /api-public/v1/user/{user}/contact-methods/phones][%d] getApiPublicV1UserUserContactMethodsPhonesForbidden ", 403)
}

func (o *GetAPIPublicV1UserUserContactMethodsPhonesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAPIPublicV1UserUserContactMethodsPhonesNotFound creates a GetAPIPublicV1UserUserContactMethodsPhonesNotFound with default headers values
func NewGetAPIPublicV1UserUserContactMethodsPhonesNotFound() *GetAPIPublicV1UserUserContactMethodsPhonesNotFound {
	return &GetAPIPublicV1UserUserContactMethodsPhonesNotFound{}
}

/*GetAPIPublicV1UserUserContactMethodsPhonesNotFound handles this case with default header values.

User not found
*/
type GetAPIPublicV1UserUserContactMethodsPhonesNotFound struct {
}

func (o *GetAPIPublicV1UserUserContactMethodsPhonesNotFound) Error() string {
	return fmt.Sprintf("[GET /api-public/v1/user/{user}/contact-methods/phones][%d] getApiPublicV1UserUserContactMethodsPhonesNotFound ", 404)
}

func (o *GetAPIPublicV1UserUserContactMethodsPhonesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAPIPublicV1UserUserContactMethodsPhonesInternalServerError creates a GetAPIPublicV1UserUserContactMethodsPhonesInternalServerError with default headers values
func NewGetAPIPublicV1UserUserContactMethodsPhonesInternalServerError() *GetAPIPublicV1UserUserContactMethodsPhonesInternalServerError {
	return &GetAPIPublicV1UserUserContactMethodsPhonesInternalServerError{}
}

/*GetAPIPublicV1UserUserContactMethodsPhonesInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetAPIPublicV1UserUserContactMethodsPhonesInternalServerError struct {
}

func (o *GetAPIPublicV1UserUserContactMethodsPhonesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api-public/v1/user/{user}/contact-methods/phones][%d] getApiPublicV1UserUserContactMethodsPhonesInternalServerError ", 500)
}

func (o *GetAPIPublicV1UserUserContactMethodsPhonesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
