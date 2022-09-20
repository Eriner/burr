package schema

import (
	"context"
	"encoding/base64"
	"fmt"
	"net"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/eriner/burr/internal/ent/hook"
	"github.com/eriner/burr/internal/ent/privacy"
)

// Sessions are authentication sessions. They can either be first-party web auth sessions or OAuth sessions.
type Session struct {
	ent.Schema
}

func (Session) Fields() []ent.Field {
	return []ent.Field{
		// "Sonyflake currently does not use the most siginifican bits of IDs, so you can convert
		// Sonyflake IDs from uint64 to int64 safely"
		// NOTE: it may be wise to only store an int64 and not a uint64, but for reasons of not wanting
		// to have to make schema changes later, it's staying a uint64 for now. YOLO.
		field.Uint64("id").
			Comment("The Session's ID").
			Immutable().
			Unique(),

		field.Bool("disabled").
			Comment("The session may be disabled by the user or by automatic security policy"),

		// NOTE: the created_at and updated_at fields are automatically created by the AuditMixin

		field.String("token").
			Comment("random 32 bytes encoded as base64").
			Unique().
			Immutable().
			Validate(func(s string) error {
				v, err := base64.RawStdEncoding.DecodeString(s)
				if err != nil {
					return err
				}
				if len(v) != 32 {
					return fmt.Errorf("invalid token size")
				}
				return nil
			}),

		// TODO: OAuth fields

		field.String("historical_ips").
			Comment("All IPs that have been associated with this session. The current IP is the first item in the slice").
			GoType([]net.IP{}),

		field.Time("expires_at").
			Comment("session expiry").
			Immutable(),
	}
}

func (Session) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").
			Unique(), // enforce globally unique ids
	}
}

func (Session) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("accounts", Account.Type).
			Unique().
			Comment("Sessions belong to Accounts"),
	}
}

func (Session) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(IDHook(), ent.OpCreate),
	}
}

func (Session) Mixin() []ent.Mixin {
	return []ent.Mixin{
		AuditMixin{},
	}
}

// TODO: enforce policy
func (Session) Policy() ent.Policy {
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
