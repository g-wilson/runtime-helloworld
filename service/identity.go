package service

import (
	"context"

	"github.com/g-wilson/runtime/logger"
)

type identityContextKey string

var ctxkey = identityContextKey("authidentity")

// Identity represents the bearer of the auth token used in a request
type Identity struct {
	UserID string
}

// NewIdentity returns a struct based on the access token claims
func NewIdentity(claims map[string]interface{}) *Identity {
	b := &Identity{}

	if sub, ok := claims["sub"].(string); ok {
		b.UserID = sub
	}

	return b
}

// GetIdentity returns the identity of the requester from the request context
func GetIdentity(ctx context.Context) *Identity {
	val := ctx.Value(ctxkey)

	if claims, ok := val.(*Identity); ok {
		return claims
	}

	return &Identity{}
}

// SetIdentity adds the identity of the requester to the request context
func SetIdentity(ctx context.Context, claims map[string]interface{}) context.Context {
	b := NewIdentity(claims)
	ctx = context.WithValue(ctx, ctxkey, b)

	return ctx
}

// IdentityLogger is a runtime context provider which adds identity details to the request log
func IdentityLogger(ctx context.Context) context.Context {
	b := GetIdentity(ctx)
	reqLogger := logger.FromContext(ctx)

	if reqLogger != nil {
		if b.UserID != "" {
			reqLogger.Update(reqLogger.Entry().WithField("user_id", b.UserID))
		}
	}

	return ctx
}
