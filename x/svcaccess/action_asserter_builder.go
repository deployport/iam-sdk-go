package svcaccess

import (
	"context"
)

// ActionAssertionBuilder builds action assertions
type ActionAssertionBuilder struct {
	assertionBuilderBase
}

// Action sets the action for the asserter
func (ab *ActionAssertionBuilder) Action(action string) *ActionAssertionBuilder {
	ab.setAction(action)
	return ab
}

// Resource adds the resource to the action asserter as root or child of a previous resource
func (ab *ActionAssertionBuilder) Resource(function, value string) *ActionAssertionBuilder {
	ab.addResource(function, value)
	return ab
}

// ResourceAll adds the resource to the action asserter
func (ab *ActionAssertionBuilder) ResourceAll(function string) *ActionAssertionBuilder {
	return ab.Resource(function, "*")
}

// Exec executes the action asserter returning [corelib.ForbiddenProblem] if the action is not allowed
func (ab *ActionAssertionBuilder) Exec(ctx context.Context) error {
	if err := ab.validate(); err != nil {
		return err
	}
	_, err := ab.asserter.AssertAction(ctx, ab.accessKeyID, ab.action, ab.root)
	return err
}
