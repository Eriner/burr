package schema

import (
	"bytes"
	"context"
	"fmt"
	"net/url"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/eriner/burr/internal/ent/hook"
	"github.com/eriner/burr/internal/ent/privacy"
	"golang.org/x/text/language"
)

const (
	urlMaxLen  = 255
	nameMaxLen = 64
	noteMaxLen = 2048
)

// projectCreationDate is used as a start time for sonyflake IDs.
var projectCreationDate, _ = time.Parse("Jan 2, 2006 at 3:04pm (MST)", "Sept 14, 2022 at 10:32am (EDT)")

type Actor struct {
	ent.Schema
}

func (Actor) Fields() []ent.Field {
	return []ent.Field{

		//
		// Required fields, universal between local, remote, and group Actor
		//

		// "Sonyflake currently does not use the most siginifican bits of IDs, so you can convert
		// Sonyflake IDs from uint64 to int64 safely"
		// NOTE: it may be wise to only store an int64 and not a uint64, but for reasons of not wanting
		// to have to make schema changes later, it's staying a uint64 for now. YOLO.
		IDField(),

		field.Enum("type").
			// It's probably worth the column to cost to store this explicitly, may turn some JOINs into SELECTs
			Comment("The type of account. All accounts are AP actors").
			Values(
				"local_user",
				"local_bot",
				"local_group",
				"local_event",
				"remote_user",
				"remote_bot",
				"remote_group",
				"remote_event",
			),

		// NOTE: the created_at and updated_at fields are automatically created by the AuditMixin

		field.String("name").
			Comment("The Actor's username").
			Immutable().
			MaxLen(nameMaxLen). // TODO: check spec
			NotEmpty(),

		field.String("display_name").
			Comment("The Actor's displayed 'friendly' name").
			MaxLen(nameMaxLen).
			NotEmpty().
			Default("unknown"),

		// read from edge. Move to Servers schema
		/*
			field.String("domain").
				Comment("The Actor's domain").
				Immutable().
				MaxLen(255).
				NotEmpty(),
		*/

		field.String("note").
			Comment("Actor note, AKA description").
			MaxLen(noteMaxLen).
			Default(""),

		field.Bool("locked").
			Comment("Actor account is locked if unconfirmed or explicitly locked").
			Default(false),

		field.Bool("memorial").
			Comment("Actor account is memorialized").
			Default(false),

		field.String("url").
			Comment("Actor URL").
			NotEmpty().
			MaxLen(urlMaxLen).
			Validate(func(s string) error {
				_, err := url.Parse(s)
				return err
			}),

		field.Bytes("pubkey").
			Comment("Actor public key").
			NotEmpty().
			Default([]byte{}),

		field.Bytes("privkey").
			Comment("Actor private key").
			Optional().
			Sensitive().
			Nillable(),

		field.String("avatar_remote_url").
			Comment("URL of the Actor's remote avatar").
			MaxLen(urlMaxLen).
			Validate(func(s string) error {
				_, err := url.Parse(s)
				return err
			}).
			Optional().
			Nillable(),

		field.String("avatar_local_file").
			Comment("The Actor's local avatar file").
			MaxLen(urlMaxLen).
			Optional().
			Nillable(),

		field.Time("avatar_updated_at").
			Comment("The time the Actor's (local) avatar was last updated").
			Optional().
			Nillable(),

		field.String("header_url").
			Comment("URL of user header").
			MaxLen(urlMaxLen).
			Nillable().
			Optional().
			Validate(func(s string) error {
				_, err := url.Parse(s)
				return err
			}),

		field.String("header_local_file").
			Comment("The Actor's local header file").
			MaxLen(urlMaxLen).
			Optional().
			Nillable(),

		field.Time("header_updated_at").
			Comment("The time the Actor's header was last updated").
			Optional().
			Nillable(),

		field.Time("last_webfinger_at").
			Comment("The time the Actor's account was last discovered with webfinger").
			Default(time.Now).
			Optional().
			Nillable(),

		field.String("inbox_url").
			Comment("The Actor's ActivityPub Inbox IRI").
			MaxLen(urlMaxLen).
			NotEmpty().
			Validate(func(s string) error {
				_, err := url.Parse(s)
				return err
			}),

		field.String("outbox_url").
			Comment("The Actor's ActivityPub Inbox IRI").
			MaxLen(urlMaxLen).
			NotEmpty().
			Validate(func(s string) error {
				_, err := url.Parse(s)
				return err
			}),

		field.String("shared_inbox_url").
			Comment("The Actor's ActivityPub Inbox IRI").
			MaxLen(urlMaxLen).
			NotEmpty().
			Validate(func(s string) error {
				_, err := url.Parse(s)
				return err
			}),

		field.String("followers_url").
			Comment("The Actor's ActivityPub Inbox IRI").
			MaxLen(urlMaxLen).
			NotEmpty().
			Validate(func(s string) error {
				_, err := url.Parse(s)
				return err
			}),

		field.Uint64("moved_to_id").
			Comment("The Actor's id, if it has moved").
			Optional().
			Nillable(),

		field.String("featured_collection_url").
			Comment("The Actor's featured items").
			Optional().
			Nillable(),

		field.Time("silenced_at").
			Comment("The time the Actor was silenced").
			Optional().
			Nillable(),

		field.Time("suspended_at").
			Comment("The time the Actor was silenced").
			Optional().
			Nillable(),

		//
		// Fields present in local account types but not remote accounts
		//

		field.Bytes("passwordHash").
			Comment("Actor bcrypt password hash").
			Sensitive().
			Nillable().
			Optional(). // only local accounts have hashes. empty hashes will fail bcrypt.ConstantTimeCompare() regardless
			MaxLen(60). // All hashes have a len of 60. MinLen not set (for non-local accounts)
			Validate(func(b []byte) error {
				if !bytes.HasPrefix(b, []byte("$2a$")) {
					return fmt.Errorf("invalid bcrypt password hash")
				}
				return nil
			}),

		field.String("recovery_code").
			// BIP words?
			Comment("local Actor password recovery code generated during account creation").
			// TODO: specify len, validate
			Sensitive().
			Nillable().
			Optional(),

		field.Uint64("role").
			Comment("local Actor role").
			Default(0).
			Optional(),

		field.Uint64("badge").
			Comment("local Actor badge").
			Default(0).
			Optional(),

		field.Enum("locale").
			Comment("local Actor locale").
			Values(
				language.AmericanEnglish.String(),
				// TODO: additional languages at some point, probably.
			),
	}
}

func (Actor) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").
			Unique(), // enforce globally unique ids
	}
}

func (Actor) Edges() []ent.Edge {
	return []ent.Edge{
		//
		// Actor Belongs to
		//
		edge.From("server", Server.Type).
			Ref("actors").
			Unique().
			Required().
			Comment("Actor belong to a Server rooted at a domain (FQDN)"),

		edge.From("events", Event.Type).
			Ref("attendees").
			Comment("Actor can belong to one or more events"),

		//
		// Belongs to Actor
		//

		edge.To("organizer_of", Event.Type).
			Annotations(entsql.Annotation{
				// When an account is deleted, delete the events they organized
				OnDelete: entsql.Cascade,
			}).
			Comment("Users can organize events"),

		edge.To("statuses", Status.Type).
			Annotations(entsql.Annotation{
				// When an account is deleted, delete the statuses
				OnDelete: entsql.Cascade,
			}).
			Comment("Actor can have zero or many statuses"),

		edge.To("following", Actor.Type).
			From("followers").
			Comment("Actor can be followed by zero or more accounts"),

		edge.To("reacted_statuses", Status.Type).
			Through("reactions", Reaction.Type).
			Comment("Actor can react to many statuses"),

		edge.To("moderating", Actor.Type).
			From("moderators").
			Comment("Group Actor can have zero or more moderators"),

		edge.To("groups", Actor.Type).
			From("members"). // TODO is this legal with the groups edge above?
			Comment("User Actor can belong to groups"),

		//
		// Belongs exclusively to User-type Actor
		//

		edge.To("sessions", Session.Type).
			Annotations(entsql.Annotation{
				// When an account is deleted, delete the sessions.
				OnDelete: entsql.Cascade,
			}).
			Comment("Sessions are owned by a single User Actor which may have many sessions"),
	}
}

func (Actor) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(IDHook(), ent.OpCreate),
	}
}

func (Actor) Mixin() []ent.Mixin {
	return []ent.Mixin{
		AuditMixin{},
	}
}

// TODO: enforce policy
func (Actor) Policy() ent.Policy {
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
