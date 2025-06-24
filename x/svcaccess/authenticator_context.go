package svcaccess

import (
	"context"

	iam "go.deployport.com/iam-sdk/client"
)

// PrincipalContextType is a type that can be used to embed a principal in a context
type PrincipalContextType[T any] struct{}

// NewPrincipalContextType returns a new PrincipalContextType
func NewPrincipalContextType[T any]() *PrincipalContextType[T] {
	return &PrincipalContextType[T]{}
}

// KeyPrincipalContextType is a PrincipalContextType for keyring.UnsealedKey
var KeyPrincipalContextType = NewPrincipalContextType[*iam.InternalAccessKey]()

func (c *PrincipalContextType[T]) WithValue(ctx context.Context, key T) context.Context {
	return context.WithValue(ctx, accessKeyIDContextKeyValue, key)
}

// AccessKeyFromContext returns the access key from the context, otherwise nil
func (c *PrincipalContextType[T]) FromContext(ctx context.Context) (val T, found bool) {
	key, found := ctx.Value(accessKeyIDContextKeyValue).(T)
	var zero T
	if !found {
		return zero, false
	}
	return key, true
}

type accessKeyIDContextKeyType struct{}

var accessKeyIDContextKeyValue = accessKeyIDContextKeyType{}
