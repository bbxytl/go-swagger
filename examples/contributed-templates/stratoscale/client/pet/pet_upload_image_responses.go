// Code generated by go-swagger; DO NOT EDIT.

package pet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/go-swagger/go-swagger/examples/contributed-templates/stratoscale/models"
)

// PetUploadImageReader is a Reader for the PetUploadImage structure.
type PetUploadImageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PetUploadImageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPetUploadImageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPetUploadImageOK creates a PetUploadImageOK with default headers values
func NewPetUploadImageOK() *PetUploadImageOK {
	return &PetUploadImageOK{}
}

/*PetUploadImageOK handles this case with default header values.

successful operation
*/
type PetUploadImageOK struct {
	Payload *models.APIResponse
}

func (o *PetUploadImageOK) Error() string {
	return fmt.Sprintf("[POST /pet/{petId}/image][%d] petUploadImageOK  %+v", 200, o.Payload)
}

func (o *PetUploadImageOK) GetPayload() *models.APIResponse {
	return o.Payload
}

func (o *PetUploadImageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
