package schema

import (
	"context"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/index"
	"github.com/eriner/burr/internal/ent/hook"
	"github.com/eriner/burr/internal/ent/privacy"
)

type Status struct {
	ent.Schema
}

func (Status) Fields() []ent.Field {
	return []ent.Field{
		IDField(),

		// NOTE: the created_at and updated_at fields are automatically created by the AuditMixin

		// TODO
	}
}

func (Status) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").
			Unique(), // enforce globally unique ids
	}
}

func (Status) Edges() []ent.Edge {
	return []ent.Edge{
		//
		// Status Belongs to
		//
		edge.From("actors", Actor.Type).
			Ref("statuses").
			Unique().
			Required().
			Comment("Statuses belong to an Account"),

		edge.From("reacted_actors", Actor.Type).
			Ref("reacted_statuses").
			Through("reactions", Reaction.Type),
	}
}

func (Status) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(IDHook(), ent.OpCreate),
	}
}

func (Status) Mixin() []ent.Mixin {
	return []ent.Mixin{
		AuditMixin{},
	}
}

// TODO: enforce policy
func (Status) Policy() ent.Policy {
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
