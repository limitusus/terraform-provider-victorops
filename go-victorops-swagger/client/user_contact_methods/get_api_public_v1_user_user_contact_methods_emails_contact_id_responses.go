// Code generated by go-swagger; DO NOT EDIT.

package user_contact_methods

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/erikvanbrakel/terraform-provider-victorops/go-victorops-swagger/models"
)

// GetAPIPublicV1UserUserContactMethodsEmailsContactIDReader is a Reader for the GetAPIPublicV1UserUserContactMethodsEmailsContactID structure.
type GetAPIPublicV1UserUserContactMethodsEmailsContactIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIPublicV1UserUserContactMethodsEmailsContactIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAPIPublicV1UserUserContactMethodsEmailsContactIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetAPIPublicV1UserUserContactMethodsEmailsContactIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetAPIPublicV1UserUserContactMethodsEmailsContactIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetAPIPublicV1UserUserContactMethodsEmailsContactIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetAPIPublicV1UserUserContactMethodsEmailsContactIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAPIPublicV1UserUserContactMethodsEmailsContactIDOK creates a GetAPIPublicV1UserUserContactMethodsEmailsContactIDOK with default headers values
func NewGetAPIPublicV1UserUserContactMethodsEmailsContactIDOK() *GetAPIPublicV1UserUserContactMethodsEmailsContactIDOK {
	return &GetAPIPublicV1UserUserContactMethodsEmailsContactIDOK{}
}

/*GetAPIPublicV1UserUserContactMethodsEmailsContactIDOK handles this case with default header values.

The indicated contact email for the user
*/
type GetAPIPublicV1UserUserContactMethodsEmailsContactIDOK struct {
	Payload []*models.UserContact
}

func (o *GetAPIPublicV1UserUserContactMethodsEmailsContactIDOK) Error() string {
	return fmt.Sprintf("[GET /api-public/v1/user/{user}/contact-methods/emails/{contactId}][%d] getApiPublicV1UserUserContactMethodsEmailsContactIdOK  %+v", 200, o.Payload)
}

func (o *GetAPIPublicV1UserUserContactMethodsEmailsContactIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIPublicV1UserUserContactMethodsEmailsContactIDUnauthorized creates a GetAPIPublicV1UserUserContactMethodsEmailsContactIDUnauthorized with default headers values
func NewGetAPIPublicV1UserUserContactMethodsEmailsContactIDUnauthorized() *GetAPIPublicV1UserUserContactMethodsEmailsContactIDUnauthorized {
	return &GetAPIPublicV1UserUserContactMethodsEmailsContactIDUnauthorized{}
}

/*GetAPIPublicV1UserUserContactMethodsEmailsContactIDUnauthorized handles this case with default header values.

Authentication parameters missing
*/
type GetAPIPublicV1UserUserContactMethodsEmailsContactIDUnauthorized struct {
}

func (o *GetAPIPublicV1UserUserContactMethodsEmailsContactIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api-public/v1/user/{user}/contact-methods/emails/{contactId}][%d] getApiPublicV1UserUserContactMethodsEmailsContactIdUnauthorized ", 401)
}

func (o *GetAPIPublicV1UserUserContactMethodsEmailsContactIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAPIPublicV1UserUserContactMethodsEmailsContactIDForbidden creates a GetAPIPublicV1UserUserContactMethodsEmailsContactIDForbidden with default headers values
func NewGetAPIPublicV1UserUserContactMethodsEmailsContactIDForbidden() *GetAPIPublicV1UserUserContactMethodsEmailsContactIDForbidden {
	return &GetAPIPublicV1UserUserContactMethodsEmailsContactIDForbidden{}
}

/*GetAPIPublicV1UserUserContactMethodsEmailsContactIDForbidden handles this case with default header values.

Authentication failed or rate-limit reached
*/
type GetAPIPublicV1UserUserContactMethodsEmailsContactIDForbidden struct {
}

func (o *GetAPIPublicV1UserUserContactMethodsEmailsContactIDForbidden) Error() string {
	return fmt.Sprintf("[GET /api-public/v1/user/{user}/contact-methods/emails/{contactId}][%d] getApiPublicV1UserUserContactMethodsEmailsContactIdForbidden ", 403)
}

func (o *GetAPIPublicV1UserUserContactMethodsEmailsContactIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAPIPublicV1UserUserContactMethodsEmailsContactIDNotFound creates a GetAPIPublicV1UserUserContactMethodsEmailsContactIDNotFound with default headers values
func NewGetAPIPublicV1UserUserContactMethodsEmailsContactIDNotFound() *GetAPIPublicV1UserUserContactMethodsEmailsContactIDNotFound {
	return &GetAPIPublicV1UserUserContactMethodsEmailsContactIDNotFound{}
}

/*GetAPIPublicV1UserUserContactMethodsEmailsContactIDNotFound handles this case with default header values.

User not found
*/
type GetAPIPublicV1UserUserContactMethodsEmailsContactIDNotFound struct {
}

func (o *GetAPIPublicV1UserUserContactMethodsEmailsContactIDNotFound) Error() string {
	return fmt.Sprintf("[GET /api-public/v1/user/{user}/contact-methods/emails/{contactId}][%d] getApiPublicV1UserUserContactMethodsEmailsContactIdNotFound ", 404)
}

func (o *GetAPIPublicV1UserUserContactMethodsEmailsContactIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAPIPublicV1UserUserContactMethodsEmailsContactIDInternalServerError creates a GetAPIPublicV1UserUserContactMethodsEmailsContactIDInternalServerError with default headers values
func NewGetAPIPublicV1UserUserContactMethodsEmailsContactIDInternalServerError() *GetAPIPublicV1UserUserContactMethodsEmailsContactIDInternalServerError {
	return &GetAPIPublicV1UserUserContactMethodsEmailsContactIDInternalServerError{}
}

/*GetAPIPublicV1UserUserContactMethodsEmailsContactIDInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetAPIPublicV1UserUserContactMethodsEmailsContactIDInternalServerError struct {
}

func (o *GetAPIPublicV1UserUserContactMethodsEmailsContactIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api-public/v1/user/{user}/contact-methods/emails/{contactId}][%d] getApiPublicV1UserUserContactMethodsEmailsContactIdInternalServerError ", 500)
}

func (o *GetAPIPublicV1UserUserContactMethodsEmailsContactIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
