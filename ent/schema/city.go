package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/index"
	"time"
)

// City holds the schema definition for the City entity.
type City struct {
	ent.Schema
}

// Fields of the City.
func (City) Fields() []ent.Field {
	return []ent.Field{
		field.Int("x").Default(0),
		field.Int("y").Default(0),
		field.String("name").NotEmpty().Default("New city"),
		field.Int("points").Default(3),

		field.Float("wood_production").Default(5),
		field.Float("stone_production").Default(0),
		field.Float("iron_production").Default(0),
		field.Float("food_production").Default(0),

		field.Float("wood_stored").Default(300),
		field.Float("stone_stored").Default(0),
		field.Float("iron_stored").Default(0),
		field.Float("food_stored").Default(0),

		field.Float("wood_limit").Default(300),
		field.Float("stone_limit").Default(0),
		field.Float("iron_limit").Default(0),
		field.Float("food_limit").Default(0),

		field.Time("queue_ends_at").Default(time.Now),
		field.Int("construction_speed").Default(1),

		field.Time("last_updated").Default(time.Now),
	}
}
func (City) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("x", "y").
			Unique(),
	}
}

// Edges of the City.
func (City) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Unique().Ref("cities"),
		edge.To("constructions", Construction.Type),
		edge.To("queue", QueueItem.Type),
	}
}
