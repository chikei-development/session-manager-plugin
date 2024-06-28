// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package savingsplansiface provides an interface to enable mocking the AWS Savings Plans service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package savingsplansiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/savingsplans"
)

// SavingsPlansAPI provides an interface to enable mocking the
// savingsplans.SavingsPlans service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// AWS Savings Plans.
//	func myFunc(svc savingsplansiface.SavingsPlansAPI) bool {
//	    // Make svc.CreateSavingsPlan request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := savingsplans.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockSavingsPlansClient struct {
//	    savingsplansiface.SavingsPlansAPI
//	}
//	func (m *mockSavingsPlansClient) CreateSavingsPlan(input *savingsplans.CreateSavingsPlanInput) (*savingsplans.CreateSavingsPlanOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockSavingsPlansClient{}
//
//	    myfunc(mockSvc)
//
//	    // Verify myFunc's functionality
//	}
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type SavingsPlansAPI interface {
	CreateSavingsPlan(*savingsplans.CreateSavingsPlanInput) (*savingsplans.CreateSavingsPlanOutput, error)
	CreateSavingsPlanWithContext(aws.Context, *savingsplans.CreateSavingsPlanInput, ...request.Option) (*savingsplans.CreateSavingsPlanOutput, error)
	CreateSavingsPlanRequest(*savingsplans.CreateSavingsPlanInput) (*request.Request, *savingsplans.CreateSavingsPlanOutput)

	DeleteQueuedSavingsPlan(*savingsplans.DeleteQueuedSavingsPlanInput) (*savingsplans.DeleteQueuedSavingsPlanOutput, error)
	DeleteQueuedSavingsPlanWithContext(aws.Context, *savingsplans.DeleteQueuedSavingsPlanInput, ...request.Option) (*savingsplans.DeleteQueuedSavingsPlanOutput, error)
	DeleteQueuedSavingsPlanRequest(*savingsplans.DeleteQueuedSavingsPlanInput) (*request.Request, *savingsplans.DeleteQueuedSavingsPlanOutput)

	DescribeSavingsPlanRates(*savingsplans.DescribeSavingsPlanRatesInput) (*savingsplans.DescribeSavingsPlanRatesOutput, error)
	DescribeSavingsPlanRatesWithContext(aws.Context, *savingsplans.DescribeSavingsPlanRatesInput, ...request.Option) (*savingsplans.DescribeSavingsPlanRatesOutput, error)
	DescribeSavingsPlanRatesRequest(*savingsplans.DescribeSavingsPlanRatesInput) (*request.Request, *savingsplans.DescribeSavingsPlanRatesOutput)

	DescribeSavingsPlans(*savingsplans.DescribeSavingsPlansInput) (*savingsplans.DescribeSavingsPlansOutput, error)
	DescribeSavingsPlansWithContext(aws.Context, *savingsplans.DescribeSavingsPlansInput, ...request.Option) (*savingsplans.DescribeSavingsPlansOutput, error)
	DescribeSavingsPlansRequest(*savingsplans.DescribeSavingsPlansInput) (*request.Request, *savingsplans.DescribeSavingsPlansOutput)

	DescribeSavingsPlansOfferingRates(*savingsplans.DescribeSavingsPlansOfferingRatesInput) (*savingsplans.DescribeSavingsPlansOfferingRatesOutput, error)
	DescribeSavingsPlansOfferingRatesWithContext(aws.Context, *savingsplans.DescribeSavingsPlansOfferingRatesInput, ...request.Option) (*savingsplans.DescribeSavingsPlansOfferingRatesOutput, error)
	DescribeSavingsPlansOfferingRatesRequest(*savingsplans.DescribeSavingsPlansOfferingRatesInput) (*request.Request, *savingsplans.DescribeSavingsPlansOfferingRatesOutput)

	DescribeSavingsPlansOfferings(*savingsplans.DescribeSavingsPlansOfferingsInput) (*savingsplans.DescribeSavingsPlansOfferingsOutput, error)
	DescribeSavingsPlansOfferingsWithContext(aws.Context, *savingsplans.DescribeSavingsPlansOfferingsInput, ...request.Option) (*savingsplans.DescribeSavingsPlansOfferingsOutput, error)
	DescribeSavingsPlansOfferingsRequest(*savingsplans.DescribeSavingsPlansOfferingsInput) (*request.Request, *savingsplans.DescribeSavingsPlansOfferingsOutput)

	ListTagsForResource(*savingsplans.ListTagsForResourceInput) (*savingsplans.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *savingsplans.ListTagsForResourceInput, ...request.Option) (*savingsplans.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*savingsplans.ListTagsForResourceInput) (*request.Request, *savingsplans.ListTagsForResourceOutput)

	ReturnSavingsPlan(*savingsplans.ReturnSavingsPlanInput) (*savingsplans.ReturnSavingsPlanOutput, error)
	ReturnSavingsPlanWithContext(aws.Context, *savingsplans.ReturnSavingsPlanInput, ...request.Option) (*savingsplans.ReturnSavingsPlanOutput, error)
	ReturnSavingsPlanRequest(*savingsplans.ReturnSavingsPlanInput) (*request.Request, *savingsplans.ReturnSavingsPlanOutput)

	TagResource(*savingsplans.TagResourceInput) (*savingsplans.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *savingsplans.TagResourceInput, ...request.Option) (*savingsplans.TagResourceOutput, error)
	TagResourceRequest(*savingsplans.TagResourceInput) (*request.Request, *savingsplans.TagResourceOutput)

	UntagResource(*savingsplans.UntagResourceInput) (*savingsplans.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *savingsplans.UntagResourceInput, ...request.Option) (*savingsplans.UntagResourceOutput, error)
	UntagResourceRequest(*savingsplans.UntagResourceInput) (*request.Request, *savingsplans.UntagResourceOutput)
}

var _ SavingsPlansAPI = (*savingsplans.SavingsPlans)(nil)
