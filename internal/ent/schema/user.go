package schema

import (
	"bytes"
	"fmt"
	"net/mail"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Comment("User full name").
			MaxLen(64).
			Default("unknown"),

		field.String("email").
			Comment("User email, unique per table").
			Sensitive().
			Unique().
			MaxLen(64).
			Validate(func(s string) error {
				if _, err := mail.ParseAddress(s); err != nil {
					return fmt.Errorf("invalid email: %w", err)
				}
				return nil
			}),

		field.Bytes("passwordHash").
			Comment("User bcrypt password hash").
			Sensitive().
			MinLen(60). // bcrypt hashes are always len(60)
			MaxLen(60).
			Validate(func(b []byte) error {
				if !bytes.HasPrefix(b, []byte("$2a$")) {
					return fmt.Errorf("invalid bcrypt password hash")
				}
				return nil
			}),
	}
}

func (User) Edges() []ent.Edge {
	return nil
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		AuditMixin{},
	}
}
