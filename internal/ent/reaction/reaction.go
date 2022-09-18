// Code generated by ent, DO NOT EDIT.

package reaction

import (
	"fmt"
	"time"

	"entgo.io/ent"
)

const (
	// Label holds the string label denoting the reaction type in the database.
	Label = "reaction"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldCreatedBy holds the string denoting the created_by field in the database.
	FieldCreatedBy = "created_by"
	// FieldUpdatedBy holds the string denoting the updated_by field in the database.
	FieldUpdatedBy = "updated_by"
	// FieldActorID holds the string denoting the actor_id field in the database.
	FieldActorID = "actor_id"
	// FieldStatusID holds the string denoting the status_id field in the database.
	FieldStatusID = "status_id"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldDat holds the string denoting the dat field in the database.
	FieldDat = "dat"
	// EdgeActors holds the string denoting the actors edge name in mutations.
	EdgeActors = "actors"
	// EdgeStatus holds the string denoting the status edge name in mutations.
	EdgeStatus = "status"
	// ActorFieldID holds the string denoting the ID field of the Actor.
	ActorFieldID = "id"
	// StatusFieldID holds the string denoting the ID field of the Status.
	StatusFieldID = "id"
	// Table holds the table name of the reaction in the database.
	Table = "reactions"
	// ActorsTable is the table that holds the actors relation/edge.
	ActorsTable = "reactions"
	// ActorsInverseTable is the table name for the Actor entity.
	// It exists in this package in order to avoid circular dependency with the "actor" package.
	ActorsInverseTable = "actors"
	// ActorsColumn is the table column denoting the actors relation/edge.
	ActorsColumn = "actor_id"
	// StatusTable is the table that holds the status relation/edge.
	StatusTable = "reactions"
	// StatusInverseTable is the table name for the Status entity.
	// It exists in this package in order to avoid circular dependency with the "status" package.
	StatusInverseTable = "status"
	// StatusColumn is the table column denoting the status relation/edge.
	StatusColumn = "status_id"
)

// Columns holds all SQL columns for reaction fields.
var Columns = []string{
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldCreatedBy,
	FieldUpdatedBy,
	FieldActorID,
	FieldStatusID,
	FieldType,
	FieldDat,
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
	Hooks  [2]ent.Hook
	Policy ent.Policy
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// Type defines the type for the "type" enum field.
type Type string

// Type values.
const (
	TypeLike  Type = "like"
	TypeBoost Type = "boost"
	TypeEmote Type = "emote"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeLike, TypeBoost, TypeEmote:
		return nil
	default:
		return fmt.Errorf("reaction: invalid enum value for type field: %q", _type)
	}
}
