// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// The input to the ListCertificatesByCA operation.
type ListCertificatesByCAInput struct {
	_ struct{} `type:"structure"`

	// Specifies the order for results. If True, the results are returned in ascending
	// order, based on the creation date.
	AscendingOrder *bool `location:"querystring" locationName:"isAscendingOrder" type:"boolean"`

	// The ID of the CA certificate. This operation will list all registered device
	// certificate that were signed by this CA certificate.
	//
	// CaCertificateId is a required field
	CaCertificateId *string `location:"uri" locationName:"caCertificateId" min:"64" type:"string" required:"true"`

	// The marker for the next set of results.
	Marker *string `location:"querystring" locationName:"marker" type:"string"`

	// The result page size.
	PageSize *int64 `location:"querystring" locationName:"pageSize" min:"1" type:"integer"`
}

// String returns the string representation
func (s ListCertificatesByCAInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListCertificatesByCAInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListCertificatesByCAInput"}

	if s.CaCertificateId == nil {
		invalidParams.Add(aws.NewErrParamRequired("CaCertificateId"))
	}
	if s.CaCertificateId != nil && len(*s.CaCertificateId) < 64 {
		invalidParams.Add(aws.NewErrParamMinLen("CaCertificateId", 64))
	}
	if s.PageSize != nil && *s.PageSize < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("PageSize", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListCertificatesByCAInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.CaCertificateId != nil {
		v := *s.CaCertificateId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "caCertificateId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.AscendingOrder != nil {
		v := *s.AscendingOrder

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "isAscendingOrder", protocol.BoolValue(v), metadata)
	}
	if s.Marker != nil {
		v := *s.Marker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "marker", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PageSize != nil {
		v := *s.PageSize

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "pageSize", protocol.Int64Value(v), metadata)
	}
	return nil
}

// The output of the ListCertificatesByCA operation.
type ListCertificatesByCAOutput struct {
	_ struct{} `type:"structure"`

	// The device certificates signed by the specified CA certificate.
	Certificates []Certificate `locationName:"certificates" type:"list"`

	// The marker for the next set of results, or null if there are no additional
	// results.
	NextMarker *string `locationName:"nextMarker" type:"string"`
}

// String returns the string representation
func (s ListCertificatesByCAOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListCertificatesByCAOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Certificates != nil {
		v := s.Certificates

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "certificates", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextMarker != nil {
		v := *s.NextMarker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextMarker", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListCertificatesByCA = "ListCertificatesByCA"

// ListCertificatesByCARequest returns a request value for making API operation for
// AWS IoT.
//
// List the device certificates signed by the specified CA certificate.
//
//    // Example sending a request using ListCertificatesByCARequest.
//    req := client.ListCertificatesByCARequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) ListCertificatesByCARequest(input *ListCertificatesByCAInput) ListCertificatesByCARequest {
	op := &aws.Operation{
		Name:       opListCertificatesByCA,
		HTTPMethod: "GET",
		HTTPPath:   "/certificates-by-ca/{caCertificateId}",
	}

	if input == nil {
		input = &ListCertificatesByCAInput{}
	}

	req := c.newRequest(op, input, &ListCertificatesByCAOutput{})
	return ListCertificatesByCARequest{Request: req, Input: input, Copy: c.ListCertificatesByCARequest}
}

// ListCertificatesByCARequest is the request type for the
// ListCertificatesByCA API operation.
type ListCertificatesByCARequest struct {
	*aws.Request
	Input *ListCertificatesByCAInput
	Copy  func(*ListCertificatesByCAInput) ListCertificatesByCARequest
}

// Send marshals and sends the ListCertificatesByCA API request.
func (r ListCertificatesByCARequest) Send(ctx context.Context) (*ListCertificatesByCAResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListCertificatesByCAResponse{
		ListCertificatesByCAOutput: r.Request.Data.(*ListCertificatesByCAOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListCertificatesByCAResponse is the response type for the
// ListCertificatesByCA API operation.
type ListCertificatesByCAResponse struct {
	*ListCertificatesByCAOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListCertificatesByCA request.
func (r *ListCertificatesByCAResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}