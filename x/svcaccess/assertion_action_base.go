package svcaccess

import (
	"fmt"

	iam "go.deployport.com/iam-sdk/client"
)

// addResourceQualifier is a helper function to add a resource qualifier to a chain
func addResourceQualifier(root, latest *iam.InternalAssertActionQualifier, function, value string) (newRoot, newLatest *iam.InternalAssertActionQualifier) {
	new := &iam.InternalAssertActionQualifier{
		Function: function,
		Value:    value,
	}
	if root == nil {
		newRoot = new
	} else {
		newRoot = root
	}
	if latest != nil {
		latest.Child = new
	}
	newLatest = new
	return
}

// assertionBuilderBase contains common fields and methods for assertion builders
type assertionBuilderBase struct {
	asserter    *Asserter
	action      string
	accessKeyID string
	root        *iam.InternalAssertActionQualifier
	latest      *iam.InternalAssertActionQualifier
}

func newAssertionBuilderBase(asserter *Asserter, accessKeyID string) assertionBuilderBase {
	return assertionBuilderBase{
		asserter:    asserter,
		accessKeyID: accessKeyID,
		root:        nil,
		latest:      nil,
		action:      "",
	}
}

// setAction sets the action for the builder
func (ab *assertionBuilderBase) setAction(action string) {
	ab.action = action
}

// addResource adds a resource to the builder
func (ab *assertionBuilderBase) addResource(function, value string) {
	ab.root, ab.latest = addResourceQualifier(ab.root, ab.latest, function, value)
}

// validate validates the builder state
func (ab *assertionBuilderBase) validate() error {
	if ab.latest == nil {
		return fmt.Errorf("builder incomplete, resource required")
	}
	if ab.action == "" {
		return fmt.Errorf("builder incomplete, action required")
	}
	return nil
}
