package svcaccess

import (
	"context"

	corelib "go.deployport.com/api-services-corelib"
	iam "go.deployport.com/iam-sdk/client"
)

// Asserter helps with asserting actions against IAM
type Asserter struct {
	iam       *iam.Client
	namespace string
	region    string
}

// NewAsserter creates a new asserter
//
// Parameters:
//   - iam: IAM client for performing authorization checks
//   - namespace: service namespace for scoping assertions
//   - region: AWS region where the service operates
func NewAsserter(iam *iam.Client, namespace, region string) *Asserter {
	return &Asserter{
		iam:       iam,
		namespace: namespace,
		region:    region,
	}
}

// AssertAction asserts an action
func (a *Asserter) AssertAction(
	ctx context.Context,
	accessKeyID string,
	action string,
	resource *iam.InternalAssertActionQualifier,
) (*iam.InternalServicesAssertActionOutput, error) {
	r, err := a.iam.Internal.Services.AssertAction(ctx, &iam.InternalServicesAssertActionInput{
		CallerNamespace:   a.namespace,
		CallerRegion:      a.region,
		CallerAccessKeyID: accessKeyID,
		CallerAction:      action,
		CallerResource:    resource,
	})
	if iam.IsInternalServicesAssertActionCallerForbiddenProblem(err) {
		return nil, &corelib.ForbiddenProblem{
			Message: err.Error(),
		}
	} else if err != nil {
		return nil, err
	}
	return r, nil
}

// AssertQuery asserts a query
func (a *Asserter) AssertQuery(
	ctx context.Context,
	accessKeyID string,
	action string,
	resource *iam.InternalAssertActionQualifier,
) (*iam.InternalServicesAssertQueryOutput, error) {
	r, err := a.iam.Internal.Services.AssertQuery(ctx, &iam.InternalServicesAssertQueryInput{
		CallerNamespace:            a.namespace,
		CallerRegion:               a.region,
		CallerAccessKeyID:          accessKeyID,
		CallerAction:               action,
		CallerResourceReplacements: resource,
	})
	if iam.IsInternalServicesAssertActionCallerForbiddenProblem(err) {
		return nil, &corelib.ForbiddenProblem{
			Message: err.Error(),
		}
	} else if err != nil {
		return nil, err
	}
	return r, nil
}

// BuildAction creates a new action asserter builder
func (a *Asserter) BuildAction(accessKeyID string) *ActionAssertionBuilder {
	return &ActionAssertionBuilder{
		assertionBuilderBase: newAssertionBuilderBase(a, accessKeyID),
	}
}

// BuildQuery creates a new query asserter builder
func (a *Asserter) BuildQuery(accessKeyID string) *QueryAssertionBuilder {
	return &QueryAssertionBuilder{
		assertionBuilderBase: newAssertionBuilderBase(a, accessKeyID),
	}
}
