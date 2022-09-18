package schema

import (
	"context"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"github.com/eriner/burr/internal/ent/privacy"
)

// Groups are logical groups, distinct from Group-type Actors.
type Group struct {
	ent.Schema
}

/*
func (Group) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("actor_id", "owner_id"),
	}
}
*/

func (Group) Fields() []ent.Field {
	return []ent.Field{
		IDField(),
	}
}

func (Group) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("actors", Actor.Type).
			Unique().
			Required(),
		edge.To("owners", Actor.Type).
			Unique().
			Required(),
	}
}

func (Group) Mixin() []ent.Mixin {
	return []ent.Mixin{
		AuditMixin{},
	}
}

// TODO: enforce policy
func (Group) Policy() ent.Policy {
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
