// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package ioteventsiface provides an interface to enable mocking the AWS IoT Events service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package ioteventsiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/iotevents"
)

// IoTEventsAPI provides an interface to enable mocking the
// iotevents.IoTEvents service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS IoT Events.
//    func myFunc(svc ioteventsiface.IoTEventsAPI) bool {
//        // Make svc.CreateDetectorModel request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := iotevents.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockIoTEventsClient struct {
//        ioteventsiface.IoTEventsAPI
//    }
//    func (m *mockIoTEventsClient) CreateDetectorModel(input *iotevents.CreateDetectorModelInput) (*iotevents.CreateDetectorModelOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockIoTEventsClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type IoTEventsAPI interface {
	CreateDetectorModel(*iotevents.CreateDetectorModelInput) (*iotevents.CreateDetectorModelOutput, error)
	CreateDetectorModelWithContext(aws.Context, *iotevents.CreateDetectorModelInput, ...request.Option) (*iotevents.CreateDetectorModelOutput, error)
	CreateDetectorModelRequest(*iotevents.CreateDetectorModelInput) (*request.Request, *iotevents.CreateDetectorModelOutput)

	CreateInput(*iotevents.CreateInputInput) (*iotevents.CreateInputOutput, error)
	CreateInputWithContext(aws.Context, *iotevents.CreateInputInput, ...request.Option) (*iotevents.CreateInputOutput, error)
	CreateInputRequest(*iotevents.CreateInputInput) (*request.Request, *iotevents.CreateInputOutput)

	DeleteDetectorModel(*iotevents.DeleteDetectorModelInput) (*iotevents.DeleteDetectorModelOutput, error)
	DeleteDetectorModelWithContext(aws.Context, *iotevents.DeleteDetectorModelInput, ...request.Option) (*iotevents.DeleteDetectorModelOutput, error)
	DeleteDetectorModelRequest(*iotevents.DeleteDetectorModelInput) (*request.Request, *iotevents.DeleteDetectorModelOutput)

	DeleteInput(*iotevents.DeleteInputInput) (*iotevents.DeleteInputOutput, error)
	DeleteInputWithContext(aws.Context, *iotevents.DeleteInputInput, ...request.Option) (*iotevents.DeleteInputOutput, error)
	DeleteInputRequest(*iotevents.DeleteInputInput) (*request.Request, *iotevents.DeleteInputOutput)

	DescribeDetectorModel(*iotevents.DescribeDetectorModelInput) (*iotevents.DescribeDetectorModelOutput, error)
	DescribeDetectorModelWithContext(aws.Context, *iotevents.DescribeDetectorModelInput, ...request.Option) (*iotevents.DescribeDetectorModelOutput, error)
	DescribeDetectorModelRequest(*iotevents.DescribeDetectorModelInput) (*request.Request, *iotevents.DescribeDetectorModelOutput)

	DescribeDetectorModelAnalysis(*iotevents.DescribeDetectorModelAnalysisInput) (*iotevents.DescribeDetectorModelAnalysisOutput, error)
	DescribeDetectorModelAnalysisWithContext(aws.Context, *iotevents.DescribeDetectorModelAnalysisInput, ...request.Option) (*iotevents.DescribeDetectorModelAnalysisOutput, error)
	DescribeDetectorModelAnalysisRequest(*iotevents.DescribeDetectorModelAnalysisInput) (*request.Request, *iotevents.DescribeDetectorModelAnalysisOutput)

	DescribeInput(*iotevents.DescribeInputInput) (*iotevents.DescribeInputOutput, error)
	DescribeInputWithContext(aws.Context, *iotevents.DescribeInputInput, ...request.Option) (*iotevents.DescribeInputOutput, error)
	DescribeInputRequest(*iotevents.DescribeInputInput) (*request.Request, *iotevents.DescribeInputOutput)

	DescribeLoggingOptions(*iotevents.DescribeLoggingOptionsInput) (*iotevents.DescribeLoggingOptionsOutput, error)
	DescribeLoggingOptionsWithContext(aws.Context, *iotevents.DescribeLoggingOptionsInput, ...request.Option) (*iotevents.DescribeLoggingOptionsOutput, error)
	DescribeLoggingOptionsRequest(*iotevents.DescribeLoggingOptionsInput) (*request.Request, *iotevents.DescribeLoggingOptionsOutput)

	GetDetectorModelAnalysisResults(*iotevents.GetDetectorModelAnalysisResultsInput) (*iotevents.GetDetectorModelAnalysisResultsOutput, error)
	GetDetectorModelAnalysisResultsWithContext(aws.Context, *iotevents.GetDetectorModelAnalysisResultsInput, ...request.Option) (*iotevents.GetDetectorModelAnalysisResultsOutput, error)
	GetDetectorModelAnalysisResultsRequest(*iotevents.GetDetectorModelAnalysisResultsInput) (*request.Request, *iotevents.GetDetectorModelAnalysisResultsOutput)

	ListDetectorModelVersions(*iotevents.ListDetectorModelVersionsInput) (*iotevents.ListDetectorModelVersionsOutput, error)
	ListDetectorModelVersionsWithContext(aws.Context, *iotevents.ListDetectorModelVersionsInput, ...request.Option) (*iotevents.ListDetectorModelVersionsOutput, error)
	ListDetectorModelVersionsRequest(*iotevents.ListDetectorModelVersionsInput) (*request.Request, *iotevents.ListDetectorModelVersionsOutput)

	ListDetectorModels(*iotevents.ListDetectorModelsInput) (*iotevents.ListDetectorModelsOutput, error)
	ListDetectorModelsWithContext(aws.Context, *iotevents.ListDetectorModelsInput, ...request.Option) (*iotevents.ListDetectorModelsOutput, error)
	ListDetectorModelsRequest(*iotevents.ListDetectorModelsInput) (*request.Request, *iotevents.ListDetectorModelsOutput)

	ListInputs(*iotevents.ListInputsInput) (*iotevents.ListInputsOutput, error)
	ListInputsWithContext(aws.Context, *iotevents.ListInputsInput, ...request.Option) (*iotevents.ListInputsOutput, error)
	ListInputsRequest(*iotevents.ListInputsInput) (*request.Request, *iotevents.ListInputsOutput)

	ListTagsForResource(*iotevents.ListTagsForResourceInput) (*iotevents.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *iotevents.ListTagsForResourceInput, ...request.Option) (*iotevents.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*iotevents.ListTagsForResourceInput) (*request.Request, *iotevents.ListTagsForResourceOutput)

	PutLoggingOptions(*iotevents.PutLoggingOptionsInput) (*iotevents.PutLoggingOptionsOutput, error)
	PutLoggingOptionsWithContext(aws.Context, *iotevents.PutLoggingOptionsInput, ...request.Option) (*iotevents.PutLoggingOptionsOutput, error)
	PutLoggingOptionsRequest(*iotevents.PutLoggingOptionsInput) (*request.Request, *iotevents.PutLoggingOptionsOutput)

	StartDetectorModelAnalysis(*iotevents.StartDetectorModelAnalysisInput) (*iotevents.StartDetectorModelAnalysisOutput, error)
	StartDetectorModelAnalysisWithContext(aws.Context, *iotevents.StartDetectorModelAnalysisInput, ...request.Option) (*iotevents.StartDetectorModelAnalysisOutput, error)
	StartDetectorModelAnalysisRequest(*iotevents.StartDetectorModelAnalysisInput) (*request.Request, *iotevents.StartDetectorModelAnalysisOutput)

	TagResource(*iotevents.TagResourceInput) (*iotevents.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *iotevents.TagResourceInput, ...request.Option) (*iotevents.TagResourceOutput, error)
	TagResourceRequest(*iotevents.TagResourceInput) (*request.Request, *iotevents.TagResourceOutput)

	UntagResource(*iotevents.UntagResourceInput) (*iotevents.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *iotevents.UntagResourceInput, ...request.Option) (*iotevents.UntagResourceOutput, error)
	UntagResourceRequest(*iotevents.UntagResourceInput) (*request.Request, *iotevents.UntagResourceOutput)

	UpdateDetectorModel(*iotevents.UpdateDetectorModelInput) (*iotevents.UpdateDetectorModelOutput, error)
	UpdateDetectorModelWithContext(aws.Context, *iotevents.UpdateDetectorModelInput, ...request.Option) (*iotevents.UpdateDetectorModelOutput, error)
	UpdateDetectorModelRequest(*iotevents.UpdateDetectorModelInput) (*request.Request, *iotevents.UpdateDetectorModelOutput)

	UpdateInput(*iotevents.UpdateInputInput) (*iotevents.UpdateInputOutput, error)
	UpdateInputWithContext(aws.Context, *iotevents.UpdateInputInput, ...request.Option) (*iotevents.UpdateInputOutput, error)
	UpdateInputRequest(*iotevents.UpdateInputInput) (*request.Request, *iotevents.UpdateInputOutput)
}

var _ IoTEventsAPI = (*iotevents.IoTEvents)(nil)
