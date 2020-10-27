package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Unique(),
		field.String("email").NotEmpty().Unique(),
		field.String("password_hash").NotEmpty(),
		field.Int("gold").Default(0),
		field.Int("diamonds").Default(0),
		field.Int("darkwood").Default(0),
		field.Int("runestone").Default(0),
		field.Int("veritium").Default(0),
		field.Int("trueseed").Default(0),
		field.Int("rank").Default(0),
		field.Int("alliance_rank").Default(0),
		field.Time("last_updated").Default(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("cities", City.Type),
		edge.To("queue", QueueItem.Type),
		edge.To("constructions", Construction.Type),
	}
}
