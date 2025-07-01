package svcaccess

import (
	"context"

	iam "go.deployport.com/iam-sdk/client"
)

// QueryAssertionBuilder builds action assertions
type QueryAssertionBuilder struct {
	assertionBuilderBase
}

// Action sets the action for the asserter
func (ab *QueryAssertionBuilder) Action(action string) *QueryAssertionBuilder {
	ab.setAction(action)
	return ab
}

// Resource adds the resource to the action asserter as root or child of a previous resource
func (ab *QueryAssertionBuilder) Resource(function string, column AssertionColumn) *QueryAssertionBuilder {
	return ab.ResourceExpression(function, column.String())
}

// ResourceExpression adds the resource to the action asserter as root or child of a previous resource
func (ab *QueryAssertionBuilder) ResourceExpression(function string, expr string) *QueryAssertionBuilder {
	ab.addResource(function, expr)
	return ab
}

// Exec executes the action asserter
func (ab *QueryAssertionBuilder) Exec(ctx context.Context) (*iam.InternalServicesAssertQueryOutput, error) {
	if err := ab.validate(); err != nil {
		return nil, err
	}
	return ab.asserter.AssertQuery(ctx, ab.accessKeyID, ab.action, ab.root)
}
