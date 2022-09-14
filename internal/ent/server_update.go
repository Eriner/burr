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
	"github.com/eriner/burr/internal/ent/predicate"
	"github.com/eriner/burr/internal/ent/server"
)

// ServerUpdate is the builder for updating Server entities.
type ServerUpdate struct {
	config
	hooks    []Hook
	mutation *ServerMutation
}

// Where appends a list predicates to the ServerUpdate builder.
func (su *ServerUpdate) Where(ps ...predicate.Server) *ServerUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetUpdatedAt sets the "updated_at" field.
func (su *ServerUpdate) SetUpdatedAt(t time.Time) *ServerUpdate {
	su.mutation.SetUpdatedAt(t)
	return su
}

// SetCreatedBy sets the "created_by" field.
func (su *ServerUpdate) SetCreatedBy(i int) *ServerUpdate {
	su.mutation.ResetCreatedBy()
	su.mutation.SetCreatedBy(i)
	return su
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (su *ServerUpdate) SetNillableCreatedBy(i *int) *ServerUpdate {
	if i != nil {
		su.SetCreatedBy(*i)
	}
	return su
}

// AddCreatedBy adds i to the "created_by" field.
func (su *ServerUpdate) AddCreatedBy(i int) *ServerUpdate {
	su.mutation.AddCreatedBy(i)
	return su
}

// ClearCreatedBy clears the value of the "created_by" field.
func (su *ServerUpdate) ClearCreatedBy() *ServerUpdate {
	su.mutation.ClearCreatedBy()
	return su
}

// SetUpdatedBy sets the "updated_by" field.
func (su *ServerUpdate) SetUpdatedBy(i int) *ServerUpdate {
	su.mutation.ResetUpdatedBy()
	su.mutation.SetUpdatedBy(i)
	return su
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (su *ServerUpdate) SetNillableUpdatedBy(i *int) *ServerUpdate {
	if i != nil {
		su.SetUpdatedBy(*i)
	}
	return su
}

// AddUpdatedBy adds i to the "updated_by" field.
func (su *ServerUpdate) AddUpdatedBy(i int) *ServerUpdate {
	su.mutation.AddUpdatedBy(i)
	return su
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (su *ServerUpdate) ClearUpdatedBy() *ServerUpdate {
	su.mutation.ClearUpdatedBy()
	return su
}

// Mutation returns the ServerMutation object of the builder.
func (su *ServerUpdate) Mutation() *ServerMutation {
	return su.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *ServerUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := su.defaults(); err != nil {
		return 0, err
	}
	if len(su.hooks) == 0 {
		affected, err = su.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ServerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			su.mutation = mutation
			affected, err = su.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(su.hooks) - 1; i >= 0; i-- {
			if su.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = su.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, su.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (su *ServerUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *ServerUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *ServerUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (su *ServerUpdate) defaults() error {
	if _, ok := su.mutation.UpdatedAt(); !ok {
		if server.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized server.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := server.UpdateDefaultUpdatedAt()
		su.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (su *ServerUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   server.Table,
			Columns: server.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: server.FieldID,
			},
		},
	}
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: server.FieldUpdatedAt,
		})
	}
	if value, ok := su.mutation.CreatedBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: server.FieldCreatedBy,
		})
	}
	if value, ok := su.mutation.AddedCreatedBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: server.FieldCreatedBy,
		})
	}
	if su.mutation.CreatedByCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: server.FieldCreatedBy,
		})
	}
	if value, ok := su.mutation.UpdatedBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: server.FieldUpdatedBy,
		})
	}
	if value, ok := su.mutation.AddedUpdatedBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: server.FieldUpdatedBy,
		})
	}
	if su.mutation.UpdatedByCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: server.FieldUpdatedBy,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{server.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// ServerUpdateOne is the builder for updating a single Server entity.
type ServerUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ServerMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (suo *ServerUpdateOne) SetUpdatedAt(t time.Time) *ServerUpdateOne {
	suo.mutation.SetUpdatedAt(t)
	return suo
}

// SetCreatedBy sets the "created_by" field.
func (suo *ServerUpdateOne) SetCreatedBy(i int) *ServerUpdateOne {
	suo.mutation.ResetCreatedBy()
	suo.mutation.SetCreatedBy(i)
	return suo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (suo *ServerUpdateOne) SetNillableCreatedBy(i *int) *ServerUpdateOne {
	if i != nil {
		suo.SetCreatedBy(*i)
	}
	return suo
}

// AddCreatedBy adds i to the "created_by" field.
func (suo *ServerUpdateOne) AddCreatedBy(i int) *ServerUpdateOne {
	suo.mutation.AddCreatedBy(i)
	return suo
}

// ClearCreatedBy clears the value of the "created_by" field.
func (suo *ServerUpdateOne) ClearCreatedBy() *ServerUpdateOne {
	suo.mutation.ClearCreatedBy()
	return suo
}

// SetUpdatedBy sets the "updated_by" field.
func (suo *ServerUpdateOne) SetUpdatedBy(i int) *ServerUpdateOne {
	suo.mutation.ResetUpdatedBy()
	suo.mutation.SetUpdatedBy(i)
	return suo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (suo *ServerUpdateOne) SetNillableUpdatedBy(i *int) *ServerUpdateOne {
	if i != nil {
		suo.SetUpdatedBy(*i)
	}
	return suo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (suo *ServerUpdateOne) AddUpdatedBy(i int) *ServerUpdateOne {
	suo.mutation.AddUpdatedBy(i)
	return suo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (suo *ServerUpdateOne) ClearUpdatedBy() *ServerUpdateOne {
	suo.mutation.ClearUpdatedBy()
	return suo
}

// Mutation returns the ServerMutation object of the builder.
func (suo *ServerUpdateOne) Mutation() *ServerMutation {
	return suo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *ServerUpdateOne) Select(field string, fields ...string) *ServerUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Server entity.
func (suo *ServerUpdateOne) Save(ctx context.Context) (*Server, error) {
	var (
		err  error
		node *Server
	)
	if err := suo.defaults(); err != nil {
		return nil, err
	}
	if len(suo.hooks) == 0 {
		node, err = suo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ServerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			suo.mutation = mutation
			node, err = suo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(suo.hooks) - 1; i >= 0; i-- {
			if suo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = suo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, suo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Server)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ServerMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (suo *ServerUpdateOne) SaveX(ctx context.Context) *Server {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *ServerUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *ServerUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (suo *ServerUpdateOne) defaults() error {
	if _, ok := suo.mutation.UpdatedAt(); !ok {
		if server.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized server.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := server.UpdateDefaultUpdatedAt()
		suo.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (suo *ServerUpdateOne) sqlSave(ctx context.Context) (_node *Server, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   server.Table,
			Columns: server.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: server.FieldID,
			},
		},
	}
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Server.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, server.FieldID)
		for _, f := range fields {
			if !server.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != server.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: server.FieldUpdatedAt,
		})
	}
	if value, ok := suo.mutation.CreatedBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: server.FieldCreatedBy,
		})
	}
	if value, ok := suo.mutation.AddedCreatedBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: server.FieldCreatedBy,
		})
	}
	if suo.mutation.CreatedByCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: server.FieldCreatedBy,
		})
	}
	if value, ok := suo.mutation.UpdatedBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: server.FieldUpdatedBy,
		})
	}
	if value, ok := suo.mutation.AddedUpdatedBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: server.FieldUpdatedBy,
		})
	}
	if suo.mutation.UpdatedByCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: server.FieldUpdatedBy,
		})
	}
	_node = &Server{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{server.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
