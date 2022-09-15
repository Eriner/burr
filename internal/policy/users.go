package policy

import (
	"context"

	"github.com/eriner/burr/internal/ent"
)

type contextKey string

var (
	SuperuserCtxKey  contextKey = "superuser"   // Should only be used when seeding the DB or performing OP actions.
	UserCtxKey       contextKey = "user"        // the current user.
	TargetUserCtxKey contextKey = "target-user" // the user target by the current user's operation.
)

// UserFromContext returns a User from a context or nil if not present.
func UserFromContext(reqCtx context.Context) *ent.User {
	// TODO: impl
	user, ok := reqCtx.Value(UserCtxKey).(*ent.User)
	if !ok {
		return nil
	}
	return user
}

func ContextWithSuperuser(ctx context.Context) context.Context {
	return context.WithValue(ctx, SuperuserCtxKey, true)
}

// ContextWithUser adds the user to a context.
func ContextWithUser(ctx context.Context, u *ent.User) context.Context {
	return context.WithValue(ctx, UserCtxKey, u)
}

// TargetUserFromContext returns a User from a context or nil if not present.
func TargetUserFromContext(reqCtx context.Context) *ent.User {
	tu, ok := reqCtx.Value(TargetUserCtxKey).(*ent.User)
	if !ok {
		return nil
	}
	return tu
}

// ContextWithTargetUser adds the user to a context.
func ContextWithTargetUser(ctx context.Context, u *ent.User) context.Context {
	return context.WithValue(ctx, UserCtxKey, u)
}
