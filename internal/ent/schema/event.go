package schema

import (
	"context"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/eriner/burr/internal/ent/hook"
	"github.com/eriner/burr/internal/ent/privacy"
)

type Event struct {
	ent.Schema
}

func (Event) Fields() []ent.Field {
	return []ent.Field{
		IDField(),

		field.String("display_name").
			Comment("The event's 'friendly' name").
			MaxLen(nameMaxLen),

		// NOTE: the created_at and updated_at fields are automatically created by the AuditMixin

		// TODO
	}
}

func (Event) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").
			Unique(), // enforce globally unique ids
	}
}

func (Event) Edges() []ent.Edge {
	return []ent.Edge{
		//
		// Belongs to Event
		//
		edge.To("attendees", Actor.Type).
			Comment("Accounts which have RSVP'd to the event"),

		//
		// Event Belongs to
		//
		edge.From("organizer", Actor.Type).
			Ref("organizer_of").
			Unique().
			Required().
			Comment("Events belong to an Account who organizes the event"),
	}
}

func (Event) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(IDHook(), ent.OpCreate),
	}
}

func (Event) Mixin() []ent.Mixin {
	return []ent.Mixin{
		AuditMixin{},
	}
}

// TODO: enforce policy
func (Event) Policy() ent.Policy {
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
