package schema

import (
	"bytes"
	"context"
	"fmt"
	"net/mail"
	"net/url"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/eriner/burr/internal/ent/hook"
	"github.com/eriner/burr/internal/ent/privacy"
	"github.com/sony/sonyflake"
)

// projectCreationDate is used as a start time for sonyflake IDs.
var projectCreationDate, _ = time.Parse("Jan 2, 2006 at 3:04pm (MST)", "Sept 14, 2022 at 10:32am (EDT)")

type Account struct {
	ent.Schema
}

func (Account) Fields() []ent.Field {
	return []ent.Field{
		// "Sonyflake currently does not use the most siginifican bits of IDs, so you can convert
		// Sonyflake IDs from uint64 to int64 safely"
		// NOTE: it may be wise to only store an int64 and not a uint64, but for reasons of not wanting
		// to have to make schema changes later, it's staying a uint64 for now. YOLO.
		field.Uint64("id").
			Comment("The Account's ID").
			Immutable().
			Unique(),

		// NOTE: the created_at and updated_at fields are automatically created by the AuditMixin

		field.String("username").
			Comment("The Account's username").
			Immutable().
			MaxLen(64). // TODO: check spec
			NotEmpty(),

		field.String("domain").
			Comment("The Account's domain").
			Immutable().
			MaxLen(255).
			NotEmpty(),

		field.String("display_name").
			Comment("Account full name").
			MaxLen(64).
			NotEmpty().
			Default("unknown"),

		field.String("email").
			Comment("Account email, unique per table").
			Sensitive().
			Unique().
			MaxLen(64). // TODO: maybe 128?
			Validate(func(s string) error {
				if _, err := mail.ParseAddress(s); err != nil {
					return fmt.Errorf("invalid email: %w", err)
				}
				return nil
			}),

		field.Bytes("passwordHash").
			Comment("Account bcrypt password hash").
			Sensitive().
			MaxLen(60). // All hashes have a len of 60. MinLen not set (for non-local accounts)
			Validate(func(b []byte) error {
				if !bytes.HasPrefix(b, []byte("$2a$")) {
					return fmt.Errorf("invalid bcrypt password hash")
				}
				return nil
			}).
			Nillable(),

		field.String("note").
			Comment("Account note, AKA description").
			Default(""),

		field.Bool("locked").
			Comment("Account account is locked if unconfirmed or explicitly locked").
			Default(false),

		field.Bool("memorial").
			Comment("Account account is memorialized").
			Default(false),

		field.String("url").
			Comment("Account URL").
			NotEmpty().
			MaxLen(255).
			Validate(func(s string) error {
				_, err := url.Parse(s)
				return err
			}),

		field.Bytes("pubkey").
			Comment("Account public key").
			NotEmpty().
			Default([]byte{}),

		field.Bytes("privkey").
			Comment("Account private key").
			Sensitive().
			Nillable(),

		field.String("avatar_remote_url").
			Comment("URL of the Account's remote avatar").
			MaxLen(512).
			Validate(func(s string) error {
				_, err := url.Parse(s)
				return err
			}).
			Nillable(),

		field.String("avatar_local_file").
			Comment("The Account's local avatar file").
			MaxLen(512).
			Nillable(),

		field.Time("avatar_updated_at").
			Comment("The time the Account's (local) avatar was last updated").
			Nillable(),

		field.String("header_url").
			Comment("URL of user header").
			MaxLen(512).
			Validate(func(s string) error {
				_, err := url.Parse(s)
				return err
			}),

		field.String("header_local_file").
			Comment("The Account's local header file").
			MaxLen(512).
			Nillable(),

		field.Time("header_updated_at").
			Comment("The time the Account's header was last updated").
			Nillable(),

		field.Time("last_webfinger_at").
			Comment("The time the Account's account was last discovered with webfinger").
			Default(time.Now).
			Nillable(),

		field.String("inbox_url").
			Comment("The Account's ActivityPub Inbox IRI").
			MaxLen(512).
			NotEmpty().
			Validate(func(s string) error {
				_, err := url.Parse(s)
				return err
			}),

		field.String("outbox_url").
			Comment("The Account's ActivityPub Inbox IRI").
			MaxLen(512).
			NotEmpty().
			Validate(func(s string) error {
				_, err := url.Parse(s)
				return err
			}),

		field.String("shared_inbox_url").
			Comment("The Account's ActivityPub Inbox IRI").
			MaxLen(512).
			NotEmpty().
			Validate(func(s string) error {
				_, err := url.Parse(s)
				return err
			}),

		field.String("followers_url").
			Comment("The Account's ActivityPub Inbox IRI").
			MaxLen(512).
			NotEmpty().
			Validate(func(s string) error {
				_, err := url.Parse(s)
				return err
			}),

		field.Uint64("moved_to_id").
			Comment("The Account's id, if it has moved").
			Nillable(),

		field.String("featured_collection_url").
			Comment("The Account's featured items").
			Nillable(),

		field.Time("silenced_at").
			Comment("The time the Account was silenced").
			Nillable(),

		field.Time("suspended_at").
			Comment("The time the Account was silenced").
			Nillable(),

		field.Int64("role").
			Comment("Accounts can be assigned roles").
			Default(0).
			Nillable(),

		field.Int64("badges").
			Comment("Accounts can have badges").
			Default(0).
			Nillable(),
	}
}

func (Account) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").
			Unique(), // enforce globally unique ids
	}
}

func (Account) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("servers", Server.Type).
			Ref("accounts").
			Comment("Accounts belong to a Server rooted at a domain (FQDN)"),

		edge.To("sessions", Session.Type).
			Annotations(entsql.Annotation{
				// When an account is deleted, delete the sessions.
				OnDelete: entsql.Cascade,
			}).
			Comment("Sessions are owned by a single Account which itself can have many sessions"),

		edge.From("followers", Account.Type).
			Ref("accounts").
			Comment("Accounts can be followed by one or more accounts"),

		edge.From("following", Account.Type).
			Ref("accounts").
			Comment("Accounts can be follow one or more accounts"),

		edge.From("roles", Role.Type).
			Ref("accounts").
			Comment("Accounts can be assigned a role"),

		edge.From("groups", Group.Type).
			Ref("accounts").
			Comment("Accounts can belong to one or more groups"),

		// NOTE: Events should belong to groups, not accounts.
		// edge.From("events", Event.Type).
		//	Ref("accounts").
		//	Comment("Accounts can belong to one or more events"),
	}
}

func (Account) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(func() ent.Hook {
			sf := sonyflake.NewSonyflake(sonyflake.Settings{
				StartTime: projectCreationDate,
				// The documentation isn't super clear, but I don't think I need to
				// define these functions. It uses the lower 16 bits of the private
				// ip by default, and I think this is just so that unique instances
				// of the app don't have ID collisions. The default implementation
				// is probably fine.
				// MachineID: func() (uint16, error) {},
				// CheckMachineID: func(u uint16) bool {},

			})
			type setter interface {
				SetID(uint64)
			}
			return func(next ent.Mutator) ent.Mutator {
				return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
					setr, ok := m.(setter)
					if !ok {
						return nil, fmt.Errorf("unexpected mutation during ID set %T", m)
					}
					id, err := sf.NextID()
					if err != nil {
						return nil, fmt.Errorf("It's been 174 years since this project was created. What a glorious error: %w", err)
					}
					setr.SetID(id)
					return next.Mutate(ctx, m)
				})
			}
		}(), ent.OpCreate),
	}
}

func (Account) Mixin() []ent.Mixin {
	return []ent.Mixin{
		AuditMixin{},
	}
}

// TODO: enforce policy
func (Account) Policy() ent.Policy {
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
