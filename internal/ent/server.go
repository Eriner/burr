// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/eriner/burr/internal/ent/server"
)

// Server is the model entity for the Server schema.
type Server struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// CreatedBy holds the value of the "created_by" field.
	CreatedBy int `json:"created_by,omitempty"`
	// UpdatedBy holds the value of the "updated_by" field.
	UpdatedBy int `json:"updated_by,omitempty"`
	// The ActivityPub server's domain
	Domain string `json:"domain,omitempty"`
	// The last time a valid response or message was recieved from this Server
	LastSeen time.Time `json:"last_seen,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ServerQuery when eager-loading is set.
	Edges ServerEdges `json:"edges"`
}

// ServerEdges holds the relations/edges for other nodes in the graph.
type ServerEdges struct {
	// Actors belong to a server
	Actors []*Actor `json:"actors,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ActorsOrErr returns the Actors value or an error if the edge
// was not loaded in eager-loading.
func (e ServerEdges) ActorsOrErr() ([]*Actor, error) {
	if e.loadedTypes[0] {
		return e.Actors, nil
	}
	return nil, &NotLoadedError{edge: "actors"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Server) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case server.FieldID, server.FieldCreatedBy, server.FieldUpdatedBy:
			values[i] = new(sql.NullInt64)
		case server.FieldDomain:
			values[i] = new(sql.NullString)
		case server.FieldCreatedAt, server.FieldUpdatedAt, server.FieldLastSeen:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Server", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Server fields.
func (s *Server) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case server.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = uint64(value.Int64)
		case server.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				s.CreatedAt = value.Time
			}
		case server.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				s.UpdatedAt = value.Time
			}
		case server.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				s.CreatedBy = int(value.Int64)
			}
		case server.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				s.UpdatedBy = int(value.Int64)
			}
		case server.FieldDomain:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field domain", values[i])
			} else if value.Valid {
				s.Domain = value.String
			}
		case server.FieldLastSeen:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_seen", values[i])
			} else if value.Valid {
				s.LastSeen = value.Time
			}
		}
	}
	return nil
}

// QueryActors queries the "actors" edge of the Server entity.
func (s *Server) QueryActors() *ActorQuery {
	return (&ServerClient{config: s.config}).QueryActors(s)
}

// Update returns a builder for updating this Server.
// Note that you need to call Server.Unwrap() before calling this method if this Server
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Server) Update() *ServerUpdateOne {
	return (&ServerClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the Server entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Server) Unwrap() *Server {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Server is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Server) String() string {
	var builder strings.Builder
	builder.WriteString("Server(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("created_at=")
	builder.WriteString(s.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(s.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", s.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", s.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("domain=")
	builder.WriteString(s.Domain)
	builder.WriteString(", ")
	builder.WriteString("last_seen=")
	builder.WriteString(s.LastSeen.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Servers is a parsable slice of Server.
type Servers []*Server

func (s Servers) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
