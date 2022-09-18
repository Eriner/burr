// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/eriner/burr/internal/ent/migrate"

	"github.com/eriner/burr/internal/ent/actor"
	"github.com/eriner/burr/internal/ent/event"
	"github.com/eriner/burr/internal/ent/group"
	"github.com/eriner/burr/internal/ent/reaction"
	"github.com/eriner/burr/internal/ent/server"
	"github.com/eriner/burr/internal/ent/session"
	"github.com/eriner/burr/internal/ent/status"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Actor is the client for interacting with the Actor builders.
	Actor *ActorClient
	// Event is the client for interacting with the Event builders.
	Event *EventClient
	// Group is the client for interacting with the Group builders.
	Group *GroupClient
	// Reaction is the client for interacting with the Reaction builders.
	Reaction *ReactionClient
	// Server is the client for interacting with the Server builders.
	Server *ServerClient
	// Session is the client for interacting with the Session builders.
	Session *SessionClient
	// Status is the client for interacting with the Status builders.
	Status *StatusClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Actor = NewActorClient(c.config)
	c.Event = NewEventClient(c.config)
	c.Group = NewGroupClient(c.config)
	c.Reaction = NewReactionClient(c.config)
	c.Server = NewServerClient(c.config)
	c.Session = NewSessionClient(c.config)
	c.Status = NewStatusClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:      ctx,
		config:   cfg,
		Actor:    NewActorClient(cfg),
		Event:    NewEventClient(cfg),
		Group:    NewGroupClient(cfg),
		Reaction: NewReactionClient(cfg),
		Server:   NewServerClient(cfg),
		Session:  NewSessionClient(cfg),
		Status:   NewStatusClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:      ctx,
		config:   cfg,
		Actor:    NewActorClient(cfg),
		Event:    NewEventClient(cfg),
		Group:    NewGroupClient(cfg),
		Reaction: NewReactionClient(cfg),
		Server:   NewServerClient(cfg),
		Session:  NewSessionClient(cfg),
		Status:   NewStatusClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Actor.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Actor.Use(hooks...)
	c.Event.Use(hooks...)
	c.Group.Use(hooks...)
	c.Reaction.Use(hooks...)
	c.Server.Use(hooks...)
	c.Session.Use(hooks...)
	c.Status.Use(hooks...)
}

// ActorClient is a client for the Actor schema.
type ActorClient struct {
	config
}

// NewActorClient returns a client for the Actor from the given config.
func NewActorClient(c config) *ActorClient {
	return &ActorClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `actor.Hooks(f(g(h())))`.
func (c *ActorClient) Use(hooks ...Hook) {
	c.hooks.Actor = append(c.hooks.Actor, hooks...)
}

// Create returns a builder for creating a Actor entity.
func (c *ActorClient) Create() *ActorCreate {
	mutation := newActorMutation(c.config, OpCreate)
	return &ActorCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Actor entities.
func (c *ActorClient) CreateBulk(builders ...*ActorCreate) *ActorCreateBulk {
	return &ActorCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Actor.
func (c *ActorClient) Update() *ActorUpdate {
	mutation := newActorMutation(c.config, OpUpdate)
	return &ActorUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ActorClient) UpdateOne(a *Actor) *ActorUpdateOne {
	mutation := newActorMutation(c.config, OpUpdateOne, withActor(a))
	return &ActorUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ActorClient) UpdateOneID(id uint64) *ActorUpdateOne {
	mutation := newActorMutation(c.config, OpUpdateOne, withActorID(id))
	return &ActorUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Actor.
func (c *ActorClient) Delete() *ActorDelete {
	mutation := newActorMutation(c.config, OpDelete)
	return &ActorDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ActorClient) DeleteOne(a *Actor) *ActorDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *ActorClient) DeleteOneID(id uint64) *ActorDeleteOne {
	builder := c.Delete().Where(actor.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ActorDeleteOne{builder}
}

// Query returns a query builder for Actor.
func (c *ActorClient) Query() *ActorQuery {
	return &ActorQuery{
		config: c.config,
	}
}

// Get returns a Actor entity by its id.
func (c *ActorClient) Get(ctx context.Context, id uint64) (*Actor, error) {
	return c.Query().Where(actor.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ActorClient) GetX(ctx context.Context, id uint64) *Actor {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryServer queries the server edge of a Actor.
func (c *ActorClient) QueryServer(a *Actor) *ServerQuery {
	query := &ServerQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(actor.Table, actor.FieldID, id),
			sqlgraph.To(server.Table, server.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, actor.ServerTable, actor.ServerColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryEvents queries the events edge of a Actor.
func (c *ActorClient) QueryEvents(a *Actor) *EventQuery {
	query := &EventQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(actor.Table, actor.FieldID, id),
			sqlgraph.To(event.Table, event.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, actor.EventsTable, actor.EventsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryOrganizerOf queries the organizer_of edge of a Actor.
func (c *ActorClient) QueryOrganizerOf(a *Actor) *EventQuery {
	query := &EventQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(actor.Table, actor.FieldID, id),
			sqlgraph.To(event.Table, event.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, actor.OrganizerOfTable, actor.OrganizerOfColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryStatuses queries the statuses edge of a Actor.
func (c *ActorClient) QueryStatuses(a *Actor) *StatusQuery {
	query := &StatusQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(actor.Table, actor.FieldID, id),
			sqlgraph.To(status.Table, status.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, actor.StatusesTable, actor.StatusesColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryFollowers queries the followers edge of a Actor.
func (c *ActorClient) QueryFollowers(a *Actor) *ActorQuery {
	query := &ActorQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(actor.Table, actor.FieldID, id),
			sqlgraph.To(actor.Table, actor.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, actor.FollowersTable, actor.FollowersPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryFollowing queries the following edge of a Actor.
func (c *ActorClient) QueryFollowing(a *Actor) *ActorQuery {
	query := &ActorQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(actor.Table, actor.FieldID, id),
			sqlgraph.To(actor.Table, actor.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, actor.FollowingTable, actor.FollowingPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryReactedStatuses queries the reacted_statuses edge of a Actor.
func (c *ActorClient) QueryReactedStatuses(a *Actor) *StatusQuery {
	query := &StatusQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(actor.Table, actor.FieldID, id),
			sqlgraph.To(status.Table, status.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, actor.ReactedStatusesTable, actor.ReactedStatusesPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryModerators queries the moderators edge of a Actor.
func (c *ActorClient) QueryModerators(a *Actor) *ActorQuery {
	query := &ActorQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(actor.Table, actor.FieldID, id),
			sqlgraph.To(actor.Table, actor.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, actor.ModeratorsTable, actor.ModeratorsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryModerating queries the moderating edge of a Actor.
func (c *ActorClient) QueryModerating(a *Actor) *ActorQuery {
	query := &ActorQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(actor.Table, actor.FieldID, id),
			sqlgraph.To(actor.Table, actor.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, actor.ModeratingTable, actor.ModeratingPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryMembers queries the members edge of a Actor.
func (c *ActorClient) QueryMembers(a *Actor) *ActorQuery {
	query := &ActorQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(actor.Table, actor.FieldID, id),
			sqlgraph.To(actor.Table, actor.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, actor.MembersTable, actor.MembersPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryGroups queries the groups edge of a Actor.
func (c *ActorClient) QueryGroups(a *Actor) *ActorQuery {
	query := &ActorQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(actor.Table, actor.FieldID, id),
			sqlgraph.To(actor.Table, actor.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, actor.GroupsTable, actor.GroupsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QuerySessions queries the sessions edge of a Actor.
func (c *ActorClient) QuerySessions(a *Actor) *SessionQuery {
	query := &SessionQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(actor.Table, actor.FieldID, id),
			sqlgraph.To(session.Table, session.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, actor.SessionsTable, actor.SessionsColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryReactions queries the reactions edge of a Actor.
func (c *ActorClient) QueryReactions(a *Actor) *ReactionQuery {
	query := &ReactionQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(actor.Table, actor.FieldID, id),
			sqlgraph.To(reaction.Table, reaction.ActorsColumn),
			sqlgraph.Edge(sqlgraph.O2M, true, actor.ReactionsTable, actor.ReactionsColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ActorClient) Hooks() []Hook {
	hooks := c.hooks.Actor
	return append(hooks[:len(hooks):len(hooks)], actor.Hooks[:]...)
}

// EventClient is a client for the Event schema.
type EventClient struct {
	config
}

// NewEventClient returns a client for the Event from the given config.
func NewEventClient(c config) *EventClient {
	return &EventClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `event.Hooks(f(g(h())))`.
func (c *EventClient) Use(hooks ...Hook) {
	c.hooks.Event = append(c.hooks.Event, hooks...)
}

// Create returns a builder for creating a Event entity.
func (c *EventClient) Create() *EventCreate {
	mutation := newEventMutation(c.config, OpCreate)
	return &EventCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Event entities.
func (c *EventClient) CreateBulk(builders ...*EventCreate) *EventCreateBulk {
	return &EventCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Event.
func (c *EventClient) Update() *EventUpdate {
	mutation := newEventMutation(c.config, OpUpdate)
	return &EventUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *EventClient) UpdateOne(e *Event) *EventUpdateOne {
	mutation := newEventMutation(c.config, OpUpdateOne, withEvent(e))
	return &EventUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *EventClient) UpdateOneID(id uint64) *EventUpdateOne {
	mutation := newEventMutation(c.config, OpUpdateOne, withEventID(id))
	return &EventUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Event.
func (c *EventClient) Delete() *EventDelete {
	mutation := newEventMutation(c.config, OpDelete)
	return &EventDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *EventClient) DeleteOne(e *Event) *EventDeleteOne {
	return c.DeleteOneID(e.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *EventClient) DeleteOneID(id uint64) *EventDeleteOne {
	builder := c.Delete().Where(event.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &EventDeleteOne{builder}
}

// Query returns a query builder for Event.
func (c *EventClient) Query() *EventQuery {
	return &EventQuery{
		config: c.config,
	}
}

// Get returns a Event entity by its id.
func (c *EventClient) Get(ctx context.Context, id uint64) (*Event, error) {
	return c.Query().Where(event.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *EventClient) GetX(ctx context.Context, id uint64) *Event {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryAttendees queries the attendees edge of a Event.
func (c *EventClient) QueryAttendees(e *Event) *ActorQuery {
	query := &ActorQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := e.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(event.Table, event.FieldID, id),
			sqlgraph.To(actor.Table, actor.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, event.AttendeesTable, event.AttendeesPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(e.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryOrganizer queries the organizer edge of a Event.
func (c *EventClient) QueryOrganizer(e *Event) *ActorQuery {
	query := &ActorQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := e.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(event.Table, event.FieldID, id),
			sqlgraph.To(actor.Table, actor.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, event.OrganizerTable, event.OrganizerColumn),
		)
		fromV = sqlgraph.Neighbors(e.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *EventClient) Hooks() []Hook {
	hooks := c.hooks.Event
	return append(hooks[:len(hooks):len(hooks)], event.Hooks[:]...)
}

// GroupClient is a client for the Group schema.
type GroupClient struct {
	config
}

// NewGroupClient returns a client for the Group from the given config.
func NewGroupClient(c config) *GroupClient {
	return &GroupClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `group.Hooks(f(g(h())))`.
func (c *GroupClient) Use(hooks ...Hook) {
	c.hooks.Group = append(c.hooks.Group, hooks...)
}

// Create returns a builder for creating a Group entity.
func (c *GroupClient) Create() *GroupCreate {
	mutation := newGroupMutation(c.config, OpCreate)
	return &GroupCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Group entities.
func (c *GroupClient) CreateBulk(builders ...*GroupCreate) *GroupCreateBulk {
	return &GroupCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Group.
func (c *GroupClient) Update() *GroupUpdate {
	mutation := newGroupMutation(c.config, OpUpdate)
	return &GroupUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *GroupClient) UpdateOne(gr *Group) *GroupUpdateOne {
	mutation := newGroupMutation(c.config, OpUpdateOne, withGroup(gr))
	return &GroupUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *GroupClient) UpdateOneID(id uint64) *GroupUpdateOne {
	mutation := newGroupMutation(c.config, OpUpdateOne, withGroupID(id))
	return &GroupUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Group.
func (c *GroupClient) Delete() *GroupDelete {
	mutation := newGroupMutation(c.config, OpDelete)
	return &GroupDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *GroupClient) DeleteOne(gr *Group) *GroupDeleteOne {
	return c.DeleteOneID(gr.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *GroupClient) DeleteOneID(id uint64) *GroupDeleteOne {
	builder := c.Delete().Where(group.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &GroupDeleteOne{builder}
}

// Query returns a query builder for Group.
func (c *GroupClient) Query() *GroupQuery {
	return &GroupQuery{
		config: c.config,
	}
}

// Get returns a Group entity by its id.
func (c *GroupClient) Get(ctx context.Context, id uint64) (*Group, error) {
	return c.Query().Where(group.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *GroupClient) GetX(ctx context.Context, id uint64) *Group {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryActors queries the actors edge of a Group.
func (c *GroupClient) QueryActors(gr *Group) *ActorQuery {
	query := &ActorQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := gr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(group.Table, group.FieldID, id),
			sqlgraph.To(actor.Table, actor.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, group.ActorsTable, group.ActorsColumn),
		)
		fromV = sqlgraph.Neighbors(gr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryOwners queries the owners edge of a Group.
func (c *GroupClient) QueryOwners(gr *Group) *ActorQuery {
	query := &ActorQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := gr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(group.Table, group.FieldID, id),
			sqlgraph.To(actor.Table, actor.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, group.OwnersTable, group.OwnersColumn),
		)
		fromV = sqlgraph.Neighbors(gr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *GroupClient) Hooks() []Hook {
	hooks := c.hooks.Group
	return append(hooks[:len(hooks):len(hooks)], group.Hooks[:]...)
}

// ReactionClient is a client for the Reaction schema.
type ReactionClient struct {
	config
}

// NewReactionClient returns a client for the Reaction from the given config.
func NewReactionClient(c config) *ReactionClient {
	return &ReactionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `reaction.Hooks(f(g(h())))`.
func (c *ReactionClient) Use(hooks ...Hook) {
	c.hooks.Reaction = append(c.hooks.Reaction, hooks...)
}

// Create returns a builder for creating a Reaction entity.
func (c *ReactionClient) Create() *ReactionCreate {
	mutation := newReactionMutation(c.config, OpCreate)
	return &ReactionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Reaction entities.
func (c *ReactionClient) CreateBulk(builders ...*ReactionCreate) *ReactionCreateBulk {
	return &ReactionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Reaction.
func (c *ReactionClient) Update() *ReactionUpdate {
	mutation := newReactionMutation(c.config, OpUpdate)
	return &ReactionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ReactionClient) UpdateOne(r *Reaction) *ReactionUpdateOne {
	mutation := newReactionMutation(c.config, OpUpdateOne)
	mutation.actors = &r.ActorID
	mutation.status = &r.StatusID
	return &ReactionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Reaction.
func (c *ReactionClient) Delete() *ReactionDelete {
	mutation := newReactionMutation(c.config, OpDelete)
	return &ReactionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Query returns a query builder for Reaction.
func (c *ReactionClient) Query() *ReactionQuery {
	return &ReactionQuery{
		config: c.config,
	}
}

// QueryActors queries the actors edge of a Reaction.
func (c *ReactionClient) QueryActors(r *Reaction) *ActorQuery {
	return c.Query().
		Where(reaction.ActorID(r.ActorID), reaction.StatusID(r.StatusID)).
		QueryActors()
}

// QueryStatus queries the status edge of a Reaction.
func (c *ReactionClient) QueryStatus(r *Reaction) *StatusQuery {
	return c.Query().
		Where(reaction.ActorID(r.ActorID), reaction.StatusID(r.StatusID)).
		QueryStatus()
}

// Hooks returns the client hooks.
func (c *ReactionClient) Hooks() []Hook {
	hooks := c.hooks.Reaction
	return append(hooks[:len(hooks):len(hooks)], reaction.Hooks[:]...)
}

// ServerClient is a client for the Server schema.
type ServerClient struct {
	config
}

// NewServerClient returns a client for the Server from the given config.
func NewServerClient(c config) *ServerClient {
	return &ServerClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `server.Hooks(f(g(h())))`.
func (c *ServerClient) Use(hooks ...Hook) {
	c.hooks.Server = append(c.hooks.Server, hooks...)
}

// Create returns a builder for creating a Server entity.
func (c *ServerClient) Create() *ServerCreate {
	mutation := newServerMutation(c.config, OpCreate)
	return &ServerCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Server entities.
func (c *ServerClient) CreateBulk(builders ...*ServerCreate) *ServerCreateBulk {
	return &ServerCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Server.
func (c *ServerClient) Update() *ServerUpdate {
	mutation := newServerMutation(c.config, OpUpdate)
	return &ServerUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ServerClient) UpdateOne(s *Server) *ServerUpdateOne {
	mutation := newServerMutation(c.config, OpUpdateOne, withServer(s))
	return &ServerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ServerClient) UpdateOneID(id uint64) *ServerUpdateOne {
	mutation := newServerMutation(c.config, OpUpdateOne, withServerID(id))
	return &ServerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Server.
func (c *ServerClient) Delete() *ServerDelete {
	mutation := newServerMutation(c.config, OpDelete)
	return &ServerDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ServerClient) DeleteOne(s *Server) *ServerDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *ServerClient) DeleteOneID(id uint64) *ServerDeleteOne {
	builder := c.Delete().Where(server.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ServerDeleteOne{builder}
}

// Query returns a query builder for Server.
func (c *ServerClient) Query() *ServerQuery {
	return &ServerQuery{
		config: c.config,
	}
}

// Get returns a Server entity by its id.
func (c *ServerClient) Get(ctx context.Context, id uint64) (*Server, error) {
	return c.Query().Where(server.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ServerClient) GetX(ctx context.Context, id uint64) *Server {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryActors queries the actors edge of a Server.
func (c *ServerClient) QueryActors(s *Server) *ActorQuery {
	query := &ActorQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(server.Table, server.FieldID, id),
			sqlgraph.To(actor.Table, actor.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, server.ActorsTable, server.ActorsColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ServerClient) Hooks() []Hook {
	hooks := c.hooks.Server
	return append(hooks[:len(hooks):len(hooks)], server.Hooks[:]...)
}

// SessionClient is a client for the Session schema.
type SessionClient struct {
	config
}

// NewSessionClient returns a client for the Session from the given config.
func NewSessionClient(c config) *SessionClient {
	return &SessionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `session.Hooks(f(g(h())))`.
func (c *SessionClient) Use(hooks ...Hook) {
	c.hooks.Session = append(c.hooks.Session, hooks...)
}

// Create returns a builder for creating a Session entity.
func (c *SessionClient) Create() *SessionCreate {
	mutation := newSessionMutation(c.config, OpCreate)
	return &SessionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Session entities.
func (c *SessionClient) CreateBulk(builders ...*SessionCreate) *SessionCreateBulk {
	return &SessionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Session.
func (c *SessionClient) Update() *SessionUpdate {
	mutation := newSessionMutation(c.config, OpUpdate)
	return &SessionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SessionClient) UpdateOne(s *Session) *SessionUpdateOne {
	mutation := newSessionMutation(c.config, OpUpdateOne, withSession(s))
	return &SessionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SessionClient) UpdateOneID(id uint64) *SessionUpdateOne {
	mutation := newSessionMutation(c.config, OpUpdateOne, withSessionID(id))
	return &SessionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Session.
func (c *SessionClient) Delete() *SessionDelete {
	mutation := newSessionMutation(c.config, OpDelete)
	return &SessionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *SessionClient) DeleteOne(s *Session) *SessionDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *SessionClient) DeleteOneID(id uint64) *SessionDeleteOne {
	builder := c.Delete().Where(session.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SessionDeleteOne{builder}
}

// Query returns a query builder for Session.
func (c *SessionClient) Query() *SessionQuery {
	return &SessionQuery{
		config: c.config,
	}
}

// Get returns a Session entity by its id.
func (c *SessionClient) Get(ctx context.Context, id uint64) (*Session, error) {
	return c.Query().Where(session.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SessionClient) GetX(ctx context.Context, id uint64) *Session {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryAccounts queries the accounts edge of a Session.
func (c *SessionClient) QueryAccounts(s *Session) *ActorQuery {
	query := &ActorQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(session.Table, session.FieldID, id),
			sqlgraph.To(actor.Table, actor.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, session.AccountsTable, session.AccountsColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *SessionClient) Hooks() []Hook {
	hooks := c.hooks.Session
	return append(hooks[:len(hooks):len(hooks)], session.Hooks[:]...)
}

// StatusClient is a client for the Status schema.
type StatusClient struct {
	config
}

// NewStatusClient returns a client for the Status from the given config.
func NewStatusClient(c config) *StatusClient {
	return &StatusClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `status.Hooks(f(g(h())))`.
func (c *StatusClient) Use(hooks ...Hook) {
	c.hooks.Status = append(c.hooks.Status, hooks...)
}

// Create returns a builder for creating a Status entity.
func (c *StatusClient) Create() *StatusCreate {
	mutation := newStatusMutation(c.config, OpCreate)
	return &StatusCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Status entities.
func (c *StatusClient) CreateBulk(builders ...*StatusCreate) *StatusCreateBulk {
	return &StatusCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Status.
func (c *StatusClient) Update() *StatusUpdate {
	mutation := newStatusMutation(c.config, OpUpdate)
	return &StatusUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *StatusClient) UpdateOne(s *Status) *StatusUpdateOne {
	mutation := newStatusMutation(c.config, OpUpdateOne, withStatus(s))
	return &StatusUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *StatusClient) UpdateOneID(id uint64) *StatusUpdateOne {
	mutation := newStatusMutation(c.config, OpUpdateOne, withStatusID(id))
	return &StatusUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Status.
func (c *StatusClient) Delete() *StatusDelete {
	mutation := newStatusMutation(c.config, OpDelete)
	return &StatusDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *StatusClient) DeleteOne(s *Status) *StatusDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *StatusClient) DeleteOneID(id uint64) *StatusDeleteOne {
	builder := c.Delete().Where(status.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &StatusDeleteOne{builder}
}

// Query returns a query builder for Status.
func (c *StatusClient) Query() *StatusQuery {
	return &StatusQuery{
		config: c.config,
	}
}

// Get returns a Status entity by its id.
func (c *StatusClient) Get(ctx context.Context, id uint64) (*Status, error) {
	return c.Query().Where(status.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *StatusClient) GetX(ctx context.Context, id uint64) *Status {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryActors queries the actors edge of a Status.
func (c *StatusClient) QueryActors(s *Status) *ActorQuery {
	query := &ActorQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(status.Table, status.FieldID, id),
			sqlgraph.To(actor.Table, actor.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, status.ActorsTable, status.ActorsColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryReactedActors queries the reacted_actors edge of a Status.
func (c *StatusClient) QueryReactedActors(s *Status) *ActorQuery {
	query := &ActorQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(status.Table, status.FieldID, id),
			sqlgraph.To(actor.Table, actor.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, status.ReactedActorsTable, status.ReactedActorsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryReactions queries the reactions edge of a Status.
func (c *StatusClient) QueryReactions(s *Status) *ReactionQuery {
	query := &ReactionQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(status.Table, status.FieldID, id),
			sqlgraph.To(reaction.Table, reaction.StatusColumn),
			sqlgraph.Edge(sqlgraph.O2M, true, status.ReactionsTable, status.ReactionsColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *StatusClient) Hooks() []Hook {
	hooks := c.hooks.Status
	return append(hooks[:len(hooks):len(hooks)], status.Hooks[:]...)
}
