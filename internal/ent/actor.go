// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/eriner/burr/internal/ent/actor"
	"github.com/eriner/burr/internal/ent/server"
)

// Actor is the model entity for the Actor schema.
type Actor struct {
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
	// The type of account. All accounts are AP actors
	Type actor.Type `json:"type,omitempty"`
	// The Actor's username
	Name string `json:"name,omitempty"`
	// The Actor's displayed 'friendly' name
	DisplayName string `json:"display_name,omitempty"`
	// Actor note, AKA description
	Note string `json:"note,omitempty"`
	// Actor account is locked if unconfirmed or explicitly locked
	Locked bool `json:"locked,omitempty"`
	// Actor account is memorialized
	Memorial bool `json:"memorial,omitempty"`
	// Actor URL
	URL string `json:"url,omitempty"`
	// Actor public key
	Pubkey []byte `json:"pubkey,omitempty"`
	// Actor private key
	Privkey *[]byte `json:"-"`
	// URL of the Actor's remote avatar
	AvatarRemoteURL *string `json:"avatar_remote_url,omitempty"`
	// The Actor's local avatar file
	AvatarLocalFile *string `json:"avatar_local_file,omitempty"`
	// The time the Actor's (local) avatar was last updated
	AvatarUpdatedAt *time.Time `json:"avatar_updated_at,omitempty"`
	// URL of user header
	HeaderURL *string `json:"header_url,omitempty"`
	// The Actor's local header file
	HeaderLocalFile *string `json:"header_local_file,omitempty"`
	// The time the Actor's header was last updated
	HeaderUpdatedAt *time.Time `json:"header_updated_at,omitempty"`
	// The time the Actor's account was last discovered with webfinger
	LastWebfingerAt *time.Time `json:"last_webfinger_at,omitempty"`
	// The Actor's ActivityPub Inbox IRI
	InboxURL string `json:"inbox_url,omitempty"`
	// The Actor's ActivityPub Inbox IRI
	OutboxURL string `json:"outbox_url,omitempty"`
	// The Actor's ActivityPub Inbox IRI
	SharedInboxURL string `json:"shared_inbox_url,omitempty"`
	// The Actor's ActivityPub Inbox IRI
	FollowersURL string `json:"followers_url,omitempty"`
	// The Actor's id, if it has moved
	MovedToID *uint64 `json:"moved_to_id,omitempty"`
	// The Actor's featured items
	FeaturedCollectionURL *string `json:"featured_collection_url,omitempty"`
	// The time the Actor was silenced
	SilencedAt *time.Time `json:"silenced_at,omitempty"`
	// The time the Actor was silenced
	SuspendedAt *time.Time `json:"suspended_at,omitempty"`
	// Actor bcrypt password hash
	PasswordHash *[]byte `json:"-"`
	// local Actor password recovery code generated during account creation
	RecoveryCode *string `json:"-"`
	// local Actor role
	Role uint64 `json:"role,omitempty"`
	// local Actor badge
	Badge uint64 `json:"badge,omitempty"`
	// local Actor locale
	Locale actor.Locale `json:"locale,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ActorQuery when eager-loading is set.
	Edges         ActorEdges `json:"edges"`
	server_actors *uint64
}

// ActorEdges holds the relations/edges for other nodes in the graph.
type ActorEdges struct {
	// Actor belong to a Server rooted at a domain (FQDN)
	Server *Server `json:"server,omitempty"`
	// Actor can belong to one or more events
	Events []*Event `json:"events,omitempty"`
	// Users can organize events
	OrganizerOf []*Event `json:"organizer_of,omitempty"`
	// Actor can have zero or many statuses
	Statuses []*Status `json:"statuses,omitempty"`
	// Actor can be followed by zero or more accounts
	Followers []*Actor `json:"followers,omitempty"`
	// Following holds the value of the following edge.
	Following []*Actor `json:"following,omitempty"`
	// Actor can react to many statuses
	ReactedStatuses []*Status `json:"reacted_statuses,omitempty"`
	// Group Actor can have zero or more moderators
	Moderators []*Actor `json:"moderators,omitempty"`
	// Moderating holds the value of the moderating edge.
	Moderating []*Actor `json:"moderating,omitempty"`
	// User Actor can belong to groups
	Members []*Actor `json:"members,omitempty"`
	// Groups holds the value of the groups edge.
	Groups []*Actor `json:"groups,omitempty"`
	// Sessions are owned by a single User Actor which may have many sessions
	Sessions []*Session `json:"sessions,omitempty"`
	// Reactions holds the value of the reactions edge.
	Reactions []*Reaction `json:"reactions,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [13]bool
}

// ServerOrErr returns the Server value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ActorEdges) ServerOrErr() (*Server, error) {
	if e.loadedTypes[0] {
		if e.Server == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: server.Label}
		}
		return e.Server, nil
	}
	return nil, &NotLoadedError{edge: "server"}
}

// EventsOrErr returns the Events value or an error if the edge
// was not loaded in eager-loading.
func (e ActorEdges) EventsOrErr() ([]*Event, error) {
	if e.loadedTypes[1] {
		return e.Events, nil
	}
	return nil, &NotLoadedError{edge: "events"}
}

// OrganizerOfOrErr returns the OrganizerOf value or an error if the edge
// was not loaded in eager-loading.
func (e ActorEdges) OrganizerOfOrErr() ([]*Event, error) {
	if e.loadedTypes[2] {
		return e.OrganizerOf, nil
	}
	return nil, &NotLoadedError{edge: "organizer_of"}
}

// StatusesOrErr returns the Statuses value or an error if the edge
// was not loaded in eager-loading.
func (e ActorEdges) StatusesOrErr() ([]*Status, error) {
	if e.loadedTypes[3] {
		return e.Statuses, nil
	}
	return nil, &NotLoadedError{edge: "statuses"}
}

// FollowersOrErr returns the Followers value or an error if the edge
// was not loaded in eager-loading.
func (e ActorEdges) FollowersOrErr() ([]*Actor, error) {
	if e.loadedTypes[4] {
		return e.Followers, nil
	}
	return nil, &NotLoadedError{edge: "followers"}
}

// FollowingOrErr returns the Following value or an error if the edge
// was not loaded in eager-loading.
func (e ActorEdges) FollowingOrErr() ([]*Actor, error) {
	if e.loadedTypes[5] {
		return e.Following, nil
	}
	return nil, &NotLoadedError{edge: "following"}
}

// ReactedStatusesOrErr returns the ReactedStatuses value or an error if the edge
// was not loaded in eager-loading.
func (e ActorEdges) ReactedStatusesOrErr() ([]*Status, error) {
	if e.loadedTypes[6] {
		return e.ReactedStatuses, nil
	}
	return nil, &NotLoadedError{edge: "reacted_statuses"}
}

// ModeratorsOrErr returns the Moderators value or an error if the edge
// was not loaded in eager-loading.
func (e ActorEdges) ModeratorsOrErr() ([]*Actor, error) {
	if e.loadedTypes[7] {
		return e.Moderators, nil
	}
	return nil, &NotLoadedError{edge: "moderators"}
}

// ModeratingOrErr returns the Moderating value or an error if the edge
// was not loaded in eager-loading.
func (e ActorEdges) ModeratingOrErr() ([]*Actor, error) {
	if e.loadedTypes[8] {
		return e.Moderating, nil
	}
	return nil, &NotLoadedError{edge: "moderating"}
}

// MembersOrErr returns the Members value or an error if the edge
// was not loaded in eager-loading.
func (e ActorEdges) MembersOrErr() ([]*Actor, error) {
	if e.loadedTypes[9] {
		return e.Members, nil
	}
	return nil, &NotLoadedError{edge: "members"}
}

// GroupsOrErr returns the Groups value or an error if the edge
// was not loaded in eager-loading.
func (e ActorEdges) GroupsOrErr() ([]*Actor, error) {
	if e.loadedTypes[10] {
		return e.Groups, nil
	}
	return nil, &NotLoadedError{edge: "groups"}
}

// SessionsOrErr returns the Sessions value or an error if the edge
// was not loaded in eager-loading.
func (e ActorEdges) SessionsOrErr() ([]*Session, error) {
	if e.loadedTypes[11] {
		return e.Sessions, nil
	}
	return nil, &NotLoadedError{edge: "sessions"}
}

// ReactionsOrErr returns the Reactions value or an error if the edge
// was not loaded in eager-loading.
func (e ActorEdges) ReactionsOrErr() ([]*Reaction, error) {
	if e.loadedTypes[12] {
		return e.Reactions, nil
	}
	return nil, &NotLoadedError{edge: "reactions"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Actor) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case actor.FieldPubkey, actor.FieldPrivkey, actor.FieldPasswordHash:
			values[i] = new([]byte)
		case actor.FieldLocked, actor.FieldMemorial:
			values[i] = new(sql.NullBool)
		case actor.FieldID, actor.FieldCreatedBy, actor.FieldUpdatedBy, actor.FieldMovedToID, actor.FieldRole, actor.FieldBadge:
			values[i] = new(sql.NullInt64)
		case actor.FieldType, actor.FieldName, actor.FieldDisplayName, actor.FieldNote, actor.FieldURL, actor.FieldAvatarRemoteURL, actor.FieldAvatarLocalFile, actor.FieldHeaderURL, actor.FieldHeaderLocalFile, actor.FieldInboxURL, actor.FieldOutboxURL, actor.FieldSharedInboxURL, actor.FieldFollowersURL, actor.FieldFeaturedCollectionURL, actor.FieldRecoveryCode, actor.FieldLocale:
			values[i] = new(sql.NullString)
		case actor.FieldCreatedAt, actor.FieldUpdatedAt, actor.FieldAvatarUpdatedAt, actor.FieldHeaderUpdatedAt, actor.FieldLastWebfingerAt, actor.FieldSilencedAt, actor.FieldSuspendedAt:
			values[i] = new(sql.NullTime)
		case actor.ForeignKeys[0]: // server_actors
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Actor", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Actor fields.
func (a *Actor) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case actor.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = uint64(value.Int64)
		case actor.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				a.CreatedAt = value.Time
			}
		case actor.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				a.UpdatedAt = value.Time
			}
		case actor.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				a.CreatedBy = int(value.Int64)
			}
		case actor.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				a.UpdatedBy = int(value.Int64)
			}
		case actor.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				a.Type = actor.Type(value.String)
			}
		case actor.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				a.Name = value.String
			}
		case actor.FieldDisplayName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field display_name", values[i])
			} else if value.Valid {
				a.DisplayName = value.String
			}
		case actor.FieldNote:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field note", values[i])
			} else if value.Valid {
				a.Note = value.String
			}
		case actor.FieldLocked:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field locked", values[i])
			} else if value.Valid {
				a.Locked = value.Bool
			}
		case actor.FieldMemorial:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field memorial", values[i])
			} else if value.Valid {
				a.Memorial = value.Bool
			}
		case actor.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				a.URL = value.String
			}
		case actor.FieldPubkey:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field pubkey", values[i])
			} else if value != nil {
				a.Pubkey = *value
			}
		case actor.FieldPrivkey:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field privkey", values[i])
			} else if value != nil {
				a.Privkey = value
			}
		case actor.FieldAvatarRemoteURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field avatar_remote_url", values[i])
			} else if value.Valid {
				a.AvatarRemoteURL = new(string)
				*a.AvatarRemoteURL = value.String
			}
		case actor.FieldAvatarLocalFile:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field avatar_local_file", values[i])
			} else if value.Valid {
				a.AvatarLocalFile = new(string)
				*a.AvatarLocalFile = value.String
			}
		case actor.FieldAvatarUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field avatar_updated_at", values[i])
			} else if value.Valid {
				a.AvatarUpdatedAt = new(time.Time)
				*a.AvatarUpdatedAt = value.Time
			}
		case actor.FieldHeaderURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field header_url", values[i])
			} else if value.Valid {
				a.HeaderURL = new(string)
				*a.HeaderURL = value.String
			}
		case actor.FieldHeaderLocalFile:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field header_local_file", values[i])
			} else if value.Valid {
				a.HeaderLocalFile = new(string)
				*a.HeaderLocalFile = value.String
			}
		case actor.FieldHeaderUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field header_updated_at", values[i])
			} else if value.Valid {
				a.HeaderUpdatedAt = new(time.Time)
				*a.HeaderUpdatedAt = value.Time
			}
		case actor.FieldLastWebfingerAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_webfinger_at", values[i])
			} else if value.Valid {
				a.LastWebfingerAt = new(time.Time)
				*a.LastWebfingerAt = value.Time
			}
		case actor.FieldInboxURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field inbox_url", values[i])
			} else if value.Valid {
				a.InboxURL = value.String
			}
		case actor.FieldOutboxURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field outbox_url", values[i])
			} else if value.Valid {
				a.OutboxURL = value.String
			}
		case actor.FieldSharedInboxURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field shared_inbox_url", values[i])
			} else if value.Valid {
				a.SharedInboxURL = value.String
			}
		case actor.FieldFollowersURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field followers_url", values[i])
			} else if value.Valid {
				a.FollowersURL = value.String
			}
		case actor.FieldMovedToID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field moved_to_id", values[i])
			} else if value.Valid {
				a.MovedToID = new(uint64)
				*a.MovedToID = uint64(value.Int64)
			}
		case actor.FieldFeaturedCollectionURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field featured_collection_url", values[i])
			} else if value.Valid {
				a.FeaturedCollectionURL = new(string)
				*a.FeaturedCollectionURL = value.String
			}
		case actor.FieldSilencedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field silenced_at", values[i])
			} else if value.Valid {
				a.SilencedAt = new(time.Time)
				*a.SilencedAt = value.Time
			}
		case actor.FieldSuspendedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field suspended_at", values[i])
			} else if value.Valid {
				a.SuspendedAt = new(time.Time)
				*a.SuspendedAt = value.Time
			}
		case actor.FieldPasswordHash:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field passwordHash", values[i])
			} else if value != nil {
				a.PasswordHash = value
			}
		case actor.FieldRecoveryCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field recovery_code", values[i])
			} else if value.Valid {
				a.RecoveryCode = new(string)
				*a.RecoveryCode = value.String
			}
		case actor.FieldRole:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field role", values[i])
			} else if value.Valid {
				a.Role = uint64(value.Int64)
			}
		case actor.FieldBadge:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field badge", values[i])
			} else if value.Valid {
				a.Badge = uint64(value.Int64)
			}
		case actor.FieldLocale:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field locale", values[i])
			} else if value.Valid {
				a.Locale = actor.Locale(value.String)
			}
		case actor.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field server_actors", value)
			} else if value.Valid {
				a.server_actors = new(uint64)
				*a.server_actors = uint64(value.Int64)
			}
		}
	}
	return nil
}

// QueryServer queries the "server" edge of the Actor entity.
func (a *Actor) QueryServer() *ServerQuery {
	return (&ActorClient{config: a.config}).QueryServer(a)
}

// QueryEvents queries the "events" edge of the Actor entity.
func (a *Actor) QueryEvents() *EventQuery {
	return (&ActorClient{config: a.config}).QueryEvents(a)
}

// QueryOrganizerOf queries the "organizer_of" edge of the Actor entity.
func (a *Actor) QueryOrganizerOf() *EventQuery {
	return (&ActorClient{config: a.config}).QueryOrganizerOf(a)
}

// QueryStatuses queries the "statuses" edge of the Actor entity.
func (a *Actor) QueryStatuses() *StatusQuery {
	return (&ActorClient{config: a.config}).QueryStatuses(a)
}

// QueryFollowers queries the "followers" edge of the Actor entity.
func (a *Actor) QueryFollowers() *ActorQuery {
	return (&ActorClient{config: a.config}).QueryFollowers(a)
}

// QueryFollowing queries the "following" edge of the Actor entity.
func (a *Actor) QueryFollowing() *ActorQuery {
	return (&ActorClient{config: a.config}).QueryFollowing(a)
}

// QueryReactedStatuses queries the "reacted_statuses" edge of the Actor entity.
func (a *Actor) QueryReactedStatuses() *StatusQuery {
	return (&ActorClient{config: a.config}).QueryReactedStatuses(a)
}

// QueryModerators queries the "moderators" edge of the Actor entity.
func (a *Actor) QueryModerators() *ActorQuery {
	return (&ActorClient{config: a.config}).QueryModerators(a)
}

// QueryModerating queries the "moderating" edge of the Actor entity.
func (a *Actor) QueryModerating() *ActorQuery {
	return (&ActorClient{config: a.config}).QueryModerating(a)
}

// QueryMembers queries the "members" edge of the Actor entity.
func (a *Actor) QueryMembers() *ActorQuery {
	return (&ActorClient{config: a.config}).QueryMembers(a)
}

// QueryGroups queries the "groups" edge of the Actor entity.
func (a *Actor) QueryGroups() *ActorQuery {
	return (&ActorClient{config: a.config}).QueryGroups(a)
}

// QuerySessions queries the "sessions" edge of the Actor entity.
func (a *Actor) QuerySessions() *SessionQuery {
	return (&ActorClient{config: a.config}).QuerySessions(a)
}

// QueryReactions queries the "reactions" edge of the Actor entity.
func (a *Actor) QueryReactions() *ReactionQuery {
	return (&ActorClient{config: a.config}).QueryReactions(a)
}

// Update returns a builder for updating this Actor.
// Note that you need to call Actor.Unwrap() before calling this method if this Actor
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Actor) Update() *ActorUpdateOne {
	return (&ActorClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the Actor entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Actor) Unwrap() *Actor {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Actor is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Actor) String() string {
	var builder strings.Builder
	builder.WriteString("Actor(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("created_at=")
	builder.WriteString(a.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(a.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", a.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", a.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", a.Type))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(a.Name)
	builder.WriteString(", ")
	builder.WriteString("display_name=")
	builder.WriteString(a.DisplayName)
	builder.WriteString(", ")
	builder.WriteString("note=")
	builder.WriteString(a.Note)
	builder.WriteString(", ")
	builder.WriteString("locked=")
	builder.WriteString(fmt.Sprintf("%v", a.Locked))
	builder.WriteString(", ")
	builder.WriteString("memorial=")
	builder.WriteString(fmt.Sprintf("%v", a.Memorial))
	builder.WriteString(", ")
	builder.WriteString("url=")
	builder.WriteString(a.URL)
	builder.WriteString(", ")
	builder.WriteString("pubkey=")
	builder.WriteString(fmt.Sprintf("%v", a.Pubkey))
	builder.WriteString(", ")
	builder.WriteString("privkey=<sensitive>")
	builder.WriteString(", ")
	if v := a.AvatarRemoteURL; v != nil {
		builder.WriteString("avatar_remote_url=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := a.AvatarLocalFile; v != nil {
		builder.WriteString("avatar_local_file=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := a.AvatarUpdatedAt; v != nil {
		builder.WriteString("avatar_updated_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := a.HeaderURL; v != nil {
		builder.WriteString("header_url=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := a.HeaderLocalFile; v != nil {
		builder.WriteString("header_local_file=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := a.HeaderUpdatedAt; v != nil {
		builder.WriteString("header_updated_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := a.LastWebfingerAt; v != nil {
		builder.WriteString("last_webfinger_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("inbox_url=")
	builder.WriteString(a.InboxURL)
	builder.WriteString(", ")
	builder.WriteString("outbox_url=")
	builder.WriteString(a.OutboxURL)
	builder.WriteString(", ")
	builder.WriteString("shared_inbox_url=")
	builder.WriteString(a.SharedInboxURL)
	builder.WriteString(", ")
	builder.WriteString("followers_url=")
	builder.WriteString(a.FollowersURL)
	builder.WriteString(", ")
	if v := a.MovedToID; v != nil {
		builder.WriteString("moved_to_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := a.FeaturedCollectionURL; v != nil {
		builder.WriteString("featured_collection_url=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := a.SilencedAt; v != nil {
		builder.WriteString("silenced_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := a.SuspendedAt; v != nil {
		builder.WriteString("suspended_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("passwordHash=<sensitive>")
	builder.WriteString(", ")
	builder.WriteString("recovery_code=<sensitive>")
	builder.WriteString(", ")
	builder.WriteString("role=")
	builder.WriteString(fmt.Sprintf("%v", a.Role))
	builder.WriteString(", ")
	builder.WriteString("badge=")
	builder.WriteString(fmt.Sprintf("%v", a.Badge))
	builder.WriteString(", ")
	builder.WriteString("locale=")
	builder.WriteString(fmt.Sprintf("%v", a.Locale))
	builder.WriteByte(')')
	return builder.String()
}

// Actors is a parsable slice of Actor.
type Actors []*Actor

func (a Actors) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}