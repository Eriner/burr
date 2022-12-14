// Code generated by ent, DO NOT EDIT.

package actor

import (
	"fmt"
	"time"

	"entgo.io/ent"
)

const (
	// Label holds the string label denoting the actor type in the database.
	Label = "actor"
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
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDisplayName holds the string denoting the display_name field in the database.
	FieldDisplayName = "display_name"
	// FieldNote holds the string denoting the note field in the database.
	FieldNote = "note"
	// FieldLocked holds the string denoting the locked field in the database.
	FieldLocked = "locked"
	// FieldMemorial holds the string denoting the memorial field in the database.
	FieldMemorial = "memorial"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// FieldPubkey holds the string denoting the pubkey field in the database.
	FieldPubkey = "pubkey"
	// FieldPrivkey holds the string denoting the privkey field in the database.
	FieldPrivkey = "privkey"
	// FieldAvatarRemoteURL holds the string denoting the avatar_remote_url field in the database.
	FieldAvatarRemoteURL = "avatar_remote_url"
	// FieldAvatarLocalFile holds the string denoting the avatar_local_file field in the database.
	FieldAvatarLocalFile = "avatar_local_file"
	// FieldAvatarUpdatedAt holds the string denoting the avatar_updated_at field in the database.
	FieldAvatarUpdatedAt = "avatar_updated_at"
	// FieldHeaderURL holds the string denoting the header_url field in the database.
	FieldHeaderURL = "header_url"
	// FieldHeaderLocalFile holds the string denoting the header_local_file field in the database.
	FieldHeaderLocalFile = "header_local_file"
	// FieldHeaderUpdatedAt holds the string denoting the header_updated_at field in the database.
	FieldHeaderUpdatedAt = "header_updated_at"
	// FieldLastWebfingerAt holds the string denoting the last_webfinger_at field in the database.
	FieldLastWebfingerAt = "last_webfinger_at"
	// FieldInboxURL holds the string denoting the inbox_url field in the database.
	FieldInboxURL = "inbox_url"
	// FieldOutboxURL holds the string denoting the outbox_url field in the database.
	FieldOutboxURL = "outbox_url"
	// FieldSharedInboxURL holds the string denoting the shared_inbox_url field in the database.
	FieldSharedInboxURL = "shared_inbox_url"
	// FieldFollowersURL holds the string denoting the followers_url field in the database.
	FieldFollowersURL = "followers_url"
	// FieldMovedToID holds the string denoting the moved_to_id field in the database.
	FieldMovedToID = "moved_to_id"
	// FieldFeaturedCollectionURL holds the string denoting the featured_collection_url field in the database.
	FieldFeaturedCollectionURL = "featured_collection_url"
	// FieldSilencedAt holds the string denoting the silenced_at field in the database.
	FieldSilencedAt = "silenced_at"
	// FieldSuspendedAt holds the string denoting the suspended_at field in the database.
	FieldSuspendedAt = "suspended_at"
	// FieldPasswordHash holds the string denoting the passwordhash field in the database.
	FieldPasswordHash = "password_hash"
	// FieldRecoveryCode holds the string denoting the recovery_code field in the database.
	FieldRecoveryCode = "recovery_code"
	// FieldRole holds the string denoting the role field in the database.
	FieldRole = "role"
	// FieldBadge holds the string denoting the badge field in the database.
	FieldBadge = "badge"
	// FieldLocale holds the string denoting the locale field in the database.
	FieldLocale = "locale"
	// EdgeServer holds the string denoting the server edge name in mutations.
	EdgeServer = "server"
	// EdgeEvents holds the string denoting the events edge name in mutations.
	EdgeEvents = "events"
	// EdgeOrganizerOf holds the string denoting the organizer_of edge name in mutations.
	EdgeOrganizerOf = "organizer_of"
	// EdgeStatuses holds the string denoting the statuses edge name in mutations.
	EdgeStatuses = "statuses"
	// EdgeFollowers holds the string denoting the followers edge name in mutations.
	EdgeFollowers = "followers"
	// EdgeFollowing holds the string denoting the following edge name in mutations.
	EdgeFollowing = "following"
	// EdgeReactedStatuses holds the string denoting the reacted_statuses edge name in mutations.
	EdgeReactedStatuses = "reacted_statuses"
	// EdgeModerators holds the string denoting the moderators edge name in mutations.
	EdgeModerators = "moderators"
	// EdgeModerating holds the string denoting the moderating edge name in mutations.
	EdgeModerating = "moderating"
	// EdgeMembers holds the string denoting the members edge name in mutations.
	EdgeMembers = "members"
	// EdgeGroups holds the string denoting the groups edge name in mutations.
	EdgeGroups = "groups"
	// EdgeSessions holds the string denoting the sessions edge name in mutations.
	EdgeSessions = "sessions"
	// EdgeReactions holds the string denoting the reactions edge name in mutations.
	EdgeReactions = "reactions"
	// Table holds the table name of the actor in the database.
	Table = "actors"
	// ServerTable is the table that holds the server relation/edge.
	ServerTable = "actors"
	// ServerInverseTable is the table name for the Server entity.
	// It exists in this package in order to avoid circular dependency with the "server" package.
	ServerInverseTable = "servers"
	// ServerColumn is the table column denoting the server relation/edge.
	ServerColumn = "server_actors"
	// EventsTable is the table that holds the events relation/edge. The primary key declared below.
	EventsTable = "event_attendees"
	// EventsInverseTable is the table name for the Event entity.
	// It exists in this package in order to avoid circular dependency with the "event" package.
	EventsInverseTable = "events"
	// OrganizerOfTable is the table that holds the organizer_of relation/edge.
	OrganizerOfTable = "events"
	// OrganizerOfInverseTable is the table name for the Event entity.
	// It exists in this package in order to avoid circular dependency with the "event" package.
	OrganizerOfInverseTable = "events"
	// OrganizerOfColumn is the table column denoting the organizer_of relation/edge.
	OrganizerOfColumn = "actor_organizer_of"
	// StatusesTable is the table that holds the statuses relation/edge.
	StatusesTable = "status"
	// StatusesInverseTable is the table name for the Status entity.
	// It exists in this package in order to avoid circular dependency with the "status" package.
	StatusesInverseTable = "status"
	// StatusesColumn is the table column denoting the statuses relation/edge.
	StatusesColumn = "actor_statuses"
	// FollowersTable is the table that holds the followers relation/edge. The primary key declared below.
	FollowersTable = "actor_following"
	// FollowingTable is the table that holds the following relation/edge. The primary key declared below.
	FollowingTable = "actor_following"
	// ReactedStatusesTable is the table that holds the reacted_statuses relation/edge. The primary key declared below.
	ReactedStatusesTable = "reactions"
	// ReactedStatusesInverseTable is the table name for the Status entity.
	// It exists in this package in order to avoid circular dependency with the "status" package.
	ReactedStatusesInverseTable = "status"
	// ModeratorsTable is the table that holds the moderators relation/edge. The primary key declared below.
	ModeratorsTable = "actor_moderating"
	// ModeratingTable is the table that holds the moderating relation/edge. The primary key declared below.
	ModeratingTable = "actor_moderating"
	// MembersTable is the table that holds the members relation/edge. The primary key declared below.
	MembersTable = "actor_groups"
	// GroupsTable is the table that holds the groups relation/edge. The primary key declared below.
	GroupsTable = "actor_groups"
	// SessionsTable is the table that holds the sessions relation/edge.
	SessionsTable = "sessions"
	// SessionsInverseTable is the table name for the Session entity.
	// It exists in this package in order to avoid circular dependency with the "session" package.
	SessionsInverseTable = "sessions"
	// SessionsColumn is the table column denoting the sessions relation/edge.
	SessionsColumn = "actor_sessions"
	// ReactionsTable is the table that holds the reactions relation/edge.
	ReactionsTable = "reactions"
	// ReactionsInverseTable is the table name for the Reaction entity.
	// It exists in this package in order to avoid circular dependency with the "reaction" package.
	ReactionsInverseTable = "reactions"
	// ReactionsColumn is the table column denoting the reactions relation/edge.
	ReactionsColumn = "actor_id"
)

// Columns holds all SQL columns for actor fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldCreatedBy,
	FieldUpdatedBy,
	FieldType,
	FieldName,
	FieldDisplayName,
	FieldNote,
	FieldLocked,
	FieldMemorial,
	FieldURL,
	FieldPubkey,
	FieldPrivkey,
	FieldAvatarRemoteURL,
	FieldAvatarLocalFile,
	FieldAvatarUpdatedAt,
	FieldHeaderURL,
	FieldHeaderLocalFile,
	FieldHeaderUpdatedAt,
	FieldLastWebfingerAt,
	FieldInboxURL,
	FieldOutboxURL,
	FieldSharedInboxURL,
	FieldFollowersURL,
	FieldMovedToID,
	FieldFeaturedCollectionURL,
	FieldSilencedAt,
	FieldSuspendedAt,
	FieldPasswordHash,
	FieldRecoveryCode,
	FieldRole,
	FieldBadge,
	FieldLocale,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "actors"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"server_actors",
}

var (
	// EventsPrimaryKey and EventsColumn2 are the table columns denoting the
	// primary key for the events relation (M2M).
	EventsPrimaryKey = []string{"event_id", "actor_id"}
	// FollowersPrimaryKey and FollowersColumn2 are the table columns denoting the
	// primary key for the followers relation (M2M).
	FollowersPrimaryKey = []string{"actor_id", "follower_id"}
	// FollowingPrimaryKey and FollowingColumn2 are the table columns denoting the
	// primary key for the following relation (M2M).
	FollowingPrimaryKey = []string{"actor_id", "follower_id"}
	// ReactedStatusesPrimaryKey and ReactedStatusesColumn2 are the table columns denoting the
	// primary key for the reacted_statuses relation (M2M).
	ReactedStatusesPrimaryKey = []string{"actor_id", "status_id"}
	// ModeratorsPrimaryKey and ModeratorsColumn2 are the table columns denoting the
	// primary key for the moderators relation (M2M).
	ModeratorsPrimaryKey = []string{"actor_id", "moderator_id"}
	// ModeratingPrimaryKey and ModeratingColumn2 are the table columns denoting the
	// primary key for the moderating relation (M2M).
	ModeratingPrimaryKey = []string{"actor_id", "moderator_id"}
	// MembersPrimaryKey and MembersColumn2 are the table columns denoting the
	// primary key for the members relation (M2M).
	MembersPrimaryKey = []string{"actor_id", "member_id"}
	// GroupsPrimaryKey and GroupsColumn2 are the table columns denoting the
	// primary key for the groups relation (M2M).
	GroupsPrimaryKey = []string{"actor_id", "member_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
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
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultDisplayName holds the default value on creation for the "display_name" field.
	DefaultDisplayName string
	// DisplayNameValidator is a validator for the "display_name" field. It is called by the builders before save.
	DisplayNameValidator func(string) error
	// DefaultNote holds the default value on creation for the "note" field.
	DefaultNote string
	// NoteValidator is a validator for the "note" field. It is called by the builders before save.
	NoteValidator func(string) error
	// DefaultLocked holds the default value on creation for the "locked" field.
	DefaultLocked bool
	// DefaultMemorial holds the default value on creation for the "memorial" field.
	DefaultMemorial bool
	// URLValidator is a validator for the "url" field. It is called by the builders before save.
	URLValidator func(string) error
	// DefaultPubkey holds the default value on creation for the "pubkey" field.
	DefaultPubkey []byte
	// PubkeyValidator is a validator for the "pubkey" field. It is called by the builders before save.
	PubkeyValidator func([]byte) error
	// AvatarRemoteURLValidator is a validator for the "avatar_remote_url" field. It is called by the builders before save.
	AvatarRemoteURLValidator func(string) error
	// AvatarLocalFileValidator is a validator for the "avatar_local_file" field. It is called by the builders before save.
	AvatarLocalFileValidator func(string) error
	// HeaderURLValidator is a validator for the "header_url" field. It is called by the builders before save.
	HeaderURLValidator func(string) error
	// HeaderLocalFileValidator is a validator for the "header_local_file" field. It is called by the builders before save.
	HeaderLocalFileValidator func(string) error
	// DefaultLastWebfingerAt holds the default value on creation for the "last_webfinger_at" field.
	DefaultLastWebfingerAt func() time.Time
	// InboxURLValidator is a validator for the "inbox_url" field. It is called by the builders before save.
	InboxURLValidator func(string) error
	// OutboxURLValidator is a validator for the "outbox_url" field. It is called by the builders before save.
	OutboxURLValidator func(string) error
	// SharedInboxURLValidator is a validator for the "shared_inbox_url" field. It is called by the builders before save.
	SharedInboxURLValidator func(string) error
	// FollowersURLValidator is a validator for the "followers_url" field. It is called by the builders before save.
	FollowersURLValidator func(string) error
	// PasswordHashValidator is a validator for the "passwordHash" field. It is called by the builders before save.
	PasswordHashValidator func([]byte) error
	// DefaultRole holds the default value on creation for the "role" field.
	DefaultRole uint64
	// DefaultBadge holds the default value on creation for the "badge" field.
	DefaultBadge uint64
)

// Type defines the type for the "type" enum field.
type Type string

// Type values.
const (
	TypeLocalUser   Type = "local_user"
	TypeLocalBot    Type = "local_bot"
	TypeLocalGroup  Type = "local_group"
	TypeLocalEvent  Type = "local_event"
	TypeRemoteUser  Type = "remote_user"
	TypeRemoteBot   Type = "remote_bot"
	TypeRemoteGroup Type = "remote_group"
	TypeRemoteEvent Type = "remote_event"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeLocalUser, TypeLocalBot, TypeLocalGroup, TypeLocalEvent, TypeRemoteUser, TypeRemoteBot, TypeRemoteGroup, TypeRemoteEvent:
		return nil
	default:
		return fmt.Errorf("actor: invalid enum value for type field: %q", _type)
	}
}

// Locale defines the type for the "locale" enum field.
type Locale string

// Locale values.
const (
	LocaleEnUS Locale = "en-US"
)

func (l Locale) String() string {
	return string(l)
}

// LocaleValidator is a validator for the "locale" field enum values. It is called by the builders before save.
func LocaleValidator(l Locale) error {
	switch l {
	case LocaleEnUS:
		return nil
	default:
		return fmt.Errorf("actor: invalid enum value for locale field: %q", l)
	}
}
