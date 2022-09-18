package schema

import (
	"context"

	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/eriner/burr/internal/ent/privacy"
)

type Reaction struct {
	ent.Schema
}

func (Reaction) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("actor_id", "status_id"),
	}
}

func (Reaction) Fields() []ent.Field {
	return []ent.Field{
		IDField(),
		field.Uint64("actor_id"),
		field.Uint64("status_id"),
		field.Enum("type").
			Comment("Reactions can be likes, boosts, bookmarks, or other 'static' responses (not quote posts)").
			Values(
				"like",
				"boost",
				"emote",
			),
		field.Uint64("dat").
			Comment("Dat is an extra argument field used by some types of rections (emotes)"),
	}
}

func (Reaction) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("actors", Actor.Type).
			Unique().
			Required().
			Field("actor_id"),
		edge.To("status", Status.Type).
			Unique().
			Required().
			Field("status_id"),
	}
}

func (Reaction) Mixin() []ent.Mixin {
	return []ent.Mixin{
		AuditMixin{},
	}
}

// TODO: enforce policy
func (Reaction) Policy() ent.Policy {
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
