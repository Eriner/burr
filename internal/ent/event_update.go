// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/eriner/burr/internal/ent/actor"
	"github.com/eriner/burr/internal/ent/event"
	"github.com/eriner/burr/internal/ent/predicate"
)

// EventUpdate is the builder for updating Event entities.
type EventUpdate struct {
	config
	hooks    []Hook
	mutation *EventMutation
}

// Where appends a list predicates to the EventUpdate builder.
func (eu *EventUpdate) Where(ps ...predicate.Event) *EventUpdate {
	eu.mutation.Where(ps...)
	return eu
}

// SetUpdatedAt sets the "updated_at" field.
func (eu *EventUpdate) SetUpdatedAt(t time.Time) *EventUpdate {
	eu.mutation.SetUpdatedAt(t)
	return eu
}

// SetCreatedBy sets the "created_by" field.
func (eu *EventUpdate) SetCreatedBy(i int) *EventUpdate {
	eu.mutation.ResetCreatedBy()
	eu.mutation.SetCreatedBy(i)
	return eu
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (eu *EventUpdate) SetNillableCreatedBy(i *int) *EventUpdate {
	if i != nil {
		eu.SetCreatedBy(*i)
	}
	return eu
}

// AddCreatedBy adds i to the "created_by" field.
func (eu *EventUpdate) AddCreatedBy(i int) *EventUpdate {
	eu.mutation.AddCreatedBy(i)
	return eu
}

// ClearCreatedBy clears the value of the "created_by" field.
func (eu *EventUpdate) ClearCreatedBy() *EventUpdate {
	eu.mutation.ClearCreatedBy()
	return eu
}

// SetUpdatedBy sets the "updated_by" field.
func (eu *EventUpdate) SetUpdatedBy(i int) *EventUpdate {
	eu.mutation.ResetUpdatedBy()
	eu.mutation.SetUpdatedBy(i)
	return eu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (eu *EventUpdate) SetNillableUpdatedBy(i *int) *EventUpdate {
	if i != nil {
		eu.SetUpdatedBy(*i)
	}
	return eu
}

// AddUpdatedBy adds i to the "updated_by" field.
func (eu *EventUpdate) AddUpdatedBy(i int) *EventUpdate {
	eu.mutation.AddUpdatedBy(i)
	return eu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (eu *EventUpdate) ClearUpdatedBy() *EventUpdate {
	eu.mutation.ClearUpdatedBy()
	return eu
}

// SetDisplayName sets the "display_name" field.
func (eu *EventUpdate) SetDisplayName(s string) *EventUpdate {
	eu.mutation.SetDisplayName(s)
	return eu
}

// AddAttendeeIDs adds the "attendees" edge to the Actor entity by IDs.
func (eu *EventUpdate) AddAttendeeIDs(ids ...uint64) *EventUpdate {
	eu.mutation.AddAttendeeIDs(ids...)
	return eu
}

// AddAttendees adds the "attendees" edges to the Actor entity.
func (eu *EventUpdate) AddAttendees(a ...*Actor) *EventUpdate {
	ids := make([]uint64, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return eu.AddAttendeeIDs(ids...)
}

// SetOrganizerID sets the "organizer" edge to the Actor entity by ID.
func (eu *EventUpdate) SetOrganizerID(id uint64) *EventUpdate {
	eu.mutation.SetOrganizerID(id)
	return eu
}

// SetOrganizer sets the "organizer" edge to the Actor entity.
func (eu *EventUpdate) SetOrganizer(a *Actor) *EventUpdate {
	return eu.SetOrganizerID(a.ID)
}

// Mutation returns the EventMutation object of the builder.
func (eu *EventUpdate) Mutation() *EventMutation {
	return eu.mutation
}

// ClearAttendees clears all "attendees" edges to the Actor entity.
func (eu *EventUpdate) ClearAttendees() *EventUpdate {
	eu.mutation.ClearAttendees()
	return eu
}

// RemoveAttendeeIDs removes the "attendees" edge to Actor entities by IDs.
func (eu *EventUpdate) RemoveAttendeeIDs(ids ...uint64) *EventUpdate {
	eu.mutation.RemoveAttendeeIDs(ids...)
	return eu
}

// RemoveAttendees removes "attendees" edges to Actor entities.
func (eu *EventUpdate) RemoveAttendees(a ...*Actor) *EventUpdate {
	ids := make([]uint64, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return eu.RemoveAttendeeIDs(ids...)
}

// ClearOrganizer clears the "organizer" edge to the Actor entity.
func (eu *EventUpdate) ClearOrganizer() *EventUpdate {
	eu.mutation.ClearOrganizer()
	return eu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (eu *EventUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := eu.defaults(); err != nil {
		return 0, err
	}
	if len(eu.hooks) == 0 {
		if err = eu.check(); err != nil {
			return 0, err
		}
		affected, err = eu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EventMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = eu.check(); err != nil {
				return 0, err
			}
			eu.mutation = mutation
			affected, err = eu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(eu.hooks) - 1; i >= 0; i-- {
			if eu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = eu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, eu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (eu *EventUpdate) SaveX(ctx context.Context) int {
	affected, err := eu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eu *EventUpdate) Exec(ctx context.Context) error {
	_, err := eu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eu *EventUpdate) ExecX(ctx context.Context) {
	if err := eu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (eu *EventUpdate) defaults() error {
	if _, ok := eu.mutation.UpdatedAt(); !ok {
		if event.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized event.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := event.UpdateDefaultUpdatedAt()
		eu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (eu *EventUpdate) check() error {
	if v, ok := eu.mutation.DisplayName(); ok {
		if err := event.DisplayNameValidator(v); err != nil {
			return &ValidationError{Name: "display_name", err: fmt.Errorf(`ent: validator failed for field "Event.display_name": %w`, err)}
		}
	}
	if _, ok := eu.mutation.OrganizerID(); eu.mutation.OrganizerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Event.organizer"`)
	}
	return nil
}

func (eu *EventUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   event.Table,
			Columns: event.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: event.FieldID,
			},
		},
	}
	if ps := eu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: event.FieldUpdatedAt,
		})
	}
	if value, ok := eu.mutation.CreatedBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: event.FieldCreatedBy,
		})
	}
	if value, ok := eu.mutation.AddedCreatedBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: event.FieldCreatedBy,
		})
	}
	if eu.mutation.CreatedByCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: event.FieldCreatedBy,
		})
	}
	if value, ok := eu.mutation.UpdatedBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: event.FieldUpdatedBy,
		})
	}
	if value, ok := eu.mutation.AddedUpdatedBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: event.FieldUpdatedBy,
		})
	}
	if eu.mutation.UpdatedByCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: event.FieldUpdatedBy,
		})
	}
	if value, ok := eu.mutation.DisplayName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldDisplayName,
		})
	}
	if eu.mutation.AttendeesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   event.AttendeesTable,
			Columns: event.AttendeesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: actor.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.RemovedAttendeesIDs(); len(nodes) > 0 && !eu.mutation.AttendeesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   event.AttendeesTable,
			Columns: event.AttendeesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: actor.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.AttendeesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   event.AttendeesTable,
			Columns: event.AttendeesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: actor.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if eu.mutation.OrganizerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   event.OrganizerTable,
			Columns: []string{event.OrganizerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: actor.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.OrganizerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   event.OrganizerTable,
			Columns: []string{event.OrganizerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: actor.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, eu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{event.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// EventUpdateOne is the builder for updating a single Event entity.
type EventUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EventMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (euo *EventUpdateOne) SetUpdatedAt(t time.Time) *EventUpdateOne {
	euo.mutation.SetUpdatedAt(t)
	return euo
}

// SetCreatedBy sets the "created_by" field.
func (euo *EventUpdateOne) SetCreatedBy(i int) *EventUpdateOne {
	euo.mutation.ResetCreatedBy()
	euo.mutation.SetCreatedBy(i)
	return euo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (euo *EventUpdateOne) SetNillableCreatedBy(i *int) *EventUpdateOne {
	if i != nil {
		euo.SetCreatedBy(*i)
	}
	return euo
}

// AddCreatedBy adds i to the "created_by" field.
func (euo *EventUpdateOne) AddCreatedBy(i int) *EventUpdateOne {
	euo.mutation.AddCreatedBy(i)
	return euo
}

// ClearCreatedBy clears the value of the "created_by" field.
func (euo *EventUpdateOne) ClearCreatedBy() *EventUpdateOne {
	euo.mutation.ClearCreatedBy()
	return euo
}

// SetUpdatedBy sets the "updated_by" field.
func (euo *EventUpdateOne) SetUpdatedBy(i int) *EventUpdateOne {
	euo.mutation.ResetUpdatedBy()
	euo.mutation.SetUpdatedBy(i)
	return euo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (euo *EventUpdateOne) SetNillableUpdatedBy(i *int) *EventUpdateOne {
	if i != nil {
		euo.SetUpdatedBy(*i)
	}
	return euo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (euo *EventUpdateOne) AddUpdatedBy(i int) *EventUpdateOne {
	euo.mutation.AddUpdatedBy(i)
	return euo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (euo *EventUpdateOne) ClearUpdatedBy() *EventUpdateOne {
	euo.mutation.ClearUpdatedBy()
	return euo
}

// SetDisplayName sets the "display_name" field.
func (euo *EventUpdateOne) SetDisplayName(s string) *EventUpdateOne {
	euo.mutation.SetDisplayName(s)
	return euo
}

// AddAttendeeIDs adds the "attendees" edge to the Actor entity by IDs.
func (euo *EventUpdateOne) AddAttendeeIDs(ids ...uint64) *EventUpdateOne {
	euo.mutation.AddAttendeeIDs(ids...)
	return euo
}

// AddAttendees adds the "attendees" edges to the Actor entity.
func (euo *EventUpdateOne) AddAttendees(a ...*Actor) *EventUpdateOne {
	ids := make([]uint64, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return euo.AddAttendeeIDs(ids...)
}

// SetOrganizerID sets the "organizer" edge to the Actor entity by ID.
func (euo *EventUpdateOne) SetOrganizerID(id uint64) *EventUpdateOne {
	euo.mutation.SetOrganizerID(id)
	return euo
}

// SetOrganizer sets the "organizer" edge to the Actor entity.
func (euo *EventUpdateOne) SetOrganizer(a *Actor) *EventUpdateOne {
	return euo.SetOrganizerID(a.ID)
}

// Mutation returns the EventMutation object of the builder.
func (euo *EventUpdateOne) Mutation() *EventMutation {
	return euo.mutation
}

// ClearAttendees clears all "attendees" edges to the Actor entity.
func (euo *EventUpdateOne) ClearAttendees() *EventUpdateOne {
	euo.mutation.ClearAttendees()
	return euo
}

// RemoveAttendeeIDs removes the "attendees" edge to Actor entities by IDs.
func (euo *EventUpdateOne) RemoveAttendeeIDs(ids ...uint64) *EventUpdateOne {
	euo.mutation.RemoveAttendeeIDs(ids...)
	return euo
}

// RemoveAttendees removes "attendees" edges to Actor entities.
func (euo *EventUpdateOne) RemoveAttendees(a ...*Actor) *EventUpdateOne {
	ids := make([]uint64, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return euo.RemoveAttendeeIDs(ids...)
}

// ClearOrganizer clears the "organizer" edge to the Actor entity.
func (euo *EventUpdateOne) ClearOrganizer() *EventUpdateOne {
	euo.mutation.ClearOrganizer()
	return euo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (euo *EventUpdateOne) Select(field string, fields ...string) *EventUpdateOne {
	euo.fields = append([]string{field}, fields...)
	return euo
}

// Save executes the query and returns the updated Event entity.
func (euo *EventUpdateOne) Save(ctx context.Context) (*Event, error) {
	var (
		err  error
		node *Event
	)
	if err := euo.defaults(); err != nil {
		return nil, err
	}
	if len(euo.hooks) == 0 {
		if err = euo.check(); err != nil {
			return nil, err
		}
		node, err = euo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EventMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = euo.check(); err != nil {
				return nil, err
			}
			euo.mutation = mutation
			node, err = euo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(euo.hooks) - 1; i >= 0; i-- {
			if euo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = euo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, euo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Event)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from EventMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (euo *EventUpdateOne) SaveX(ctx context.Context) *Event {
	node, err := euo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (euo *EventUpdateOne) Exec(ctx context.Context) error {
	_, err := euo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (euo *EventUpdateOne) ExecX(ctx context.Context) {
	if err := euo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (euo *EventUpdateOne) defaults() error {
	if _, ok := euo.mutation.UpdatedAt(); !ok {
		if event.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized event.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := event.UpdateDefaultUpdatedAt()
		euo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (euo *EventUpdateOne) check() error {
	if v, ok := euo.mutation.DisplayName(); ok {
		if err := event.DisplayNameValidator(v); err != nil {
			return &ValidationError{Name: "display_name", err: fmt.Errorf(`ent: validator failed for field "Event.display_name": %w`, err)}
		}
	}
	if _, ok := euo.mutation.OrganizerID(); euo.mutation.OrganizerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Event.organizer"`)
	}
	return nil
}

func (euo *EventUpdateOne) sqlSave(ctx context.Context) (_node *Event, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   event.Table,
			Columns: event.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: event.FieldID,
			},
		},
	}
	id, ok := euo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Event.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := euo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, event.FieldID)
		for _, f := range fields {
			if !event.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != event.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := euo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := euo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: event.FieldUpdatedAt,
		})
	}
	if value, ok := euo.mutation.CreatedBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: event.FieldCreatedBy,
		})
	}
	if value, ok := euo.mutation.AddedCreatedBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: event.FieldCreatedBy,
		})
	}
	if euo.mutation.CreatedByCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: event.FieldCreatedBy,
		})
	}
	if value, ok := euo.mutation.UpdatedBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: event.FieldUpdatedBy,
		})
	}
	if value, ok := euo.mutation.AddedUpdatedBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: event.FieldUpdatedBy,
		})
	}
	if euo.mutation.UpdatedByCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: event.FieldUpdatedBy,
		})
	}
	if value, ok := euo.mutation.DisplayName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldDisplayName,
		})
	}
	if euo.mutation.AttendeesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   event.AttendeesTable,
			Columns: event.AttendeesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: actor.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.RemovedAttendeesIDs(); len(nodes) > 0 && !euo.mutation.AttendeesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   event.AttendeesTable,
			Columns: event.AttendeesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: actor.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.AttendeesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   event.AttendeesTable,
			Columns: event.AttendeesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: actor.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if euo.mutation.OrganizerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   event.OrganizerTable,
			Columns: []string{event.OrganizerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: actor.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.OrganizerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   event.OrganizerTable,
			Columns: []string{event.OrganizerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: actor.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Event{config: euo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, euo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{event.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
