package schema

import (
	"context"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/eriner/burr/internal/ent/privacy"
)

type Like struct {
	ent.Schema
}

func (Like) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("account_id", "status_id"),
	}
}

func (Like) Fields() []ent.Field {
	return []ent.Field{
		field.Time("liked_at").
			Immutable().
			Default(time.Now),
		field.Uint64("account_id").
			Immutable(),
		field.Uint64("status_id").
			Immutable(),
	}
}

func (Like) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("accounts", Account.Type).
			Unique().
			Required().
			Field("account_id"),
		edge.To("status", Status.Type).
			Unique().
			Required().
			Field("status_id"),
	}
}

func (Like) Mixin() []ent.Mixin {
	return []ent.Mixin{
		AuditMixin{},
	}
}

// TODO: enforce policy
func (Like) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			privacy.ContextQueryMutationRule(func(ctx context.Context) error {
				return privacy.Allow
			}),
		},
		Query: privacy.QueryPolicy{
			privacy.ContextQueryMutationRule(func(ctx context.Context) error {
				return privacy.Allow
			}),
		},
	}
}
