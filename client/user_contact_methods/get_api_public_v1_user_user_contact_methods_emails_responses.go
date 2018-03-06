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

// GetAPIPublicV1UserUserContactMethodsEmailsReader is a Reader for the GetAPIPublicV1UserUserContactMethodsEmails structure.
type GetAPIPublicV1UserUserContactMethodsEmailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIPublicV1UserUserContactMethodsEmailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAPIPublicV1UserUserContactMethodsEmailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetAPIPublicV1UserUserContactMethodsEmailsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetAPIPublicV1UserUserContactMethodsEmailsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetAPIPublicV1UserUserContactMethodsEmailsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetAPIPublicV1UserUserContactMethodsEmailsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAPIPublicV1UserUserContactMethodsEmailsOK creates a GetAPIPublicV1UserUserContactMethodsEmailsOK with default headers values
func NewGetAPIPublicV1UserUserContactMethodsEmailsOK() *GetAPIPublicV1UserUserContactMethodsEmailsOK {
	return &GetAPIPublicV1UserUserContactMethodsEmailsOK{}
}

/*GetAPIPublicV1UserUserContactMethodsEmailsOK handles this case with default header values.

The list of contact emails for the user
*/
type GetAPIPublicV1UserUserContactMethodsEmailsOK struct {
	Payload models.GetAPIPublicV1UserUserContactMethodsEmailsOKBody
}

func (o *GetAPIPublicV1UserUserContactMethodsEmailsOK) Error() string {
	return fmt.Sprintf("[GET /api-public/v1/user/{user}/contact-methods/emails][%d] getApiPublicV1UserUserContactMethodsEmailsOK  %+v", 200, o.Payload)
}

func (o *GetAPIPublicV1UserUserContactMethodsEmailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIPublicV1UserUserContactMethodsEmailsUnauthorized creates a GetAPIPublicV1UserUserContactMethodsEmailsUnauthorized with default headers values
func NewGetAPIPublicV1UserUserContactMethodsEmailsUnauthorized() *GetAPIPublicV1UserUserContactMethodsEmailsUnauthorized {
	return &GetAPIPublicV1UserUserContactMethodsEmailsUnauthorized{}
}

/*GetAPIPublicV1UserUserContactMethodsEmailsUnauthorized handles this case with default header values.

Authentication parameters missing
*/
type GetAPIPublicV1UserUserContactMethodsEmailsUnauthorized struct {
}

func (o *GetAPIPublicV1UserUserContactMethodsEmailsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api-public/v1/user/{user}/contact-methods/emails][%d] getApiPublicV1UserUserContactMethodsEmailsUnauthorized ", 401)
}

func (o *GetAPIPublicV1UserUserContactMethodsEmailsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAPIPublicV1UserUserContactMethodsEmailsForbidden creates a GetAPIPublicV1UserUserContactMethodsEmailsForbidden with default headers values
func NewGetAPIPublicV1UserUserContactMethodsEmailsForbidden() *GetAPIPublicV1UserUserContactMethodsEmailsForbidden {
	return &GetAPIPublicV1UserUserContactMethodsEmailsForbidden{}
}

/*GetAPIPublicV1UserUserContactMethodsEmailsForbidden handles this case with default header values.

Authentication failed or rate-limit reached
*/
type GetAPIPublicV1UserUserContactMethodsEmailsForbidden struct {
}

func (o *GetAPIPublicV1UserUserContactMethodsEmailsForbidden) Error() string {
	return fmt.Sprintf("[GET /api-public/v1/user/{user}/contact-methods/emails][%d] getApiPublicV1UserUserContactMethodsEmailsForbidden ", 403)
}

func (o *GetAPIPublicV1UserUserContactMethodsEmailsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAPIPublicV1UserUserContactMethodsEmailsNotFound creates a GetAPIPublicV1UserUserContactMethodsEmailsNotFound with default headers values
func NewGetAPIPublicV1UserUserContactMethodsEmailsNotFound() *GetAPIPublicV1UserUserContactMethodsEmailsNotFound {
	return &GetAPIPublicV1UserUserContactMethodsEmailsNotFound{}
}

/*GetAPIPublicV1UserUserContactMethodsEmailsNotFound handles this case with default header values.

User not found
*/
type GetAPIPublicV1UserUserContactMethodsEmailsNotFound struct {
}

func (o *GetAPIPublicV1UserUserContactMethodsEmailsNotFound) Error() string {
	return fmt.Sprintf("[GET /api-public/v1/user/{user}/contact-methods/emails][%d] getApiPublicV1UserUserContactMethodsEmailsNotFound ", 404)
}

func (o *GetAPIPublicV1UserUserContactMethodsEmailsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAPIPublicV1UserUserContactMethodsEmailsInternalServerError creates a GetAPIPublicV1UserUserContactMethodsEmailsInternalServerError with default headers values
func NewGetAPIPublicV1UserUserContactMethodsEmailsInternalServerError() *GetAPIPublicV1UserUserContactMethodsEmailsInternalServerError {
	return &GetAPIPublicV1UserUserContactMethodsEmailsInternalServerError{}
}

/*GetAPIPublicV1UserUserContactMethodsEmailsInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetAPIPublicV1UserUserContactMethodsEmailsInternalServerError struct {
}

func (o *GetAPIPublicV1UserUserContactMethodsEmailsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api-public/v1/user/{user}/contact-methods/emails][%d] getApiPublicV1UserUserContactMethodsEmailsInternalServerError ", 500)
}

func (o *GetAPIPublicV1UserUserContactMethodsEmailsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
