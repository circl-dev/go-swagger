// Code generated by go-swagger; DO NOT EDIT.

package uploads

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/strfmt"
	"github.com/circl-dev/runtime"
)

// UploadFileReader is a Reader for the UploadFile structure.
type UploadFileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UploadFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUploadFileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUploadFileOK creates a UploadFileOK with default headers values
func NewUploadFileOK() *UploadFileOK {
	return &UploadFileOK{}
}

/* UploadFileOK describes a response with status code 200, with default header values.

OK
*/
type UploadFileOK struct {
}

func (o *UploadFileOK) Error() string {
	return fmt.Sprintf("[POST /upload][%d] uploadFileOK ", 200)
}

func (o *UploadFileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
