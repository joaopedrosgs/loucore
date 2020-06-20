package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/index"
)

// Queue holds the schema definition for the Queue entity.
type QueueItem struct {
	ent.Schema
}

// Fields of the QueueItem.
func (QueueItem) Fields() []ent.Field {
	return []ent.Field{
		field.Time("start_at"),
		field.Int("duration"),
		field.Time("completion"),
		field.Int("action"),
		field.Int("order"),
	}
}

// Edges of the QueueItem.
func (QueueItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Ref("queue").Unique(),
		edge.From("city", City.Type).Ref("queue").Unique(),
		edge.From("construction", Construction.Type).Ref("queue").Unique(),
	}
}

func (QueueItem) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("order").
			Edges("city").
			Unique(),
	}
}
