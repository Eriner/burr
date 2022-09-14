package schema

import "entgo.io/ent"

// Server holds the schema definition for the Server entity.
type Server struct {
	ent.Schema
}

func (Server) Fields() []ent.Field {
	return nil
}

func (Server) Edges() []ent.Edge {
	return nil
}

func (Server) Mixin() []ent.Mixin {
	return []ent.Mixin{
		AuditMixin{},
	}
}
