// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/eriner/burr/internal/ent/actor"
	"github.com/eriner/burr/internal/ent/reaction"
	"github.com/eriner/burr/internal/ent/status"
)

// Reaction is the model entity for the Reaction schema.
type Reaction struct {
	config `json:"-"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// CreatedBy holds the value of the "created_by" field.
	CreatedBy int `json:"created_by,omitempty"`
	// UpdatedBy holds the value of the "updated_by" field.
	UpdatedBy int `json:"updated_by,omitempty"`
	// ActorID holds the value of the "actor_id" field.
	ActorID uint64 `json:"actor_id,omitempty"`
	// StatusID holds the value of the "status_id" field.
	StatusID uint64 `json:"status_id,omitempty"`
	// Reactions can be likes, boosts, bookmarks, or other 'static' responses (not quote posts)
	Type reaction.Type `json:"type,omitempty"`
	// Dat is an extra argument field used by some types of rections (emotes)
	Dat uint64 `json:"dat,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ReactionQuery when eager-loading is set.
	Edges ReactionEdges `json:"edges"`
}

// ReactionEdges holds the relations/edges for other nodes in the graph.
type ReactionEdges struct {
	// Actors holds the value of the actors edge.
	Actors *Actor `json:"actors,omitempty"`
	// Status holds the value of the status edge.
	Status *Status `json:"status,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ActorsOrErr returns the Actors value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ReactionEdges) ActorsOrErr() (*Actor, error) {
	if e.loadedTypes[0] {
		if e.Actors == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: actor.Label}
		}
		return e.Actors, nil
	}
	return nil, &NotLoadedError{edge: "actors"}
}

// StatusOrErr returns the Status value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ReactionEdges) StatusOrErr() (*Status, error) {
	if e.loadedTypes[1] {
		if e.Status == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: status.Label}
		}
		return e.Status, nil
	}
	return nil, &NotLoadedError{edge: "status"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Reaction) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case reaction.FieldCreatedBy, reaction.FieldUpdatedBy, reaction.FieldActorID, reaction.FieldStatusID, reaction.FieldDat:
			values[i] = new(sql.NullInt64)
		case reaction.FieldType:
			values[i] = new(sql.NullString)
		case reaction.FieldCreatedAt, reaction.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Reaction", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Reaction fields.
func (r *Reaction) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case reaction.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				r.CreatedAt = value.Time
			}
		case reaction.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				r.UpdatedAt = value.Time
			}
		case reaction.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				r.CreatedBy = int(value.Int64)
			}
		case reaction.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				r.UpdatedBy = int(value.Int64)
			}
		case reaction.FieldActorID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field actor_id", values[i])
			} else if value.Valid {
				r.ActorID = uint64(value.Int64)
			}
		case reaction.FieldStatusID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status_id", values[i])
			} else if value.Valid {
				r.StatusID = uint64(value.Int64)
			}
		case reaction.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				r.Type = reaction.Type(value.String)
			}
		case reaction.FieldDat:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field dat", values[i])
			} else if value.Valid {
				r.Dat = uint64(value.Int64)
			}
		}
	}
	return nil
}

// QueryActors queries the "actors" edge of the Reaction entity.
func (r *Reaction) QueryActors() *ActorQuery {
	return (&ReactionClient{config: r.config}).QueryActors(r)
}

// QueryStatus queries the "status" edge of the Reaction entity.
func (r *Reaction) QueryStatus() *StatusQuery {
	return (&ReactionClient{config: r.config}).QueryStatus(r)
}

// Update returns a builder for updating this Reaction.
// Note that you need to call Reaction.Unwrap() before calling this method if this Reaction
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Reaction) Update() *ReactionUpdateOne {
	return (&ReactionClient{config: r.config}).UpdateOne(r)
}

// Unwrap unwraps the Reaction entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Reaction) Unwrap() *Reaction {
	_tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Reaction is not a transactional entity")
	}
	r.config.driver = _tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Reaction) String() string {
	var builder strings.Builder
	builder.WriteString("Reaction(")
	builder.WriteString("created_at=")
	builder.WriteString(r.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(r.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", r.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", r.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("actor_id=")
	builder.WriteString(fmt.Sprintf("%v", r.ActorID))
	builder.WriteString(", ")
	builder.WriteString("status_id=")
	builder.WriteString(fmt.Sprintf("%v", r.StatusID))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", r.Type))
	builder.WriteString(", ")
	builder.WriteString("dat=")
	builder.WriteString(fmt.Sprintf("%v", r.Dat))
	builder.WriteByte(')')
	return builder.String()
}

// Reactions is a parsable slice of Reaction.
type Reactions []*Reaction

func (r Reactions) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}
