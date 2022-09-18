package schema

import (
	"context"
	"fmt"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/sony/sonyflake"
)

func IDField() ent.Field {
	return field.Uint64("id").
		Immutable().
		Unique()
}

func IDHook() ent.Hook {
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
}
