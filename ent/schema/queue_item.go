package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/index"
)

// Queue holds the schema definition for the Queue entity.
type QueueItem struct {
	ent.Schema
}

// Fields of the QueueItem.
func (QueueItem) Fields() []ent.Field {
	return []ent.Field{
		field.Int("duration"),
		field.Int("action"),
		field.Int("position"),
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

// Index of the QueueItem.
func (QueueItem) Index() []ent.Index {
	return []ent.Index{
		// non-unique index.
		index.Fields("position"),
		// unique index.
		index.Fields("position").
			Edges("city").
			Unique(),
	}
}
