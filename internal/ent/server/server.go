// Code generated by ent, DO NOT EDIT.

package server

import (
	"time"

	"entgo.io/ent"
)

const (
	// Label holds the string label denoting the server type in the database.
	Label = "server"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldCreatedBy holds the string denoting the created_by field in the database.
	FieldCreatedBy = "created_by"
	// FieldUpdatedBy holds the string denoting the updated_by field in the database.
	FieldUpdatedBy = "updated_by"
	// FieldDomain holds the string denoting the domain field in the database.
	FieldDomain = "domain"
	// FieldLastSeen holds the string denoting the last_seen field in the database.
	FieldLastSeen = "last_seen"
	// EdgeActors holds the string denoting the actors edge name in mutations.
	EdgeActors = "actors"
	// Table holds the table name of the server in the database.
	Table = "servers"
	// ActorsTable is the table that holds the actors relation/edge.
	ActorsTable = "actors"
	// ActorsInverseTable is the table name for the Actor entity.
	// It exists in this package in order to avoid circular dependency with the "actor" package.
	ActorsInverseTable = "actors"
	// ActorsColumn is the table column denoting the actors relation/edge.
	ActorsColumn = "server_actors"
)

// Columns holds all SQL columns for server fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldCreatedBy,
	FieldUpdatedBy,
	FieldDomain,
	FieldLastSeen,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/eriner/burr/internal/ent/runtime"
var (
	Hooks  [3]ent.Hook
	Policy ent.Policy
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DomainValidator is a validator for the "domain" field. It is called by the builders before save.
	DomainValidator func(string) error
	// DefaultLastSeen holds the default value on creation for the "last_seen" field.
	DefaultLastSeen func() time.Time
)
