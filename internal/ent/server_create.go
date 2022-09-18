// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/eriner/burr/internal/ent/actor"
	"github.com/eriner/burr/internal/ent/server"
)

// ServerCreate is the builder for creating a Server entity.
type ServerCreate struct {
	config
	mutation *ServerMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (sc *ServerCreate) SetCreatedAt(t time.Time) *ServerCreate {
	sc.mutation.SetCreatedAt(t)
	return sc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sc *ServerCreate) SetNillableCreatedAt(t *time.Time) *ServerCreate {
	if t != nil {
		sc.SetCreatedAt(*t)
	}
	return sc
}

// SetUpdatedAt sets the "updated_at" field.
func (sc *ServerCreate) SetUpdatedAt(t time.Time) *ServerCreate {
	sc.mutation.SetUpdatedAt(t)
	return sc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (sc *ServerCreate) SetNillableUpdatedAt(t *time.Time) *ServerCreate {
	if t != nil {
		sc.SetUpdatedAt(*t)
	}
	return sc
}

// SetCreatedBy sets the "created_by" field.
func (sc *ServerCreate) SetCreatedBy(i int) *ServerCreate {
	sc.mutation.SetCreatedBy(i)
	return sc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (sc *ServerCreate) SetNillableCreatedBy(i *int) *ServerCreate {
	if i != nil {
		sc.SetCreatedBy(*i)
	}
	return sc
}

// SetUpdatedBy sets the "updated_by" field.
func (sc *ServerCreate) SetUpdatedBy(i int) *ServerCreate {
	sc.mutation.SetUpdatedBy(i)
	return sc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (sc *ServerCreate) SetNillableUpdatedBy(i *int) *ServerCreate {
	if i != nil {
		sc.SetUpdatedBy(*i)
	}
	return sc
}

// SetDomain sets the "domain" field.
func (sc *ServerCreate) SetDomain(s string) *ServerCreate {
	sc.mutation.SetDomain(s)
	return sc
}

// SetLastSeen sets the "last_seen" field.
func (sc *ServerCreate) SetLastSeen(t time.Time) *ServerCreate {
	sc.mutation.SetLastSeen(t)
	return sc
}

// SetNillableLastSeen sets the "last_seen" field if the given value is not nil.
func (sc *ServerCreate) SetNillableLastSeen(t *time.Time) *ServerCreate {
	if t != nil {
		sc.SetLastSeen(*t)
	}
	return sc
}

// SetID sets the "id" field.
func (sc *ServerCreate) SetID(u uint64) *ServerCreate {
	sc.mutation.SetID(u)
	return sc
}

// AddActorIDs adds the "actors" edge to the Actor entity by IDs.
func (sc *ServerCreate) AddActorIDs(ids ...uint64) *ServerCreate {
	sc.mutation.AddActorIDs(ids...)
	return sc
}

// AddActors adds the "actors" edges to the Actor entity.
func (sc *ServerCreate) AddActors(a ...*Actor) *ServerCreate {
	ids := make([]uint64, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return sc.AddActorIDs(ids...)
}

// Mutation returns the ServerMutation object of the builder.
func (sc *ServerCreate) Mutation() *ServerMutation {
	return sc.mutation
}

// Save creates the Server in the database.
func (sc *ServerCreate) Save(ctx context.Context) (*Server, error) {
	var (
		err  error
		node *Server
	)
	if err := sc.defaults(); err != nil {
		return nil, err
	}
	if len(sc.hooks) == 0 {
		if err = sc.check(); err != nil {
			return nil, err
		}
		node, err = sc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ServerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sc.check(); err != nil {
				return nil, err
			}
			sc.mutation = mutation
			if node, err = sc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(sc.hooks) - 1; i >= 0; i-- {
			if sc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, sc.mutation)
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

// SaveX calls Save and panics if Save returns an error.
func (sc *ServerCreate) SaveX(ctx context.Context) *Server {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *ServerCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *ServerCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *ServerCreate) defaults() error {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		if server.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized server.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := server.DefaultCreatedAt()
		sc.mutation.SetCreatedAt(v)
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		if server.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized server.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := server.DefaultUpdatedAt()
		sc.mutation.SetUpdatedAt(v)
	}
	if _, ok := sc.mutation.LastSeen(); !ok {
		if server.DefaultLastSeen == nil {
			return fmt.Errorf("ent: uninitialized server.DefaultLastSeen (forgotten import ent/runtime?)")
		}
		v := server.DefaultLastSeen()
		sc.mutation.SetLastSeen(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (sc *ServerCreate) check() error {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Server.created_at"`)}
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Server.updated_at"`)}
	}
	if _, ok := sc.mutation.Domain(); !ok {
		return &ValidationError{Name: "domain", err: errors.New(`ent: missing required field "Server.domain"`)}
	}
	if v, ok := sc.mutation.Domain(); ok {
		if err := server.DomainValidator(v); err != nil {
			return &ValidationError{Name: "domain", err: fmt.Errorf(`ent: validator failed for field "Server.domain": %w`, err)}
		}
	}
	if _, ok := sc.mutation.LastSeen(); !ok {
		return &ValidationError{Name: "last_seen", err: errors.New(`ent: missing required field "Server.last_seen"`)}
	}
	return nil
}

func (sc *ServerCreate) sqlSave(ctx context.Context) (*Server, error) {
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	return _node, nil
}

func (sc *ServerCreate) createSpec() (*Server, *sqlgraph.CreateSpec) {
	var (
		_node = &Server{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: server.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: server.FieldID,
			},
		}
	)
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := sc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: server.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := sc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: server.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := sc.mutation.CreatedBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: server.FieldCreatedBy,
		})
		_node.CreatedBy = value
	}
	if value, ok := sc.mutation.UpdatedBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: server.FieldUpdatedBy,
		})
		_node.UpdatedBy = value
	}
	if value, ok := sc.mutation.Domain(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: server.FieldDomain,
		})
		_node.Domain = value
	}
	if value, ok := sc.mutation.LastSeen(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: server.FieldLastSeen,
		})
		_node.LastSeen = value
	}
	if nodes := sc.mutation.ActorsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   server.ActorsTable,
			Columns: []string{server.ActorsColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ServerCreateBulk is the builder for creating many Server entities in bulk.
type ServerCreateBulk struct {
	config
	builders []*ServerCreate
}

// Save creates the Server entities in the database.
func (scb *ServerCreateBulk) Save(ctx context.Context) ([]*Server, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Server, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ServerMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *ServerCreateBulk) SaveX(ctx context.Context) []*Server {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *ServerCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *ServerCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
