package schema

import (
	"context"
	"encoding/base64"
	"fmt"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/eriner/burr/internal/ent/hook"
	"github.com/eriner/burr/internal/ent/privacy"
	"lukechampine.com/frand"
)

// Sessions are authentication sessions. They can either be first-party web auth sessions or OAuth sessions.
// Sessions should persist in the database for 1yr after expiry, but with the "disabled" boolean set to true.
type Session struct {
	ent.Schema
}

func (Session) Fields() []ent.Field {
	return []ent.Field{
		IDField(),

		field.Enum("type").
			Comment("Sessions can derrive from the local (password auth), oauth, or app_password").
			Values(
				"local",
				"oauth",
				"app_password",
			).
			Immutable(),

		field.Bool("disabled").
			Comment("The session may be disabled by the user or by automatic security policy"),

		// NOTE: the created_at and updated_at fields are automatically created by the AuditMixin
		// Session expiry can be determined by the application at runtime based on the created_at field.

		field.String("token").
			Comment("random 32 bytes encoded as base64").
			Unique().
			Immutable().
			DefaultFunc(func() string {
				b := make([]byte, 20)
				_, _ = frand.Read(b)
				out := make([]byte, 27)
				base64.RawStdEncoding.Encode(out, b)
				return string(out)
			}).
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
		field.String("user_agent").
			Comment("The last known user-agent").
			Optional(),

		field.String("ips").
			Comment("All IPs that have been associated with this session. Reverse-chornological order. The current IP is the first item in the slice"),
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
		edge.To("accounts", Actor.Type).
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
