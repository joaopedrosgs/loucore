package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
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
