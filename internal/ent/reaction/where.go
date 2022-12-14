// Code generated by ent, DO NOT EDIT.

package reaction

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/eriner/burr/internal/ent/predicate"
)

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Reaction {
	return predicate.Reaction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Reaction {
	return predicate.Reaction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v int) predicate.Reaction {
	return predicate.Reaction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedBy), v))
	})
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v int) predicate.Reaction {
	return predicate.Reaction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedBy), v))
	})
}

// ActorID applies equality check predicate on the "actor_id" field. It's identical to ActorIDEQ.
func ActorID(v uint64) predicate.Reaction {
	return predicate.Reaction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldActorID), v))
	})
}

// StatusID applies equality check predicate on the "status_id" field. It's identical to StatusIDEQ.
func StatusID(v uint64) predicate.Reaction {
	return predicate.Reaction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatusID), v))
	})
}

// Dat applies equality check predicate on the "dat" field. It's identical to DatEQ.
func Dat(v uint64) predicate.Reaction {
	return predicate.Reaction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDat), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Reaction {
	return predicate.Reaction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Reaction {
	return predicate.Reaction(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Reaction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Reaction(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Reaction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Reaction(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Reaction {
	return predicate.Reaction(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Reaction {
	return predicate.Reaction(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Reaction {
	return predicate.Reaction(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Reaction {
	return predicate.Reaction(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Reaction {
	return predicate.Reaction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Reaction {
	return predicate.Reaction(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Reaction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Reaction(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Reaction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Reaction(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Reaction {
	return predicate.Reaction(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Reaction {
	return predicate.Reaction(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Reaction {
	return predicate.Reaction(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Reaction {
	return predicate.Reaction(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v int) predicate.Reaction {
	return predicate.Reaction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedBy), v))
	})
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v int) predicate.Reaction {
	return predicate.Reaction(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedBy), v))
	})
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...int) predicate.Reaction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Reaction(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedBy), v...))
	})
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...int) predicate.Reaction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Reaction(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedBy), v...))
	})
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v int) predicate.Reaction {
	return predicate.Reaction(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedBy), v))
	})
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v int) predicate.Reaction {
	return predicate.Reaction(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedBy), v))
	})
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v int) predicate.Reaction {
	return predicate.Reaction(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedBy), v))
	})
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v int) predicate.Reaction {
	return predicate.Reaction(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedBy), v))
	})
}

// CreatedByIsNil applies the IsNil predicate on the "created_by" field.
func CreatedByIsNil() predicate.Reaction {
	return predicate.Reaction(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCreatedBy)))
	})
}

// CreatedByNotNil applies the NotNil predicate on the "created_by" field.
func CreatedByNotNil() predicate.Reaction {
	return predicate.Reaction(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCreatedBy)))
	})
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v int) predicate.Reaction {
	return predicate.Reaction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedBy), v))
	})
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v int) predicate.Reaction {
	return predicate.Reaction(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedBy), v))
	})
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...int) predicate.Reaction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Reaction(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedBy), v...))
	})
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...int) predicate.Reaction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Reaction(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedBy), v...))
	})
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v int) predicate.Reaction {
	return predicate.Reaction(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedBy), v))
	})
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v int) predicate.Reaction {
	return predicate.Reaction(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedBy), v))
	})
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v int) predicate.Reaction {
	return predicate.Reaction(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedBy), v))
	})
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v int) predicate.Reaction {
	return predicate.Reaction(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedBy), v))
	})
}

// UpdatedByIsNil applies the IsNil predicate on the "updated_by" field.
func UpdatedByIsNil() predicate.Reaction {
	return predicate.Reaction(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUpdatedBy)))
	})
}

// UpdatedByNotNil applies the NotNil predicate on the "updated_by" field.
func UpdatedByNotNil() predicate.Reaction {
	return predicate.Reaction(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUpdatedBy)))
	})
}

// ActorIDEQ applies the EQ predicate on the "actor_id" field.
func ActorIDEQ(v uint64) predicate.Reaction {
	return predicate.Reaction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldActorID), v))
	})
}

// ActorIDNEQ applies the NEQ predicate on the "actor_id" field.
func ActorIDNEQ(v uint64) predicate.Reaction {
	return predicate.Reaction(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldActorID), v))
	})
}

// ActorIDIn applies the In predicate on the "actor_id" field.
func ActorIDIn(vs ...uint64) predicate.Reaction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Reaction(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldActorID), v...))
	})
}

// ActorIDNotIn applies the NotIn predicate on the "actor_id" field.
func ActorIDNotIn(vs ...uint64) predicate.Reaction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Reaction(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldActorID), v...))
	})
}

// StatusIDEQ applies the EQ predicate on the "status_id" field.
func StatusIDEQ(v uint64) predicate.Reaction {
	return predicate.Reaction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatusID), v))
	})
}

// StatusIDNEQ applies the NEQ predicate on the "status_id" field.
func StatusIDNEQ(v uint64) predicate.Reaction {
	return predicate.Reaction(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatusID), v))
	})
}

// StatusIDIn applies the In predicate on the "status_id" field.
func StatusIDIn(vs ...uint64) predicate.Reaction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Reaction(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldStatusID), v...))
	})
}

// StatusIDNotIn applies the NotIn predicate on the "status_id" field.
func StatusIDNotIn(vs ...uint64) predicate.Reaction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Reaction(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldStatusID), v...))
	})
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v Type) predicate.Reaction {
	return predicate.Reaction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldType), v))
	})
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v Type) predicate.Reaction {
	return predicate.Reaction(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldType), v))
	})
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...Type) predicate.Reaction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Reaction(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldType), v...))
	})
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...Type) predicate.Reaction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Reaction(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldType), v...))
	})
}

// DatEQ applies the EQ predicate on the "dat" field.
func DatEQ(v uint64) predicate.Reaction {
	return predicate.Reaction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDat), v))
	})
}

// DatNEQ applies the NEQ predicate on the "dat" field.
func DatNEQ(v uint64) predicate.Reaction {
	return predicate.Reaction(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDat), v))
	})
}

// DatIn applies the In predicate on the "dat" field.
func DatIn(vs ...uint64) predicate.Reaction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Reaction(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDat), v...))
	})
}

// DatNotIn applies the NotIn predicate on the "dat" field.
func DatNotIn(vs ...uint64) predicate.Reaction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Reaction(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDat), v...))
	})
}

// DatGT applies the GT predicate on the "dat" field.
func DatGT(v uint64) predicate.Reaction {
	return predicate.Reaction(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDat), v))
	})
}

// DatGTE applies the GTE predicate on the "dat" field.
func DatGTE(v uint64) predicate.Reaction {
	return predicate.Reaction(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDat), v))
	})
}

// DatLT applies the LT predicate on the "dat" field.
func DatLT(v uint64) predicate.Reaction {
	return predicate.Reaction(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDat), v))
	})
}

// DatLTE applies the LTE predicate on the "dat" field.
func DatLTE(v uint64) predicate.Reaction {
	return predicate.Reaction(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDat), v))
	})
}

// HasActors applies the HasEdge predicate on the "actors" edge.
func HasActors() predicate.Reaction {
	return predicate.Reaction(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, ActorsColumn),
			sqlgraph.To(ActorsInverseTable, ActorFieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, ActorsTable, ActorsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasActorsWith applies the HasEdge predicate on the "actors" edge with a given conditions (other predicates).
func HasActorsWith(preds ...predicate.Actor) predicate.Reaction {
	return predicate.Reaction(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, ActorsColumn),
			sqlgraph.To(ActorsInverseTable, ActorFieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, ActorsTable, ActorsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasStatus applies the HasEdge predicate on the "status" edge.
func HasStatus() predicate.Reaction {
	return predicate.Reaction(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, StatusColumn),
			sqlgraph.To(StatusInverseTable, StatusFieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, StatusTable, StatusColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStatusWith applies the HasEdge predicate on the "status" edge with a given conditions (other predicates).
func HasStatusWith(preds ...predicate.Status) predicate.Reaction {
	return predicate.Reaction(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, StatusColumn),
			sqlgraph.To(StatusInverseTable, StatusFieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, StatusTable, StatusColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Reaction) predicate.Reaction {
	return predicate.Reaction(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Reaction) predicate.Reaction {
	return predicate.Reaction(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Reaction) predicate.Reaction {
	return predicate.Reaction(func(s *sql.Selector) {
		p(s.Not())
	})
}
