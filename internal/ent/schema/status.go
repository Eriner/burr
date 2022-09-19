package schema

import (
	"context"
	"fmt"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/eriner/burr/internal/ent/hook"
	"github.com/eriner/burr/internal/ent/privacy"
	"github.com/sony/sonyflake"
)

type Status struct {
	ent.Schema
}

func (Status) Fields() []ent.Field {
	return []ent.Field{
		// "Sonyflake currently does not use the most siginifican bits of IDs, so you can convert
		// Sonyflake IDs from uint64 to int64 safely"
		// NOTE: it may be wise to only store an int64 and not a uint64, but for reasons of not wanting
		// to have to make schema changes later, it's staying a uint64 for now. YOLO.
		field.Uint64("id").
			Comment("The Status's ID").
			Immutable().
			Unique(),

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
		edge.From("accounts", Account.Type).
			Ref("statuses").
			Unique().
			Required().
			Comment("Statuses belong to an Account"),

		edge.From("liked_accounts", Account.Type).
			Ref("liked_statuses").
			Through("likes", Like.Type),
	}
}

func (Status) Hooks() []ent.Hook {
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
