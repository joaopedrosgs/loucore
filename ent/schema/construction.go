package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/index"
	"time"
)

// Construction holds the schema definition for the Construction entity.
type Construction struct {
	ent.Schema
}

// Fields of the Construction.
func (Construction) Fields() []ent.Field {
	return []ent.Field{
		field.Int("x"),
		field.Int("y"),
		field.Float("raw_production").Default(0),
		field.Float("production").Default(0),
		field.Int("type").Default(0),
		field.Int("level").Default(0),
		field.Float("modifier").Default(1),
		field.Time("last_updated").Default(time.Now),
		field.Bool("need_refresh").Default(true),
	}
}

// Edges of the Construction.
func (Construction) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("city", City.Type).Ref("constructions").Unique().Required(),
		edge.From("owner", User.Type).Ref("constructions").Unique(),
		edge.To("queue", QueueItem.Type),
		edge.To("affects", Construction.Type),
		edge.From("affected_by", Construction.Type).Ref("affects"),
	}
}
func (Construction) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("x", "y").
			Edges("city").
			Unique(),
	}
}
