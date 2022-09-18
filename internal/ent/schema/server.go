package schema

import (
	"context"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/eriner/burr/internal/ent/hook"
	"github.com/eriner/burr/internal/ent/privacy"
)

type Server struct {
	ent.Schema
}

func (Server) Fields() []ent.Field {
	return []ent.Field{
		IDField(),

		// NOTE: the created_at and updated_at fields are automatically created by the AuditMixin

		field.String("domain").
			Comment("The ActivityPub server's domain").
			Immutable().
			Unique().
			MaxLen(urlMaxLen),

		field.Time("last_seen").
			Comment("The last time a valid response or message was recieved from this Server").
			Default(time.Now),
	}
}

func (Server) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id", "domain").
			Unique(),
	}
}

func (Server) Edges() []ent.Edge {
	return []ent.Edge{
		//
		// Belongs to Server
		//
		edge.To("actors", Actor.Type).
			Comment("Actors belong to a server"),
	}
}

func (Server) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(IDHook(), ent.OpCreate),
	}
}

func (Server) Mixin() []ent.Mixin {
	return []ent.Mixin{
		AuditMixin{},
	}
}

// TODO: enforce policy
func (Server) Policy() ent.Policy {
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
