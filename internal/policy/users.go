package policy

import (
	"context"

	"github.com/eriner/burr/internal/ent"
)

type contextKey string

var (
	SuperuserCtxKey   contextKey = "superuser"    // Should only be used when seeding the DB or performing OP actions.
	ActorCtxKey       contextKey = "actor"        // the current authenticated actor.
	TargetActorCtxKey contextKey = "target-actor" // the user target by the current user's operation.
)

// ActorFromContext returns a Actor from a context or nil if not present.
func ActorFromContext(reqCtx context.Context) *ent.Actor {
	// TODO: impl
	user, ok := reqCtx.Value(ActorCtxKey).(*ent.Actor)
	if !ok {
		return nil
	}
	return user
}

func ContextWithSuperuser(ctx context.Context) context.Context {
	return context.WithValue(ctx, SuperuserCtxKey, true)
}

// ContextWithActor adds the user to a context.
func ContextWithActor(ctx context.Context, u *ent.Actor) context.Context {
	return context.WithValue(ctx, ActorCtxKey, u)
}

// TargetActorFromContext returns a Actor from a context or nil if not present.
func TargetActorFromContext(reqCtx context.Context) *ent.Actor {
	tu, ok := reqCtx.Value(TargetActorCtxKey).(*ent.Actor)
	if !ok {
		return nil
	}
	return tu
}

// ContextWithTargetActor adds the user to a context.
func ContextWithTargetActor(ctx context.Context, u *ent.Actor) context.Context {
	return context.WithValue(ctx, ActorCtxKey, u)
}
